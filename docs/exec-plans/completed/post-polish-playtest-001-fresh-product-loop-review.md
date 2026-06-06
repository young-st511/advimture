# POST-POLISH-PLAYTEST-001 — Fresh Product Loop Review

Slice-ID: POST-POLISH-PLAYTEST-001
Created: 2026-06-06
Status: completed
Scope-Mode: review
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/FORWARD_PLAN.md
- docs/roadmap/CHANGES.md
- docs/roadmap/CONTENT_EVIDENCE_BUNDLE_001.md
- docs/roadmap/POST_POLISH_PLAYTEST_2026-06-06.md
- docs/exec-plans/active/post-polish-playtest-001-fresh-product-loop-review.md
- docs/exec-plans/completed/post-polish-playtest-001-fresh-product-loop-review.md

## 목표

`CONTENT-ARC-POLISH-001`, `JUDGMENT-DRILL-REVIEW-001`, `UI-CONSOLE-POLISH-001` 완료 후 현재 playable loop evidence를 fresh review하고, 다음 병목을 하나로 좁힌다.

## 후보

- Option A: line reuse applied drill
- Option B: search-then-scope applied drill
- Option C: bracket-pair hardening
- Option D: 추가 UI console polish

## 선택 기준

- 새 engine, 새 dependency, progress 저장 포맷 변경 없이 열 수 있는가?
- command-choice의 핵심 정체성인 "상황에 맞는 Vim 도구 선택"을 더 선명하게 만드는가?
- 기존 evidence와 E2E가 품질 저하 없이 좁게 확장될 수 있는가?

## 수용 기준

- first tour, first dispatch, judgment drill, review loop evidence가 현재 제품 방향을 같은 언어로 설명한다.
- P0/P1 blocker가 없고, 다음 slice가 release candidate가 아니라 좁은 품질 보강으로 판정된다.
- 후보 비교 결과와 선택 이유가 `docs/roadmap/POST_POLISH_PLAYTEST_2026-06-06.md`에 남는다.
- 선택한 후속 slice의 scope가 progress schema, content schema, dependency, 새 Vim engine을 건드리지 않는다.
- `git diff --check`, focused evidence E2E, `go test ./...`가 후속 slice 검증에 포함된다.

## Step 1: Evidence Inventory

- [x] 기존 evidence bundle과 roadmap 상태를 읽는다.
- [x] first tour, first dispatch, judgment drill, review loop의 현재 대표 E2E를 확정한다.

## Step 2: Option Review

- [x] line reuse, search-then-scope, bracket-pair hardening, UI polish를 비교한다.
- [x] 이번 iteration의 최선 후보를 하나 고른다.

## Step 3: Review Artifact

- [x] fresh review 문서를 작성한다.
- [x] 후속 slice의 만족 조건을 명확히 연결한다.

## Step 4: Closeout

- [x] roadmap 문서를 최신 상태로 맞춘다.
- [x] ExecPlan을 completed로 이동한다.
