package editor

import "testing"

func TestCommand_Quit(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	for _, cmd := range []string{"q", "q!"} {
		r := ExecuteCommand(cmd, buf, cur, undo)
		if !r.Quit {
			t.Errorf(":%s should quit", cmd)
		}
	}
}

func TestCommand_Save(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("w", buf, cur, undo)
	if !r.Save || r.Quit {
		t.Error(":w should save but not quit")
	}
}

func TestCommand_SaveQuit(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	for _, cmd := range []string{"wq", "wq!", "x"} {
		r := ExecuteCommand(cmd, buf, cur, undo)
		if !r.Save || !r.Quit {
			t.Errorf(":%s should save and quit", cmd)
		}
	}
}

func TestCommand_GotoLine(t *testing.T) {
	buf := NewBuffer([]string{"a", "b", "c", "d"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("3", buf, cur, undo)
	if r.GotoLine != 2 {
		t.Errorf(":3 should go to line 2 (0-based), got %d", r.GotoLine)
	}
}

func TestCommand_RangeDelete(t *testing.T) {
	buf := NewBuffer([]string{"a", "b", "c", "d"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("2,3d", buf, cur, undo)
	if r.Error != "" {
		t.Fatalf("expected no error, got '%s'", r.Error)
	}
	if buf.LineCount() != 2 {
		t.Errorf("expected 2 lines, got %d", buf.LineCount())
	}
	if buf.GetLine(0) != "a" || buf.GetLine(1) != "d" {
		t.Errorf("expected [a, d], got [%s, %s]", buf.GetLine(0), buf.GetLine(1))
	}

	// Undo
	undo.Undo(buf)
	if buf.LineCount() != 4 {
		t.Errorf("after undo: expected 4 lines, got %d", buf.LineCount())
	}
}

func TestCommand_Substitute(t *testing.T) {
	buf := NewBuffer([]string{"hello world hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	// Without g flag: replace first occurrence only
	r := ExecuteCommand("s/hello/hi/", buf, cur, undo)
	if r.Error != "" {
		t.Fatalf("unexpected error: %s", r.Error)
	}
	if buf.GetLine(0) != "hi world hello" {
		t.Errorf("expected 'hi world hello', got '%s'", buf.GetLine(0))
	}
}

func TestCommand_SubstituteGlobal(t *testing.T) {
	buf := NewBuffer([]string{"hello world hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("s/hello/hi/g", buf, cur, undo)
	if r.Error != "" {
		t.Fatalf("unexpected error: %s", r.Error)
	}
	if buf.GetLine(0) != "hi world hi" {
		t.Errorf("expected 'hi world hi', got '%s'", buf.GetLine(0))
	}

	// Undo
	undo.Undo(buf)
	if buf.GetLine(0) != "hello world hello" {
		t.Errorf("after undo: expected 'hello world hello', got '%s'", buf.GetLine(0))
	}
}

func TestCommand_GlobalSubstitute(t *testing.T) {
	buf := NewBuffer([]string{"aa bb", "aa cc", "dd ee"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("%s/aa/XX/g", buf, cur, undo)
	if r.Error != "" {
		t.Fatalf("unexpected error: %s", r.Error)
	}
	if buf.GetLine(0) != "XX bb" || buf.GetLine(1) != "XX cc" {
		t.Errorf("expected XX replacements, got ['%s', '%s']", buf.GetLine(0), buf.GetLine(1))
	}
	if buf.GetLine(2) != "dd ee" {
		t.Errorf("line 2 should be unchanged, got '%s'", buf.GetLine(2))
	}

	// Undo restores all
	undo.Undo(buf)
	if buf.GetLine(0) != "aa bb" || buf.GetLine(1) != "aa cc" {
		t.Errorf("after undo: expected originals, got ['%s', '%s']", buf.GetLine(0), buf.GetLine(1))
	}
}

func TestCommand_SubstituteNotFound(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("s/xyz/abc/g", buf, cur, undo)
	if r.Error == "" {
		t.Error("expected error for pattern not found")
	}
}

func TestCommand_Unknown(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("foobar", buf, cur, undo)
	if r.Error == "" {
		t.Error("expected error for unknown command")
	}
}

func TestCommand_Empty(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("", buf, cur, undo)
	if r.Error != "" || r.Quit || r.Save {
		t.Error("empty command should be no-op")
	}
}

func TestCommand_RangeDelete_InvalidRange(t *testing.T) {
	buf := NewBuffer([]string{"a", "b"})
	cur := NewCursor()
	undo := NewUndoManager()

	r := ExecuteCommand("3,1d", buf, cur, undo)
	if r.Error == "" {
		t.Error("expected error for invalid range")
	}
}
