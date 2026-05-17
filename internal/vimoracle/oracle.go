package vimoracle

import (
	"context"
	"fmt"
	"reflect"

	"github.com/young-st511/advimture/internal/vimengine"
)

type Case struct {
	Name    string
	Initial vimengine.State
	Keys    []string
}

type Executor interface {
	RunOracle(context.Context, Case) (vimengine.State, error)
}

type Result struct {
	CaseName   string
	Matched    bool
	Engine     vimengine.State
	Oracle     vimengine.State
	Mismatches []Mismatch
}

type Mismatch struct {
	Field  string
	Engine string
	Oracle string
}

func RunCase(ctx context.Context, executor Executor, testCase Case) (Result, error) {
	engineResult := vimengine.ApplyKeys(testCase.Initial, testCase.Keys)
	oracleState, err := executor.RunOracle(ctx, copyCase(testCase))
	if err != nil {
		return Result{}, err
	}
	result := Compare(engineResult.State, oracleState)
	result.CaseName = testCase.Name
	return result, nil
}

func Compare(engineState vimengine.State, oracleState vimengine.State) Result {
	engine := vimengine.NewWithState(engineState).State()
	oracle := vimengine.NewWithState(oracleState).State()
	mismatches := make([]Mismatch, 0)

	if !reflect.DeepEqual(engine.Lines, oracle.Lines) {
		mismatches = append(mismatches, Mismatch{
			Field:  "lines",
			Engine: fmt.Sprintf("%q", engine.Lines),
			Oracle: fmt.Sprintf("%q", oracle.Lines),
		})
	}
	if engine.Cursor.Row != oracle.Cursor.Row || engine.Cursor.Col != oracle.Cursor.Col {
		mismatches = append(mismatches, Mismatch{
			Field:  "cursor",
			Engine: fmt.Sprintf("%d,%d", engine.Cursor.Row, engine.Cursor.Col),
			Oracle: fmt.Sprintf("%d,%d", oracle.Cursor.Row, oracle.Cursor.Col),
		})
	}
	if engine.Mode != oracle.Mode {
		mismatches = append(mismatches, Mismatch{
			Field:  "mode",
			Engine: string(engine.Mode),
			Oracle: string(oracle.Mode),
		})
	}

	return Result{
		Matched:    len(mismatches) == 0,
		Engine:     copyState(engine),
		Oracle:     copyState(oracle),
		Mismatches: mismatches,
	}
}

func copyCase(testCase Case) Case {
	return Case{
		Name:    testCase.Name,
		Initial: copyState(testCase.Initial),
		Keys:    copyStrings(testCase.Keys),
	}
}

func copyState(state vimengine.State) vimengine.State {
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
