package playable

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"testing"
)

func TestRecoveryBatchStopsAtSuccess(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest()})
	m, _ = updateWithKey(t, m, "llj")
	if m.State().Status != "succeeded" || m.currentExerciseID() != "normal-motion-basic-001" || m.State().Cursor.Col != 2 {
		t.Fatalf("batch result: %+v", m.State())
	}
}

func TestRecoveryIgnoresTerminalReply(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest()})
	for _, key := range []tea.KeyMsg{{Type: tea.KeyRunes, Runes: []rune("]"), Alt: true}, {Type: tea.KeyRunes, Runes: []rune("11;rgb:0000/0000/0000")}, {Type: tea.KeyRunes, Runes: []rune("\\"), Alt: true}} {
		next, _ := m.Update(key)
		m = next.(Model)
	}
	if len(m.run.State().Runtime.KeyTrace) != 0 {
		t.Fatal(m.run.State())
	}
	m, _ = updateWithKey(t, m, "ll")
	if m.State().Status != "succeeded" {
		t.Fatal(m.State())
	}
}

func TestRecoveryFailedModeActions(t *testing.T) {
	for _, key := range []string{"enter", "r", "?"} {
		t.Run(key, func(t *testing.T) {
			m := New(Options{ContentRoot: contentRootForTest(), Progress: progressCompleteBefore(t, "insert-mode-entry-001")})
			for _, k := range []string{"i", "z", "z", "z"} {
				m, _ = updateWithKey(t, m, k)
			}
			if m.State().Status != "failed" {
				t.Fatal(m.State())
			}
			m, _ = updateWithKey(t, m, key)
			if key == "?" {
				if !strings.Contains(m.View(), "힌트 내용") {
					t.Fatal(m.View())
				}
			} else if m.State().Status != "running" || m.State().Buffer[0] != "ac" {
				t.Fatal(m.State())
			}
		})
	}
}

func TestRecoveryCommandBackspace(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest(), Progress: progressCompleteBefore(t, "vim-ex-command-substitute-002")})
	m, _ = updateWithKey(t, m, ":%s/TODO/DONX")
	m, _ = updateWithSpecialKey(t, m, tea.KeyBackspace)
	if got := m.run.State().Runtime.Vim.CommandLine; got != "%s/TODO/DON" {
		t.Fatalf("command=%q", got)
	}
	m, _ = updateWithSpecialKey(t, m, tea.KeyLeft)
	if strings.Contains(m.run.State().Runtime.Vim.CommandLine, "left") {
		t.Fatal(m.View())
	}
}

func TestRecoveryFailureFooterFits(t *testing.T) {
	m := New(Options{ContentRoot: contentRootForTest()})
	m, _ = updateWithWindowSize(t, m, 80, 24)
	m, _ = updateWithSpecialKey(t, m, tea.KeyRight)
	m, _ = updateWithKey(t, m, "?")
	for _, want := range []string{"다시 시도:", "힌트: ?", "종료: q"} {
		if !strings.Contains(m.View(), want) {
			t.Fatalf("missing %s: %s", want, m.View())
		}
	}
}
