# ExecPlan: Adventure Beat Console Layer

Slice-ID: UI-ADVENTURE-BEAT-001
Created: 2026-06-20
Status: completed
Scope-Mode: strict
Allowed-Paths:
- `docs/exec-plans/completed/ui-adventure-beat-001.md`
- `docs/gameplay/tui-screen-contract.md`
- `docs/gameplay/tui-ux-direction.md`
- `docs/gameplay/spec.md`
- `docs/roadmap/UX_BACKLOG_001.md`
- `internal/playable/**`
- `internal/playableview/**`
- `test/e2e/**`

## 영향 도메인

- Gameplay: running 화면의 adventure feedback, input response, animation tick을 다룬다.
- Verification: renderer/unit test와 focused E2E로 text-only 환경에서도 의미가 유지되는지 검증한다.

## 공통 제한

- `go.mod`, `go.sum` 변경 없음.
- `internal/progress/` 저장 포맷 변경 없음.
- content schema 변경 없음.
- 새 Vim engine capability 추가 없음.
- 새 dependency 추가 없음.
- 특정 상용 게임의 UI, 문구, asset을 복제하지 않는다. 텍스트 어드벤처/퍼즐 게임의 “짧은 장면 전환, 선택 반응, 상태 연출” 원칙만 Advimture 톤으로 번역한다.

## 배경

Modal hierarchy는 개선됐지만 running 화면은 여전히 설명형 텍스트가 많고, 입력에 반응하는 게임적 표식이 부족하다. 사용자는 ASCII art 기반 애니메이션과 텍스트 인터랙션을 포함해 adventure 게임에 맞는 UI 방향을 요구했다.

이번 slice는 큰 레이아웃 재작성 전에 1줄짜리 `adventure signal rail`을 추가한다. Mission HUD와 Runbook Console 사이의 핵심 정보 흐름을 유지하면서, 플레이어 입력에 짧게 반응하고 terminal tick으로 신호가 움직이는 느낌을 준다.

## Todo 1: Contract and Direction

- 상태: completed
- 목표: adventure beat의 위치와 금지 경계를 screen contract/UX direction에 기록한다.
- 충족 기준:
  - signal rail은 현재 mission보다 앞서지 않고, console 접근을 늦추지 않는 1줄 보조 연출로 정의된다.
  - input echo는 저장 포맷과 runtime key trace가 아니라 화면 반응 전용 상태임을 명시한다.

## Todo 2: Render Adventure Signal Rail

- 상태: completed
- 목표: HUD running 화면에 animated ASCII signal rail과 input echo를 표시한다.
- 충족 기준:
  - `AnimationFrame`에 따라 ASCII marker 위치가 바뀐다.
  - `InputEcho`가 있으면 짧은 입력 반응을 표시한다.
  - 80 column에서 line overflow 없이 `SIGNAL` rail이 보인다.
  - color/style이 없어도 의미를 읽을 수 있다.

## Todo 3: Wire Bubble Tea Interaction

- 상태: completed
- 목표: playable model이 tick으로 animation frame을 갱신하고, key/hint/ignored input을 echo한다.
- 충족 기준:
  - `Init()`은 animation tick command를 반환한다.
  - `Update()`는 tick마다 다음 tick을 예약한다.
  - key input 후 `InputEcho`가 renderer에 전달된다.
  - progress 저장, content schema, engine capability는 변경하지 않는다.

## Todo 4: Verify and Close

- 상태: completed
- 목표: focused test/E2E와 전체 테스트 후 completed plan으로 이동한다.
- 충족 기준:
  - `go test ./internal/playable ./internal/playableview` 통과.
  - focused E2E가 running signal/input echo를 확인한다.
  - `go test ./...` 통과.
  - `git diff --check` 통과.

## 검증

- `go test ./internal/playable ./internal/playableview`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_adventure_signal_input_echo.yaml`
- `go test ./...`
- `git diff --check`
