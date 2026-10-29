package scoring

import (
	"github.com/young-st511/advimture/internal/runtime"
	"testing"
)

func TestEquivalentEfficientSolutionCanEarnS(t *testing.T) {
	got := Evaluate(Input{Status: runtime.StatusSucceeded, KeyTrace: []string{"y", "y", "j", "p"}, ExpectedKeys: []string{"V", "y", "j", "p"}, Attempts: 1})
	if got.Grade != GradeS {
		t.Fatal(got)
	}
}
