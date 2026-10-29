package playable

import "testing"

func TestFairReuseAlternatives(t *testing.T) {
	for _, tc := range []struct{ id, keys string }{{"command-choice-line-reuse-001", "yyjp"}, {"command-choice-line-reuse-001", "Vyjp"}, {"command-choice-quote-reuse-001", "yi\"j$P"}, {"command-choice-quote-reuse-001", "yi\"G$P"}, {"incident-reuse-002", "yyp"}, {"incident-reuse-002", "Vyp"}} {
		t.Run(tc.id+tc.keys, func(t *testing.T) {
			m := New(Options{ContentRoot: contentRootForTest(), Progress: progressCompleteBefore(t, tc.id)})
			m, _ = updateWithKey(t, m, tc.keys)
			if m.State().Status != "succeeded" || m.State().Score.Grade != "S" {
				t.Fatal(m.State())
			}
		})
	}
}
