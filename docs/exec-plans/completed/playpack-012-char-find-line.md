# PLAYPACK-012 — Char Find Line Tutorial

Slice-ID: PLAYPACK-012
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/exec-plans/active/playpack-012-char-find-line.md
- docs/exec-plans/completed/vim-030-char-find-engine.md
- content/
- test/e2e/
- internal/content/

## 목표

VIM-030에서 구현한 `char-find-line`을 4~6문항 tutorial playpack으로 연결한다.

## 포함

- `f=`, `t,`, `df,`, `ct"` 중심 tutorial content
- 필요한 경우 `cf` 또는 `dt` 보강 문항
- replay gate와 coverage gate
- full playlist focused E2E

## 제외

- 새 engine 기능
- `F/T`, `;`, `,`, count prefix
- progress 저장 포맷 변경
- applied incident 확장

## 수용 기준

- approved + implemented tutorial playlist는 8문항 이하로 유지된다.
- `f`, `t`, `df`, `dt`, `cf`, `ct` coverage가 optimal trace에 등장한다.
- content replay와 `playable_char_find_full.yaml` E2E가 통과한다.
- `make e2e-playable`을 통과한다.

## Step 1: Content Design

- [x] exercise set 설계
- [x] scenario copy 작성
- [x] playlist wiring

## Step 2: Tests and E2E

- [x] content loader/replay tests 갱신
- [x] focused E2E 추가
- [x] Makefile suite 연결

## Step 3: Verification

- [x] content tests
- [x] focused E2E
- [x] `go test ./...`
- [x] `make e2e-playable`
- [x] `git diff --check`
