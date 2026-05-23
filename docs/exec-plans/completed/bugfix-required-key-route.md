# BUGFIX-REQUIRED-KEY-ROUTE — Required Key Route Gate

Slice-ID: BUGFIX-REQUIRED-KEY-ROUTE
Created: 2026-05-23
Completed: 2026-05-23
Status: completed
Scope-Mode: bugfix
Allowed-Paths:
- docs/exec-plans/active/bugfix-required-key-route.md
- docs/exec-plans/completed/bugfix-required-key-route.md
- docs/roadmap/CHANGES.md
- docs/gameplay/spec.md
- internal/runtime/session.go
- internal/runtime/session_test.go
- internal/content/loader_test.go
- internal/playable/model_test.go
- content/exercises/first-five-minutes.yaml
- content/scenarios/first-five-minutes.yaml
- test/e2e/playable_constraint_required_key.yaml
- test/e2e/playable_small_edits_full.yaml
- test/e2e/playable_full_first_five_minute.yaml
- test/e2e/playable_word_motion_required_route.yaml

## 버그

- `ctrl+r`를 배우는 tutorial에서 `ctrl+r`를 눌러도 바로 clear되지 않고 숨은 `h` 입력까지 요구했다.
- `w`, `e`를 배우는 word-motion stage에서 `l`/`h` 반복으로 목표 위치에 도달해도 clear될 수 있었다.

## 원인

- `undo-redo-basic-002`는 redo 학습이 목표인데 early goal failure를 피하려고 target cursor와 optimal trace에 `h`를 섞어 두었다.
- `word-motion-basic-*` exercise에는 `constraints.required_keys`와 strict route 제약이 없어 cursor goal만 만족하면 성공할 수 있었다.
- runtime은 목표에 닿는 순간 required key가 없으면 즉시 실패시키므로, redo처럼 중간에 목표 상태와 같은 상태를 지나가는 문항을 자연스럽게 만들기 어려웠다.

## 완료 내용

- runtime이 required key 없이 목표에 잠깐 도달하더라도 입력 여지가 남고 이후 목표를 벗어날 수 있으면 계속 진행하도록 했다.
- 목표 상태에 머문 채 required key를 나중에 덧붙이는 우회는 `required_keys_missing`으로 실패하도록 막았다.
- `undo-redo-basic-002`를 `x`, `u`, `ctrl+r` 3입력으로 clear되도록 content와 E2E를 갱신했다.
- `word-motion-basic-001/002/003`에 각각 `w/b/e` required key와 h/l forbidden route를 추가했다.
- `playable_word_motion_required_route.yaml` E2E를 추가해 `l` 우회 실패 후 `w,w` 성공을 검증했다.

## 검증 결과

- `go test ./internal/runtime/... ./internal/content/... ./internal/playable/...`
- `go test ./...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_small_edits_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_word_motion_required_route.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_full_first_five_minute.yaml`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`
