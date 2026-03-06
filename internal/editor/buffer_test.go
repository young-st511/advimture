package editor

import (
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := NewBuffer([]string{"hello", "world"})
	if b.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", b.LineCount())
	}
	if b.GetLine(0) != "hello" {
		t.Errorf("expected 'hello', got '%s'", b.GetLine(0))
	}
}

func TestNewBufferEmpty(t *testing.T) {
	b := NewBuffer(nil)
	if b.LineCount() != 1 {
		t.Errorf("expected 1 line for empty buffer, got %d", b.LineCount())
	}
	if b.GetLine(0) != "" {
		t.Errorf("expected empty string, got '%s'", b.GetLine(0))
	}
}

func TestNewBufferFromText(t *testing.T) {
	b := NewBufferFromText("line1\nline2\nline3")
	if b.LineCount() != 3 {
		t.Errorf("expected 3 lines, got %d", b.LineCount())
	}
	if b.GetLine(2) != "line3" {
		t.Errorf("expected 'line3', got '%s'", b.GetLine(2))
	}
}

func TestInsertChar(t *testing.T) {
	b := NewBuffer([]string{"hllo"})
	b.InsertChar(0, 1, 'e')
	if b.GetLine(0) != "hello" {
		t.Errorf("expected 'hello', got '%s'", b.GetLine(0))
	}
}

func TestInsertCharUnicode(t *testing.T) {
	b := NewBuffer([]string{"안하세요"})
	b.InsertChar(0, 1, '녕')
	if b.GetLine(0) != "안녕하세요" {
		t.Errorf("expected '안녕하세요', got '%s'", b.GetLine(0))
	}
}

func TestDeleteChar(t *testing.T) {
	b := NewBuffer([]string{"helllo"})
	deleted := b.DeleteChar(0, 3)
	if deleted != 'l' {
		t.Errorf("expected deleted 'l', got '%c'", deleted)
	}
	if b.GetLine(0) != "hello" {
		t.Errorf("expected 'hello', got '%s'", b.GetLine(0))
	}
}

func TestDeleteCharEmptyLine(t *testing.T) {
	b := NewBuffer([]string{""})
	deleted := b.DeleteChar(0, 0)
	if deleted != 0 {
		t.Errorf("expected 0, got '%c'", deleted)
	}
}

func TestInsertLine(t *testing.T) {
	b := NewBuffer([]string{"first", "third"})
	b.InsertLine(1, "second")
	if b.LineCount() != 3 {
		t.Errorf("expected 3 lines, got %d", b.LineCount())
	}
	if b.GetLine(1) != "second" {
		t.Errorf("expected 'second', got '%s'", b.GetLine(1))
	}
}

func TestDeleteLine(t *testing.T) {
	b := NewBuffer([]string{"first", "second", "third"})
	deleted := b.DeleteLine(1)
	if deleted != "second" {
		t.Errorf("expected 'second', got '%s'", deleted)
	}
	if b.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", b.LineCount())
	}
}

func TestDeleteLineLastLine(t *testing.T) {
	b := NewBuffer([]string{"only"})
	b.DeleteLine(0)
	if b.LineCount() != 1 {
		t.Errorf("buffer should keep at least 1 line, got %d", b.LineCount())
	}
	if b.GetLine(0) != "" {
		t.Errorf("expected empty line, got '%s'", b.GetLine(0))
	}
}

func TestSplitLine(t *testing.T) {
	b := NewBuffer([]string{"helloworld"})
	b.SplitLine(0, 5)
	if b.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", b.LineCount())
	}
	if b.GetLine(0) != "hello" {
		t.Errorf("expected 'hello', got '%s'", b.GetLine(0))
	}
	if b.GetLine(1) != "world" {
		t.Errorf("expected 'world', got '%s'", b.GetLine(1))
	}
}

func TestJoinLines(t *testing.T) {
	b := NewBuffer([]string{"hello", "world"})
	b.JoinLines(0)
	if b.LineCount() != 1 {
		t.Errorf("expected 1 line, got %d", b.LineCount())
	}
	if b.GetLine(0) != "helloworld" {
		t.Errorf("expected 'helloworld', got '%s'", b.GetLine(0))
	}
}

func TestDeleteRange_SameLine(t *testing.T) {
	b := NewBuffer([]string{"hello world"})
	deleted := b.DeleteRange(0, 5, 0, 11)
	if deleted != " world" {
		t.Errorf("expected ' world', got '%s'", deleted)
	}
	if b.GetLine(0) != "hello" {
		t.Errorf("expected 'hello', got '%s'", b.GetLine(0))
	}
}

func TestDeleteRange_MultiLine(t *testing.T) {
	b := NewBuffer([]string{"aaa", "bbb", "ccc", "ddd"})
	deleted := b.DeleteRange(0, 2, 2, 1)
	if deleted != "a\nbbb\nc" {
		t.Errorf("expected 'a\\nbbb\\nc', got '%s'", deleted)
	}
	if b.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", b.LineCount())
	}
	if b.GetLine(0) != "aacc" {
		t.Errorf("expected 'aacc', got '%s'", b.GetLine(0))
	}
}

func TestDeleteLines(t *testing.T) {
	b := NewBuffer([]string{"aaa", "bbb", "ccc", "ddd"})
	deleted := b.DeleteLines(1, 2)
	if deleted != "bbb\nccc" {
		t.Errorf("expected 'bbb\\nccc', got '%s'", deleted)
	}
	if b.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", b.LineCount())
	}
}

func TestDeleteLines_All(t *testing.T) {
	b := NewBuffer([]string{"only"})
	b.DeleteLines(0, 0)
	if b.LineCount() != 1 {
		t.Errorf("should keep at least 1 line, got %d", b.LineCount())
	}
	if b.GetLine(0) != "" {
		t.Errorf("expected empty, got '%s'", b.GetLine(0))
	}
}

func TestGetText(t *testing.T) {
	b := NewBuffer([]string{"hello", "world"})
	if b.GetText() != "hello\nworld" {
		t.Errorf("expected 'hello\\nworld', got '%s'", b.GetText())
	}
}

func TestLineRuneLen(t *testing.T) {
	b := NewBuffer([]string{"안녕하세요"})
	if b.LineRuneLen(0) != 5 {
		t.Errorf("expected 5 runes, got %d", b.LineRuneLen(0))
	}
}

func TestSetLine(t *testing.T) {
	b := NewBuffer([]string{"old"})
	b.SetLine(0, "new")
	if b.GetLine(0) != "new" {
		t.Errorf("expected 'new', got '%s'", b.GetLine(0))
	}
	// Out of bounds: no panic
	b.SetLine(5, "noop")
	b.SetLine(-1, "noop")
}

func TestLines_ReturnsCopy(t *testing.T) {
	b := NewBuffer([]string{"a", "b"})
	lines := b.Lines()
	lines[0] = "modified"
	if b.GetLine(0) != "a" {
		t.Errorf("modifying Lines() copy should not affect buffer")
	}
}

func TestInsertChar_OutOfBounds(t *testing.T) {
	b := NewBuffer([]string{"abc"})
	b.InsertChar(-1, 0, 'x') // no panic
	b.InsertChar(5, 0, 'x')  // no panic
	b.InsertChar(0, -1, 'x') // clamped to 0
	if b.GetLine(0) != "xabc" {
		t.Errorf("expected 'xabc', got '%s'", b.GetLine(0))
	}
}

func TestDeleteChar_OutOfBounds(t *testing.T) {
	b := NewBuffer([]string{"abc"})
	r := b.DeleteChar(0, -1)
	if r != 0 {
		t.Errorf("expected 0 for negative col, got '%c'", r)
	}
	r = b.DeleteChar(0, 10)
	if r != 0 {
		t.Errorf("expected 0 for out of bounds col, got '%c'", r)
	}
}

func TestJoinLines_LastLine(t *testing.T) {
	b := NewBuffer([]string{"only"})
	b.JoinLines(0) // last line, no line below — no-op
	if b.LineCount() != 1 || b.GetLine(0) != "only" {
		t.Errorf("JoinLines on last line should be no-op")
	}
}

func TestDeleteRange_NegativeCol(t *testing.T) {
	b := NewBuffer([]string{"hello"})
	deleted := b.DeleteRange(0, -1, 0, 3)
	if deleted != "hel" {
		t.Errorf("expected 'hel', got '%s'", deleted)
	}
}

func TestDeleteChar_SingleCharLine(t *testing.T) {
	b := NewBuffer([]string{"x"})
	deleted := b.DeleteChar(0, 0)
	if deleted != 'x' {
		t.Errorf("expected 'x', got '%c'", deleted)
	}
	if b.GetLine(0) != "" {
		t.Errorf("expected empty line, got '%s'", b.GetLine(0))
	}
}

func TestInsertText_SingleLine(t *testing.T) {
	b := NewBuffer([]string{"hd"})
	b.InsertText(0, 1, "ello worl")
	if b.GetLine(0) != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", b.GetLine(0))
	}
}

func TestInsertText_MultiLine(t *testing.T) {
	b := NewBuffer([]string{"start end"})
	b.InsertText(0, 6, "middle\nnew line\n")
	if b.LineCount() != 3 {
		t.Errorf("expected 3 lines, got %d", b.LineCount())
	}
	if b.GetLine(0) != "start middle" {
		t.Errorf("line 0: expected 'start middle', got '%s'", b.GetLine(0))
	}
	if b.GetLine(1) != "new line" {
		t.Errorf("line 1: expected 'new line', got '%s'", b.GetLine(1))
	}
	if b.GetLine(2) != "end" {
		t.Errorf("line 2: expected 'end', got '%s'", b.GetLine(2))
	}
}
