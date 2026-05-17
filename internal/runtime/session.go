package runtime

import "github.com/young-st511/advimture/internal/vimengine"

type Status string

const (
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
)

type Exercise struct {
	ID      string
	Initial vimengine.State
	Goal    Goal
	Hints   []Hint
}

type Hint struct {
	AfterKeys int
	Text      string
}

type Goal struct {
	Cursor *vimengine.Cursor
	Mode   *vimengine.Mode
	Lines  []string
}

type State struct {
	ExerciseID string
	Status     Status
	Vim        vimengine.State
	KeyTrace   []string
	Attempts   int
}

type StepResult struct {
	State       State
	Vim         vimengine.Result
	MatchedGoal bool
}

type Session struct {
	exercise Exercise
	initial  vimengine.State
	engine   *vimengine.Engine
	status   Status
	keyTrace []string
	attempts int
}

func CursorGoal(row int, col int) *vimengine.Cursor {
	return &vimengine.Cursor{
		Row:        row,
		Col:        col,
		DesiredCol: col,
	}
}

func ModeGoal(mode vimengine.Mode) *vimengine.Mode {
	next := mode
	return &next
}

func NewSession(exercise Exercise) *Session {
	initial := vimengine.NewWithState(exercise.Initial).State()
	session := &Session{
		exercise: copyExercise(exercise),
		initial:  copyVimState(initial),
		engine:   vimengine.NewWithState(initial),
		status:   StatusRunning,
		attempts: 1,
	}
	if session.exercise.Goal.Matches(initial) {
		session.status = StatusSucceeded
	}
	return session
}

func (s *Session) State() State {
	return State{
		ExerciseID: s.exercise.ID,
		Status:     s.status,
		Vim:        copyVimState(s.engine.State()),
		KeyTrace:   copyStrings(s.keyTrace),
		Attempts:   s.attempts,
	}
}

func (s *Session) ApplyKey(key string) StepResult {
	if s.status != StatusRunning {
		return StepResult{
			State:       s.State(),
			Vim:         vimengine.Result{State: copyVimState(s.engine.State())},
			MatchedGoal: s.status == StatusSucceeded,
		}
	}

	vimResult := s.engine.Apply(key)
	s.keyTrace = append(s.keyTrace, key)
	matched := s.exercise.Goal.Matches(vimResult.State)
	if matched {
		s.status = StatusSucceeded
	}

	return StepResult{
		State:       s.State(),
		Vim:         copyVimResult(vimResult),
		MatchedGoal: matched,
	}
}

func (s *Session) Retry() State {
	s.engine = vimengine.NewWithState(s.initial)
	s.status = StatusRunning
	s.keyTrace = nil
	s.attempts++
	if s.exercise.Goal.Matches(s.engine.State()) {
		s.status = StatusSucceeded
	}
	return s.State()
}

func (s *Session) CurrentHint() (string, bool) {
	keyCount := len(s.keyTrace)
	var selected string
	for _, hint := range s.exercise.Hints {
		if hint.AfterKeys <= keyCount {
			selected = hint.Text
		}
	}
	if selected == "" {
		return "", false
	}
	return selected, true
}

func (g Goal) Matches(state vimengine.State) bool {
	normalized := vimengine.NewWithState(state).State()
	if g.Cursor != nil {
		if normalized.Cursor.Row != g.Cursor.Row || normalized.Cursor.Col != g.Cursor.Col {
			return false
		}
	}
	if g.Mode != nil && normalized.Mode != *g.Mode {
		return false
	}
	if g.Lines != nil && !sameStrings(normalized.Lines, g.Lines) {
		return false
	}
	return true
}

func copyExercise(exercise Exercise) Exercise {
	next := exercise
	next.Initial = copyVimState(exercise.Initial)
	next.Goal = copyGoal(exercise.Goal)
	next.Hints = copyHints(exercise.Hints)
	return next
}

func copyGoal(goal Goal) Goal {
	next := goal
	if goal.Cursor != nil {
		cursor := *goal.Cursor
		next.Cursor = &cursor
	}
	if goal.Mode != nil {
		mode := *goal.Mode
		next.Mode = &mode
	}
	next.Lines = copyStrings(goal.Lines)
	return next
}

func copyHints(hints []Hint) []Hint {
	if hints == nil {
		return nil
	}
	next := make([]Hint, len(hints))
	copy(next, hints)
	return next
}

func copyVimResult(result vimengine.Result) vimengine.Result {
	return vimengine.Result{
		State:  copyVimState(result.State),
		Events: copyVimEvents(result.Events),
	}
}

func copyVimEvents(events []vimengine.Event) []vimengine.Event {
	if events == nil {
		return nil
	}
	next := make([]vimengine.Event, len(events))
	copy(next, events)
	return next
}

func copyVimState(state vimengine.State) vimengine.State {
	next := state
	next.Lines = copyStrings(state.Lines)
	return next
}

func copyStrings(values []string) []string {
	if values == nil {
		return nil
	}
	next := make([]string, len(values))
	copy(next, values)
	return next
}

func sameStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
