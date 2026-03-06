package editor

import "testing"

func TestUndoRedo_InsertChar(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	um := NewUndoManager()

	buf.InsertChar(0, 5, '!')
	um.Record(Operation{
		Type: OpInsertChar, Row: 0, Col: 5, Char: '!',
		CursorRow: 0, CursorCol: 4,
	})

	if buf.GetLine(0) != "hello!" {
		t.Fatalf("after insert: expected 'hello!', got '%s'", buf.GetLine(0))
	}

	row, col, ok := um.Undo(buf)
	if !ok {
		t.Fatal("undo should succeed")
	}
	if buf.GetLine(0) != "hello" {
		t.Errorf("after undo: expected 'hello', got '%s'", buf.GetLine(0))
	}
	if row != 0 || col != 4 {
		t.Errorf("undo cursor: expected (0,4), got (%d,%d)", row, col)
	}

	_, _, ok = um.Redo(buf)
	if !ok {
		t.Fatal("redo should succeed")
	}
	if buf.GetLine(0) != "hello!" {
		t.Errorf("after redo: expected 'hello!', got '%s'", buf.GetLine(0))
	}
}

func TestUndoRedo_DeleteChar(t *testing.T) {
	buf := NewBuffer([]string{"hello"})
	um := NewUndoManager()

	deleted := buf.DeleteChar(0, 4)
	um.Record(Operation{
		Type: OpDeleteChar, Row: 0, Col: 4, Char: deleted,
		CursorRow: 0, CursorCol: 4,
	})

	if buf.GetLine(0) != "hell" {
		t.Fatalf("after delete: expected 'hell', got '%s'", buf.GetLine(0))
	}

	um.Undo(buf)
	if buf.GetLine(0) != "hello" {
		t.Errorf("after undo: expected 'hello', got '%s'", buf.GetLine(0))
	}
}

func TestUndoRedo_DeleteLine(t *testing.T) {
	buf := NewBuffer([]string{"aaa", "bbb", "ccc"})
	um := NewUndoManager()

	deleted := buf.DeleteLine(1)
	um.Record(Operation{
		Type: OpDeleteLine, Row: 1, Text: deleted,
		CursorRow: 0, CursorCol: 0,
	})

	if buf.LineCount() != 2 {
		t.Fatalf("after delete: expected 2 lines, got %d", buf.LineCount())
	}

	um.Undo(buf)
	if buf.LineCount() != 3 || buf.GetLine(1) != "bbb" {
		t.Errorf("after undo: expected 3 lines with 'bbb', got %d lines, line1='%s'",
			buf.LineCount(), buf.GetLine(1))
	}
}

func TestUndoRedo_SetLine(t *testing.T) {
	buf := NewBuffer([]string{"old text"})
	um := NewUndoManager()

	old := buf.GetLine(0)
	buf.SetLine(0, "new text")
	um.Record(Operation{
		Type: OpSetLine, Row: 0, Text: "new text", OldText: old,
		CursorRow: 0, CursorCol: 0,
	})

	um.Undo(buf)
	if buf.GetLine(0) != "old text" {
		t.Errorf("after undo: expected 'old text', got '%s'", buf.GetLine(0))
	}

	um.Redo(buf)
	if buf.GetLine(0) != "new text" {
		t.Errorf("after redo: expected 'new text', got '%s'", buf.GetLine(0))
	}
}

func TestUndoRedo_SplitJoin(t *testing.T) {
	buf := NewBuffer([]string{"helloworld"})
	um := NewUndoManager()

	buf.SplitLine(0, 5)
	um.Record(Operation{
		Type: OpSplitLine, Row: 0, Col: 5,
		CursorRow: 0, CursorCol: 4,
	})

	if buf.LineCount() != 2 {
		t.Fatalf("after split: expected 2 lines, got %d", buf.LineCount())
	}

	um.Undo(buf)
	if buf.LineCount() != 1 || buf.GetLine(0) != "helloworld" {
		t.Errorf("after undo: expected 'helloworld', got '%s'", buf.GetLine(0))
	}

	um.Redo(buf)
	if buf.LineCount() != 2 {
		t.Errorf("after redo: expected 2 lines, got %d", buf.LineCount())
	}
}

func TestUndoRedo_DeleteRange(t *testing.T) {
	buf := NewBuffer([]string{"hello world"})
	um := NewUndoManager()

	deleted := buf.DeleteRange(0, 5, 0, 11)
	um.Record(Operation{
		Type: OpDeleteRange, Row: 0, Col: 5, EndRow: 0, EndCol: 11,
		Text: deleted, CursorRow: 0, CursorCol: 5,
	})

	if buf.GetLine(0) != "hello" {
		t.Fatalf("after delete range: expected 'hello', got '%s'", buf.GetLine(0))
	}

	um.Undo(buf)
	if buf.GetLine(0) != "hello world" {
		t.Errorf("after undo: expected 'hello world', got '%s'", buf.GetLine(0))
	}
}

func TestUndoRedo_DeleteLines(t *testing.T) {
	buf := NewBuffer([]string{"aaa", "bbb", "ccc", "ddd"})
	um := NewUndoManager()

	deleted := buf.DeleteLines(1, 2)
	um.Record(Operation{
		Type: OpDeleteLines, Row: 1, EndRow: 2,
		Text: deleted, CursorRow: 0, CursorCol: 0,
	})

	if buf.LineCount() != 2 {
		t.Fatalf("after delete lines: expected 2, got %d", buf.LineCount())
	}

	um.Undo(buf)
	if buf.LineCount() != 4 || buf.GetLine(1) != "bbb" || buf.GetLine(2) != "ccc" {
		t.Errorf("after undo: expected 4 lines, got %d", buf.LineCount())
	}
}

func TestUndoRedo_Composite(t *testing.T) {
	buf := NewBuffer([]string{"hello world"})
	um := NewUndoManager()

	// Simulate "cw" (delete word + enter insert): delete "hello" then set line
	old := buf.GetLine(0)
	buf.SetLine(0, " world")
	child1 := Operation{Type: OpSetLine, Row: 0, Text: " world", OldText: old}

	um.Record(Operation{
		Type: OpComposite, Children: []Operation{child1},
		CursorRow: 0, CursorCol: 0,
	})

	um.Undo(buf)
	if buf.GetLine(0) != "hello world" {
		t.Errorf("after undo composite: expected 'hello world', got '%s'", buf.GetLine(0))
	}
}

func TestUndo_ClearsRedoStack(t *testing.T) {
	buf := NewBuffer([]string{"abc"})
	um := NewUndoManager()

	buf.DeleteChar(0, 2)
	um.Record(Operation{Type: OpDeleteChar, Row: 0, Col: 2, Char: 'c', CursorRow: 0, CursorCol: 2})

	um.Undo(buf)
	if !um.CanRedo() {
		t.Fatal("should be able to redo")
	}

	// New edit clears redo stack
	buf.InsertChar(0, 0, 'X')
	um.Record(Operation{Type: OpInsertChar, Row: 0, Col: 0, Char: 'X', CursorRow: 0, CursorCol: 0})

	if um.CanRedo() {
		t.Error("redo stack should be cleared after new edit")
	}
}

func TestUndo_EmptyStack(t *testing.T) {
	buf := NewBuffer([]string{"abc"})
	um := NewUndoManager()

	_, _, ok := um.Undo(buf)
	if ok {
		t.Error("undo on empty stack should return false")
	}

	_, _, ok = um.Redo(buf)
	if ok {
		t.Error("redo on empty stack should return false")
	}
}
