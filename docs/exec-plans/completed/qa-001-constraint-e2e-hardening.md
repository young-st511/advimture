# ExecPlan: Constraint E2E hardening

Slice-ID: QA-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/qa-001-constraint-e2e-hardening.md
- docs/exec-plans/completed/qa-001-constraint-e2e-hardening.md
- test/e2e/playable_constraint_forbidden_retry.yaml
- test/e2e/playable_constraint_required_key.yaml
- test/e2e/playable_constraint_max_inputs.yaml
- Makefile

## 목표

Constraint/failure/scoring 루프가 실제 TUI E2E에서 충분히 감시되는지 검증하고, 회귀 방지 루프를 명시한다.

## 수용 기준

- `make e2e-playable`은 forbidden input, max input 초과, required key 누락, retry 복구를 모두 포함한다.
- forbidden input E2E는 실패 후 `enter` retry로 같은 exercise를 성공 완료하는 흐름을 검증한다.
- required key E2E는 목표에 도착해도 의도 입력이 없으면 실패하고 progress를 저장하지 않음을 검증한다.
- max input E2E는 입력 제한 초과 시 실패하고 progress를 저장하지 않음을 검증한다.
- 전체 playable E2E가 통과한다.

## 범위

- 포함: E2E coverage audit
- 포함: Makefile 루프 확인
- 제외: 새 runtime 기능 구현

## 검증 계획

- `make e2e-playable`
- `go test ./...`
- `git diff --check`

## 검증 결과

- `make e2e-playable`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 작업 항목

- [x] E2E 시나리오 coverage를 확인한다.
- [x] `make e2e-playable`을 실행한다.
- [x] `go test ./...`를 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
