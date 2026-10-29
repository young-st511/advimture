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
	m, _ = updateWithKey(t, m, "/error\n")
	if strings.Contains(m.View(), "timeout") || !strings.Contains(m.View(), "상태값을 복구") {
		t.Fatal(m.View())
	}
	m, _ = updateWithKey(t, m, "m")
	m, _ = updateWithKey(t, m, "2")
	if m.currentExerciseID() != "normal-motion-basic-002" {
		t.Fatal(m.currentExerciseID())
	}
}

func TestLibraryListsAllEpisodesWithoutWritingProgress(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest()})
	before := len(m.progress.Missions)
	m, _ = updateWithKey(t, m, "m")
	m, _ = updateWithKey(t, m, "6")
	if len(m.chapterStarts()) != 27 {
		t.Fatal(len(m.chapterStarts()))
	}
	m, _ = updateWithKey(t, m, "j")
	m, _ = updateWithKey(t, m, "esc")
	m, _ = updateWithKey(t, m, "esc")
	if m.routeMenu || len(m.progress.Missions) != before {
		t.Fatal("cancel mutated progress or menu state")
	}
}
