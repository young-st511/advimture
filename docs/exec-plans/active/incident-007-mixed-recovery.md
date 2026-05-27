# INCIDENT-007 — Mixed Recovery Run

Slice-ID: INCIDENT-007
Created: 2026-05-28
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/exec-plans/active/incident-007-mixed-recovery.md
- content/exercises/
- content/scenarios/
- content/playlists/
- test/e2e/
- internal/content/loader_test.go
- Makefile

## 목표

이미 배운 search, text object, substitute, visual-line, char-find를 하나의 incident run에서 조합한다. 한 beat에 너무 많은 command를 몰지 않고, 각 beat가 “상황에 맞는 도구 선택”을 하나씩 요구하도록 만든다.

## 포함

- 4~5 beat mixed incident
- focused E2E
- loader count/playable list update
- `make e2e-playable` 연결

## 제외

- 새 Vim engine 기능
- 새 schema
- progress 저장 포맷 변경
- visual block/count/register 확장

## 수용 기준

- 각 beat는 이미 implemented command만 사용한다.
- 최소 beat 구성은 `/`, `ci"`, `V...d`, `ct,`, `:%s` 중 4개 이상을 포함한다.
- content replay, focused E2E, `go test ./...`, `make e2e-playable`, `git diff --check`를 통과한다.

## Step 1: Content

- [ ] exercise set 작성
- [ ] scenario copy 작성
- [ ] playlist wiring

## Step 2: E2E

- [ ] focused E2E 작성
- [ ] Makefile suite 연결
- [ ] loader tests 갱신

## Step 3: Verification

- [ ] content tests
- [ ] focused E2E
- [ ] full regression
