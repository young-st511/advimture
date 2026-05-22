package vimengine

func enterVisualMode(state State, key string) Result {
	kind := SelectionCharwise
	if key == KeyShiftV {
		kind = SelectionLinewise
	}
	next := copyState(state)
	next.Mode = ModeVisual
	next.PendingKey = ""
	next.Selection = normalizedSelectionForLines(Selection{
		Active: true,
		Kind:   kind,
		Anchor: next.Cursor,
		Head:   next.Cursor,
	}, next.Lines)
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventMoved,
			Key:  key,
		}},
	}
}

func applyVisualKey(state State, key string) Result {
	if state.PendingKey == KeyG {
		if key == KeyG {
			next := copyState(state)
			next.PendingKey = ""
			return applyVisualMotion(next, key, func(next State) Result {
				return moveDocumentStart(next, key)
			})
		}
		next := copyState(state)
		next.PendingKey = ""
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "key is not supported after g in visual mode",
			}},
		}
	}
	switch key {
	case KeyV, KeyShiftV:
		next := copyState(state)
		next.Mode = ModeNormal
		next.PendingKey = ""
		next.Selection = nil
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventModeReset,
				Key:  key,
			}},
		}
	case KeyD:
		return deleteVisualSelection(state, key)
	case KeyY:
		return yankVisualSelection(state, key)
	case KeyG:
		next := copyState(state)
		next.PendingKey = key
		return Result{
			State: copyState(next),
			Events: []Event{{
				Type: EventPendingKey,
				Key:  key,
			}},
		}
	case KeyH:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveHorizontal(next, key, -1)
		})
	case KeyL:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveHorizontal(next, key, 1)
		})
	case KeyJ:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveVertical(next, key, 1)
		})
	case KeyK:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveVertical(next, key, -1)
		})
	case KeyW:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveWordForward(next, key)
		})
	case KeyB:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveWordBackward(next, key)
		})
	case KeyE:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveWordEnd(next, key)
		})
	case KeyShiftG:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveDocumentEnd(next, key)
		})
	case KeyZero:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveLineStart(next, key)
		})
	case KeyDollar:
		return applyVisualMotion(state, key, func(next State) Result {
			return moveLineEnd(next, key)
		})
	default:
		return Result{
			State: copyState(state),
			Events: []Event{{
				Type:    EventUnsupportedKey,
				Key:     key,
				Message: "key is not supported in visual mode",
			}},
		}
	}
}

func deleteVisualSelection(state State, key string) Result {
	if state.Selection != nil && state.Selection.Kind == SelectionLinewise {
		return deleteLinewiseVisualSelection(state, key)
	}
	row, start, end, ok := visualLineRange(state)
	if !ok {
		return unsupportedVisualOperator(state, key)
	}
	next := pushUndo(state)
	line := []rune(next.Lines[row])
	if start >= end {
		return boundary(state, key)
	}
	next.Register = Register{
		Text: string(line[start:end]),
	}
	line = append(line[:start], line[end:]...)
	next.Lines[row] = string(line)
	next.Mode = ModeNormal
	next.Selection = nil
	next.Cursor.Row = row
	next.Cursor.Col = clampCol(start, next.Lines[row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func yankVisualSelection(state State, key string) Result {
	if state.Selection != nil && state.Selection.Kind == SelectionLinewise {
		return yankLinewiseVisualSelection(state, key)
	}
	row, start, end, ok := visualLineRange(state)
	if !ok {
		return unsupportedVisualOperator(state, key)
	}
	next := copyState(state)
	line := []rune(next.Lines[row])
	if start >= end {
		return boundary(state, key)
	}
	next.Register = Register{
		Text: string(line[start:end]),
	}
	next.Mode = ModeNormal
	next.Selection = nil
	next.Cursor.Row = row
	next.Cursor.Col = clampCol(start, next.Lines[row])
	next.Cursor.DesiredCol = next.Cursor.Col
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventYanked,
			Key:  key,
		}},
	}
}

func deleteLinewiseVisualSelection(state State, key string) Result {
	start, end, ok := visualRowRange(state)
	if !ok {
		return unsupportedVisualOperator(state, key)
	}
	next := pushUndo(state)
	next.Register = Register{
		Lines:    copyLines(next.Lines[start : end+1]),
		Linewise: true,
	}
	next.Lines = append(next.Lines[:start], next.Lines[end+1:]...)
	if len(next.Lines) == 0 {
		next.Lines = []string{""}
	}
	row := start
	if row >= len(next.Lines) {
		row = len(next.Lines) - 1
	}
	next.Mode = ModeNormal
	next.PendingKey = ""
	next.Selection = nil
	next.Cursor.Row = row
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventChanged,
			Key:  key,
		}},
	}
}

func yankLinewiseVisualSelection(state State, key string) Result {
	start, end, ok := visualRowRange(state)
	if !ok {
		return unsupportedVisualOperator(state, key)
	}
	next := copyState(state)
	next.Register = Register{
		Lines:    copyLines(next.Lines[start : end+1]),
		Linewise: true,
	}
	next.Mode = ModeNormal
	next.PendingKey = ""
	next.Selection = nil
	next.Cursor.Row = start
	next.Cursor.Col = 0
	next.Cursor.DesiredCol = 0
	return Result{
		State: copyState(next),
		Events: []Event{{
			Type: EventYanked,
			Key:  key,
		}},
	}
}

func visualRowRange(state State) (int, int, bool) {
	if state.Selection == nil || !state.Selection.Active || state.Selection.Kind != SelectionLinewise {
		return 0, 0, false
	}
	start := state.Selection.Start.Row
	end := state.Selection.End.Row
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start >= len(state.Lines) {
		start = len(state.Lines) - 1
	}
	if end >= len(state.Lines) {
		end = len(state.Lines) - 1
	}
	return start, end, start <= end
}

func visualLineRange(state State) (int, int, int, bool) {
	if state.Selection == nil || !state.Selection.Active || state.Selection.Kind != SelectionCharwise {
		return 0, 0, 0, false
	}
	if state.Selection.Start.Row != state.Selection.End.Row {
		return 0, 0, 0, false
	}
	row := state.Selection.Start.Row
	if row < 0 || row >= len(state.Lines) {
		return 0, 0, 0, false
	}
	line := []rune(state.Lines[row])
	start := state.Selection.Start.Col
	end := state.Selection.End.Col + 1
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	return row, start, end, start < end
}

func unsupportedVisualOperator(state State, key string) Result {
	return Result{
		State: copyState(state),
		Events: []Event{{
			Type:    EventUnsupportedKey,
			Key:     key,
			Message: "visual operator is not supported for this selection",
		}},
	}
}

func applyVisualMotion(state State, key string, move func(State) Result) Result {
	next := copyState(state)
	if next.Selection == nil || !next.Selection.Active {
		next.Selection = normalizedSelectionForLines(Selection{
			Active: true,
			Kind:   SelectionCharwise,
			Anchor: next.Cursor,
			Head:   next.Cursor,
		}, next.Lines)
	}
	anchor := next.Selection.Anchor
	kind := next.Selection.Kind
	result := move(next)
	result.State.Mode = ModeVisual
	result.State.PendingKey = ""
	result.State.Selection = normalizedSelectionForLines(Selection{
		Active: true,
		Kind:   kind,
		Anchor: anchor,
		Head:   result.State.Cursor,
	}, result.State.Lines)
	return result
}
