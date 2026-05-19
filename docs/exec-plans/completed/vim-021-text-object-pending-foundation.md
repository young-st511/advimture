# VIM-021 — Text Object Pending Foundation

Slice-ID: VIM-021
Created: 2026-05-19
Completed: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/exec-plans/active/vim-021-text-object-pending-foundation.md
- docs/exec-plans/completed/vim-021-text-object-pending-foundation.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- internal/vimengine/
- internal/tuiadapter/

## 목표

`diw`, `ciw`, `yiw` 구현 전 단계로 `operator -> i -> w` 3-key sequence를 엔진이 안정적으로 받을 수 있게 한다.

## 결과

- `d/c/y` pending 뒤 `i` 입력 시 `di`, `ci`, `yi` inner text object pending 상태로 전환한다.
- unsupported text object sequence는 buffer/register를 변경하지 않고 unsupported event로 닫는다.
- `innerWordRange` helper는 keyword/symbol run을 현재 text object 범위로 계산한다.
- 공백과 빈 줄은 selection failure로 처리한다.

## 제외

- `diw`, `ciw`, `yiw` 실제 mutation/register semantics는 VIM-022로 넘긴다.
- quote/pair text object, around object, count prefix는 제외했다.

## 검증

- `go test ./internal/vimengine/...`
- `go test ./internal/tuiadapter/...`
