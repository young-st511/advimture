package review

import (
	"path/filepath"
	"testing"

	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/progress"
)

func TestCandidatesPrioritizeIncompleteThenLowGradeThenKeyCount(t *testing.T) {
	lib := loadLibrary(t)
	progressState := progress.Progress{
		Missions: map[string]progress.MissionProgress{
			"normal-motion-basic-001": {Completed: true, BestGrade: "S", BestKeystrokes: 2},
			"normal-motion-basic-002": {Completed: true, BestGrade: "B", BestKeystrokes: 1},
			"normal-motion-basic-003": {Completed: true, BestGrade: "S", BestKeystrokes: 4},
		},
	}

	candidates := Candidates(lib, progressState, Options{
		OrderedExerciseIDs: []string{
			"normal-motion-basic-001",
			"normal-motion-basic-002",
			"normal-motion-basic-003",
			"normal-motion-basic-004",
		},
		Limit: 3,
	})

	assertCandidates(t, candidates, []Candidate{
		{ExerciseID: "normal-motion-basic-004", Reason: ReasonIncomplete},
		{ExerciseID: "normal-motion-basic-002", Reason: ReasonLowGrade},
		{ExerciseID: "normal-motion-basic-003", Reason: ReasonKeyCount},
	})
}

func TestCandidatesReturnsStableOrderedIncompleteExercises(t *testing.T) {
	lib := loadLibrary(t)
	progressState := progress.Progress{Missions: map[string]progress.MissionProgress{}}

	candidates := Candidates(lib, progressState, Options{
		OrderedExerciseIDs: []string{"search-basic-002", "search-basic-001"},
		Limit:              2,
	})

	assertCandidates(t, candidates, []Candidate{
		{ExerciseID: "search-basic-002", Reason: ReasonIncomplete},
		{ExerciseID: "search-basic-001", Reason: ReasonIncomplete},
	})
}

func TestCandidateSummaryExplainsReason(t *testing.T) {
	got := Candidate{
		Title:          "다음 timeout 찾기",
		Reason:         ReasonKeyCount,
		BestKeystrokes: 12,
		OptimalKeys:    10,
	}.Summary()
	want := "다음 timeout 찾기: 복구 입력 12/10 keys"
	if got != want {
		t.Fatalf("Summary() = %q, want %q", got, want)
	}
}

func loadLibrary(t *testing.T) content.Library {
	t.Helper()
	lib, err := content.LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}
	return lib
}

func assertCandidates(t *testing.T, got []Candidate, want []Candidate) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len(candidates) = %d, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i].ExerciseID != want[i].ExerciseID || got[i].Reason != want[i].Reason {
			t.Fatalf("candidate[%d] = %s/%s, want %s/%s", i, got[i].ExerciseID, got[i].Reason, want[i].ExerciseID, want[i].Reason)
		}
	}
}
