package review

import (
	"fmt"
	"sort"

	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/progress"
)

type Reason string

const (
	ReasonIncomplete Reason = "incomplete"
	ReasonLowGrade   Reason = "low_grade"
	ReasonKeyCount   Reason = "key_count"
)

type Candidate struct {
	ExerciseID       string
	Title            string
	CommandClusterID string
	Reason           Reason
	BestGrade        string
	BestKeystrokes   int
	OptimalKeys      int
	Priority         int
}

type Options struct {
	OrderedExerciseIDs []string
	Limit              int
}

func Candidates(library content.Library, progressState progress.Progress, options Options) []Candidate {
	orderedIDs := orderedExerciseIDs(library, options.OrderedExerciseIDs)
	limit := options.Limit
	if limit <= 0 {
		limit = len(orderedIDs)
	}

	candidates := make([]Candidate, 0)
	for _, id := range orderedIDs {
		exercise, ok := library.Exercises[id]
		if !ok || !isReviewableExercise(exercise) {
			continue
		}
		if candidate, ok := candidateForExercise(exercise, progressState); ok {
			candidates = append(candidates, candidate)
		}
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].Priority != candidates[j].Priority {
			return candidates[i].Priority < candidates[j].Priority
		}
		return false
	})
	if len(candidates) > limit {
		return append([]Candidate(nil), candidates[:limit]...)
	}
	return append([]Candidate(nil), candidates...)
}

func (c Candidate) Summary() string {
	switch c.Reason {
	case ReasonIncomplete:
		return fmt.Sprintf("%s: 미복구", c.Title)
	case ReasonLowGrade:
		grade := c.BestGrade
		if grade == "" {
			grade = "-"
		}
		return fmt.Sprintf("%s: 복구 등급 %s", c.Title, grade)
	case ReasonKeyCount:
		return fmt.Sprintf("%s: 복구 입력 %d/%d keys", c.Title, c.BestKeystrokes, c.OptimalKeys)
	default:
		return c.Title
	}
}

func (c Candidate) DailyRouteLabel() string {
	switch c.Reason {
	case ReasonIncomplete:
		return fmt.Sprintf("%s(미복구)", c.Title)
	case ReasonLowGrade:
		grade := c.BestGrade
		if grade == "" {
			grade = "-"
		}
		return fmt.Sprintf("%s(등급 %s)", c.Title, grade)
	case ReasonKeyCount:
		return fmt.Sprintf("%s(%d/%d keys)", c.Title, c.BestKeystrokes, c.OptimalKeys)
	default:
		return c.Title
	}
}

func orderedExerciseIDs(library content.Library, preferred []string) []string {
	seen := make(map[string]bool)
	ids := make([]string, 0, len(library.Exercises))
	for _, id := range preferred {
		if _, ok := library.Exercises[id]; ok && !seen[id] {
			seen[id] = true
			ids = append(ids, id)
		}
	}
	if len(ids) > 0 {
		return ids
	}
	var rest []string
	for id := range library.Exercises {
		rest = append(rest, id)
	}
	sort.Strings(rest)
	ids = append(ids, rest...)
	return ids
}

func isReviewableExercise(exercise content.ExerciseDocument) bool {
	return (exercise.Status == content.StatusApproved || exercise.Status == content.StatusImplemented) &&
		exercise.EngineSupport == content.EngineSupportImplemented &&
		exercise.ReplayStatus == content.ReplayStatusPass
}

func candidateForExercise(exercise content.ExerciseDocument, progressState progress.Progress) (Candidate, bool) {
	mission, ok := progressState.Missions[exercise.ID]
	base := Candidate{
		ExerciseID:       exercise.ID,
		Title:            exercise.Title,
		CommandClusterID: exercise.CommandCluster,
		BestGrade:        mission.BestGrade,
		BestKeystrokes:   mission.BestKeystrokes,
		OptimalKeys:      exercise.Grading.OptimalKeyCount,
	}
	if !ok || !mission.Completed {
		base.Reason = ReasonIncomplete
		base.Priority = 0
		return base, true
	}
	if mission.BestGrade != "" && mission.BestGrade != "S" {
		base.Reason = ReasonLowGrade
		base.Priority = 1
		return base, true
	}
	if exercise.Grading.OptimalKeyCount > 0 && mission.BestKeystrokes > exercise.Grading.OptimalKeyCount {
		base.Reason = ReasonKeyCount
		base.Priority = 2
		return base, true
	}
	return Candidate{}, false
}
