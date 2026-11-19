// zeo-app-memory-guard is a stateless, unprivileged memory-policy observer.
// It never starts/stops processes or contacts a control API. The App units bind
// their lifecycle to this observer through systemd BindsTo= and After=.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const mib = uint64(1024 * 1024)

func available(raw []byte) (uint64, error) {
	if len(raw) > 65536 {
		return 0, errors.New("meminfo exceeds bound")
	}
	var value uint64
	found := false
	for _, line := range strings.Split(string(raw), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 || fields[0] != "MemAvailable:" {
			continue
		}
		if found || len(fields) != 3 || fields[2] != "kB" {
			return 0, errors.New("ambiguous meminfo")
		}
		amount, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil || amount > math.MaxUint64/1024 {
			return 0, errors.New("invalid meminfo amount")
		}
		value = amount * 1024
		found = true
	}
	if !found {
		return 0, errors.New("MemAvailable missing")
	}
	return value, nil
}

func readMeminfo() ([]byte, error) {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(io.LimitReader(f, 65537))
}

func evaluate(read func() ([]byte, error), minimum uint64, out io.Writer) int {
	raw, err := read()
	value := uint64(0)
	if err == nil {
		value, err = available(raw)
	}
	if err != nil {
		_ = json.NewEncoder(out).Encode(map[string]string{"state": "memory_observation_unavailable"})
		return 2 // No usable observation is never healthy.
	}
	if value < minimum {
		_ = json.NewEncoder(out).Encode(map[string]any{"state": "app_memory_floor_crossed", "available_bytes": value, "minimum_bytes": minimum})
		return 1
	}
	return 0
}

func monitor(ctx context.Context, read func() ([]byte, error), minimum uint64, ticks <-chan time.Time, out io.Writer) int {
	if code := evaluate(read, minimum, out); code != 0 {
		return code
	}
	for {
		select {
		case <-ctx.Done():
			return 0 // A planned stop still deactivates the bound App units.
		case _, ok := <-ticks:
			if !ok {
				return 2
			}
			if code := evaluate(read, minimum, out); code != 0 {
				return code
			}
		}
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer, read func() ([]byte, error)) int {
	f := flag.NewFlagSet("zeo-app-memory-guard", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	minimum := f.Uint64("minimum-mib", 600, "App-only abort floor; installation profile owns adoption")
	interval := f.Duration("interval", 200*time.Millisecond, "observation interval")
	check := f.Bool("check", false, "perform one observation")
	err := f.Parse(args)
	if errors.Is(err, flag.ErrHelp) {
		_, _ = io.WriteString(stdout, "zeo-app-memory-guard [--check] [--minimum-mib 600] [--interval 200ms]\n")
		return 0
	}
	if err != nil || f.NArg() != 0 || *minimum == 0 || *minimum > math.MaxUint64/mib || *interval < 50*time.Millisecond || *interval > 10*time.Second {
		_, _ = io.WriteString(stderr, "invalid memory-guard arguments\n")
		return 2
	}
	if *check {
		return evaluate(read, *minimum*mib, stdout)
	}
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()
	return monitor(ctx, read, *minimum*mib, ticker.C, stdout)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr, readMeminfo))
}
