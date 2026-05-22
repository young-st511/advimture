# COMMAND-CHOICE-001 — Mixed Command Choice Drill Design

Slice-ID: COMMAND-CHOICE-001
Created: 2026-05-23
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/command-choice-001-mixed-drill-design.md
- docs/gameplay/command-choice-drills.md
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md

## 목표

이미 배운 Vim command를 섞어 두고 플레이어가 적절한 도구를 고르는 mixed drill 설계 기준을 고정한다.

## 수용 기준

- 새 Vim engine 기능을 추가하지 않는다.
- 새 content schema 필드를 추가하지 않는다.
- command choice drill을 command cluster가 아니라 적용 레이어로 정의한다.
- 최소 3개 이상의 drill type을 정의한다.
- 첫 playable 후보와 승격 gate를 문서화한다.
- 정답 key 암기가 아니라 선택 이유를 강화하는 기준을 둔다.

## 제외 항목

- playable YAML 추가
- loader schema 변경
- progress schema 변경
- 새 command cluster 추가
- review/daily UI 변경

## 완료 내용

- `docs/gameplay/command-choice-drills.md`에 command choice drill 목적, 원칙, drill type, authoring rubric, 첫 후보, playable gate를 추가했다.
- 첫 후보는 `scope-choice`, `reuse-choice`, `search-then-act`, `range-choice`로 나눴다.
- 실제 playable 승격은 별도 ExecPlan에서 처리하도록 분리했다.

## 검증 계획

- `git diff --check`

## 검증 결과

- passed: `git diff --check`
