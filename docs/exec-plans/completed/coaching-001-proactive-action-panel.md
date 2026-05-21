# COACHING-001 — Proactive Action Panel Coaching

Slice-ID: COACHING-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/coaching-001-proactive-action-panel.md
- docs/exec-plans/completed/coaching-001-proactive-action-panel.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- internal/runtime/
- internal/scenario/
- internal/playable/
- test/e2e/
- Makefile

## 목표

저장 포맷과 content schema를 바꾸지 않고, strict constraint가 있는 문항에서 플레이어가 실패 전에 훈련 의도를 볼 수 있게 한다.

## 범위

- 포함:
  - running/failed `ACTION` 패널에 아직 쓰지 않은 required key coaching 표시
  - `?` hint 요청 결과를 `ACTION` 패널에 표시
  - focused model test와 E2E scenario
  - `make e2e-playable`에 focused coaching E2E 추가
- 제외:
  - Practice/Challenge schema 분리
  - progress 저장 포맷 변경
  - content YAML hint schema 변경
  - 새 command나 Vim engine 기능

## 수용 기준

- completed: 첫 strict tutorial 문항 시작 화면은 `ACTION` 패널 안에 훈련 키를 표시한다.
- completed: 목표에는 도착했지만 required key가 빠져 실패한 화면도 아직 쓰지 않은 훈련 키를 표시한다.
- completed: `?`를 눌러 얻은 hint는 화면에 보이며, hint 요청 전까지는 grade/progress를 바꾸지 않는다.
- completed: command/search mode action panel에는 일반 hint/quit 안내를 섞지 않는다.
- completed: 저장 포맷과 content schema는 변경하지 않는다.

## 검증 결과

- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/runtime/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./internal/playable/...`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml`
- passed: `GOCACHE=/Users/young/github.com/young-st511/advimture/artifacts/go-build-cache go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
