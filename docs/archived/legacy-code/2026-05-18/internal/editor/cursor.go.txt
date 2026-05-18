package editor

// Cursor manages the cursor position within a buffer.
type Cursor struct {
	Row        int
	Col        int
	DesiredCol int // remembered column for vertical movement across lines of different lengths
}

// NewCursor creates a cursor at position (0, 0).
func NewCursor() *Cursor {
	return &Cursor{Row: 0, Col: 0, DesiredCol: 0}
}

// MoveTo sets the cursor to an absolute position and updates DesiredCol.
func (c *Cursor) MoveTo(row, col int) {
	c.Row = row
	c.Col = col
	c.DesiredCol = col
}

// MoveToKeepDesired sets the cursor position without updating DesiredCol.
// Used for vertical movement (j/k) where we want to remember the original column.
func (c *Cursor) MoveToKeepDesired(row, col int) {
	c.Row = row
	c.Col = col
}

// ClampToBuffer ensures cursor position is within the buffer bounds.
// In Normal mode, the cursor can't go past the last character (len-1).
// In Insert mode, the cursor can be at len (after last character).
func (c *Cursor) ClampToBuffer(buf *Buffer, insertMode bool) {
	// Clamp row
	if c.Row < 0 {
		c.Row = 0
	}
	if c.Row >= buf.LineCount() {
		c.Row = buf.LineCount() - 1
	}

	// Clamp col
	lineLen := buf.LineRuneLen(c.Row)
	maxCol := lineLen - 1
	if insertMode {
		maxCol = lineLen
	}
	if maxCol < 0 {
		maxCol = 0
	}
	if c.Col < 0 {
		c.Col = 0
	}
	if c.Col > maxCol {
		c.Col = maxCol
	}
}

// ClampColForVerticalMove adjusts the column when moving vertically between lines.
// Uses DesiredCol to restore the original column when possible.
func (c *Cursor) ClampColForVerticalMove(buf *Buffer) {
	lineLen := buf.LineRuneLen(c.Row)
	maxCol := lineLen - 1
	if maxCol < 0 {
		maxCol = 0
	}
	if c.DesiredCol <= maxCol {
		c.Col = c.DesiredCol
	} else {
		c.Col = maxCol
	}
}
