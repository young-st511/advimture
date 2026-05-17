package scoring

import exerciseruntime "github.com/young-st511/advimture/internal/runtime"

type Grade string

const (
	GradeS Grade = "S"
	GradeA Grade = "A"
	GradeB Grade = "B"
	GradeC Grade = "C"
	GradeF Grade = "F"
)

type Input struct {
	Status       exerciseruntime.Status
	KeyTrace     []string
	ExpectedKeys []string
	Attempts     int
	HintsUsed    int
}

type Result struct {
	Passed           bool
	ExactKeys        bool
	Efficiency       float64
	Penalty          float64
	Score            float64
	Grade            Grade
	KeyCount         int
	ExpectedKeyCount int
}

func Evaluate(input Input) Result {
	keyTrace := copyStrings(input.KeyTrace)
	expectedKeys := copyStrings(input.ExpectedKeys)
	passed := input.Status == exerciseruntime.StatusSucceeded
	exactKeys := keysMatch(keyTrace, expectedKeys)
	efficiency := calculateEfficiency(len(expectedKeys), len(keyTrace))
	penalty := calculatePenalty(input.Attempts, input.HintsUsed)

	if !passed {
		return Result{
			Passed:           false,
			ExactKeys:        exactKeys,
			Efficiency:       efficiency,
			Penalty:          penalty,
			Score:            0,
			Grade:            GradeF,
			KeyCount:         len(keyTrace),
			ExpectedKeyCount: len(expectedKeys),
		}
	}

	score := efficiency - penalty
	if score < 0 {
		score = 0
	}
	if !exactKeys && score > 0.89 {
		score = 0.89
	}

	return Result{
		Passed:           true,
		ExactKeys:        exactKeys,
		Efficiency:       efficiency,
		Penalty:          penalty,
		Score:            score,
		Grade:            gradeForScore(score),
		KeyCount:         len(keyTrace),
		ExpectedKeyCount: len(expectedKeys),
	}
}

func calculateEfficiency(expectedCount int, actualCount int) float64 {
	if expectedCount == 0 {
		return 1
	}
	if actualCount <= 0 {
		return 0
	}
	efficiency := float64(expectedCount) / float64(actualCount)
	if efficiency > 1 {
		return 1
	}
	return efficiency
}

func calculatePenalty(attempts int, hintsUsed int) float64 {
	var penalty float64
	if attempts > 1 {
		penalty += float64(attempts-1) * 0.15
	}
	if hintsUsed > 0 {
		penalty += float64(hintsUsed) * 0.10
	}
	if penalty > 1 {
		return 1
	}
	return penalty
}

func gradeForScore(score float64) Grade {
	switch {
	case score >= 0.95:
		return GradeS
	case score >= 0.85:
		return GradeA
	case score >= 0.65:
		return GradeB
	case score >= 0.45:
		return GradeC
	default:
		return GradeF
	}
}

func keysMatch(left []string, right []string) bool {
	if len(right) == 0 {
		return true
	}
	if len(left) != len(right) {
		return false
	}
	for index := range right {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func copyStrings(values []string) []string {
	if values == nil {
		return nil
	}
	next := make([]string, len(values))
	copy(next, values)
	return next
}
