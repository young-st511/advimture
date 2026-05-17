# ExecPlan: App state summary assertion

Slice-ID: E2E-003
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- cmd/e2e-runner/**
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/verification/**
- docs/exec-plans/active/e2e-003-app-state-summary-assertion.md

## 목표

첫 playable slice에서 TUI 화면 텍스트만이 아니라 app state를 검증할 수 있도록 E2E runner에 app state summary assertion을 추가한다.

## 범위

- 포함: scenario YAML의 `assert.app_state` schema
- 포함: test HOME 내부 app state summary JSON 읽기
- 포함: buffer/cursor/mode/status/score/progress assertion
- 포함: summary evidence에 app state path와 load 여부 기록
- 제외: 실제 Bubble Tea app이 state summary를 쓰는 구현
- 제외: 첫 playable slice 구현

## State Summary 계약

기본 위치는 test HOME 기준 `.advimture/e2e_state.json`이다. 첫 playable slice는 `ADVIMTURE_E2E=1`일 때 이 파일을 쓸 수 있다.

```json
{
  "buffer": ["abc"],
  "cursor": {"row": 0, "col": 2},
  "mode": "normal",
  "status": "succeeded",
  "score": {"grade": "S", "passed": true},
  "progress": {"mission_id": "mission-1", "completed": true}
}
```

## 검증 계획

- `go test ./cmd/e2e-runner`
- `go test ./...`
- `make e2e-smoke`

## 승인 체크

- [x] app state summary JSON을 읽을 수 있다.
- [x] buffer/cursor/mode/status/score/progress assertion이 동작한다.
- [x] app state가 없는데 assertion이 있으면 실패한다.
- [x] summary evidence에 app state load 여부가 기록된다.
- [x] 전체 테스트와 smoke E2E가 통과한다.

## 의사결정 로그

- 2026-05-18: runner는 기본 위치 `.advimture/e2e_state.json`을 읽는다. 실제 앱이 state summary를 쓰는 구현은 PLAY-001에서 연결한다.
- 2026-05-18: app state assertion은 화면 텍스트와 별개로 buffer/cursor/mode/status/score/progress를 검증한다. 첫 playable slice의 성공 여부를 화면만으로 판단하지 않기 위함이다.
