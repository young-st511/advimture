# REUSE-CHOICE-001 — Reuse Judgment Drill Design

Slice-ID: REUSE-CHOICE-001
Created: 2026-05-28
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/reuse-choice-001-reuse-judgment.md

## 목표

직접 다시 입력하지 않고 이미 있는 텍스트를 재사용해야 하는 상황을 command-choice drill로 설계한다. 후보는 `yi" + P`, `yy/p`, `.`, retype이며, 첫 playable 구현 후보를 하나로 고른다.

## 포함

- reuse-choice 후보 2~3개
- intended command와 부적합한 대안의 이유
- forbidden shortcut 기준
- 다음 구현 slice 분리

## 제외

- content YAML 구현
- 새 Vim engine 기능
- 새 schema
- progress 저장 포맷 변경

## 수용 기준

- 첫 playable 후보는 이미 implemented command만 사용한다.
- 성공/실패 copy는 “왜 재사용해야 하는가”를 설명한다.
- 다음 구현 slice `CHOICE-PLAY-003`의 include/exclude/검증 계획이 정의된다.

## Step 1: Candidate Design

- [ ] reuse-choice 후보 작성
- [ ] 우회/금지 입력 기준 작성

## Step 2: Selection

- [ ] 구현 리스크 비교
- [ ] 첫 playable 후보 결정

## Step 3: Next Slice

- [ ] CHOICE-PLAY-003 구현 범위 정의
- [ ] 검증 계획 정의
