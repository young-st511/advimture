# UX-BACKLOG-001 — Foundation UX Backlog

Slice-ID: UX-BACKLOG-001
Created: 2026-05-26
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/tui-ux-direction.md
- docs/verification/tui-ui-qa-contract.md
- docs/exec-plans/active/ux-backlog-001-foundation-backlog.md
- docs/roadmap/UX_BACKLOG_001.md

## 목표

Foundation playtest/polish에서 발견한 UI/UX 개선 후보를 다음 content breadth 작업 전에 우선순위와 검증 방식까지 정리한다.

## 범위

- 포함:
  - 완료된 HUD/help/choice/success 개선의 남은 리스크 정리
  - 다음 UI/UX 후보의 우선순위, trigger, 검증 방식 정의
  - content breadth로 넘길 준비 조건 정리
- 제외:
  - 코드 변경
  - 새 E2E runner 기능 추가
  - 저장 포맷 변경

## 수용 기준

- backlog는 P0/P1/P2로 나뉘며 각 항목은 “왜 중요한가 / 언제 열 것인가 / 어떻게 검증할 것인가”를 가진다.
- content breadth 이전에 당장 막는 P0가 있는지 판단한다.
- `docs/gameplay/tui-ux-direction.md` 또는 `docs/verification/tui-ui-qa-contract.md`에 반복 가능한 원칙이 승격된다.
- `git diff --check`를 통과한다.

## Step 1: Evidence Review

- [x] 최근 UX slice 결과 확인
- [x] 남은 UX 리스크 분류

## Step 2: Backlog Draft

- [x] P0/P1/P2 backlog 작성
- [x] QA 계약/UX 방향 문서에 반복 원칙 승격

## Step 3: Verification

- [x] docs review
- [x] `git diff --check`

## 결과

- `docs/roadmap/UX_BACKLOG_001.md`에 P0/P1/P2 backlog와 content breadth readiness를 정리했다.
- 현재 content breadth를 막는 P0 UX blocker는 없다고 판단했다.
- 반복 가능한 원칙은 TUI UX direction과 TUI UI QA contract에 승격했다.
