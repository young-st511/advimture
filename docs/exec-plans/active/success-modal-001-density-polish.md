# SUCCESS-MODAL-001 — Success Modal Density Polish

Slice-ID: SUCCESS-MODAL-001
Created: 2026-05-26
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/tui-screen-contract.md
- docs/exec-plans/active/success-modal-001-density-polish.md
- internal/playable/
- internal/playableview/
- test/e2e/

## 목표

성공 floating modal이 `RUNBOOK SEALED`와 `STEP SEALED`, record/debrief/action line을 중복적으로 보여주는 문제를 줄이고, 다음 행동과 학습 피드백이 더 빠르게 읽히도록 정리한다.

## 범위

- 포함:
  - success modal heading/density 점검
  - renderer/model 테스트 갱신
  - 대표 success E2E expectation 갱신
- 제외:
  - 실패 modal 재설계
  - progress 저장 포맷 변경
  - score 계산 변경
  - 화면 전체 레이아웃 재설계

## 수용 기준

- success modal은 같은 화면 안에서 `RUNBOOK SEALED`와 `STEP SEALED`를 중복 heading처럼 노출하지 않는다.
- 성공 기록, best record, runbook 완료 수, 잔류 리스크, next action은 유지된다.
- next action line은 terminal width 안에서 잘리지 않는다.
- renderer/playable tests, focused E2E, `make e2e-playable`을 통과한다.

## Step 1: Current Modal Review

- [ ] success modal renderer와 playable success lines 확인
- [ ] 중복 heading과 필수 정보 분리

## Step 2: Implementation

- [ ] success modal density 조정
- [ ] renderer/playable tests 갱신
- [ ] E2E expectation 갱신

## Step 3: Verification

- [ ] focused tests
- [ ] focused E2E
- [ ] `go test ./...`
- [ ] `make e2e-playable`
- [ ] `git diff --check`
