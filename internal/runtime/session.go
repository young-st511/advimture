package runtime

import "github.com/young-st511/advimture/internal/vimengine"

type Status string

const (
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

type FailureReason string

const (
	FailureNone                FailureReason = ""
	FailureForbiddenInput      FailureReason = "forbidden_input"
	FailureMaxInputsExceeded   FailureReason = "max_inputs_exceeded"
	FailureRequiredKeysMissing FailureReason = "required_keys_missing"
)

type Exercise struct {
	ID          string
	Initial     vimengine.State
	Goal        Goal
	Hints       []Hint
	Constraints Constraints
}

type Hint struct {
	AfterKeys int
	Text      string
}

type Constraints struct {
	MaxInputs     int
	RequiredKeys  []string
	ForbiddenKeys []string
	AttemptLimit  int
}

type Goal struct {
	Cursor  *vimengine.Cursor
	Mode    *vimengine.Mode
	Lines   []string
	Command *string
}

type State struct {
	ExerciseID   string
	Status       Status
	Vim          vimengine.State
	KeyTrace     []string
	RequiredKeys []string
	Attempts     int
	AttemptLimit int
	MaxInputs    int
	InputsLeft   int
	Failure      FailureReason
	Message      string
}

type StepResult struct {
	State       State
	Vim         vimengine.Result
	MatchedGoal bool
}

type Session struct {
	exercise                       Exercise
	initial                        vimengine.State
	engine                         *vimengine.Engine
	status                         Status
	keyTrace                       []string
	attempts                       int
	failure                        FailureReason
	message                        string
	matchedGoalWithMissingRequired bool
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

func CommandGoal(command string) *string {
	return &command
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
	if session.canSucceedWithCurrentTrace(initial) {
		session.status = StatusSucceeded
	}
	return session
}

func (s *Session) State() State {
	return State{
		ExerciseID:   s.exercise.ID,
		Status:       s.status,
		Vim:          copyVimState(s.engine.State()),
		KeyTrace:     copyStrings(s.keyTrace),
		RequiredKeys: copyStrings(s.exercise.Constraints.RequiredKeys),
		Attempts:     s.attempts,
		AttemptLimit: s.exercise.Constraints.AttemptLimit,
		MaxInputs:    s.exercise.Constraints.MaxInputs,
		InputsLeft:   s.inputsLeft(),
		Failure:      s.failure,
		Message:      s.message,
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

	s.keyTrace = append(s.keyTrace, key)
	if s.isForbidden(key) {
		s.fail(FailureForbiddenInput, "이 입력은 이번 문항에서 사용할 수 없습니다.")
		return StepResult{
			State:       s.State(),
			Vim:         vimengine.Result{State: copyVimState(s.engine.State())},
			MatchedGoal: false,
		}
	}
	if s.exceededMaxInputs() {
		s.fail(FailureMaxInputsExceeded, "입력 제한을 초과했습니다.")
		return StepResult{
			State:       s.State(),
			Vim:         vimengine.Result{State: copyVimState(s.engine.State())},
			MatchedGoal: false,
		}
	}

	vimResult := s.engine.Apply(key)
	matched := s.exercise.Goal.Matches(vimResult.State)
	if !matched {
		s.matchedGoalWithMissingRequired = false
	}
	if matched {
		if missing := s.missingRequiredKeys(); len(missing) > 0 {
			if s.shouldFailMissingRequiredKeys() {
				s.fail(FailureRequiredKeysMissing, "목표에는 도착했지만 이번 문항의 의도한 입력을 사용하지 않았습니다.")
			} else {
				s.matchedGoalWithMissingRequired = true
			}
			matched = false
		} else if s.matchedGoalWithMissingRequired {
			s.fail(FailureRequiredKeysMissing, "목표에는 도착했지만 이번 문항의 의도한 입력을 사용하지 않았습니다.")
			matched = false
		} else {
			s.status = StatusSucceeded
		}
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
	s.failure = FailureNone
	s.message = ""
	s.matchedGoalWithMissingRequired = false
	if s.canSucceedWithCurrentTrace(s.engine.State()) {
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
	if g.Command != nil && normalized.LastCommand != *g.Command {
		return false
	}
	return true
}

func (s *Session) isForbidden(key string) bool {
	for _, forbidden := range s.exercise.Constraints.ForbiddenKeys {
		if forbidden == key {
			return true
		}
	}
	return false
}

func (s *Session) exceededMaxInputs() bool {
	maxInputs := s.exercise.Constraints.MaxInputs
	return maxInputs > 0 && len(s.keyTrace) > maxInputs
}

func (s *Session) missingRequiredKeys() []string {
	var missing []string
	for _, required := range s.exercise.Constraints.RequiredKeys {
		if !containsString(s.keyTrace, required) {
			missing = append(missing, required)
		}
	}
	return missing
}

func (s *Session) shouldFailMissingRequiredKeys() bool {
	maxInputs := s.exercise.Constraints.MaxInputs
	return maxInputs <= 0 || len(s.keyTrace) >= maxInputs
}

func (s *Session) canSucceedWithCurrentTrace(state vimengine.State) bool {
	if !s.exercise.Goal.Matches(state) {
		return false
	}
	return len(s.missingRequiredKeys()) == 0
}

func (s *Session) inputsLeft() int {
	maxInputs := s.exercise.Constraints.MaxInputs
	if maxInputs <= 0 {
		return -1
	}
	left := maxInputs - len(s.keyTrace)
	if left < 0 {
		return 0
	}
	return left
}

func (s *Session) fail(reason FailureReason, message string) {
	s.status = StatusFailed
	s.failure = reason
	s.message = message
}

func copyExercise(exercise Exercise) Exercise {
	next := exercise
	next.Initial = copyVimState(exercise.Initial)
	next.Goal = copyGoal(exercise.Goal)
	next.Hints = copyHints(exercise.Hints)
	next.Constraints = copyConstraints(exercise.Constraints)
	return next
}

func copyConstraints(constraints Constraints) Constraints {
	return Constraints{
		MaxInputs:     constraints.MaxInputs,
		RequiredKeys:  copyStrings(constraints.RequiredKeys),
		ForbiddenKeys: copyStrings(constraints.ForbiddenKeys),
		AttemptLimit:  constraints.AttemptLimit,
	}
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
	if goal.Command != nil {
		command := *goal.Command
		next.Command = &command
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

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
