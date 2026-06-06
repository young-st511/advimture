# Midterm Todo

> 현재 중기 보드만 둔다. 앞으로의 2~8주 방향은 `docs/roadmap/FORWARD_PLAN.md`를 본다. 과거 중기 플랜의 긴 히스토리는 `docs/roadmap/archive/history/MIDTERM_TODO_2026-05-29.md`와 `docs/exec-plans/completed/`를 본다.

## 현재 상태

Status: first dispatch / judgment drill / console polish completed

현재 active slice는 없다. `PLAYABLE-QUALITY-BASELINE-001` 이후의 `CONTENT-ARC-POLISH-001 -> JUDGMENT-DRILL-REVIEW-001 -> UI-CONSOLE-POLISH-001` 중기 보강도 완료됐고, 바로 출시 후보를 포장하지 않는다는 방향은 유지한다.

## 현재 중기 플랜

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | PLAYABLE-QUALITY-BASELINE-001 | completed | 세계관/UX/UI/콘텐츠 목표/엔진 모듈화를 release-quality 기준으로 점검하고 개선 |
| 2 | CONTENT-ARC-POLISH-001 | completed | incident 001~003을 원인 신호 추적 -> 구조 재동기화 -> 오염 구간 격리 첫 dispatch arc로 정리 |
| 3 | JUDGMENT-DRILL-REVIEW-001 | completed | command-choice 5 beat를 Vim 도구 선택 판단 중심으로 보강 |
| 4 | UI-CONSOLE-POLISH-001 | completed | action line 화면 label과 machine-readable action id 분리 |
| 5 | POST-POLISH-PLAYTEST-001 | next candidate | fresh evidence review로 다음 병목 판정 |
| 6 | RELEASE-CANDIDATE-001 | later | 실제 공개 후보를 묶기로 결정했을 때 release note/evidence/tag 후보 정리 |

권장은 release candidate 포장이 아니라 quality baseline 이후의 후속 보강을 필요할 때 좁게 여는 것이다. 큰 구조 변경은 허용하지만, progress 저장 포맷/content schema/새 의존성은 별도 확인 전까지 열지 않는다.

## PLAYABLE-QUALITY-BASELINE-001 출구 조건

| 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|-----------|-------------|------|---------------|
| pre-RC hardening 완료 | 세계관 기준, UX/UI 기준, 콘텐츠 목표/기획 backlog, 엔진/모듈화 판정, 대표 evidence review | `go test ./...`, 필요한 focused E2E, `git diff --check` | 신규 대형 콘텐츠 구현 금지, 위험한 schema/dependency/progress 변경은 별도 checkpoint |

## PLAYTEST-GATE-001 검사 축

| 축 | 질문 | evidence |
|----|------|----------|
| README 실행성 | 처음 온 사람이 README만 보고 실행/검증할 수 있는가? | README, `make release-check` output |
| 첫 5분 | tutorial 0~3 초반이 command 학습 목적을 잃지 않는가? | `playable_ftue_first_five_route`, final/timeline |
| tutorial 확장 | 중반 tutorial이 지치거나 command memory가 부족하지 않은가? | open-line/repeat/search/quote/visual/char-find full E2E |
| incident 감각 | incident 3개 이상이 “복구 작전”처럼 읽히는가? | incident full route final/timeline |
| 진행/복습 | debrief, review queue, next dispatch가 반복 동기를 주는가? | review/debrief E2E, app_state |
| terminal UI | modal, HUD, 긴 한국어 문구가 터미널 크기에서 무너지지 않는가? | screen_final, 필요 시 viewport smoke |
| 세계관 정합성 | tutorial, incident, review language가 Runbook Dispatch 프레임 안에서 이어지는가? | world-frame, scenario-tone, incident final/timeline |
| 엔진/모듈화 | 품질 개선을 좁은 slice로 열 수 있는 분리가 되어 있는가? | ENGINE_TODO, package boundaries, 관련 tests |

## 다음 결정 기준

- **새 engine을 열 때**: `docs/roadmap/ENGINE_TODO.md`의 hardening candidates와 `docs/gameplay/vim-curriculum-map.md`의 coverage gaps를 먼저 본다.
- **게임성을 보강할 때**: `docs/roadmap/UX_BACKLOG_001.md`, `docs/roadmap/PLATFORM_RFC_001.md`, 최신 playtest review를 먼저 본다.
- **content를 늘릴 때**: 기존 implemented engine만 사용할 수 있는지 확인하고, long route에는 final/timeline E2E evidence를 남긴다.

## 최근 완료된 중기 플랜

| ID | 완료일 | 요약 |
|----|--------|------|
| CONTENT-ARC-POLISH-001 | 2026-06-06 | incident 001~003 copy/framing을 첫 Runbook Dispatch arc로 정리. 새 YAML/schema 없음 |
| JUDGMENT-DRILL-REVIEW-001 | 2026-06-06 | command-choice 5 beat success/failure copy와 drill docs를 선택 이유 중심으로 최신화 |
| UI-CONSOLE-POLISH-001 | 2026-06-06 | FocusPanel action label과 app_state action id를 분리하고 80x24 modal focused E2E 통과 |
| PLAYABLE-QUALITY-BASELINE-001 | 2026-06-02 | release-quality 기준, 콘텐츠 기획, 모듈/엔진 충분성, header/mode/modal UX 보강, completion audit 완료 |
| PRE-RC-HARDENING-001 | 2026-06-02 | P0/P1 없음. Incident hint cue wrapping과 long incident app_state evidence gap 보강 |
| FIRST-RUN-POLISH-001 | 2026-06-01 | cue 중복 제거, modal 주변 detailed review/daily line 정리, 80x24 viewport smoke와 mid tutorial evidence 보강 |
| PLAYTEST-GATE-001 | 2026-05-30 | P0/P1 blocker 없음. 후속 first-run polish 완료 |
| RELEASE-READINESS-001 | 2026-05-30 | README, Makefile release gate, known limitations 정리 |
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
