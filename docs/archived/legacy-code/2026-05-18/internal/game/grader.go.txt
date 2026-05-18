package game

import "github.com/young-st511/advimture/internal/data"

// GradeResult holds the result of a completed mission.
type GradeResult struct {
	Grade       string // "S", "A", "B", "C"
	EffKeys     int    // effective keystrokes
	OptimalKeys int    // mission's optimal keystrokes
	TimeMs      int64  // elapsed time in milliseconds
	Accuracy    float64 // EffKeys/TotalKeys * 100
	Success     bool    // false if text comparison failed
	Diffs       []data.DiffLine
	TotalDiff   int
}

// CalcGrade returns the grade letter based on effective vs optimal keystrokes.
//   S: effKeys <= optimalKeys * 1.0
//   A: effKeys <= optimalKeys * 1.5
//   B: effKeys <= optimalKeys * 2.5
//   C: completed (any key count)
func CalcGrade(effKeys, optimalKeys int) string {
	if optimalKeys <= 0 {
		return "C"
	}
	switch {
	case effKeys <= optimalKeys:
		return "S"
	case float64(effKeys) <= float64(optimalKeys)*1.5:
		return "A"
	case float64(effKeys) <= float64(optimalKeys)*2.5:
		return "B"
	default:
		return "C"
	}
}

// CalcAccuracy returns effective/total * 100. Returns 100 when totalKeys == 0.
func CalcAccuracy(effKeys, totalKeys int) float64 {
	if totalKeys == 0 {
		return 100.0
	}
	v := float64(effKeys) / float64(totalKeys) * 100.0
	if v > 100.0 {
		v = 100.0
	}
	return v
}

// MentorMessage returns Vi 선배's congratulatory message for the given grade.
func MentorMessage(grade string) string {
	switch grade {
	case "S":
		return "완벽해! 진정한 Vim 마스터의 손놀림이야. 이 속도라면 어떤 파일도 무서울 게 없어."
	case "A":
		return "훌륭해! 조금만 더 효율적인 커맨드를 쓰면 S등급도 충분히 가능해."
	case "B":
		return "잘 했어. 아직 개선의 여지가 있어. 더 짧은 키스트로크를 찾아봐."
	default: // C
		return "완료는 했으니 됐어. 하지만 다음에는 더 빠른 방법을 고민해봐."
	}
}

// gradeRank maps grade letter to numeric rank for comparison (higher = better).
func gradeRank(grade string) int {
	switch grade {
	case "S":
		return 4
	case "A":
		return 3
	case "B":
		return 2
	case "C":
		return 1
	default:
		return 0
	}
}

// IsBetterGrade returns true if newGrade is strictly better than oldGrade.
func IsBetterGrade(newGrade, oldGrade string) bool {
	return gradeRank(newGrade) > gradeRank(oldGrade)
}
