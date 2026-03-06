package editor

import "unicode"

// ExecuteMotion applies a motion to the cursor on the given buffer.
// count is the number of times to repeat the motion.
// Returns the new (row, col) position.
func ExecuteMotion(buf *Buffer, cur *Cursor, motion MotionType, count int, ch rune) {
	for i := 0; i < count; i++ {
		switch motion {
		case MotionLeft:
			motionLeft(cur)
		case MotionRight:
			motionRight(buf, cur)
		case MotionDown:
			motionDown(buf, cur)
		case MotionUp:
			motionUp(buf, cur)
		case MotionWordForward:
			motionWordForward(buf, cur)
		case MotionWordBackward:
			motionWordBackward(buf, cur)
		case MotionWordEnd:
			motionWordEnd(buf, cur)
		case MotionLineStart:
			motionLineStart(cur)
		case MotionLineEnd:
			motionLineEnd(buf, cur)
		case MotionFileTop:
			motionFileTop(cur)
		case MotionFileBottom:
			motionFileBottom(buf, cur)
		case MotionFindChar:
			motionFindChar(buf, cur, ch)
		case MotionTillChar:
			// For count > 1, use findChar for intermediate steps, till only on last
			if i < count-1 {
				motionFindChar(buf, cur, ch)
			} else {
				motionTillChar(buf, cur, ch)
			}
		}
	}
}

// MotionRange returns the range (startRow, startCol, endRow, endCol) that a motion covers
// from the current cursor position. endCol is exclusive. Used by operators (d, c, y).
func MotionRange(buf *Buffer, cur *Cursor, motion MotionType, count int, ch rune) (int, int, int, int) {
	startRow, startCol := cur.Row, cur.Col

	// Create a temporary cursor to find the end position
	tmp := &Cursor{Row: cur.Row, Col: cur.Col, DesiredCol: cur.DesiredCol}
	ExecuteMotion(buf, tmp, motion, count, ch)
	endRow, endCol := tmp.Row, tmp.Col

	// Determine if motion went backward
	backward := endRow < startRow || (endRow == startRow && endCol < startCol)

	// Ensure start <= end
	if backward {
		startRow, startCol, endRow, endCol = endRow, endCol, startRow, startCol
	}

	// Inclusive motions: the destination character is part of the range
	// Exclusive motions: the destination character is NOT part of the range
	if isInclusiveMotion(motion) {
		endCol++
	} else if backward {
		// For backward exclusive motions (e.g., db), don't include the original cursor position
		// startCol..endCol already covers the right range
	} else {
		// For forward exclusive motions (e.g., dw), endCol is already the right boundary
	}

	return startRow, startCol, endRow, endCol
}

// isInclusiveMotion returns true for motions where the destination character is included.
func isInclusiveMotion(motion MotionType) bool {
	switch motion {
	case MotionWordEnd, MotionFindChar, MotionTillChar, MotionLineEnd, MotionFileBottom:
		return true
	}
	return false
}

func motionLeft(cur *Cursor) {
	if cur.Col > 0 {
		cur.Col--
	}
	cur.DesiredCol = cur.Col
}

func motionRight(buf *Buffer, cur *Cursor) {
	maxCol := buf.LineRuneLen(cur.Row) - 1
	if maxCol < 0 {
		maxCol = 0
	}
	if cur.Col < maxCol {
		cur.Col++
	}
	cur.DesiredCol = cur.Col
}

func motionDown(buf *Buffer, cur *Cursor) {
	if cur.Row < buf.LineCount()-1 {
		cur.Row++
		cur.ClampColForVerticalMove(buf)
	}
}

func motionUp(buf *Buffer, cur *Cursor) {
	if cur.Row > 0 {
		cur.Row--
		cur.ClampColForVerticalMove(buf)
	}
}

func motionLineStart(cur *Cursor) {
	cur.Col = 0
	cur.DesiredCol = 0
}

func motionLineEnd(buf *Buffer, cur *Cursor) {
	lineLen := buf.LineRuneLen(cur.Row)
	if lineLen > 0 {
		cur.Col = lineLen - 1
	} else {
		cur.Col = 0
	}
	cur.DesiredCol = cur.Col
}

func motionFileTop(cur *Cursor) {
	cur.Row = 0
	cur.Col = 0
	cur.DesiredCol = 0
}

func motionFileBottom(buf *Buffer, cur *Cursor) {
	cur.Row = buf.LineCount() - 1
	lineLen := buf.LineRuneLen(cur.Row)
	if cur.Col >= lineLen && lineLen > 0 {
		cur.Col = lineLen - 1
	}
	if lineLen == 0 {
		cur.Col = 0
	}
	cur.DesiredCol = cur.Col
}

// Word boundary classification
type charClass int

const (
	classSpace charClass = iota
	classWord            // alphanumeric + underscore
	classPunct           // everything else (punctuation)
)

func classifyRune(r rune) charClass {
	if unicode.IsSpace(r) {
		return classSpace
	}
	if r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) {
		return classWord
	}
	return classPunct
}

func motionWordForward(buf *Buffer, cur *Cursor) {
	row, col := cur.Row, cur.Col
	runes := buf.GetLineRunes(row)

	// If line is empty or at end, go to next line
	if len(runes) == 0 {
		if row < buf.LineCount()-1 {
			cur.Row = row + 1
			cur.Col = 0
			cur.DesiredCol = 0
		}
		return
	}

	if col >= len(runes) {
		col = len(runes) - 1
	}
	startClass := classifyRune(runes[col])

	// Skip current word (same class)
	for col < len(runes) && classifyRune(runes[col]) == startClass {
		col++
	}

	// Skip whitespace
	for col < len(runes) && classifyRune(runes[col]) == classSpace {
		col++
	}

	if col >= len(runes) {
		// Move to next line
		if row < buf.LineCount()-1 {
			cur.Row = row + 1
			nextRunes := buf.GetLineRunes(row + 1)
			// Skip leading whitespace on next line
			cur.Col = 0
			for cur.Col < len(nextRunes) && classifyRune(nextRunes[cur.Col]) == classSpace {
				cur.Col++
			}
			if cur.Col >= len(nextRunes) {
				cur.Col = 0
			}
		} else {
			// Last line: go to end
			cur.Col = len(runes) - 1
		}
	} else {
		cur.Col = col
	}
	cur.DesiredCol = cur.Col
}

func motionWordBackward(buf *Buffer, cur *Cursor) {
	row, col := cur.Row, cur.Col

	for {
		runes := buf.GetLineRunes(row)

		if col <= 0 {
			if row > 0 {
				row--
				prevRunes := buf.GetLineRunes(row)
				if len(prevRunes) > 0 {
					col = len(prevRunes) - 1
					// Continue loop to find word start from end of prev line
					continue
				}
				col = 0
			}
			break
		}

		col--

		// Skip whitespace backward
		for col > 0 && classifyRune(runes[col]) == classSpace {
			col--
		}

		if col == 0 && classifyRune(runes[col]) == classSpace {
			if row > 0 {
				row--
				prevRunes := buf.GetLineRunes(row)
				if len(prevRunes) > 0 {
					col = len(prevRunes) - 1
					continue
				}
				col = 0
			} else {
				col = 0
			}
			break
		}

		// Skip same-class characters backward
		cc := classifyRune(runes[col])
		for col > 0 && classifyRune(runes[col-1]) == cc {
			col--
		}
		break
	}

	cur.Row = row
	cur.Col = col
	cur.DesiredCol = cur.Col
}

func motionWordEnd(buf *Buffer, cur *Cursor) {
	row, col := cur.Row, cur.Col

	for {
		runes := buf.GetLineRunes(row)

		if len(runes) == 0 {
			if row < buf.LineCount()-1 {
				row++
				col = 0
				continue
			}
			break
		}

		if col >= len(runes)-1 {
			if row < buf.LineCount()-1 {
				row++
				col = 0
				continue
			}
			break
		}

		col++ // move at least one position

		// Skip whitespace
		for col < len(runes) && classifyRune(runes[col]) == classSpace {
			col++
		}

		if col >= len(runes) {
			if row < buf.LineCount()-1 {
				row++
				col = 0
				continue
			}
			col = len(runes) - 1
			break
		}

		// Skip same-class characters to find end
		cc := classifyRune(runes[col])
		for col < len(runes)-1 && classifyRune(runes[col+1]) == cc {
			col++
		}
		break
	}

	cur.Row = row
	cur.Col = col
	cur.DesiredCol = cur.Col
}

func motionFindChar(buf *Buffer, cur *Cursor, ch rune) {
	runes := buf.GetLineRunes(cur.Row)
	for i := cur.Col + 1; i < len(runes); i++ {
		if runes[i] == ch {
			cur.Col = i
			cur.DesiredCol = cur.Col
			return
		}
	}
	// Not found: cursor stays
}

func motionTillChar(buf *Buffer, cur *Cursor, ch rune) {
	runes := buf.GetLineRunes(cur.Row)
	for i := cur.Col + 1; i < len(runes); i++ {
		if runes[i] == ch {
			cur.Col = i - 1
			cur.DesiredCol = cur.Col
			return
		}
	}
	// Not found: cursor stays
}
