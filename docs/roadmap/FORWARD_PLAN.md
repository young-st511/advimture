# Forward Plan — Foundation to First Release

Last reviewed: 2026-05-30
Status: current rolling plan

## 목적

이 문서는 앞으로 2~8주 동안 어떤 순서로 Advimture를 출시 가능한 Vim 학습 게임으로 다듬을지 정리한다.

- `PROGRAM.md`: 지금 무엇이 active인지 확인한다.
- `MIDTERM_TODO.md`: 현재 중기 보드를 확인한다.
- `FORWARD_PLAN.md`: 왜 그 순서로 가는지, 다음 몇 주의 방향을 확인한다.

작업 시작 전에는 `PROGRAM.md -> MIDTERM_TODO.md -> FORWARD_PLAN.md` 순서로 읽는다.

## 현재 판단

Foundation engine과 E2E loop는 충분히 튼튼해졌다. 다음 병목은 새 Vim command 수가 아니라 출시 가능한 게임 루프다.

현재 상태:

- Vim engine: tutorial/incident를 만들 수 있을 만큼 충분히 닫힘
- Content: tutorial coverage와 incident 001~007이 있음
- E2E: long route final/timeline evidence까지 보강됨
- UI/UX: Mission HUD, Runbook Console, floating modal 기반은 있음
- 출시감: mission/review loop, 반복 동기, UI polish, release readiness가 아직 부족함

따라서 다음 순서는 **Foundation exit review -> game loop/platform polish -> content breadth 또는 small engine hardening -> release readiness**로 간다.

## 0. 운영 원칙

- 새 기능보다 먼저 현재 evidence를 본다.
- E2E가 부족하다고 느껴지면 구현을 멈추고 verification을 보강한다.
- 새 engine은 하나의 command contract만 다룬다.
- progress 저장 포맷은 사용자 승인 전까지 변경하지 않는다.
- long route E2E는 `screen_timeline.txt`와 `screen_final.txt`를 남긴다.
- 문서가 stale해질 수 있는 변경을 하면 `PROGRAM.md`, `MIDTERM_TODO.md`, 이 문서를 함께 확인한다.

## 1. Immediate Plan

### FOUNDATION-EXIT-001 — Foundation Exit Review

목표: 지금의 foundation이 첫 공개 가능한 학습 게임 루프로 충분한지 판정한다.

산출물:

- 첫 출시 후보 범위
- 반드시 고쳐야 할 P0/P1
- 다음 중기 플랜 확정
- release gate 초안

검증:

- 최신 `make e2e-playable` 결과 또는 최신 evidence 확인
- long incident `screen_final.txt` spot review
- UI/UX, learning design, architecture 관점 review

Decision gate:

- game loop 부족이 크면 `PLATFORM-REVIEW-003`으로 간다.
- content 양이 부족하면 `CONTENT-BREADTH-002`로 간다.
- engine coverage가 병목이면 `QUOTE-PAIR-HARDEN-001`로 간다.

## 2. Recommended Midterm Sequence

### 1. PLATFORM-REVIEW-003 — Mission/Review Game Loop

권장 우선순위: 높음

목표: 저장 포맷 변경 없이 review queue, mission progress, daily route를 더 게임답게 묶는다.

포함:

- 첫 화면/완료 화면에서 다음 행동이 더 선명하게 보이게 함
- `재점검`, `잔류 리스크`, `오늘의 복구 루트`를 mission loop로 정리
- 실패/낮은 grade/높은 key count를 다음 플레이 이유로 연결

제외:

- progress schema v2
- daily streak 저장
- spaced review due date 저장

### 2. CONTENT-BREADTH-002 — Applied Content Expansion

권장 우선순위: 중간

목표: 새 engine 없이 기존 command를 조합하는 applied incident와 command-choice를 늘린다.

후보:

- line reuse choice: 검증된 줄 전체를 `V` + `y` + `p`로 재사용
- repeat-change choice: 같은 변경을 `.`로 반복할지 판단
- search-then-act incident: `/`, `n`, `N`으로 찾고 적절한 edit command 선택
- mixed incident 008: 3~5 beat 이하로 제한한 생존 어드벤처 run

품질 기준:

- 한 beat는 하나의 판단만 요구한다.
- 새 command를 소개하지 않는다.
- long route에는 final/timeline evidence를 남긴다.

### 3. QUOTE-PAIR-HARDEN-001 — Quote/Pair Text Object Hardening

권장 우선순위: 중간

목표: 기존 `i"` text object를 `i'`, `i(`, `i{`로 확장한다.

포함:

- `ci'`, `di'`, `yi'`
- `ci(` 또는 `ci{` 중 작은 첫 scope
- config/JSON/function-argument style exercise

제외:

- nested pair
- escaped quote
- around object
- count/register prefix

### 4. UI-POLISH-002 — Release UI Polish

권장 우선순위: Foundation exit 결과에 따라 결정

목표: 출시 전 화면을 개발 UI가 아니라 Vim adventure console처럼 읽히게 다듬는다.

후보:

- color/emphasis pass
- learned command memory
- wide layout side rail
- pre-start briefing modal

주의:

- 화면 문구보다 `app_state` 검증을 우선한다.
- color 없는 환경에서도 의미가 보존되어야 한다.

## 3. Release Readiness

첫 공개 전 필요 항목:

- `README.md`에 설치/실행/테스트 안내
- 첫 실행 경험 검증
- 터미널 크기별 smoke
- progress 파일 안전성 점검
- release build command
- known limitations 정리

첫 공개 기준:

- tutorial route가 막힘 없이 진행된다.
- incident route가 3개 이상 게임처럼 읽힌다.
- 실패/재시도/힌트가 플레이를 막지 않는다.
- `make e2e-playable`이 통과한다.
- long incident evidence가 남는다.
- progress schema 변경 없이 저장/재개가 안전하다.

## 4. Long-Run Candidates

아래는 출시 전 필수가 아니다.

- progress schema v2
- spaced review due date
- daily streak/history
- macros/register/count prefix
- visual block
- regex search/highlight/history
- terminal cell-grid viewport parser

이 후보들은 실제 플레이 evidence로 병목이 확인될 때만 연다.

## 5. 문서 업데이트 규칙

각 slice 종료 시:

1. `PROGRAM.md`: active/recent completed/next 후보 갱신
2. `MIDTERM_TODO.md`: 현재 중기 보드 상태 갱신
3. `FORWARD_PLAN.md`: 추천 순서나 gate가 바뀌었으면 `Last reviewed`와 관련 섹션 갱신
4. `CHANGES.md`: 가정 변경은 append-only로 기록
5. 오래된 review/health 문서는 `docs/roadmap/archive/`로 이동

이 규칙을 지키지 못하면 다음 작업 전에 docs cleanup slice를 먼저 연다.
