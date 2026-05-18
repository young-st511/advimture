package editor

import (
	"testing"
)

func TestCursorMoveTo(t *testing.T) {
	c := NewCursor()
	c.MoveTo(3, 5)
	if c.Row != 3 || c.Col != 5 || c.DesiredCol != 5 {
		t.Errorf("expected (3,5,5), got (%d,%d,%d)", c.Row, c.Col, c.DesiredCol)
	}
}

func TestCursorMoveToKeepDesired(t *testing.T) {
	c := NewCursor()
	c.MoveTo(0, 10) // DesiredCol = 10
	c.MoveToKeepDesired(1, 5)
	if c.DesiredCol != 10 {
		t.Errorf("expected DesiredCol 10, got %d", c.DesiredCol)
	}
	if c.Col != 5 {
		t.Errorf("expected Col 5, got %d", c.Col)
	}
}

func TestClampToBuffer_NormalMode(t *testing.T) {
	buf := NewBuffer([]string{"hello", "hi", ""})
	c := NewCursor()

	// Col past end of line in Normal mode: clamp to last char
	c.MoveTo(0, 10)
	c.ClampToBuffer(buf, false)
	if c.Col != 4 {
		t.Errorf("expected col 4, got %d", c.Col)
	}

	// Empty line: col should be 0
	c.MoveTo(2, 5)
	c.ClampToBuffer(buf, false)
	if c.Col != 0 {
		t.Errorf("expected col 0 on empty line, got %d", c.Col)
	}

	// Row past buffer end
	c.MoveTo(10, 0)
	c.ClampToBuffer(buf, false)
	if c.Row != 2 {
		t.Errorf("expected row 2, got %d", c.Row)
	}

	// Negative row
	c.MoveTo(-1, 0)
	c.ClampToBuffer(buf, false)
	if c.Row != 0 {
		t.Errorf("expected row 0, got %d", c.Row)
	}
}

func TestClampToBuffer_InsertMode(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	c := NewCursor()

	// Insert mode: can be at len (after last char)
	c.MoveTo(0, 5)
	c.ClampToBuffer(buf, true)
	if c.Col != 5 {
		t.Errorf("expected col 5 in insert mode, got %d", c.Col)
	}

	// But not beyond
	c.MoveTo(0, 10)
	c.ClampToBuffer(buf, true)
	if c.Col != 5 {
		t.Errorf("expected col 5, got %d", c.Col)
	}
}

func TestClampColForVerticalMove(t *testing.T) {
	buf := NewBuffer([]string{
		"long line here",   // len=14
		"hi",               // len=2
		"another long one", // len=16
	})
	c := NewCursor()

	// Start at col 10 on long line
	c.MoveTo(0, 10)

	// Move to short line: col clamped but DesiredCol preserved
	c.Row = 1
	c.ClampColForVerticalMove(buf)
	if c.Col != 1 {
		t.Errorf("expected col 1 on short line, got %d", c.Col)
	}
	if c.DesiredCol != 10 {
		t.Errorf("expected DesiredCol 10, got %d", c.DesiredCol)
	}

	// Move to another long line: col restored to DesiredCol
	c.Row = 2
	c.ClampColForVerticalMove(buf)
	if c.Col != 10 {
		t.Errorf("expected col 10 restored, got %d", c.Col)
	}
}

func TestClampColForVerticalMove_EmptyLine(t *testing.T) {
	buf := NewBuffer([]string{"hello", "", "world"})
	c := NewCursor()

	c.MoveTo(0, 3)
	c.Row = 1
	c.ClampColForVerticalMove(buf)
	if c.Col != 0 {
		t.Errorf("expected col 0 on empty line, got %d", c.Col)
	}

	// Restore when moving to non-empty line
	c.Row = 2
	c.ClampColForVerticalMove(buf)
	if c.Col != 3 {
		t.Errorf("expected col 3, got %d", c.Col)
	}
}
