package playable

import (
	"strings"
	"testing"
)

func TestGuidanceIsBoundedAndConceptsAreFree(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest(), Progress: progressCompleteBefore(t, "incident-hotfix-001")})
	for i := 0; i < 2; i++ {
		m, _ = updateWithKey(t, m, "?")
	}
	if m.run.State().HintsUsed != 0 {
		t.Fatal("concept guidance charged")
	}
	for i := 0; i < 8; i++ {
		m, _ = updateWithKey(t, m, "?")
	}
	if m.hintRequests != 3 || m.run.State().HintsUsed != 1 {
		t.Fatal(m.hintRequests, m.run.State().HintsUsed)
	}
	m, _ = updateWithKey(t, m, "/error\n")
	if m.State().Status != "succeeded" || m.State().Score.Grade == "F" || !strings.Contains(m.View(), "도움받아 완료") {
		t.Fatal(m.View())
	}
}

func TestReviewCompletesWithoutReplayingMasteredTutorial(t *testing.T) {
	p := progressWithAllPlayableCompleted(t)
	first := p.Missions["normal-motion-basic-001"]
	first.BestGrade = "A"
	p.Missions["normal-motion-basic-001"] = first
	m := New(Options{ContentRoot: contentRootForTest(), Progress: p})
	m, _ = updateWithKey(t, m, "m")
	m, _ = updateWithKey(t, m, "4")
	if m.currentExerciseID() != "normal-motion-basic-001" {
		t.Fatal(m.currentExerciseID())
	}
	m, _ = updateWithKey(t, m, "ll")
	m, _ = updateWithKey(t, m, "enter")
	if m.currentExerciseID() == "normal-motion-basic-002" || !strings.Contains(m.View(), "출격 경로") {
		t.Fatal(m.View())
	}
}
