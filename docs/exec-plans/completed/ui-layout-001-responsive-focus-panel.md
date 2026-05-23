# UI-LAYOUT-001 — Responsive FocusPanel Layout

Slice-ID: UI-LAYOUT-001
Created: 2026-05-23
Status: completed
Completed: 2026-05-23
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/gameplay/tui-ux-direction.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 영향 도메인

- Gameplay: 사용자 노출 문구는 한국어 톤을 유지하고, 저장 포맷과 content schema는 변경하지 않는다.
- Verification: 화면 레이아웃 변경 후에도 key trace, progress, app_state assertion은 유지한다.

## 목표

`UI-MODAL-001`에서 FocusPanel을 console 위로 올렸지만, 렌더러는 아직 고정 폭 박스를 그대로 사용한다. 이 slice는 Bubble Tea `WindowSizeMsg`를 playable model에 저장하고 renderer에 전달해, FocusPanel이 터미널 폭에 맞춰 중앙 정렬되는 modal-like 레이아웃으로 보이게 한다.

## 범위

- 포함:
  - playable model의 terminal width/height 저장
  - `playableview.Screen` width/height 필드 추가
  - FocusPanel width clamp와 horizontal centering
  - renderer/model 단위 테스트
  - 대표 E2E 회귀 확인
- 제외:
  - modal open 상태 입력 차단
  - vertical center overlay
  - screen 전체 재배치
  - 색상/theme 대규모 변경
  - 새 dependency 추가

## 수용 기준

- `tea.WindowSizeMsg` 수신 후 playable view는 terminal width를 renderer에 전달한다.
- FocusPanel은 terminal width가 충분하면 화면 중앙에 배치된다.
- FocusPanel은 좁은 화면에서 terminal width를 넘지 않도록 width를 줄인다.
- terminal width가 없거나 0이면 기존 fallback 폭으로 렌더링한다.
- 기존 key trace, progress, app_state UI focus_panel assertion은 유지된다.

## Step 1: Responsive FocusPanel

- 목표: FocusPanel을 width-aware centered panel로 렌더링한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - `internal/playable/model_test.go`
- 충족 기준: WindowSizeMsg 전달, centered/clamped FocusPanel
- Boundaries 주의: 저장 포맷, content schema, dependency 변경 금지
- 상세 작업:
  - [x] `Screen.Width/Height` 추가
  - [x] `Model`에 terminal size 저장
  - [x] FocusPanel width 계산과 horizontal centering 구현
  - [x] fallback 렌더링 유지

## Step 2: 검증과 완료 처리

- 목표: 대표 E2E와 문서 상태를 갱신한다.
- 변경 파일:
  - `docs/roadmap/PROGRAM.md`
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
- 테스트 파일: n/a
- 충족 기준: docs가 responsive FocusPanel 현재 동작과 후속 vertical overlay 제외 범위를 설명한다.
- Boundaries 주의: true vertical overlay는 후속 slice로 남긴다.
- 상세 작업:
  - [x] gameplay docs에 width-aware FocusPanel 반영
  - [x] ExecPlan completed 이동
  - [x] `go test ./...`, 대표 E2E 실행

## 검증 결과

- `go test ./internal/playableview ./internal/playable`

---

## 실행 규칙

각 Step은 Spec 기반 테스트 작성 → 구현 → 변경 범위 테스트 → 필요한 E2E → 문서 동기화 → 커밋/푸시 순서로 진행한다.
