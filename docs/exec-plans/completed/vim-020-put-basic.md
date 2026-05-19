# VIM-020 — Put Basic

Slice-ID: VIM-020
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-020-put-basic.md
- docs/exec-plans/completed/vim-020-put-basic.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/ENGINE_TODO.md
- internal/vimengine/
- internal/tuiadapter/
- internal/runtime/

## 목표

VIM-019의 unnamed register를 실제 편집에 사용할 수 있도록 `p`, `P`를 구현한다. charwise register는 현재 줄의 cursor 앞/뒤에 붙이고, linewise register는 현재 줄 위/아래에 붙인다.

## 범위

- 포함:
  - `KeyP`, `KeyShiftP`
  - charwise `p`: cursor 뒤에 register text 삽입
  - charwise `P`: cursor 앞에 register text 삽입
  - linewise `p`: 현재 줄 아래에 register lines 삽입
  - linewise `P`: 현재 줄 위에 register lines 삽입
  - put mutation의 undo/redo
  - runtime replay smoke
  - `tuiadapter` normal mode `p`, `P` mapping
- 제외:
  - named register
  - count prefix
  - multi-register history
  - text object
  - YAML content/E2E 연결

## 구현 결과

- `KeyP`, `KeyShiftP`를 추가했다.
- `putRegister`, `putCharwiseRegister`, `putLinewiseRegister` helper를 추가했다.
- charwise put은 삽입된 text의 마지막 글자로 cursor를 이동한다.
- linewise put은 첫 삽입 줄의 첫 column으로 cursor를 이동한다.
- put mutation은 기존 undo/redo stack을 사용한다.
- runtime session smoke에서 `yy`, `p` trace로 linewise put 성공을 검증했다.

## 검증 결과

- `go test ./internal/vimengine/...`: pass
- `go test ./internal/tuiadapter/...`: pass
- `go test ./internal/runtime/...`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## E2E Evidence

콘텐츠 연결 전 엔진 루프이므로 E2E는 추가하지 않았다. `PLAYPACK-004`에서 `y/p` content와 함께 추가한다.

## 의사결정 로그

- 2026-05-19: empty register put은 boundary event로 처리한다.
- 2026-05-19: 첫 linewise put은 Vim의 세부 register metadata보다 학습용 행 삽입 semantics를 우선한다.
