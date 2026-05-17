package vimoracle

import (
	"context"
	"errors"
	"testing"

	"github.com/young-st511/advimture/internal/vimengine"
)

func TestRunCaseMatchesOracleState(t *testing.T) {
	testCase := Case{
		Name:    "move-right",
		Initial: vimengine.NewState([]string{"abc"}),
		Keys:    []string{vimengine.KeyL, vimengine.KeyL},
	}
	executor := fakeExecutor{
		state: vimengine.State{
			Mode:  vimengine.ModeNormal,
			Lines: []string{"abc"},
			Cursor: vimengine.Cursor{
				Row: 0,
				Col: 2,
			},
		},
	}

	result, err := RunCase(context.Background(), executor, testCase)
	if err != nil {
		t.Fatalf("RunCase returned error: %v", err)
	}
	if !result.Matched {
		t.Fatalf("Matched = false, mismatches = %+v", result.Mismatches)
	}
}

func TestRunCaseReportsMismatches(t *testing.T) {
	testCase := Case{
		Name:    "move-right",
		Initial: vimengine.NewState([]string{"abc"}),
		Keys:    []string{vimengine.KeyL},
	}
	executor := fakeExecutor{
		state: vimengine.State{
			Mode:  vimengine.ModeInsert,
			Lines: []string{"abc"},
			Cursor: vimengine.Cursor{
				Row: 0,
				Col: 2,
			},
		},
	}

	result, err := RunCase(context.Background(), executor, testCase)
	if err != nil {
		t.Fatalf("RunCase returned error: %v", err)
	}
	if result.Matched {
		t.Fatal("Matched = true, want false")
	}
	assertMismatch(t, result.Mismatches, "cursor")
	assertMismatch(t, result.Mismatches, "mode")
}

func TestRunCasePropagatesOracleError(t *testing.T) {
	wantErr := errors.New("oracle unavailable")
	_, err := RunCase(context.Background(), fakeExecutor{err: wantErr}, Case{
		Name:    "oracle-error",
		Initial: vimengine.NewState([]string{"abc"}),
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}
}

func TestCompareReportsLineMismatch(t *testing.T) {
	result := Compare(vimengine.State{
		Mode:  vimengine.ModeNormal,
		Lines: []string{"abc"},
	}, vimengine.State{
		Mode:  vimengine.ModeNormal,
		Lines: []string{"xyz"},
	})

	if result.Matched {
		t.Fatal("Matched = true, want false")
	}
	assertMismatch(t, result.Mismatches, "lines")
}

type fakeExecutor struct {
	state vimengine.State
	err   error
}

func (f fakeExecutor) RunOracle(context.Context, Case) (vimengine.State, error) {
	return f.state, f.err
}

func assertMismatch(t *testing.T, mismatches []Mismatch, field string) {
	t.Helper()

	for _, mismatch := range mismatches {
		if mismatch.Field == field {
			return
		}
	}
	t.Fatalf("missing mismatch field %q in %+v", field, mismatches)
}
