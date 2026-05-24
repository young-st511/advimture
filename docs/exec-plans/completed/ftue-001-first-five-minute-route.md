# FTUE-001 — First 5-Minute Route

Slice-ID: FTUE-001
Created: 2026-05-25
Status: completed
Completed: 2026-05-25
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/first-five-minute-route.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- Makefile
- test/e2e/
- artifacts/e2e/

## 목표

처음 실행한 플레이어가 5분 안에 Advimture의 핵심 효용을 느끼는 canonical route를 정의한다. 첫 루프는 새 Vim command를 추가하지 않고, 현재 구현된 tutorial과 UI-HUD-003 화면을 기준으로 `이동 → 생존 → 빠른 이동 → 작은 수정` 흐름을 검증한다.

## 범위

- 포함:
  - 현재 playable 첫 실행 순서 점검
  - Tutorial 0~2와 작은 수정 초입의 pacing 검토
  - 첫 5분 route 후보 정의
  - 대표 E2E fixture 또는 기존 fixture evidence 재사용 계획
  - 긴 안내/불필요한 보조 정보/초반 실패 피드백의 UX 점검
  - `docs/gameplay/spec.md`의 첫 5분 미확인 항목 해소
- 제외:
  - 새 Vim engine command 추가
  - progress 저장 포맷 변경
  - command-choice playable 승격
  - split cockpit wide layout
  - daily route 저장 구조 변경

## 수용 기준

- 첫 5분 route가 tutorial/playlist ID 기준으로 명시된다.
- route가 플레이어에게 주는 학습 감각이 `이동`, `생존`, `빠른 이동`, `작은 수정` 중 무엇인지 설명된다.
- route의 대표 성공/실패/완료 화면 evidence를 확인할 수 있다.
- 첫 5분 route를 검증하는 E2E 전략이 명시된다.
- 너무 이른 정보 노출, 긴 briefing, modal overflow, retry/next 흐름 문제를 점검한다.
- `docs/gameplay/spec.md`의 `첫 5분 플레이 루프` 미확인 항목이 결정 상태로 갱신된다.

## Step 1: Current Route Audit

- 목표: 현재 첫 실행 흐름과 evidence를 확인한다.
- 상세 작업:
  - [x] current playable playlist order 확인
  - [x] Tutorial 0~2의 총 문항 수, 학습 목표, 예상 플레이 시간을 점검
  - [x] 기존 E2E fixture 중 첫 루프 검증에 쓸 수 있는 것 표시

## Step 2: First 5-Minute Contract

- 목표: 첫 5분 route를 제품 계약으로 고정한다.
- 상세 작업:
  - [x] route include/exclude 결정
  - [x] 플레이어 감정 곡선과 학습 목표 작성
  - [x] `docs/gameplay/spec.md` 미확인 항목 갱신

## Step 3: Evidence and Follow-up

- 목표: FTUE가 이후 UI/playable/content 작업의 기준이 되게 한다.
- 상세 작업:
  - [x] E2E evidence 경로 기록
  - [x] 필요 시 새 focused E2E fixture 초안 작성
  - [x] UI-PLAYTEST-001 입력 목록 작성
  - [x] 검증: `make e2e-playable` 또는 focused E2E, `git diff --check`

## 결정

- 첫 5분 canonical route는 `tutorial-0-movement` 전체, `tutorial-1-survival` 전체, `tutorial-2-fast-navigation` 전체, `tutorial-3-small-edits` 첫 문항까지다.
- 기존 `playable_full_first_five_minute.yaml`은 Ex command 고급 튜토리얼까지 포함하므로 regression fixture로 유지하되, FTUE 대표 evidence는 `playable_ftue_first_five_route.yaml`로 분리한다.
- 첫 5분 루프는 새 command/schema/progress 변경 없이 현재 playable 순서를 제품 계약으로 고정한다.

## 검증 결과

- `go run ./cmd/e2e-runner --scenario test/e2e/playable_ftue_first_five_route.yaml`
- `git diff --check`

## 실행 규칙

문서 계약을 먼저 고정하고, 새 content/schema/progress 변경은 하지 않는다. 콘텐츠 수정이 필요하다고 판단되면 별도 ExecPlan으로 분리한다.
