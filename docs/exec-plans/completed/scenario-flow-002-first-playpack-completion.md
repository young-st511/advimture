# ExecPlan: First playpack scenario completion

Slice-ID: SCENARIO-FLOW-002
Created: 2026-05-18
Status: active
Scope-Mode: normal
Allowed-Paths:
- content/exercises/first-five-minutes.yaml
- content/scenarios/first-five-minutes.yaml
- content/playlists/first-five-minutes.yaml
- docs/gameplay/spec.md
- docs/gameplay/command-catalog.md
- docs/gameplay/content-requirements.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- cmd/e2e-runner/**
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/scenario-flow-002-first-playpack-completion.md
- docs/exec-plans/completed/scenario-flow-002-first-playpack-completion.md
- internal/content/loader_test.go
- internal/playable/model_test.go
- test/e2e/**

## 목표

첫 플레이팩의 scenario/content 완성도를 playable 기준으로 끌어올린다. SCENARIO-FLOW-001에서 남긴 `normal-motion-basic`의 `h/k` coverage gap을 닫고, 첫 플레이팩이 17개 approved + replay-pass exercise로 처음부터 끝까지 완주되게 한다.

## 영향 도메인

- Gameplay: `h/j/k/l` 기본 이동 cluster의 optimal trace coverage를 완성한다.
- Scenario: 기본 이동 4방향을 같은 DevOps/터미널 상황 톤으로 묶는다.
- Verification: loader coverage, playable model, full E2E count와 progress assertion을 갱신한다.

## 수용 기준

- `normal-motion-basic`의 `coverage_required: ["h", "j", "k", "l"]`가 모두 approved + replay-pass exercise의 optimal trace에 등장한다.
- 새 `h`, `k` exercise는 target state, optimal keys, allowed keys, hints, grading, E2E assertions를 가진다.
- 새 scenario는 exercise target/keys/pass condition을 바꾸지 않고 `does_not_change`를 유지한다.
- 첫 플레이팩 full E2E는 17개 exercise를 순서대로 완주한다.
- 실제 progress 저장 포맷과 content schema는 변경하지 않는다.

## 범위

- 포함: `normal-motion-basic` h/k exercise 추가
- 포함: h/k scenario 추가
- 포함: playlist와 E2E 순서 동기화
- 포함: docs bank와 current spec 갱신
- 제외: 새 command cluster 추가
- 제외: 엔진 동작 변경
- 제외: progress 저장 포맷 변경
- 제외: multi-playlist selection 구현

## 검증 계획

- `go test ./internal/content/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] current docs/content/E2E 기준을 확인한다.
- [x] `h` exercise와 scenario를 추가한다.
- [x] `k` exercise와 scenario를 추가한다.
- [x] playlist 순서와 E2E assertion을 17개 exercise 기준으로 동기화한다.
- [x] E2E runner wait가 과거 화면 로그에 오탐하지 않도록 보강한다.
- [x] docs bank, spec, PROGRAM을 갱신한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 커밋/푸시한다.

## 완료 결과

- `normal-motion-basic-003`으로 `h` 왼쪽 이동 문항과 scenario를 추가했다.
- `normal-motion-basic-004`로 `k` 위쪽 이동 문항과 scenario를 추가했다.
- `normal-motion-basic` coverage report가 `h/j/k/l` 전체 covered 상태가 되었다.
- 첫 플레이팩 full E2E를 17개 exercise 완주 경로로 갱신했다.
- E2E runner의 `wait_screen_contains`가 이전 wait 이후 새 출력만 보게 하여 `Next: enter` 오탐을 막았다.

## 검증 결과

- `go test ./cmd/e2e-runner` 통과
- `go test ./internal/content/...` 통과
- `go test ./internal/playable/...` 통과
- `go test ./...` 통과
- `make e2e-smoke` 통과
- `make e2e-playable` 통과
- `git diff --check` 통과
