# DOCS-CLEANUP-001 — Roadmap Staleness Cleanup

Slice-ID: DOCS-CLEANUP-001
Created: 2026-05-30
Completed: 2026-05-30
Status: completed
Scope-Mode: docs-only
Allowed-Paths:
- docs/README.md
- docs/exec-plans/active/docs-cleanup-001-roadmap-staleness.md
- docs/exec-plans/completed/docs-cleanup-001-roadmap-staleness.md
- docs/gameplay/content-requirements.md
- docs/gameplay/tui-ux-direction.md
- docs/gameplay/vim-curriculum-map.md
- docs/roadmap/ENGINE_TODO.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/PLATFORM_RFC_001.md
- docs/roadmap/UX_BACKLOG_001.md
- docs/roadmap/archive/
- docs/roadmap/CHANGES.md
- docs/verification/selection-app-state-contract.md
- docs/verification/tui-ui-qa-contract.md

## 목표

현재 참조 가능한 위치에 남아 있는 오래된 next 후보와 이미 완료된 backlog 항목을 정리한다. 다음 Agent가 stale roadmap을 active 계획처럼 읽지 않도록 canonical/current/history 경계를 명확히 한다.

## 범위

- 포함:
  - `PROGRAM.md`를 현재 phase/active/recent completed 중심으로 축소
  - `MIDTERM_TODO.md`를 현재 보드 중심으로 재작성하고 이전 긴 히스토리 archive
  - `ENGINE_TODO.md`, `vim-curriculum-map.md`, UX/QA contract의 stale next 후보 갱신
  - 오래된 health/review 문서를 `docs/roadmap/archive/`로 이동
  - 앞으로의 문서 정리 규칙을 `docs/README.md`에 추가
- 제외:
  - Go 코드 변경
  - content YAML 변경
  - progress 저장 포맷 변경
  - 새 기능/새 엔진 계획 승인

## 수용 기준

- `PROGRAM.md`는 현재 phase, active slice, 다음 권장 후보, 최근 완료만 보여준다.
- `MIDTERM_TODO.md`는 현재 중기 보드와 다음 후보만 보여주고, 과거 긴 히스토리는 archive에 있다.
- root `docs/roadmap/`에는 오래된 health/post-visual review가 current 후보처럼 남지 않는다.
- `ENGINE_TODO.md`와 `vim-curriculum-map.md`는 완료된 `Inline Target`, `incident-006`, `reuse-choice`를 next로 표시하지 않는다.
- `screen_final.txt`, linewise visual처럼 이미 구현된 항목이 QA/backlog에서 next처럼 보이지 않는다.
- 앞으로 stale 문서를 방지할 운영 규칙이 문서화된다.

## 검증 계획

- `rg`로 주요 stale phrase 재검색
- `git diff --check`

## Step 1: Active Docs

- [x] `PROGRAM.md` 축소
- [x] `MIDTERM_TODO.md` 현재 보드화
- [x] roadmap history archive 이동

## Step 2: Domain/Contract Sync

- [x] `ENGINE_TODO.md` 갱신
- [x] `vim-curriculum-map.md` 갱신
- [x] UX/QA/selection contract 갱신
- [x] platform/content requirement 상태 메모 갱신

## Step 3: Operating Rule

- [x] `docs/README.md`에 freshness policy 추가
- [x] 검증 후 completed로 이동

## 검증 결과

- stale phrase search: intended README policy hit only
- archived old roadmap direct-path reference search: pass
- `git diff --check`: pass

## 의사결정 로그

- 2026-05-30: active plan이 없는 시점에서 문서 cleanup을 먼저 수행한다.
- 2026-05-30: 과거 판단 문서는 삭제하지 않고 archive로 이동해 히스토리는 보존한다.
- 2026-05-30: `PROGRAM.md`와 `MIDTERM_TODO.md`는 긴 완료 로그가 아니라 현재 판단 보드로 유지한다.

## 미해결 질문

- 없음.
