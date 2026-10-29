package playable

import (
	"strings"
	"testing"
)

func TestFirstRecoveryRouteIsShortAndPreservesLibrary(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest()})
	m, _ = updateWithKey(t, m, "m")
	if !strings.Contains(m.View(), "첫 복구") {
		t.Fatal(m.View())
	}
	m, _ = updateWithKey(t, m, "1")
	for _, keys := range []string{"ll", "ww", "/error\n", "ciwup"} {
		m, _ = updateWithKey(t, m, keys)
		if strings.HasPrefix(keys, "ciw") {
			m, _ = updateWithKey(t, m, "esc")
		}
		if m.State().Status != "succeeded" {
			t.Fatalf("%s: %+v", keys, m.State())
		}
		m, _ = updateWithKey(t, m, "enter")
	}
	if m.currentExerciseID() != "incident-hotfix-001" {
		t.Fatal(m.currentExerciseID())
	}
	m, _ = updateWithKey(t, m, "m")
	m, _ = updateWithKey(t, m, "2")
	if m.currentExerciseID() != "normal-motion-basic-002" {
		t.Fatal(m.currentExerciseID())
	}
}
