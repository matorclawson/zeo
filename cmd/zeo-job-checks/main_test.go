package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Reproduce the two checks lost during the historical v8 manual repin. These
// names are stable obligations, independent of the source commit or method name.
func fixture(t *testing.T) (string, []string, catalog) {
	t.Helper()
	dir := t.TempDir()
	body := []byte("def test_old_stamp_refused(): pass\nassert state.source == HEAD\n")
	if err := os.WriteFile(filepath.Join(dir, "owned.txt"), body, 0600); err != nil {
		t.Fatal(err)
	}
	c := catalog{Schema: "zeo.job-checks/v1", Sources: []source{{Path: "owned.txt", SHA256: digest(body)}}, Checks: []check{
		{ID: "app.old-stamp-refusal", Source: "owned.txt", Contains: "def test_old_stamp_refused", Argv: []string{"python3", "-m", "unittest", "scanner"}},
		{ID: "app.consumer-source-binding", Source: "owned.txt", Contains: "assert state.source == HEAD", Argv: []string{"node", "harness.mjs", "${OUT}"}},
	}}
	b, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"baseline.json", "current.json"} {
		if err := os.WriteFile(filepath.Join(dir, name), b, 0600); err != nil {
			t.Fatal(err)
		}
	}
	args := []string{"--baseline", filepath.Join(dir, "baseline.json"), "--baseline-sha256", digest(b), "--catalog", filepath.Join(dir, "current.json"), "--inputs", dir}
	return dir, args, c
}

func writeCurrent(t *testing.T, dir string, c catalog) {
	t.Helper()
	b, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "current.json"), b, 0600); err != nil {
		t.Fatal(err)
	}
}

func TestLostScannerAndConsumerChecksFailBeforeAnyOutput(t *testing.T) {
	for _, removed := range []int{0, 1} {
		t.Run([]string{"old-stamp-refusal", "consumer-binding"}[removed], func(t *testing.T) {
			dir, args, c := fixture(t)
			c.Checks = append(c.Checks[:removed:removed], c.Checks[removed+1:]...)
			writeCurrent(t, dir, c)
			var out, errs bytes.Buffer
			if run(args, &out, &errs) != 1 || out.Len() != 0 || !strings.Contains(errs.String(), "dropped approved check") {
				t.Fatal("lost check accepted or partial job emitted", out.String(), errs.String())
			}
		})
	}
}

func TestRenamingAnObligationNeedsExplicitRetirement(t *testing.T) {
	dir, args, c := fixture(t)
	c.Checks[0].ID = "app.renamed-refusal"
	writeCurrent(t, dir, c)
	var out, errs bytes.Buffer
	if run(args, &out, &errs) != 1 || out.Len() != 0 {
		t.Fatal("renamed ID silently retired approved obligation")
	}
}

func TestRetirementRequiresNamedMissingIDAndReason(t *testing.T) {
	for _, reason := range []string{"", "   ", "Retired by owner decision example-17; superseded by mandatory source binding"} {
		t.Run(reason, func(t *testing.T) {
			dir, args, c := fixture(t)
			id := c.Checks[0].ID
			c.Checks = c.Checks[1:]
			c.Retired = []retirement{{ID: id, Reason: reason}}
			writeCurrent(t, dir, c)
			var out, errs bytes.Buffer
			code := run(args, &out, &errs)
			if strings.TrimSpace(reason) == "" {
				if code != 1 || out.Len() != 0 {
					t.Fatal("empty retirement reason accepted")
				}
			} else if code != 0 || !strings.Contains(out.String(), "consumer-source-binding") {
				t.Fatal("explicit retirement failed", errs.String())
			}
		})
	}
	for _, id := range []string{"app.unknown", "app.consumer-source-binding"} {
		dir, args, c := fixture(t)
		c.Retired = []retirement{{ID: id, Reason: "example"}}
		writeCurrent(t, dir, c)
		var out, errs bytes.Buffer
		if run(args, &out, &errs) != 1 || out.Len() != 0 {
			t.Fatal("unknown or still-active retirement accepted")
		}
	}
}

func TestRepinnedOwnerFileStillMustContainEachApprovedAnchor(t *testing.T) {
	dir, args, c := fixture(t)
	b := []byte("def test_old_stamp_refused(): pass\n")
	if err := os.WriteFile(filepath.Join(dir, "owned.txt"), b, 0600); err != nil {
		t.Fatal(err)
	}
	c.Sources[0].SHA256 = digest(b)
	writeCurrent(t, dir, c)
	var out, errs bytes.Buffer
	if run(args, &out, &errs) != 1 || out.Len() != 0 || !strings.Contains(errs.String(), "missing pinned check anchor") {
		t.Fatal("new owner file silently lost consumer assertion", errs.String())
	}
}

func TestChangedOwnerBytesAndChangedBaselineAreRefused(t *testing.T) {
	for _, name := range []string{"owned.txt", "baseline.json"} {
		dir, args, _ := fixture(t)
		if err := os.WriteFile(filepath.Join(dir, name), []byte("replacement"), 0600); err != nil {
			t.Fatal(err)
		}
		var out, errs bytes.Buffer
		if run(args, &out, &errs) != 1 || out.Len() != 0 {
			t.Fatal("changed pinned input accepted", name)
		}
	}
}

func TestGenerateAllChecksOncePerInvocationWithoutExecutingThem(t *testing.T) {
	dir, args, c := fixture(t)
	marker := filepath.Join(dir, "must-not-exist")
	c.Checks[0].Argv = []string{"touch", marker, "a'b", "$(must-stay-literal)"}
	extra := c.Checks[0]
	extra.ID = "app.another-refusal"
	c.Checks = append(c.Checks, extra)
	writeCurrent(t, dir, c)
	var out, errs bytes.Buffer
	if run(args, &out, &errs) != 0 {
		t.Fatal(errs.String())
	}
	s := out.String()
	if strings.Count(s, "'touch'") != 1 || !strings.Contains(s, "# check app.another-refusal") || !strings.Contains(s, `"${OUT:?required by check catalog}"`) || !strings.Contains(s, `'a'"'"'b'`) || !strings.Contains(s, "'$(must-stay-literal)'") {
		t.Fatal("incomplete or unsafe invocation fragment", s)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("generator executed a check")
	}
}

func TestDuplicateIDsAndUnknownCatalogFieldsRefuse(t *testing.T) {
	dir, args, c := fixture(t)
	c.Checks = append(c.Checks, c.Checks[0])
	writeCurrent(t, dir, c)
	var out, errs bytes.Buffer
	if run(args, &out, &errs) != 1 || out.Len() != 0 {
		t.Fatal("duplicate obligation accepted")
	}
	_, _, valid := fixture(t)
	b, _ := json.Marshal(valid)
	b = append([]byte(`{"ignored":true,`), b[1:]...)
	if _, err := decode(b); err == nil {
		t.Fatal("unknown catalog field accepted")
	}
}
