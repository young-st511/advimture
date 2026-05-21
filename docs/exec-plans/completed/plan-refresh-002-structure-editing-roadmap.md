# PLAN-REFRESH-002 — Structure Editing Roadmap

Slice-ID: PLAN-REFRESH-002
Created: 2026-05-22
Completed: 2026-05-22
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/exec-plans/completed/plan-refresh-002-structure-editing-roadmap.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/vim-curriculum-map.md

## 목표

완료된 `Review Queue and Incident Runs` 이후의 다음 중기 플랜을 문서에 고정하고, stale next candidate를 현재 방향에 맞게 갱신한다.

## 범위

- 포함:
  - 새 중기 플랜 `Structure Editing and Applied Survival` 문서화
  - 다음 활성 slice를 `TEXT-PAIR-GAP-001`로 설정
  - curriculum map의 Next Playpack Candidate를 quote/pair text object로 갱신
- 제외:
  - Go 코드 변경
  - content YAML 변경
  - progress 저장 포맷 변경

## 수용 기준

- completed: `docs/roadmap/MIDTERM_TODO.md`에 새 중기 플랜과 출구 조건이 있다.
- completed: `docs/roadmap/PROGRAM.md`의 활성 slice가 다음 실행 루프를 가리킨다.
- completed: `docs/gameplay/vim-curriculum-map.md`의 next candidate가 완료된 open-line-edit에서 quote/pair text object로 바뀐다.

## 검증 결과

- passed: `git diff --check`
