package editor

import (
	"slices"
	"strings"
)

// Buffer manages text content as a slice of lines.
type Buffer struct {
	lines []string
}

// NewBuffer creates a buffer with the given text lines.
func NewBuffer(lines []string) *Buffer {
	if len(lines) == 0 {
		lines = []string{""}
	}
	copied := make([]string, len(lines))
	copy(copied, lines)
	return &Buffer{lines: copied}
}

// NewBufferFromText creates a buffer from a single text string, splitting by newlines.
func NewBufferFromText(text string) *Buffer {
	lines := strings.Split(text, "\n")
	// Remove trailing empty line from split if text ends with newline
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return NewBuffer(lines)
}

// LineCount returns the number of lines in the buffer.
func (b *Buffer) LineCount() int {
	return len(b.lines)
}

// GetLine returns the line at the given index. Returns "" if out of bounds.
func (b *Buffer) GetLine(row int) string {
	if row < 0 || row >= len(b.lines) {
		return ""
	}
	return b.lines[row]
}

// GetLineRunes returns the runes of the line at the given index.
func (b *Buffer) GetLineRunes(row int) []rune {
	return []rune(b.GetLine(row))
}

// LineRuneLen returns the rune length of the line at the given index.
func (b *Buffer) LineRuneLen(row int) int {
	return len(b.GetLineRunes(row))
}

// SetLine replaces the line at the given index.
func (b *Buffer) SetLine(row int, text string) {
	if row >= 0 && row < len(b.lines) {
		b.lines[row] = text
	}
}

// InsertChar inserts a character at the given position (rune-based column).
func (b *Buffer) InsertChar(row, col int, ch rune) {
	if row < 0 || row >= len(b.lines) {
		return
	}
	runes := []rune(b.lines[row])
	if col < 0 {
		col = 0
	}
	if col > len(runes) {
		col = len(runes)
	}
	runes = slices.Insert(runes, col, ch)
	b.lines[row] = string(runes)
}

// DeleteChar deletes the character at the given position (rune-based column).
// Returns the deleted rune, or 0 if nothing was deleted.
func (b *Buffer) DeleteChar(row, col int) rune {
	if row < 0 || row >= len(b.lines) {
		return 0
	}
	runes := []rune(b.lines[row])
	if col < 0 || col >= len(runes) {
		return 0
	}
	deleted := runes[col]
	runes = append(runes[:col], runes[col+1:]...)
	b.lines[row] = string(runes)
	return deleted
}

// InsertLine inserts a new line at the given row index, pushing existing lines down.
func (b *Buffer) InsertLine(row int, text string) {
	if row < 0 {
		row = 0
	}
	if row > len(b.lines) {
		row = len(b.lines)
	}
	b.lines = slices.Insert(b.lines, row, text)
}

// DeleteLine removes the line at the given row index and returns it.
// The buffer always keeps at least one empty line.
func (b *Buffer) DeleteLine(row int) string {
	if row < 0 || row >= len(b.lines) {
		return ""
	}
	deleted := b.lines[row]
	b.lines = append(b.lines[:row], b.lines[row+1:]...)
	if len(b.lines) == 0 {
		b.lines = []string{""}
	}
	return deleted
}

// SplitLine splits the line at (row, col) into two lines.
// Characters before col stay on current line, characters from col move to a new line below.
func (b *Buffer) SplitLine(row, col int) {
	if row < 0 || row >= len(b.lines) {
		return
	}
	runes := []rune(b.lines[row])
	if col < 0 {
		col = 0
	}
	if col > len(runes) {
		col = len(runes)
	}
	before := string(runes[:col])
	after := string(runes[col:])
	b.lines[row] = before
	b.InsertLine(row+1, after)
}

// JoinLines joins line at row with the line below it.
func (b *Buffer) JoinLines(row int) {
	if row < 0 || row >= len(b.lines)-1 {
		return
	}
	b.lines[row] = b.lines[row] + b.lines[row+1]
	b.lines = append(b.lines[:row+1], b.lines[row+2:]...)
}

// GetText returns the full buffer content as a single string.
func (b *Buffer) GetText() string {
	return strings.Join(b.lines, "\n")
}

// Lines returns a copy of all lines.
func (b *Buffer) Lines() []string {
	copied := make([]string, len(b.lines))
	copy(copied, b.lines)
	return copied
}

// DeleteRange deletes text from (startRow, startCol) to (endRow, endCol) exclusive.
// Returns the deleted text.
func (b *Buffer) DeleteRange(startRow, startCol, endRow, endCol int) string {
	if startRow < 0 || startRow >= len(b.lines) {
		return ""
	}
	if endRow < 0 || endRow >= len(b.lines) {
		endRow = len(b.lines) - 1
		endCol = b.LineRuneLen(endRow)
	}
	if startCol < 0 {
		startCol = 0
	}
	if endCol < 0 {
		endCol = 0
	}

	if startRow == endRow {
		runes := []rune(b.lines[startRow])
		if startCol > len(runes) {
			startCol = len(runes)
		}
		if endCol > len(runes) {
			endCol = len(runes)
		}
		deleted := string(runes[startCol:endCol])
		b.lines[startRow] = string(runes[:startCol]) + string(runes[endCol:])
		return deleted
	}

	// Multi-line delete
	var deleted strings.Builder

	// First line: keep before startCol
	firstRunes := []rune(b.lines[startRow])
	if startCol > len(firstRunes) {
		startCol = len(firstRunes)
	}
	deleted.WriteString(string(firstRunes[startCol:]))

	// Middle lines: delete entirely
	for i := startRow + 1; i < endRow; i++ {
		deleted.WriteRune('\n')
		deleted.WriteString(b.lines[i])
	}

	// Last line: keep from endCol onwards
	lastRunes := []rune(b.lines[endRow])
	if endCol > len(lastRunes) {
		endCol = len(lastRunes)
	}
	deleted.WriteRune('\n')
	deleted.WriteString(string(lastRunes[:endCol]))

	// Combine first line prefix + last line suffix
	b.lines[startRow] = string(firstRunes[:startCol]) + string(lastRunes[endCol:])

	// Remove lines between startRow+1 and endRow (inclusive)
	b.lines = append(b.lines[:startRow+1], b.lines[endRow+1:]...)

	return deleted.String()
}

// DeleteLines removes lines from startRow to endRow (inclusive) and returns them joined by newline.
func (b *Buffer) DeleteLines(startRow, endRow int) string {
	if startRow < 0 {
		startRow = 0
	}
	if endRow >= len(b.lines) {
		endRow = len(b.lines) - 1
	}
	if startRow > endRow {
		return ""
	}

	deleted := strings.Join(b.lines[startRow:endRow+1], "\n")
	b.lines = append(b.lines[:startRow], b.lines[endRow+1:]...)
	if len(b.lines) == 0 {
		b.lines = []string{""}
	}
	return deleted
}

// InsertText inserts text (possibly multi-line) at the given position.
func (b *Buffer) InsertText(row, col int, text string) {
	if row < 0 || row >= len(b.lines) {
		return
	}
	insertLines := strings.Split(text, "\n")
	runes := []rune(b.lines[row])
	if col > len(runes) {
		col = len(runes)
	}

	if len(insertLines) == 1 {
		b.lines[row] = string(runes[:col]) + insertLines[0] + string(runes[col:])
		return
	}

	suffix := string(runes[col:])
	b.lines[row] = string(runes[:col]) + insertLines[0]

	for i := 1; i < len(insertLines)-1; i++ {
		b.InsertLine(row+i, insertLines[i])
	}

	lastIdx := row + len(insertLines) - 1
	b.InsertLine(lastIdx, insertLines[len(insertLines)-1]+suffix)
}
