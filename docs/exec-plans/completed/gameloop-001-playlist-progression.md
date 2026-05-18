# ExecPlan: Playlist progression and autosave loop

Slice-ID: GAMELOOP-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/playable/**
- internal/content/**
- test/e2e/**
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/active/gameloop-001-playlist-progression.md
- docs/exec-plans/completed/gameloop-001-playlist-progression.md

## 목표

file-backed exercise를 하나씩만 실행하던 playable을 playlist 기반 반복 학습 루프로 확장한다. 성공 후 다음 문제로 넘어가고, 완료 progress를 자동 저장하며, 기존 E2E smoke가 첫 성공 상태를 계속 확인할 수 있게 한다.

## 영향 도메인

- Gameplay: sequential unlock, retry, hint, success 후 next 흐름을 제공한다.
- Progress: 저장 포맷은 바꾸지 않고 기존 `Missions` map에 exercise ID를 key로 저장한다.
- Verification: 기존 smoke는 첫 문제 성공 상태를 유지하고, unit test에서 next/autosave/restart를 검증한다.

## 수용 기준

- playable은 playlist beat 순서로 implemented + replay-pass exercise를 실행한다.
- 현재 exercise 성공 시 해당 exercise ID로 progress가 자동 저장된다.
- 성공 상태에서 `enter`를 누르면 다음 exercise로 이동한다.
- 이미 완료된 exercise가 progress에 있으면 첫 미완료 exercise에서 시작한다.
- 마지막 exercise까지 성공하면 완료 화면을 유지하고 더 이상 다음 문제로 넘어가지 않는다.
- 기존 `make e2e-smoke`와 `make e2e-playable`은 계속 통과한다.

## 범위

- 포함: playlist 순서 기반 run selection
- 포함: success 후 `enter` next action
- 포함: 기존 progress format 재사용
- 제외: 새 progress schema
- 제외: 별도 campaign/lesson selection UI
- 제외: unlock animation 또는 보상 화면

## 검증 계획

- `go test ./internal/playable/...`
- `go test ./internal/content/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] playlist progression Red test를 작성한다.
- [x] playable model에 playlist state와 next transition을 추가한다.
- [x] progress 기반 첫 미완료 exercise 선택을 구현한다.
- [x] success/complete 화면 문구를 정리한다.
- [x] 검증을 통과시킨다.
- [x] 문서를 completed 상태로 동기화한다.
