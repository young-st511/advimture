# UI-PLAYTEST-001 — HUD Usability Pass

Slice-ID: UI-PLAYTEST-001
Created: 2026-05-25
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/tui-ux-direction.md
- docs/gameplay/first-five-minute-route.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- docs/exec-plans/active/ui-playtest-001-hud-usability-pass.md
- docs/exec-plans/completed/ftue-001-first-five-minute-route.md
- internal/playableview/
- internal/playable/
- test/e2e/
- artifacts/e2e/

## 목표

FTUE-001에서 고정한 첫 5분 route를 기준으로 `Mission HUD + Runbook Console + floating modal` 화면이 실제로 읽기 좋은지 점검하고, 새 레이아웃의 polish backlog와 필요한 최소 수정을 닫는다.

## 범위

- 포함:
  - FTUE running/failed/succeeded 화면 evidence 확인
  - status line의 개발식 표현 점검
  - modal 긴 문장 overflow와 action line 보존 확인
  - command mode prompt와 floating modal의 충돌 여부 확인
  - 필요한 경우 renderer 단위의 작은 copy/spacing 개선
- 제외:
  - split cockpit wide layout
  - 실제 modal input-state 차단
  - progress 저장 포맷 변경
  - 새 command/content schema 추가

## 수용 기준

- FTUE 대표 evidence에서 running/failed/succeeded 화면의 정보 우선순위가 명확하다.
- `MISSION` cue line이 과도하게 길거나 콘솔보다 더 강하게 시선을 빼앗는 경우 개선 후보로 기록한다.
- floating modal은 실패/성공 이유와 다음 행동을 함께 보존한다.
- status line에 남은 개발식 label은 후속 또는 이번 slice의 작은 개선으로 분류한다.
- 검증은 renderer/playable test와 focused E2E 또는 `make e2e-playable`로 수행한다.

## Step 1: Evidence Review

- [ ] `playable_ftue_first_five_route` evidence 확인
- [ ] failure focused evidence 확인
- [ ] command mode evidence 확인

## Step 2: Minimal Polish

- [ ] status/cue/modal에서 즉시 고칠 작은 문제 분류
- [ ] renderer 테스트 추가 또는 갱신
- [ ] 필요한 최소 구현

## Step 3: Verification

- [ ] focused E2E
- [ ] `go test ./...`
- [ ] `git diff --check`
