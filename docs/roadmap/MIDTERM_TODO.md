# Midterm Todo

> 현재 중기 보드만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 과거 중기 플랜의 긴 히스토리는 `docs/roadmap/archive/history/MIDTERM_TODO_2026-05-29.md`와 `docs/exec-plans/completed/`를 본다.

## 현재 상태

Status: next slice proposed

현재 active slice는 없다. `UI-POLISH-002`로 command memory cue까지 연결했다. 다음 단계는 `RELEASE-READINESS-001`로 첫 공개 전 문서와 검증 기준을 정리하는 것이다.

## 다음 중기 플랜 후보

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | RELEASE-READINESS-001 | proposed | 첫 공개 전 설치/실행/터미널 크기/known limitations/release build gate 정리 |

권장은 `RELEASE-READINESS-001`을 열어 release 문서와 known limitations를 현재 구현에 맞추는 것이다. 현재 rolling plan은 `docs/roadmap/FORWARD_PLAN.md`를 따른다.

## RELEASE-READINESS-001 출구 조건

| 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|-----------|-------------|------|---------------|
| UI polish 완료 | README/usage/test/known limitations/release build gate 정리 | `go test ./...`, `make e2e-playable`, `go build .` | 문서가 구현보다 앞서거나 stale next plan을 만들지 않음 |

## 다음 결정 기준

- **새 engine을 열 때**: `docs/roadmap/ENGINE_TODO.md`의 hardening candidates와 `docs/gameplay/vim-curriculum-map.md`의 coverage gaps를 먼저 본다.
- **게임성을 보강할 때**: `docs/roadmap/UX_BACKLOG_001.md`, `docs/roadmap/PLATFORM_RFC_001.md`, 최신 playtest review를 먼저 본다.
- **content를 늘릴 때**: 기존 implemented engine만 사용할 수 있는지 확인하고, long route에는 final/timeline E2E evidence를 남긴다.

## 최근 완료된 중기 플랜

| ID | 완료일 | 요약 |
|----|--------|------|
| UI-POLISH-002 | 2026-05-30 | tutorial/incident command memory cue 추가. schema 변경 없음 |
| QUOTE-PAIR-HARDEN-001 | 2026-05-30 | `di'`, `ci'`, `yi'` single quote text object를 tutorial/E2E까지 연결 |
| CONTENT-BREADTH-002 | 2026-05-30 | `incident-005-command-choice` fifth beat로 repeat-change choice 추가 |
| PLATFORM-REVIEW-003 | 2026-05-30 | 성공 debrief와 마지막 dispatch를 review queue로 연결. progress schema 변경 없음 |
| FOUNDATION-EXIT-001 | 2026-05-30 | Foundation 조건부 통과. 다음 병목은 새 engine보다 mission/review/game loop로 판정 |
| PLAN-REFRESH-009 | 2026-05-30 | 앞으로 2~8주 rolling plan과 문서 참고 규칙 작성 |
| E2E-EVIDENCE-008 | 2026-05-29 | 긴 incident와 command-choice applied route의 final/timeline evidence bundle 고정 |
| Applied Mastery Runs | 2026-05-28 | `incident-006`, quote value reuse-choice, `incident-007` 완료 |
| Inline Target Motions | 2026-05-26 | `f/t`, `df/dt/cf/ct`, tutorial, command-choice 적용 beat 완료 |
| Foundation Playtest / UX Polish | 2026-05-26 | HUD/help/choice/success modal polish와 command-choice content breadth 완료 |

## 운영 규칙

- 이 파일은 현재/다음 중기 계획만 유지한다.
- 앞으로의 방향이 바뀌면 `docs/roadmap/FORWARD_PLAN.md`도 함께 갱신한다.
- 완료된 긴 계획표는 `docs/roadmap/archive/history/`로 이동한다.
- "다음 중기 플랜" 섹션이 여러 개 쌓이면 stale 위험으로 보고 정리한다.
