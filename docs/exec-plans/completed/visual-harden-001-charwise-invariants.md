# VISUAL-HARDEN-001 — Charwise Visual Invariants

Slice-ID: VISUAL-HARDEN-001
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: test-hardening
Allowed-Paths:
- docs/exec-plans/active/visual-harden-001-charwise-invariants.md
- docs/exec-plans/completed/visual-harden-001-charwise-invariants.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/vimengine/engine_test.go
- internal/runtime/session_test.go

## 목표

새 visual 기능을 추가하지 않고, 현재 같은 줄 charwise visual selection의 edge/invariant를 테스트로 고정한다.

## 수용 기준

- 줄 끝 boundary motion 후 selection이 깨지지 않는지 테스트한다.
- 빈 줄에서 visual operator가 buffer를 변경하지 않고 visual state를 보존하는지 테스트한다.
- visual `y`가 undo stack을 만들지 않는지 테스트한다.
- visual `d` 후 undo가 buffer를 복구하고 register 학습 동작을 보존하는지 테스트한다.
- multi-line visual operator unsupported 동작을 실제 visual motion trace로 테스트한다.
- visual state normalization/clamp가 out-of-range selection을 안전하게 정규화하는지 테스트한다.

## 결과

- engine 테스트로 boundary, empty line, visual yank undo behavior, visual delete undo/register behavior, multi-line unsupported, selection clamp를 고정했다.
- runtime 테스트로 unsupported multi-line visual operator가 session success로 오판되지 않음을 고정했다.

## 제외한 것

- behavior 변경
- linewise `V`, visual block, multi-line operator 구현
- content/E2E 추가
- progress 저장 포맷 변경

## 검증

- `go test ./internal/vimengine/...`
- `go test ./internal/runtime/...`
- `go test ./...`
- `git diff --check`
