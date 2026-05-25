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

func TestRenderFocusPanelOverlayDoesNotMoveConsoleWhenHeightIsKnown(t *testing.T) {
	base := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "커서 위치 맞추기",
		Message:     "목표까지 이동하세요.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "running",
	})
	withOverlay := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "커서 위치 맞추기",
		Message:     "목표까지 이동하세요.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "running",
		FocusPanel: &FocusPanel{
			Kind:  "failure",
			Title: "RECOVERY REQUIRED",
			Lines: []string{
				"실패 원인입니다.",
				"Inputs left: 1/2",
				"Attempts: 1/unlimited",
				"Retry: r or enter",
			},
		},
	})

	if lineIndex(base, "RUNBOOK CONSOLE") != lineIndex(withOverlay, "RUNBOOK CONSOLE") {
		t.Fatalf("base = %q\nwithOverlay = %q\nwant console line unchanged", base, withOverlay)
	}
}

func TestRenderHUDPlacesMissionBeforeConsoleWhenSizeIsKnown(t *testing.T) {
	view := Render(Screen{
		Width:            100,
		Height:           30,
		PlaylistTitle:    "Tutorial 2",
		PlaylistCategory: "tutorial",
		ReviewSummary:    "재점검 대상: 단어 시작점으로 뛰어가기: 미복구",
		DailyRoute:       "오늘의 복구 루트: 3건 대기",
		ReviewCount:      3,
		ReviewPrimary:    "단어 시작점으로 뛰어가기",
		Title:            "서비스 이름 찾기",
		Message:          "backend로 바로 이동하세요.",
		BufferLines:      []string{"service api backend enabled"},
		Mode:             "normal",
		Status:           "running",
		ExerciseTotal:    7,
		FocusPanel: &FocusPanel{
			Kind:  "training",
			Title: "TRAINING BRIEF",
			Lines: []string{"Coach: 훈련 키 w", "?: hint  q: quit"},
		},
	})

	missionIndex := strings.Index(view, "MISSION")
	consoleIndex := strings.Index(view, "RUNBOOK CONSOLE")
	bufferIndex := strings.Index(view, "> [s]ervice")
	if missionIndex == -1 || consoleIndex == -1 || bufferIndex == -1 {
		t.Fatalf("Render output = %q, want mission HUD and console", view)
	}
	if missionIndex > consoleIndex || consoleIndex > bufferIndex {
		t.Fatalf("Render output = %q, want mission -> console -> buffer order", view)
	}
	if strings.Contains(view, "\n복구 현황\n") {
		t.Fatalf("Render output = %q, should not render recovery status as a large pre-console section", view)
	}
	if !strings.Contains(view, "복구 메모: 재점검 3건 · 다음: 단어 시작점으로 뛰어가기") {
		t.Fatalf("Render output = %q, want compact recovery memo in mission HUD", view)
	}
	if strings.Contains(view, "오늘의 복구 루트: 3건 대기") {
		t.Fatalf("Render output = %q, should not expose detailed daily route in running tutorial HUD", view)
	}
	cueIndex := strings.Index(view, "TRAINING BRIEF")
	recoveryIndex := strings.Index(view, "복구 메모:")
	if cueIndex == -1 || recoveryIndex == -1 || cueIndex > recoveryIndex {
		t.Fatalf("Render output = %q, want action cue before recovery memo", view)
	}
	if !strings.Contains(view, "NORMAL · running · cursor 0:0") {
		t.Fatalf("Render output = %q, want polished HUD status line", view)
	}
	if strings.Contains(view, "Mode: normal") || strings.Contains(view, "Cursor: 0,0") {
		t.Fatalf("Render output = %q, should not show debug status labels in HUD", view)
	}
}

func TestRenderHUDUsesIncidentRecoverySummary(t *testing.T) {
	view := Render(Screen{
		Width:            100,
		Height:           30,
		PlaylistTitle:    "릴레이 기지 001",
		PlaylistCategory: "incident",
		ReviewSummary:    "재점검 대상: timeout 위치 추적: 미복구",
		DailyRoute:       "오늘의 복구 루트: timeout 위치 추적(미복구) 외 2건 대기",
		ReviewCount:      3,
		ReviewPrimary:    "timeout 위치 추적",
		Title:            "timeout 위치 추적",
		Message:          "장애 로그에서 timeout marker를 찾아 복구 지점을 고정하세요.",
		BufferLines:      []string{"ok", "timeout"},
		Mode:             "normal",
		Status:           "running",
		FocusPanel: &FocusPanel{
			Kind:  "incident",
			Title: "OPERATOR JUDGMENT",
			Lines: []string{"판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요.", "?: hint  q: quit"},
		},
	})

	if !strings.Contains(view, "복구 현황: 재점검 3건 · 잔류: timeout 위치 추적") {
		t.Fatalf("Render output = %q, want compact incident recovery status", view)
	}
	if strings.Contains(view, "오늘의 복구 루트: timeout 위치 추적") {
		t.Fatalf("Render output = %q, should not expose detailed daily route in running incident HUD", view)
	}
}

func TestRenderHUDWrapsLongBriefingBeforeConsole(t *testing.T) {
	view := Render(Screen{
		Width:            72,
		Height:           30,
		PlaylistCategory: "incident",
		Title:            "복구 범위 판별",
		Message:          "릴레이 기지의 라우팅 파일에 안전한 route 값과 quarantine 블록이 함께 보입니다. 값 하나를 고치지 말고 오염된 줄 묶음을 골라 제거하세요.",
		BufferLines:      []string{"route=\"safe\"", "quarantine=temp"},
		Mode:             "normal",
		Status:           "running",
		FocusPanel: &FocusPanel{
			Kind:  "incident",
			Title: "OPERATOR JUDGMENT",
			Lines: []string{"판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요.", "?: hint  q: quit"},
		},
	})

	lines := strings.Split(view, "\n")
	for _, line := range lines {
		if strings.Contains(line, "릴레이 기지의") && displayWidth(line) > hudTextWidth(72) {
			t.Fatalf("briefing line width = %d, want <= %d: %q", displayWidth(line), hudTextWidth(72), line)
		}
	}
	if strings.Contains(view, "오염된 줄 묶음을 골\n") || strings.Contains(view, "오염된 줄 묶음을 골라\n") {
		t.Fatalf("Render output = %q, should not leave incomplete briefing fragment", view)
	}
	if lineIndex(view, "OPERATOR JUDGMENT") > lineIndex(view, "RUNBOOK CONSOLE") {
		t.Fatalf("Render output = %q, want cue before console", view)
	}
}

func TestRenderHUDKeepsCompactRecoverySummaryInModePanel(t *testing.T) {
	view := Render(Screen{
		Width:            100,
		Height:           30,
		PlaylistCategory: "incident",
		ReviewSummary:    "재점검 대상: 복구 범위 판별: 미복구",
		DailyRoute:       "오늘의 복구 루트: 복구 범위 판별(미복구)",
		ReviewCount:      1,
		ReviewPrimary:    "복구 범위 판별",
		Title:            "복구 범위 판별",
		Message:          "줄 묶음을 고르세요.",
		BufferLines:      []string{"quarantine=temp"},
		Mode:             "visual",
		Status:           "running",
		FocusPanel: &FocusPanel{
			Kind:  "mode",
			Title: "VISUAL CHANNEL",
			Lines: []string{"Keys: motion expands selection  esc/v: normal"},
		},
	})

	if !strings.Contains(view, "복구 현황: 재점검 1건 · 잔류: 복구 범위 판별") {
		t.Fatalf("Render output = %q, want compact recovery summary in mode panel", view)
	}
	if strings.Contains(view, "오늘의 복구 루트: 복구 범위 판별") {
		t.Fatalf("Render output = %q, should not expose detailed daily route in mode panel", view)
	}
}

func TestRenderHUDFailureModalAppearsInsideConsoleAfterBuffer(t *testing.T) {
	view := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "서비스 이름 찾기",
		Message:     "backend로 바로 이동하세요.",
		BufferLines: []string{"service api backend enabled"},
		Mode:        "normal",
		Status:      "failed",
		FocusPanel: &FocusPanel{
			Kind:  "failure",
			Title: "RECOVERY REQUIRED",
			Lines: []string{
				"한 글자씩 가면 늦습니다. 다음 단어의 시작으로 이동하는 motion을 사용하세요.",
				"Inputs left: 1/2",
				"Attempts: 1/unlimited",
				"Coach: 훈련 키 w",
				"Retry: r or enter",
			},
		},
	})

	consoleIndex := strings.Index(view, "RUNBOOK CONSOLE")
	bufferIndex := strings.Index(view, "> [s]ervice")
	modalIndex := strings.Index(view, "RECOVERY CHECK")
	if consoleIndex == -1 || bufferIndex == -1 || modalIndex == -1 {
		t.Fatalf("Render output = %q, want console buffer and floating modal", view)
	}
	if !(consoleIndex < bufferIndex && bufferIndex < modalIndex) {
		t.Fatalf("Render output = %q, want floating modal after buffer inside console core", view)
	}
	for _, want := range []string{"RECOVERY REQUIRED", "Mistake", "Next", "Coach: 훈련 키 w", "Retry: r or enter"} {
		if !strings.Contains(view, want) {
			t.Fatalf("Render output = %q, want %q", view, want)
		}
	}
}

func TestRenderFocusPanelOverlayKeepsActionLineWhenContentOverflows(t *testing.T) {
	base := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "되돌아온 커서 정렬",
		Message:     "표시 지점을 오른쪽으로 지나쳤습니다.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "succeeded",
	})
	view := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "되돌아온 커서 정렬",
		Message:     "표시 지점을 오른쪽으로 지나쳤습니다.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "succeeded",
		FocusPanel: &FocusPanel{
			Kind:  "success",
			Title: "STEP SEALED",
			Lines: []string{
				"좋습니다. h는 오른쪽으로 지나쳤을 때 커서를 왼쪽으로 되돌리는 기본 이동입니다.",
				"복구 기록: grade S, 2 keys",
				"최단 복구 기록: grade S, 2 keys",
				"Runbook: 3/4 복구 완료",
				"잔류 리스크: 위쪽 로그 줄로 복귀하기: 미복구",
				"오늘의 복구 루트: 3건 대기",
				"Next: enter",
			},
		},
	})

	if !strings.Contains(view, "Next: enter") {
		t.Fatalf("Render output = %q, want action line preserved", view)
	}
	if strings.Contains(view, "STEP SEALED") {
		t.Fatalf("Render output = %q, should not duplicate success modal heading", view)
	}
	if !strings.Contains(view, "RUNBOOK SEALED") {
		t.Fatalf("Render output = %q, want success modal heading", view)
	}
	if lineIndex(view, "RUNBOOK CONSOLE") != lineIndex(base, "RUNBOOK CONSOLE") {
		t.Fatalf("Render output = %q, want fixed console line", view)
	}
}

func TestRenderFocusPanelOverlayPrioritizesRetryOverQuitWhenFailureOverflows(t *testing.T) {
	view := Render(Screen{
		Width:       100,
		Height:      30,
		Title:       "금지 입력 복구",
		Message:     "한 글자씩 가면 늦습니다.",
		BufferLines: []string{"abc"},
		Mode:        "normal",
		Status:      "failed",
		FocusPanel: &FocusPanel{
			Kind:  "failure",
			Title: "RECOVERY REQUIRED",
			Lines: []string{
				"한 글자씩 가면 늦습니다. 다음 단어의 시작으로 이동하는 motion을 사용하세요. 이 입력은 이번 문항에서 사용할 수 없습니다.",
				"Inputs left: 0/2",
				"Attempts: 1/unlimited",
				"Coach: 훈련 키 w",
				"Retry: r or enter",
				"?: hint  q: quit",
			},
		},
	})

	if !strings.Contains(view, "Retry: r or enter") {
		t.Fatalf("Render output = %q, want retry action preserved", view)
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
		t.Fatalf("Render output = %q, want current task, recovery status, and console sections", view)
	}
	if !strings.Contains(view, "복구 현황") {
		t.Fatalf("Render output = %q, want recovery status section label", view)
	}
	if strings.Contains(view, "OPS") {
		t.Fatalf("Render output = %q, should not expose OPS label", view)
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

func lineIndex(text string, needle string) int {
	for i, line := range strings.Split(text, "\n") {
		if strings.Contains(line, needle) {
			return i
		}
	}
	return -1
}
