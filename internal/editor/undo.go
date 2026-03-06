package editor

// OpType represents the type of an undoable operation.
type OpType int

const (
	OpInsertChar OpType = iota
	OpDeleteChar
	OpInsertLine
	OpDeleteLine
	OpSetLine
	OpSplitLine
	OpJoinLines
	OpDeleteRange
	OpInsertText
	OpDeleteLines
	OpComposite // groups multiple operations as one undo unit
)

// Operation records a single undoable action.
type Operation struct {
	Type OpType

	// Position
	Row int
	Col int

	// For char operations
	Char rune

	// For line operations
	Text string

	// For range operations
	EndRow int
	EndCol int

	// For SetLine
	OldText string

	// For composite operations
	Children []Operation

	// Cursor state before the operation (for undo restore)
	CursorRow int
	CursorCol int
}

// UndoManager manages undo/redo stacks using operation logs.
type UndoManager struct {
	undoStack []Operation
	redoStack []Operation
}

// NewUndoManager creates a new undo manager.
func NewUndoManager() *UndoManager {
	return &UndoManager{}
}

// Record adds an operation to the undo stack and clears the redo stack.
func (u *UndoManager) Record(op Operation) {
	u.undoStack = append(u.undoStack, op)
	u.redoStack = nil
}

// CanUndo returns true if there are operations to undo.
func (u *UndoManager) CanUndo() bool {
	return len(u.undoStack) > 0
}

// CanRedo returns true if there are operations to redo.
func (u *UndoManager) CanRedo() bool {
	return len(u.redoStack) > 0
}

// Undo reverses the last operation and returns cursor position to restore.
func (u *UndoManager) Undo(buf *Buffer) (row, col int, ok bool) {
	if !u.CanUndo() {
		return 0, 0, false
	}
	op := u.undoStack[len(u.undoStack)-1]
	u.undoStack = u.undoStack[:len(u.undoStack)-1]

	undoOp(buf, &op)
	u.redoStack = append(u.redoStack, op)
	return op.CursorRow, op.CursorCol, true
}

// Redo re-applies the last undone operation and returns cursor position.
func (u *UndoManager) Redo(buf *Buffer) (row, col int, ok bool) {
	if !u.CanRedo() {
		return 0, 0, false
	}
	op := u.redoStack[len(u.redoStack)-1]
	u.redoStack = u.redoStack[:len(u.redoStack)-1]

	redoOp(buf, &op)
	u.undoStack = append(u.undoStack, op)
	return op.Row, op.Col, true
}

func undoOp(buf *Buffer, op *Operation) {
	switch op.Type {
	case OpInsertChar:
		buf.DeleteChar(op.Row, op.Col)
	case OpDeleteChar:
		buf.InsertChar(op.Row, op.Col, op.Char)
	case OpInsertLine:
		buf.DeleteLine(op.Row)
	case OpDeleteLine:
		buf.InsertLine(op.Row, op.Text)
	case OpSetLine:
		buf.SetLine(op.Row, op.OldText)
	case OpSplitLine:
		buf.JoinLines(op.Row)
	case OpJoinLines:
		buf.SplitLine(op.Row, op.Col)
	case OpDeleteRange:
		buf.InsertText(op.Row, op.Col, op.Text)
	case OpInsertText:
		// Calculate how many lines the inserted text spans
		buf.DeleteRange(op.Row, op.Col, op.EndRow, op.EndCol)
	case OpDeleteLines:
		// Re-insert deleted lines
		lines := splitLines(op.Text)
		for i, line := range lines {
			buf.InsertLine(op.Row+i, line)
		}
	case OpComposite:
		// Undo children in reverse order
		for i := len(op.Children) - 1; i >= 0; i-- {
			undoOp(buf, &op.Children[i])
		}
	}
}

func redoOp(buf *Buffer, op *Operation) {
	switch op.Type {
	case OpInsertChar:
		buf.InsertChar(op.Row, op.Col, op.Char)
	case OpDeleteChar:
		buf.DeleteChar(op.Row, op.Col)
	case OpInsertLine:
		buf.InsertLine(op.Row, op.Text)
	case OpDeleteLine:
		buf.DeleteLine(op.Row)
	case OpSetLine:
		buf.SetLine(op.Row, op.Text)
	case OpSplitLine:
		buf.SplitLine(op.Row, op.Col)
	case OpJoinLines:
		buf.JoinLines(op.Row)
	case OpDeleteRange:
		buf.DeleteRange(op.Row, op.Col, op.EndRow, op.EndCol)
	case OpInsertText:
		buf.InsertText(op.Row, op.Col, op.Text)
	case OpDeleteLines:
		buf.DeleteLines(op.Row, op.EndRow)
	case OpComposite:
		for i := range op.Children {
			redoOp(buf, &op.Children[i])
		}
	}
}

func splitLines(text string) []string {
	if text == "" {
		return []string{""}
	}
	var lines []string
	start := 0
	for i, ch := range text {
		if ch == '\n' {
			lines = append(lines, text[start:i])
			start = i + 1
		}
	}
	lines = append(lines, text[start:])
	return lines
}
