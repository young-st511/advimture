package playableview

import (
	"strings"
	"testing"

	"github.com/young-st511/advimture/internal/tuiadapter"
)

func TestRenderLineShowsVisualSelection(t *testing.T) {
	line := RenderLine("abcd", 0, 0, 3, &tuiadapter.SelectionView{
		Active: true,
		Kind:   "charwise",
		Start:  tuiadapter.CursorView{Row: 0, Col: 1},
		End:    tuiadapter.CursorView{Row: 0, Col: 3},
	})

	if line != "> a{b}{c}[d]" {
		t.Fatalf("RenderLine = %q, want visual selection", line)
	}
}

func TestRenderLineShowsLinewiseVisualSelection(t *testing.T) {
	line := RenderLine("abcd", 0, 0, 2, &tuiadapter.SelectionView{
		Active: true,
		Kind:   "linewise",
		Start:  tuiadapter.CursorView{Row: 0, Col: 0},
		End:    tuiadapter.CursorView{Row: 0, Col: 3},
	})

	if line != "> {a}{b}[c]{d}" {
		t.Fatalf("RenderLine = %q, want linewise visual selection", line)
	}
}

func TestRenderScreenIncludesFocusPanelBeforeConsole(t *testing.T) {
	view := Render(Screen{
		PlaylistTitle: "Tutorial 0",
		Title:         "커서 위치 맞추기",
		Message:       "목표까지 이동하세요.",
		BufferLines:   []string{"abc"},
		Mode:          "normal",
		Status:        "running",
		CursorCol:     1,
		ExerciseTotal: 4,
		FocusPanel: &FocusPanel{
			Kind:  "training",
			Title: "TRAINING BRIEF",
			Lines: []string{"Coach: 훈련 키 l", "?: hint  q: quit"},
		},
	})

	for _, want := range []string{"Tutorial 0", "커서 위치 맞추기", "> a[b]c", "Exercise: 1/4", "TRAINING BRIEF"} {
		if !strings.Contains(view, want) {
			t.Fatalf("Render output = %q, want %q", view, want)
		}
	}
	panelIndex := strings.Index(view, "TRAINING BRIEF")
	consoleIndex := strings.Index(view, "RUNBOOK CONSOLE")
	if panelIndex == -1 || consoleIndex == -1 || panelIndex > consoleIndex {
		t.Fatalf("Render output = %q, want focus panel before console", view)
	}
}

func TestRenderCentersFocusPanelWhenWidthIsKnown(t *testing.T) {
	view := Render(Screen{
		Width:       100,
		Title:       "커서 위치 맞추기",
		Message:     "목표까지 이동하세요.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "running",
		FocusPanel: &FocusPanel{
			Kind:  "training",
			Title: "TRAINING BRIEF",
			Lines: []string{"?: hint  q: quit"},
		},
	})

	borderIndex := strings.Index(view, "┌")
	consoleIndex := strings.Index(view, "RUNBOOK CONSOLE")
	if borderIndex == -1 || consoleIndex == -1 {
		t.Fatalf("Render output = %q, want focus panel and console", view)
	}
	panelLineStart := strings.LastIndex(view[:borderIndex], "\n") + 1
	if borderIndex-panelLineStart == 0 {
		t.Fatalf("Render output = %q, want centered focus panel with leading spaces", view)
	}
}

func TestRenderShrinksFocusPanelForNarrowWidth(t *testing.T) {
	view := Render(Screen{
		Width:       36,
		Title:       "커서 위치 맞추기",
		Message:     "목표까지 이동하세요.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "running",
		FocusPanel: &FocusPanel{
			Kind:  "training",
			Title: "TRAINING BRIEF",
			Lines: []string{"?: hint  q: quit"},
		},
	})

	for _, line := range strings.Split(view, "\n") {
		if strings.Contains(line, "TRAINING BRIEF") && len([]rune(line)) > 36 {
			t.Fatalf("focus panel line width = %d, want <= 36: %q", len([]rune(line)), line)
		}
	}
}

func TestRenderPrioritizesCurrentTaskBeforeOpsSummary(t *testing.T) {
	view := Render(Screen{
		PlaylistTitle: "Tutorial 0",
		ReviewSummary: "재점검 대상: 목표 문자까지 이동하기: 미복구",
		DailyRoute:    "오늘의 복구 루트: 3건 대기",
		Title:         "커서 위치 맞추기",
		Message:       "목표까지 이동하세요.",
		BufferLines:   []string{"abc"},
		Mode:          "normal",
		Status:        "running",
		ExerciseTotal: 4,
		FocusPanel: &FocusPanel{
			Kind:  "training",
			Title: "TRAINING BRIEF",
			Lines: []string{"?: hint  q: quit"},
		},
	})

	titleIndex := strings.Index(view, "커서 위치 맞추기")
	reviewIndex := strings.Index(view, "재점검 대상:")
	consoleIndex := strings.Index(view, "RUNBOOK CONSOLE")
	bufferIndex := strings.Index(view, "> [a]bc")
	if titleIndex == -1 || reviewIndex == -1 || consoleIndex == -1 || bufferIndex == -1 {
		t.Fatalf("Render output = %q, want current task, ops, and console sections", view)
	}
	if titleIndex > reviewIndex {
		t.Fatalf("Render output = %q, want current task before ops summary", view)
	}
	if consoleIndex > bufferIndex {
		t.Fatalf("Render output = %q, want console label before buffer", view)
	}
	if !strings.Contains(view, "ADVIMTURE | Tutorial 0 | Exercise: 1/4 | Status: running") {
		t.Fatalf("Render output = %q, want compact header", view)
	}
}
