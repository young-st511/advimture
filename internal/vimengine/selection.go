package vimengine

type SelectionKind string

const (
	SelectionCharwise SelectionKind = "charwise"
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
