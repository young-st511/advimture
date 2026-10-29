package runtime

import (
	"github.com/young-st511/advimture/internal/vimengine"
	"testing"
)

func TestReuseRequiresActualYankAndPut(t *testing.T) {
	ex := Exercise{ID: "reuse", Initial: vimengine.NewState([]string{"abc", ""}), Goal: Goal{Lines: []string{"abc", "abc"}, Mode: ModeGoal(vimengine.ModeNormal)}, Constraints: Constraints{RequireReuse: true, RequiredKeys: []string{"y", "p"}}}
	s := NewSession(ex)
	for _, k := range []string{"i", "y", "p", "esc", "u", "j", "i", "a", "b", "c", "esc"} {
		s.ApplyKey(k)
	}
	if s.State().Status == StatusSucceeded {
		t.Fatal("typed key names falsely satisfy reuse")
	}
	s = NewSession(ex)
	for _, k := range []string{"y", "y", "p"} {
		s.ApplyKey(k)
	}
	// The result is a new line rather than replacement of the empty line.
	if !s.yankPut {
		t.Fatal("actual yank/put not recognized")
	}
	s.Retry()
	if s.yankPut {
		t.Fatal("retry retained provenance")
	}
}
