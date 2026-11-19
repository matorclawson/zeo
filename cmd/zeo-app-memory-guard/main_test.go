package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

func mem(mibValue uint64) []byte {
	return []byte(fmt.Sprintf("MemTotal: 4000000 kB\nMemAvailable:\t%d kB\n", mibValue*1024))
}

func TestMemAvailableIsBytesAndDoesNotUseMemFree(t *testing.T) {
	v, err := available(append([]byte("MemFree: 1 kB\n"), mem(600)...))
	if err != nil || v != 600*mib {
		t.Fatal(v, err)
	}
}

func TestMissingDuplicateMalformedAndOverflowObservationRefuse(t *testing.T) {
	for _, raw := range []string{"MemFree: 700000 kB\n", "MemAvailable: 1 kB\nMemAvailable: 2 kB\n", "MemAvailable: -1 kB\n", "MemAvailable: 4 MB\n", "MemAvailable: 18446744073709551615 kB\n", "MemAvailable: 4 kB extra\n", strings.Repeat("x", 65537)} {
		if _, err := available([]byte(raw)); err == nil {
			t.Fatal("accepted malformed/missing observation")
		}
	}
}

func TestCheckThresholdBoundaryAndNoObservationIsNotHealthy(t *testing.T) {
	for _, value := range []uint64{599, 600, 601} {
		var out, errout bytes.Buffer
		code := run(context.Background(), []string{"--check", "--minimum-mib=600"}, &out, &errout, func() ([]byte, error) { return mem(value), nil })
		if (value < 600 && code != 1) || (value >= 600 && code != 0) {
			t.Fatal(value, code)
		}
	}
	var out bytes.Buffer
	if code := evaluate(func() ([]byte, error) { return nil, errors.New("private path detail") }, 600*mib, &out); code != 2 || strings.Contains(out.String(), "private path detail") {
		t.Fatal("unreadable observation leaked details or became healthy")
	}
}

func TestHealthyThenLowObservationEndsMonitor(t *testing.T) {
	ticks := make(chan time.Time, 1)
	ticks <- time.Time{}
	reads := 0
	var out bytes.Buffer
	code := monitor(context.Background(), func() ([]byte, error) {
		reads++
		if reads == 1 {
			return mem(700), nil
		}
		return mem(599), nil
	}, 600*mib, ticks, &out)
	if code != 1 || reads != 2 || !strings.Contains(out.String(), "app_memory_floor_crossed") {
		t.Fatal(code, reads, out.String())
	}
}

func TestHealthyThenLostObservationFailsClosed(t *testing.T) {
	ticks := make(chan time.Time, 1)
	ticks <- time.Time{}
	reads := 0
	var out bytes.Buffer
	code := monitor(context.Background(), func() ([]byte, error) {
		reads++
		if reads == 1 {
			return mem(700), nil
		}
		return []byte("MemFree: 1 kB"), nil
	}, 600*mib, ticks, &out)
	if code != 2 || reads != 2 {
		t.Fatal(code, reads)
	}
}

func TestPlannedCancellationStopsWithoutAnotherObservation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	reads := 0
	var out bytes.Buffer
	code := monitor(ctx, func() ([]byte, error) { reads++; return mem(700), nil }, 600*mib, make(chan time.Time), &out)
	if code != 0 || reads != 1 {
		t.Fatal(code, reads)
	}
}

func TestClosedObservationClockDoesNotReportHealthy(t *testing.T) {
	ticks := make(chan time.Time)
	close(ticks)
	var out bytes.Buffer
	if monitor(context.Background(), func() ([]byte, error) { return mem(700), nil }, 600*mib, ticks, &out) != 2 {
		t.Fatal("lost monitor clock reported healthy")
	}
}

func TestInvalidBudgetOrUnboundedIntervalRefusesBeforeRead(t *testing.T) {
	for _, args := range [][]string{{"--minimum-mib=0"}, {"--minimum-mib=18446744073709551615"}, {"--interval=1ms"}, {"--interval=11s"}, {"unexpected"}} {
		var out, errout bytes.Buffer
		reads := 0
		if run(context.Background(), args, &out, &errout, func() ([]byte, error) { reads++; return mem(700), nil }) != 2 || reads != 0 {
			t.Fatal("invalid profile read/entered monitor", args)
		}
	}
}
