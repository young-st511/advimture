package editor

// MatchPos represents a single search match position in the buffer.
type MatchPos struct {
	Row, Col, Len int
}

// SearchState holds the current search state.
type SearchState struct {
	Pattern    string
	Matches    []MatchPos
	CurrentIdx int
}

// FindMatches searches all lines for the literal pattern and returns all match positions.
// Returns nil if pattern is empty.
func FindMatches(lines []string, pattern string) []MatchPos {
	if pattern == "" {
		return nil
	}
	patternRunes := []rune(pattern)
	pLen := len(patternRunes)

	var matches []MatchPos
	for row, line := range lines {
		lineRunes := []rune(line)
		lLen := len(lineRunes)
		if lLen < pLen {
			continue
		}
		for col := 0; col <= lLen-pLen; col++ {
			matched := true
			for i := 0; i < pLen; i++ {
				if lineRunes[col+i] != patternRunes[i] {
					matched = false
					break
				}
			}
			if matched {
				matches = append(matches, MatchPos{Row: row, Col: col, Len: pLen})
			}
		}
	}
	return matches
}
