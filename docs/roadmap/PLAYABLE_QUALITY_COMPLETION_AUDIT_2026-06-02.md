# Playable Quality Completion Audit — 2026-06-02

## 목적

`PLAYABLE-QUALITY-BASELINE-001`이 사용자 목표를 실제로 만족하는지 요구사항별로 감사한다.

이 감사의 완료 기준은 "바로 출시한다"가 아니다. 현재 playable loop가 출시 가능한 품질 기준에 가까워졌고, 남은 작업이 P0/P1 blocker가 아니라 후속 backlog나 별도 ExecPlan으로 분리되었는지를 확인한다.

## 요구사항 감사

| 요구사항 | 완료 증거 | 판정 |
|----------|-----------|------|
| 세계관 정합성 기준 설정 | `docs/gameplay/world-frame.md`, `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md` | 완료 |
| 세계관 기준에 맞는 개선 | Header track label이 `Tutorial` / `Runbook Dispatch`를 첫 viewport에 표시한다. `playable_incident_001_full`, `playable_command_choice_scope` evidence가 이를 검증한다. | 완료 |
| UX/UI release-quality 목표 설정 | `docs/gameplay/tui-ux-direction.md`, `docs/gameplay/tui-screen-contract.md`, `docs/roadmap/UX_BACKLOG_001.md` | 완료 |
| UX/UI 개선 | mode cue/title 한국어화, success/failure modal 보조 label 한국어화, header track label 추가. 관련 unit test와 focused E2E가 통과했다. | 완료 |
| 콘텐츠 출시 품질 목표 설정 | `docs/roadmap/CONTENT_QUALITY_PLAN_001.md`의 First Tour, Core Toolbelt, First Dispatch Arc, Judgment Drill, Review Loop | 완료 |
| 현재 콘텐츠 점검 | 현재 tutorial, incident, command-choice, review loop를 minimum release-quality content shape에 매핑했다. | 완료 |
| 콘텐츠는 기획까지만 진행 | 새 YAML content, content schema, engine, progress 저장 포맷을 변경하지 않았다. | 완료 |
| 모듈화와 엔진 충분성 검증 | `docs/roadmap/MODULE_QUALITY_REVIEW_2026-06-02.md`, `go list ./internal/...`, import flow review | 완료 |
| 큰 구조 변경 허용 원칙 | 즉시 진행 / checkpoint / 사용자 승인 경로로 blast radius를 나눴다. 큰 수정 자체가 아니라 evidence 없는 제품 계약 변경을 막는다. | 완료 |
| 대표 evidence 기반 P0/P1 blocker 판정 | first tour, incident 001, command-choice, review queue, viewport modal evidence를 spot review했다. | 완료 |

## Current-State Evidence

대표 E2E evidence:

- `artifacts/e2e/playable_ftue_first_five_route/screen_final.txt`
- `artifacts/e2e/playable_ftue_first_five_route/screen_timeline.txt`
- `artifacts/e2e/playable_ftue_first_five_route/app_state.json`
- `artifacts/e2e/playable_incident_001_full/screen_final.txt`
- `artifacts/e2e/playable_incident_001_full/screen_timeline.txt`
- `artifacts/e2e/playable_incident_001_full/app_state.json`
- `artifacts/e2e/playable_command_choice_scope/screen_final.txt`
- `artifacts/e2e/playable_command_choice_scope/screen_timeline.txt`
- `artifacts/e2e/playable_command_choice_scope/app_state.json`
- `artifacts/e2e/playable_review_queue/screen_final.txt`
- `artifacts/e2e/playable_review_queue/screen_timeline.txt`
- `artifacts/e2e/playable_review_queue/app_state.json`

Spot review 판정:

- First Tour evidence는 `Tutorial` track, 현재 목표, command memory, success modal, residual risk, next action을 유지한다.
- Incident 001 evidence는 `Runbook Dispatch` track, relay station title, search/change/open/yank/substitute 조합, runbook completion, residual risk를 유지한다.
- Command-choice evidence는 `Runbook Dispatch` track과 선택 이유 중심의 성공 feedback을 유지한다.
- Review queue evidence는 progress v1만으로 `잔류 리스크`와 `다음 출격` 반복 동기를 설명한다.

## Guardrail Audit

변경하지 않은 경계:

- `go.mod`, `go.sum`
- `internal/progress/`
- `content/`

의도적으로 열지 않은 작업:

- release tag 또는 release candidate packaging
- 새 tutorial/incident YAML 구현
- content schema 변경
- progress 저장 포맷 변경
- 새 의존성 추가
- pre-start modal input boundary
- 새 Vim engine capability

## Residual Work

남은 항목은 완료 blocker가 아니라 후속 backlog다.

| 후보 | 이유 | 다음 경로 |
|------|------|-----------|
| `UI-ACTION-LANGUAGE-001` | `Retry` / `Next runbook` action line은 아직 화면 label과 E2E marker를 겸한다. | 화면 label과 machine-readable action id 분리 시 별도 slice |
| `UI-STYLE-001` | ASCII fallback 중심이라 제품 감각 polish 여지가 있다. | 주요 flow 안정 후 style-only pass |
| `UI-RAIL-001` | wide terminal에서 side rail이 있으면 장기 콘텐츠가 더 안정된다. | review/daily가 mission briefing을 밀 때 |
| `UI-PRESTART-001` | first dispatch 전환을 더 극적으로 만들 수 있다. | HUD 두 줄이 부족하다는 evidence가 반복될 때 |
| `CONTENT-ARC-POLISH-001` | incident 001~003 arc framing을 더 매끄럽게 만들 수 있다. | 새 content 구현 없이 scenario copy polish만 별도 ExecPlan |
| `JUDGMENT-DRILL-REVIEW-001` | command-choice 선택 이유를 더 정밀하게 감사할 수 있다. | evidence bundle 기반 review |

## 결론

`PLAYABLE-QUALITY-BASELINE-001`은 목표 범위 안에서 완료 상태다.

현재 playable loop는 바로 release candidate로 포장된 상태는 아니지만, 세계관/UX/UI/콘텐츠 기획/엔진 모듈화가 출시 가능한 품질 기준에 맞춰 정렬되었다. 남은 항목은 P0/P1 blocker가 아니라 후속 품질 backlog 또는 별도 ExecPlan 후보로 분리되었다.
