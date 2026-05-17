package content

import (
	"fmt"
	"strings"
	"unicode/utf8"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/vimengine"
)

type ExerciseSpec struct {
	ID               string
	CommandClusterID string
	Title            string
	Initial          StateSpec
	Goal             GoalSpec
	Hints            []HintSpec
	ExpectedKeys     []string
	AllowedKeys      []string
}

type StateSpec struct {
	Lines  []string
	Cursor *CursorSpec
	Mode   string
}

type GoalSpec struct {
	Cursor *CursorSpec
	Mode   string
	Lines  []string
}

type CursorSpec struct {
	Row int
	Col int
}

type HintSpec struct {
	AfterKeys int
	Text      string
}

type CompiledExercise struct {
	Exercise         exerciseruntime.Exercise
	CommandClusterID string
	Title            string
	ExpectedKeys     []string
	AllowedKeys      []string
}

func CursorSpecPtr(row int, col int) *CursorSpec {
	return &CursorSpec{Row: row, Col: col}
}

func CompileExercise(spec ExerciseSpec) (CompiledExercise, error) {
	if err := validateExerciseSpec(spec); err != nil {
		return CompiledExercise{}, err
	}

	initialMode, err := parseMode(spec.Initial.Mode)
	if err != nil {
		return CompiledExercise{}, fmt.Errorf("initial mode: %w", err)
	}
	initial := vimengine.NewState(spec.Initial.Lines)
	initial.Mode = initialMode
	if spec.Initial.Cursor != nil {
		initial.Cursor = vimCursor(*spec.Initial.Cursor)
	}
	initial = vimengine.NewWithState(initial).State()

	goal, err := compileGoal(spec.Goal)
	if err != nil {
		return CompiledExercise{}, err
	}

	hints := make([]exerciseruntime.Hint, len(spec.Hints))
	for index, hint := range spec.Hints {
		hints[index] = exerciseruntime.Hint{
			AfterKeys: hint.AfterKeys,
			Text:      hint.Text,
		}
	}

	return CompiledExercise{
		Exercise: exerciseruntime.Exercise{
			ID:      spec.ID,
			Initial: initial,
			Goal:    goal,
			Hints:   hints,
		},
		CommandClusterID: spec.CommandClusterID,
		Title:            spec.Title,
		ExpectedKeys:     copyStrings(spec.ExpectedKeys),
		AllowedKeys:      copyStrings(spec.AllowedKeys),
	}, nil
}

func validateExerciseSpec(spec ExerciseSpec) error {
	if strings.TrimSpace(spec.ID) == "" {
		return fmt.Errorf("exercise id is required")
	}
	if len(spec.Initial.Lines) == 0 {
		return fmt.Errorf("initial lines are required")
	}
	if _, err := parseMode(spec.Initial.Mode); err != nil {
		return fmt.Errorf("initial mode: %w", err)
	}
	if spec.Initial.Cursor != nil && !cursorInLines(*spec.Initial.Cursor, spec.Initial.Lines) {
		return fmt.Errorf("initial cursor is out of range")
	}
	if !spec.Goal.hasAnyTarget() {
		return fmt.Errorf("goal must include cursor, mode, or lines")
	}
	if spec.Goal.Mode != "" {
		if _, err := parseMode(spec.Goal.Mode); err != nil {
			return fmt.Errorf("goal mode: %w", err)
		}
	}
	if spec.Goal.Cursor != nil {
		goalLines := spec.Initial.Lines
		if spec.Goal.Lines != nil {
			goalLines = spec.Goal.Lines
		}
		if !cursorInLines(*spec.Goal.Cursor, goalLines) {
			return fmt.Errorf("goal cursor is out of range")
		}
	}
	for index, hint := range spec.Hints {
		if hint.AfterKeys < 0 {
			return fmt.Errorf("hint %d after_keys must be non-negative", index)
		}
		if strings.TrimSpace(hint.Text) == "" {
			return fmt.Errorf("hint %d text is required", index)
		}
	}
	return nil
}

func compileGoal(spec GoalSpec) (exerciseruntime.Goal, error) {
	goal := exerciseruntime.Goal{
		Lines: copyStrings(spec.Lines),
	}
	if spec.Cursor != nil {
		goal.Cursor = exerciseruntime.CursorGoal(spec.Cursor.Row, spec.Cursor.Col)
	}
	if spec.Mode != "" {
		mode, err := parseMode(spec.Mode)
		if err != nil {
			return exerciseruntime.Goal{}, err
		}
		goal.Mode = exerciseruntime.ModeGoal(mode)
	}
	return goal, nil
}

func parseMode(value string) (vimengine.Mode, error) {
	switch value {
	case "", string(vimengine.ModeNormal):
		return vimengine.ModeNormal, nil
	case string(vimengine.ModeInsert):
		return vimengine.ModeInsert, nil
	case string(vimengine.ModeCommand):
		return vimengine.ModeCommand, nil
	default:
		return "", fmt.Errorf("unsupported mode %q", value)
	}
}

func (g GoalSpec) hasAnyTarget() bool {
	return g.Cursor != nil || g.Mode != "" || g.Lines != nil
}

func vimCursor(cursor CursorSpec) vimengine.Cursor {
	return vimengine.Cursor{
		Row:        cursor.Row,
		Col:        cursor.Col,
		DesiredCol: cursor.Col,
	}
}

func cursorInLines(cursor CursorSpec, lines []string) bool {
	if cursor.Row < 0 || cursor.Row >= len(lines) || cursor.Col < 0 {
		return false
	}
	return cursor.Col <= maxCursorCol(lines[cursor.Row])
}

func maxCursorCol(line string) int {
	length := utf8.RuneCountInString(line)
	if length == 0 {
		return 0
	}
	return length - 1
}

func copyStrings(values []string) []string {
	if values == nil {
		return nil
	}
	next := make([]string, len(values))
	copy(next, values)
	return next
}
