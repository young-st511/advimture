# VIM-028 — Charwise Visual Delete/Yank

Slice-ID: VIM-028
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: engine-runtime
Allowed-Paths:
- docs/exec-plans/active/vim-028-charwise-visual-operators.md
- docs/exec-plans/completed/vim-028-charwise-visual-operators.md
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- internal/vimengine/engine.go
- internal/vimengine/engine_test.go
- internal/runtime/session_test.go

## 목표

같은 줄 charwise visual selection에 `d`와 `y`를 적용한다.

## 완료 내용

- visual mode `d`가 같은 줄 inclusive selection range를 삭제하고 normal mode로 돌아간다.
- visual mode `d`가 삭제한 텍스트를 unnamed charwise register에 저장한다.
- visual mode `y`가 같은 줄 inclusive selection range를 unnamed charwise register에 저장하고 normal mode로 돌아간다.
- 성공한 `d/y`는 selection을 clear한다.
- 성공한 `d/y` 이후 cursor는 normalized selection start에 둔다.
- multi-line charwise selection의 `d/y`는 unsupported event로 유지한다.
- runtime에서 visual delete trace가 goal과 required key를 만족하는지 검증했다.

## 제외한 것

- linewise `V`
- visual block `<C-v>`
- multi-line charwise delete/yank
- `c`, `>`, `<`, `p`, `P` visual mode operator
- count/register prefix
- dot repeat 연계
- visual tutorial content

## 검증 결과

- passed: `go test ./internal/vimengine/...`
- passed: `go test ./internal/runtime/...`
- passed: `go test ./...`
- passed: `make e2e-playable`
- passed: `git diff --check`
