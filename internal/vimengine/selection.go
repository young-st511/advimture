package vimengine

type SelectionKind string

const (
	SelectionCharwise SelectionKind = "charwise"
	SelectionLinewise SelectionKind = "linewise"
)

type Selection struct {
	Active bool
	Kind   SelectionKind
	Anchor Cursor
	Head   Cursor
	Start  Cursor
	End    Cursor
}

func normalizedSelection(selection Selection) *Selection {
	if selection.Kind == "" {
		selection.Kind = SelectionCharwise
	}
	selection.Active = true
	selection.Start, selection.End = orderedCursors(selection.Anchor, selection.Head)
	return &selection
}

func normalizedSelectionForLines(selection Selection, lines []string) *Selection {
	if selection.Kind == SelectionLinewise {
		return normalizedLinewiseSelection(selection, lines)
	}
	return normalizedSelection(selection)
}

func normalizedLinewiseSelection(selection Selection, lines []string) *Selection {
	selection.Active = true
	selection.Kind = SelectionLinewise
	startRow := selection.Anchor.Row
	endRow := selection.Head.Row
	if startRow > endRow {
		startRow, endRow = endRow, startRow
	}
	if startRow < 0 {
		startRow = 0
	}
	if endRow < 0 {
		endRow = 0
	}
	if startRow >= len(lines) {
		startRow = len(lines) - 1
	}
	if endRow >= len(lines) {
		endRow = len(lines) - 1
	}
	selection.Start = Cursor{Row: startRow, Col: 0, DesiredCol: 0}
	selection.End = Cursor{Row: endRow, Col: lineEndCol(lines[endRow]), DesiredCol: lineEndCol(lines[endRow])}
	return &selection
}

func lineEndCol(line string) int {
	if lineRuneLen(line) == 0 {
		return 0
	}
	return lineRuneLen(line) - 1
}

func orderedCursors(left Cursor, right Cursor) (Cursor, Cursor) {
	if left.Row < right.Row || (left.Row == right.Row && left.Col <= right.Col) {
		return left, right
	}
	return right, left
}

func clampSelectionCursor(cursor Cursor, lines []string) Cursor {
	if cursor.Row < 0 {
		cursor.Row = 0
	}
	if cursor.Row >= len(lines) {
		cursor.Row = len(lines) - 1
	}
	cursor.Col = clampCol(cursor.Col, lines[cursor.Row])
	cursor.DesiredCol = cursor.Col
	return cursor
}
