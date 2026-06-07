# NEXT-PLAYTEST-REVIEW-001 — Fresh Direction Review

Slice-ID: NEXT-PLAYTEST-REVIEW-001
Created: 2026-06-07
Status: completed
Scope-Mode: review
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/roadmap/NEXT_PLAYTEST_REVIEW_2026-06-07.md
- docs/exec-plans/active/next-playtest-review-001-fresh-direction.md
- docs/exec-plans/completed/next-playtest-review-001-fresh-direction.md

## 목표

`SEARCH-THEN-SCOPE-APPLIED-001`과 `BRACKET-PAIR-HARDEN-001` 완료 후 현재 playable loop evidence를 다시 읽고, 다음 보강 축을 하나로 좁힌다.

## 후보

- Option A: command-choice breadth
- Option B: tutorial pacing / mid-tour compression
- Option C: review-loop motivation polish
- Option D: deeper pair hardening
- Option E: release candidate prep

## 선택 기준

- 바로 출시하거나 tag를 묶지 않는다.
- 새 dependency, progress 저장 포맷, content schema 변경 없이 열 수 있는가?
- 이미 구현한 Vim command를 "상황에 맞게 고르는 판단"으로 더 선명하게 만드는가?
- 최근 E2E evidence가 실제 병목을 보여주는가?
- 좁은 ExecPlan으로 열고 focused E2E로 닫을 수 있는가?

## 수용 기준

- first tour, mid tutorial, bracket pair tutorial, first dispatch, judgment drill, incident 008, review loop evidence를 spot review한다.
- P0/P1 blocker 유무를 명시한다.
- 다음 slice를 하나 추천하고, 이번에 열지 않을 후보도 이유와 함께 남긴다.
- `docs/roadmap/NEXT_PLAYTEST_REVIEW_2026-06-07.md`에 evidence 표와 후보 비교를 남긴다.
- `PROGRAM.md`, `MIDTERM_TODO.md`, `FORWARD_PLAN.md`, `CHANGES.md`가 같은 다음 방향을 가리킨다.
- `git diff --check`와 `go test ./...`를 통과한다.
- `go.mod`, `go.sum`, `internal/progress/`, content schema 변경이 없다.

## Step 1: Evidence Inventory

- [x] 대표 E2E artifact를 확정한다.
- [x] 각 artifact의 final screen, timeline, app_state를 확인한다.

## Step 2: Option Review

- [x] 후보를 선택 기준으로 비교한다.
- [x] 이번 iteration의 최선 후보를 하나 고른다.

## Step 3: Review Artifact

- [x] fresh review 문서를 작성한다.
- [x] 다음 slice의 완료 기준을 명확히 제안한다.

## Step 4: Closeout

- [x] roadmap 문서를 최신 상태로 맞춘다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동한다.
