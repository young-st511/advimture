# INCIDENT-006 — Inline Target Repair Run

Slice-ID: INCIDENT-006
Created: 2026-05-28
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/exec-plans/active/incident-006-inline-target-repair.md
- content/exercises/
- content/scenarios/
- content/playlists/
- test/e2e/
- internal/content/loader_test.go
- Makefile

## 목표

`/target`으로 손상된 줄을 찾고, `ct,`로 comma 앞 값만 교체하는 applied incident를 구현한다. 새 command 소개가 아니라 `search-basic`과 `char-find-line`을 실제 복구 상황에서 조합하게 만든다.

## 포함

- `incident-inline-target-*` exercise set
- matching scenario set
- `incident-006-inline-target-repair` playlist
- focused E2E
- loader count/playable list update
- `make e2e-playable` 연결

## 제외

- 새 Vim engine 기능
- progress 저장 포맷 변경
- `F/T`, `;`, `,`, count prefix
- 새 command-choice schema

## 수용 기준

- 최소 한 문항은 `/target` + `ct,`를 같은 optimal trace에서 사용한다.
- `ct,`가 delimiter를 보존해야 하는 이유가 scenario success/failure copy에 드러난다.
- content replay, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Content

- [x] exercise set 작성
- [x] scenario copy 작성
- [x] playlist wiring

## Step 2: E2E

- [x] focused E2E 작성
- [x] Makefile suite 연결
- [x] loader tests 갱신

## Step 3: Verification

- [x] content tests
- [x] focused E2E
- [x] full regression
