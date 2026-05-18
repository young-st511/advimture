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

func TestRunCaseCanLockWordMotionFixture(t *testing.T) {
	testCase := Case{
		Name:    "word-motion-starts",
		Initial: vimengine.NewState([]string{"service api backend enabled"}),
		Keys:    []string{vimengine.KeyW, vimengine.KeyW, vimengine.KeyE},
	}
	executor := fakeExecutor{
		state: vimengine.State{
			Mode:  vimengine.ModeNormal,
			Lines: []string{"service api backend enabled"},
			Cursor: vimengine.Cursor{
				Row: 0,
				Col: 18,
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

func TestRunCaseCanLockNavigationMotionFixture(t *testing.T) {
	initial := vimengine.NewWithState(vimengine.State{
		Mode:  vimengine.ModeNormal,
		Lines: []string{"server {", "route api", "status ok"},
		Cursor: vimengine.Cursor{
			Row:        1,
			Col:        5,
			DesiredCol: 5,
		},
	}).State()
	testCase := Case{
		Name:    "navigation-line-and-file",
		Initial: initial,
		Keys:    []string{vimengine.KeyDollar, vimengine.KeyG, vimengine.KeyG, vimengine.KeyShiftG},
	}
	executor := fakeExecutor{
		state: vimengine.State{
			Mode:  vimengine.ModeNormal,
			Lines: []string{"server {", "route api", "status ok"},
			Cursor: vimengine.Cursor{
				Row: 2,
				Col: 0,
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
