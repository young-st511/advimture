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

func TestRenderScreenIncludesActionPanel(t *testing.T) {
	view := Render(Screen{
		PlaylistTitle: "Tutorial 0",
		Title:         "커서 위치 맞추기",
		Message:       "목표까지 이동하세요.",
		BufferLines:   []string{"abc"},
		Mode:          "normal",
		Status:        "running",
		CursorCol:     1,
		ExerciseTotal: 4,
		ActionLines:   []string{"ACTION", "?: hint  q: quit"},
	})

	for _, want := range []string{"Tutorial 0", "커서 위치 맞추기", "> a[b]c", "Exercise: 1/4", "ACTION"} {
		if !strings.Contains(view, want) {
			t.Fatalf("Render output = %q, want %q", view, want)
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
		ActionLines:   []string{"ACTION"},
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
