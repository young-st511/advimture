package scenario

import (
	"fmt"
	"strings"

	"github.com/young-st511/advimture/internal/content"
	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scoring"
)

type Spec struct {
	ID          string
	Title       string
	Briefing    string
	SuccessText string
	FailureText string
	Exercise    content.CompiledExercise
}

type State struct {
	ScenarioID string
	Title      string
	Briefing   string
	Message    string
	Status     exerciseruntime.Status
	Runtime    exerciseruntime.State
	Score      *scoring.Result
	HintsUsed  int
}

type StepResult struct {
	State   State
	Runtime exerciseruntime.StepResult
}

type Run struct {
	spec      Spec
	session   *exerciseruntime.Session
	message   string
	score     *scoring.Result
	hintsUsed int
}

func NewRun(spec Spec) (*Run, error) {
	if strings.TrimSpace(spec.ID) == "" {
		return nil, fmt.Errorf("scenario id is required")
	}
	if strings.TrimSpace(spec.Exercise.Exercise.ID) == "" {
		return nil, fmt.Errorf("scenario exercise is required")
	}

	run := &Run{
		spec:    copySpec(spec),
		session: exerciseruntime.NewSession(spec.Exercise.Exercise),
		message: spec.Briefing,
	}
	if run.message == "" {
		run.message = spec.Title
	}
	run.scoreIfSucceeded()
	return run, nil
}

func (r *Run) State() State {
	return State{
		ScenarioID: r.spec.ID,
		Title:      r.spec.Title,
		Briefing:   r.briefing(),
		Message:    r.message,
		Status:     r.session.State().Status,
		Runtime:    r.session.State(),
		Score:      copyScore(r.score),
		HintsUsed:  r.hintsUsed,
	}
}

func (r *Run) briefing() string {
	if r.spec.Briefing != "" {
		return r.spec.Briefing
	}
	return r.spec.Title
}

func (r *Run) ApplyKey(key string) StepResult {
	step := r.session.ApplyKey(key)
	if step.State.Status == exerciseruntime.StatusSucceeded {
		r.message = r.spec.SuccessText
		r.scoreIfSucceeded()
	} else if step.State.Status == exerciseruntime.StatusFailed {
		r.message = r.failureMessage(step.State.Message)
		r.scoreIfFailed()
	}
	return StepResult{
		State:   r.State(),
		Runtime: step,
	}
}

func (r *Run) RequestHint() (string, bool) {
	hint, ok := r.session.CurrentHint()
	if !ok {
		return "", false
	}
	r.hintsUsed++
	return hint, true
}

func (r *Run) Retry() State {
	r.session.Retry()
	r.message = r.spec.Briefing
	if r.message == "" {
		r.message = r.spec.Title
	}
	r.score = nil
	r.hintsUsed = 0
	return r.State()
}

func (r *Run) scoreIfSucceeded() {
	state := r.session.State()
	if state.Status != exerciseruntime.StatusSucceeded {
		return
	}
	result := scoring.Evaluate(scoring.Input{
		Status:       state.Status,
		Failure:      state.Failure,
		KeyTrace:     state.KeyTrace,
		ExpectedKeys: r.spec.Exercise.ExpectedKeys,
		Attempts:     state.Attempts,
		HintsUsed:    r.hintsUsed,
	})
	r.score = &result
}

func (r *Run) scoreIfFailed() {
	state := r.session.State()
	result := scoring.Evaluate(scoring.Input{
		Status:       state.Status,
		Failure:      state.Failure,
		KeyTrace:     state.KeyTrace,
		ExpectedKeys: r.spec.Exercise.ExpectedKeys,
		Attempts:     state.Attempts,
		HintsUsed:    r.hintsUsed,
	})
	r.score = &result
}

func (r *Run) failureMessage(runtimeMessage string) string {
	if strings.TrimSpace(r.spec.FailureText) != "" {
		if strings.TrimSpace(runtimeMessage) != "" {
			return r.spec.FailureText + " " + runtimeMessage
		}
		return r.spec.FailureText
	}
	if strings.TrimSpace(runtimeMessage) != "" {
		return runtimeMessage
	}
	return "아직 목표를 달성하지 못했습니다. 다시 시도하세요."
}

func copySpec(spec Spec) Spec {
	next := spec
	next.Exercise = copyCompiledExercise(spec.Exercise)
	return next
}

func copyCompiledExercise(exercise content.CompiledExercise) content.CompiledExercise {
	next := exercise
	next.ExpectedKeys = copyStrings(exercise.ExpectedKeys)
	next.AllowedKeys = copyStrings(exercise.AllowedKeys)
	return next
}

func copyScore(score *scoring.Result) *scoring.Result {
	if score == nil {
		return nil
	}
	next := *score
	return &next
}

func copyStrings(values []string) []string {
	if values == nil {
		return nil
	}
	next := make([]string, len(values))
	copy(next, values)
	return next
}
