# ExecPlan: Action Bar and Dense Running HUD

Slice-ID: UI-ACTION-HUD-001
Created: 2026-06-23
Status: completed
Scope-Mode: strict
Allowed-Paths:
- `docs/exec-plans/completed/ui-action-hud-001.md`
- `docs/gameplay/spec.md`
- `docs/gameplay/tui-screen-contract.md`
- `docs/gameplay/tui-ux-direction.md`
- `docs/roadmap/PROGRAM.md`
- `docs/roadmap/UX_BACKLOG_001.md`
- `docs/roadmap/CHANGES.md`
- `internal/playable/**`
- `internal/playableview/**`
- `test/e2e/**`
- `artifacts/e2e/**`

## 영향 도메인

- Gameplay: running HUD, action affordance, mode-specific action language를 다룬다.
- Verification: renderer tests와 focused E2E가 80x24 화면에서 action bar와 dense HUD를 검증한다.

## Domain Contract 핵심 제약

- 사용자 노출 문구는 한국어 톤을 우선한다.
- 저장 포맷 변경은 사용자 승인 없이 하지 않는다.
- 새 content schema의 ID와 파일명 변경은 기획 승인 없이 하지 않는다.
- 새 의존성은 추가하지 않는다.
- TUI 조작은 key trace로 재현 가능해야 한다.

## 수용 기준 참조

- `docs/gameplay/spec.md`:
  - terminal size가 있는 running HUD는 `MISSION`, `GOAL`, `TOOLS` 또는 `JUDGMENT`, `SIGNAL`, `ACTIONS`를 본문과 구분해 표시한다.
  - running action은 `ACTIONS  [key] label` grammar로 표시한다.
  - mode-specific 화면은 hint/quit보다 현재 mode의 실행/취소 action을 우선한다.
  - 80-column running 화면에서 detailed review/daily line은 현재 목표와 console 접근을 밀어내지 않는다.

## 공통 제한

- release/tag/push 작업은 하지 않는다.
- `go.mod`, `go.sum` 변경 없음.
- `internal/progress/` 저장 포맷 변경 없음.
- content schema 변경 없음.
- 새 Vim engine capability 추가 없음.
- success/failure recovery report 구조 변경은 `UI-REPORT-001`로 분리한다.

## Step 1: Dense Running HUD Contract

- 상태: completed
- 목표: running 화면의 정보 구조를 `MISSION/GOAL/TOOLS or JUDGMENT/SIGNAL/ACTIONS`로 고정한다.
- 변경 파일:
  - `docs/gameplay/spec.md`
  - `docs/gameplay/tui-screen-contract.md`
  - `docs/gameplay/tui-ux-direction.md`
  - `docs/exec-plans/completed/ui-action-hud-001.md`
- 테스트 파일:
  - `internal/playableview/render_test.go`
- 충족 기준:
  - 80x24 tutorial running viewport에서 `MISSION`, `GOAL`, `TOOLS`, `SIGNAL`, `ACTIONS`가 보인다.
  - 80x24 incident running viewport에서 `MISSION`, `GOAL` 또는 `SITUATION`, `JUDGMENT`, `SIGNAL`, `ACTIONS`가 보인다.
  - running 화면의 review/daily detail은 현재 목표보다 높은 계층에 올라오지 않는다.
- Boundaries 주의:
  - progress/content schema 변경 금지.

## Step 2: Action Bar Renderer and Mode Actions

- 상태: completed
- 목표: `FocusPanel.actions`를 running 화면에서 `ACTIONS [key] label` action bar로 렌더링하고, mode-specific action을 명확히 표시한다.
- 변경 파일:
  - `internal/playableview/render.go`
  - `internal/playableview/render_test.go`
  - `internal/playable/model.go`
  - `internal/playable/model_test.go`
- 테스트 파일:
  - `internal/playableview/render_test.go`
  - `internal/playable/model_test.go`
- 충족 기준:
  - `hint`는 `[?] 힌트 - grade 영향`, `quit`은 `[q] 종료`로 표시된다.
  - command/search/insert/visual mode는 현재 mode의 실행/취소/normal 복귀 action을 hint/quit보다 우선 표시한다.
  - 기존 `FocusPanel.actions[].id` app_state 계약은 유지된다.
- Boundaries 주의:
  - action id를 추가하더라도 저장 포맷에는 반영하지 않는다.

## Step 3: Focused E2E and Close

- 상태: completed
- 목표: focused E2E와 전체 검증 후 plan을 completed로 이동한다.
- 변경 파일:
  - `test/e2e/**`
  - `docs/roadmap/**`
  - `docs/exec-plans/completed/ui-action-hud-001.md`
- 테스트 파일:
  - `test/e2e/playable_action_hud_running.yaml`
  - 필요 시 기존 focused viewport E2E
- 충족 기준:
  - focused E2E에서 running action bar와 dense HUD가 검증된다.
  - `go test ./internal/playable ./internal/playableview` 통과.
  - `go test ./...` 통과.
  - `git diff --check` 통과.

---

## 실행 규칙 (하네스 모드)

각 Step은 아래 사이클을 순서대로 수행한다.

### 1. Spec 기반 테스트 작성

구현 전, spec.md 수용 기준을 기반으로 테스트를 먼저 작성한다.

### 2. 구현

테스트를 통과하도록 해당 Step의 작업을 수행한다.

### 3. 검증

- `go test ./internal/playable ./internal/playableview`
- focused E2E
- `go test ./...`
- `git diff --check`

### 4. 문서 동기화

- spec/contract/backlog/changelog를 현재 동작과 맞춘다.

### 5. 커밋

검증 통과 후 문서와 코드를 함께 커밋한다.

## 완료 기록

- 2026-06-23: running HUD를 `MISSION`, `GOAL`, `TOOLS`/`JUDGMENT`, `SIGNAL`, `ACTIONS` grammar로 재정리했다.
- 2026-06-23: `FocusPanel.actions`를 running 화면의 `ACTIONS  [key] label` action bar로 렌더링하고, command/search/insert/visual mode action이 hint/quit보다 우선 보이게 했다.
- 2026-06-23: 기존 E2E의 긴 briefing/old mode 문자열 wait를 새 HUD/action grammar 기준으로 갱신했다.

## 검증

- `go test ./internal/playableview ./internal/playable`
- `go test ./...`
- `make e2e-playable`
- `git diff --check`

## 완료 판단

`UI-ACTION-HUD-001`은 completed다. Running 화면은 현재 mission과 다음 조작을 먼저 읽히게 하고, 실패/성공 recovery report 구조는 후속 `UI-REPORT-001` 범위로 남긴다. Progress schema, content schema, dependency, engine capability는 변경하지 않았다.
