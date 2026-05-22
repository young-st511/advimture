# PROGRESS-LANGUAGE-001 — Recovery Language Pass

Slice-ID: PROGRESS-LANGUAGE-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: ui-copy-and-e2e
Allowed-Paths:
- docs/exec-plans/active/progress-language-001-recovery-language-pass.md
- docs/exec-plans/completed/progress-language-001-recovery-language-pass.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/playable/model.go
- internal/playable/model_test.go
- internal/review/review.go
- internal/review/review_test.go
- test/e2e/playable_debrief_success.yaml
- test/e2e/playable_review_queue.yaml
- test/e2e/playable_coaching_panel.yaml

## 목표

저장 포맷 변경 없이 review queue, success debrief, best record, playlist completion 문구를 `원격 시설 복구국 / Runbook Dispatch` 프레임에 맞춘다.

## 완료 내용

- review queue prefix를 `재점검 대상`으로 변경했다.
- incomplete candidate를 `미복구`로 표시한다.
- low grade candidate를 `복구 등급 <grade>`로 표시한다.
- key count candidate를 `복구 입력 <best>/<optimal> keys`로 표시한다.
- success debrief를 `복구 기록: grade <grade>, <n> keys`로 표시한다.
- best record를 `최단 복구 기록: grade <grade>, <n> keys`로 표시한다.
- playlist completion을 `Runbook: <completed>/<total> 복구 완료`로 표시한다.
- residual recommendation은 `잔류 리스크`로 유지했다.

## 제외한 것

- progress schema 변경
- review/daily/mastery 저장 구조 추가
- 새로운 replay menu
- scoring 계산 변경
- content target/optimal key 변경

## 검증 결과

- passed: `go test ./internal/review/...`
- passed: `go test ./internal/playable/...`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_debrief_success.yaml`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_review_queue.yaml`
- passed: `go run ./cmd/e2e-runner --scenario test/e2e/playable_coaching_panel.yaml`
