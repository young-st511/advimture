# PLAN-REFRESH-008 — Applied Mastery Roadmap

Slice-ID: PLAN-REFRESH-008
Created: 2026-05-28
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/plan-refresh-008-applied-mastery-roadmap.md

## 목표

`char-find-line` 완료 이후 새 Vim engine을 바로 늘리지 않고, 이미 배운 command를 incident와 command-choice에서 조합하는 중기 플랜을 고정한다.

## 결정

다음 중기 플랜은 `Applied Mastery Runs`로 한다.

순서:

1. `PLAN-REFRESH-008`: 중기 플랜 문서화
2. `INCIDENT-006`: `/target` + `ct,` applied incident 구현
3. `REUSE-CHOICE-001`: 재사용 판단 드릴 설계
4. `CHOICE-PLAY-003`: reuse-choice playable 구현
5. `INCIDENT-007`: mixed recovery run 확장
6. `PLAYTEST-REVIEW-002`: 품질 점검과 다음 후보 정리

## 제외

- 새 Vim engine 기능
- progress 저장 포맷 변경
- 새 dependency 추가
- UI layout 대개편

## 수용 기준

- 중기 플랜은 `PROGRAM.md`와 `MIDTERM_TODO.md`에 반영된다.
- 첫 구현 slice는 별도 active ExecPlan으로 분리된다.
- 이후 slice들은 새 engine 없이 기존 implemented command만 사용한다.

## 검증

- [x] `git diff --check`
