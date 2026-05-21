# QA-PLATFORM-001 — E2E Suite Refresh

Slice-ID: QA-PLATFORM-001
Created: 2026-05-21
Completed: 2026-05-21
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/qa-platform-001-e2e-suite-refresh.md
- docs/exec-plans/completed/qa-platform-001-e2e-suite-refresh.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- test/e2e/
- Makefile

## 목표

중간점검 중 발견된 stale E2E 기대값을 현재 playable tutorial 흐름에 맞추고, `make e2e-playable`이 최신 playpack full E2E를 실행하도록 갱신한다.

## 범위

- 포함:
  - `playable_operator_grammar_full.yaml`, `playable_yank_put_full.yaml`, `playable_text_object_inner_word_full.yaml` 마지막 action expectation 갱신
  - `playable_open_line_full.yaml`, `playable_repeat_last_change_full.yaml` 마지막 action expectation 갱신
  - `playable_full_first_five_minute.yaml` legacy first-cut expectation 갱신
  - `make e2e-playable`에 operator/yank/text-object/open-line/repeat/search full E2E 추가
  - verification docs의 E2E suite 설명 갱신
- 제외:
  - runtime/playable 동작 변경
  - progress 저장 포맷 변경
  - 새 E2E runner assertion 추가

## 수용 기준

- completed: 중간 tutorial full E2E는 뒤의 tutorial이 있음을 반영해 `Next tutorial: enter`를 검증한다.
- completed: search full E2E는 마지막 tutorial로 `Playlist complete`를 유지한다.
- completed: `make e2e-playable`은 최신 full playpack E2E를 포함한다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_operator_grammar_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_yank_put_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_text_object_inner_word_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_open_line_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_repeat_last_change_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_search_basic_full.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_full_first_five_minute.yaml`
- passed: `make e2e-playable`
- passed: `git diff --check`

최종 Go 회귀 검증은 중간점검 문서 작성 후 함께 수행한다.
