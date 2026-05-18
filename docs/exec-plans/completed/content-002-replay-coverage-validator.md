# ExecPlan: Replay and coverage validator

Slice-ID: CONTENT-002
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/content/**
- content/**
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/content-002-replay-coverage-validator.md

## 목표

approved/implemented content가 실제 플레이 가능한 상태인지 loader 단계에서 검증한다. content authoring이 늘어나도 replay가 깨진 문항이나 coverage 기준이 빠진 cluster가 playable 후보로 승격되지 않게 한다.

## 영향 도메인

- Gameplay: content schema ID/파일명은 변경하지 않는다. 설계 순서는 `Vim command -> Exercise -> Scenario`를 유지한다.
- Verification: E2E 부족이 느껴지면 app wiring을 멈추고 validator/unit test를 먼저 보강한다.

## 수용 기준

- approved + implemented command cluster는 비어 있지 않은 `coverage_required`를 가져야 한다.
- approved + implemented exercise는 `replay_status: pass`가 아니면 load에 실패한다.
- `replay_status: pass`인 approved + implemented exercise는 `optimal_keys`를 실제 runtime에 재생했을 때 목표 상태와 E2E assertion을 만족해야 한다.
- `PlayableExercises()`는 replay gate를 통과한 exercise만 반환한다.
- 기존 첫 playable content는 validator를 통과하고 PLAY-002 E2E를 계속 통과한다.

## 범위

- 포함: content loader validation 강화
- 포함: first playable fixture의 replay status 승격
- 포함: content unit test
- 제외: 자동 replay status 파일 갱신 도구
- 제외: cluster 전체 coverage를 모두 채우는 신규 exercise 제작
- 제외: app/game loop 변경

## 검증 계획

- `go test ./internal/content/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`

## 작업 항목

- [x] validator 수용 기준을 Red test로 작성한다.
- [x] replay status와 optimal key replay 검증을 구현한다.
- [x] `PlayableExercises()`가 replay gate를 반영하게 한다.
- [x] root content fixture를 gate 통과 상태로 승격한다.
- [x] 문서를 completed 상태로 동기화한다.

## 의사결정 로그

- 2026-05-18: `replay_status: pass`는 사람이 표시하지만, loader가 실제 `optimal_keys` replay와 E2E assertion으로 다시 검증한다.
- 2026-05-18: approved + implemented exercise는 `e2e_assertions.buffer/cursor/mode/status`를 모두 요구한다.
- 2026-05-18: runtime이 성공 이후 입력을 무시하므로, replay key trace가 `optimal_keys` 전체와 정확히 일치하는지도 검사한다.

## 완료 결과

- approved + implemented command cluster에 `coverage_required` 필수 gate를 추가했다.
- approved + implemented exercise에 `replay_status: pass` 필수 gate를 추가했다.
- pass exercise는 runtime replay, key trace, E2E assertion을 loader 단계에서 검증한다.
- 첫 playable fixture `normal-motion-basic-001`을 `replay_status: pass`로 승격했다.
