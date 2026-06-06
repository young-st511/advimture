# Playable Quality Baseline Review — 2026-06-02

## 목적

Advimture를 바로 출시하지는 않는다. 대신 현재 playable loop가 "출시 가능한 품질"에 가까워지도록 세계관, UX/UI, 콘텐츠 목표, 엔진/모듈화 기준을 하나의 baseline으로 묶는다.

이 문서는 release candidate 포장 문서가 아니다. 다음 개발을 계속하되, 매 slice가 제품 품질을 실제로 끌어올리는지 판단하기 위한 기준이다.

## Release-Quality Definition

현재 단계의 release-quality는 "모든 기능이 많은 게임"이 아니라 아래 상태를 뜻한다.

- 처음 실행한 사람이 자신이 누구이고 무엇을 복구하는지 1분 안에 이해한다.
- Tutorial은 Vim command를 명확히 가르치고, incident는 배운 command를 조합하는 복구 작전처럼 읽힌다.
- 실패, hint, 성공, review queue는 다음 행동을 숨기지 않는다.
- 작은 터미널에서도 핵심 action line과 hint가 잘리지 않는다.
- 대표 route evidence는 사람이 읽어도 게임의 흐름과 UI 의도를 확인할 수 있다.
- 콘텐츠 부족분은 "아직 없음"으로 방치하지 않고 다음 playable milestone 기획으로 정리된다.
- 엔진/모듈 구조는 다음 품질 개선을 좁은 slice로 열 수 있을 만큼 분리되어 있다.

## Axis 1. World

Release-quality 기준:

- 플레이어 역할은 `Runbook Dispatch`의 원격 복구 오퍼레이터다.
- 사건 단위는 "시설 하나, 콘솔 하나, 장애 하나"를 넘지 않는다.
- 세계관 명사는 `릴레이`, `신호`, `runbook`, `잔류 리스크`, `재점검`, `복구 기록`처럼 짧고 반복 가능한 어휘를 쓴다.
- 브리핑은 상황 1문장 + Vim 조작 목표 1문장으로 끝난다.
- 성공/실패 피드백은 세계관 감탄보다 어떤 Vim 조작이 좋았거나 필요한지 먼저 말한다.
- story가 target state, optimal keys, constraints를 바꾸면 실패다.

현재 판정:

- `docs/gameplay/world-frame.md`, `scenario-tone.md`, incident 001~007은 같은 세계관 프레임을 공유한다.
- review/debrief 언어도 `잔류 리스크`, `재점검`, `다음 출격`으로 정리되어 있다.
- P0/P1 blocker는 없다.
- P2: 초반 tutorial은 아직 "Runbook Dispatch"보다 Vim 첫 투어 성격이 강하다. 이는 의도된 단계지만, tutorial에서 incident로 넘어갈 때 플레이어 역할 전환을 더 부드럽게 만들 여지는 있다.
- Done: Header가 현재 track을 표시한다. Tutorial 화면은 `Tutorial`, incident 화면은 `Runbook Dispatch`를 playlist title 앞에 보여준다.

다음 후보:

- Tutorial 마지막에서 incident 첫 진입으로 넘어가는 `Next runbook` copy를 spot review한다.
- Incident 005 command-choice는 판단 훈련이므로, "범위 판별"이 세계관보다 Vim 선택 이유를 더 선명하게 말하는지 다시 본다.

## Axis 2. UX/UI

Release-quality 기준:

- 화면의 첫 시선은 항상 현재 mission title, 목표, runbook buffer로 간다.
- Tutorial running 화면은 새 command를 숨기지 않는다.
- Incident running 화면은 정답 sequence를 처음부터 노출하지 않고, hint/failure에서 command memory를 점진 공개한다.
- 실패 modal은 왜 실패했는지와 `Retry`를 유지한다.
- 성공 modal은 복구 기록, 최단 복구 기록, 잔류 리스크, 다음 행동을 유지한다.
- review/daily 정보는 현재 목표보다 위에 오지 않는다.
- 80x24/80x30 계열 viewport에서 action/hint line이 clipping으로 사라지지 않는다.

현재 판정:

- `Mission HUD -> Runbook Console -> floating modal -> status line` 정보 구조가 잡혀 있다.
- `FocusPanel`과 `app_state.ui.focus_panel`로 UI intent를 stable evidence로 검증한다.
- 80x24 success/failure modal smoke와 80x30 incident hint wrapping smoke가 있다.
- P0/P1 blocker는 없다.
- P2: 현재 화면은 검증 가능성이 강한 ASCII fallback 위주라, 제품 감각 측면의 color/emphasis pass는 아직 열려 있다.
- P2: wide terminal에서 side rail이 아직 없으므로 콘텐츠가 더 길어지면 HUD가 다시 밀릴 수 있다.
- Done: insert/search/command/visual mode cue는 영어 `Keys: ...`와 `* CHANNEL` 문구 대신 한국어 mode title/action label을 사용한다.
- Done: success/failure floating modal의 renderer 보조 label은 `Mistake`, `Learned`, `Result`, failure `Next` 대신 `실수`, `배운 점`, `기록`, `힌트`를 사용한다.

다음 후보:

- `UI-STYLE-001`: color/emphasis를 의미 보존 방식으로 얇게 적용한다.
- `UI-RAIL-001`: review/daily/command memory가 mission briefing을 다시 밀기 시작하면 wide layout side rail을 연다.
- `UI-PRESTART-001`: episode 전 brief가 HUD 2줄로 부족해질 때 runtime key trace와 분리된 pre-start modal을 검토한다.
- `UI-ACTION-LANGUAGE-001`: `Retry`/`Next runbook` 같은 action line을 한국어화하려면 화면 label과 E2E/action id 계약을 분리한다.

## Axis 3. Content

Release-quality 기준:

- 첫 플레이 루트는 Tutorial 0~3 초반까지 자연스럽게 이어진다.
- 중반 tutorial은 operator grammar, yank/put, text object, open line, repeat, search, visual, char find의 핵심을 작게 쪼개 보여준다.
- Incident는 새 command 소개가 아니라 이미 배운 command를 조합하는 적용 런이다.
- 첫 출시 가능한 콘텐츠 목표는 "짧은 튜토리얼 묶음 + mixed incident 3개 이상 + command-choice 적용 run + review queue"다.
- 추가 콘텐츠는 이 goal 안에서 구현하지 않고, command coverage와 incident arc 단위로 기획한다.

현재 판정:

- 현재 playable은 tutorial episode 0~9, 90~94와 incident 001~007, command-choice run을 가진다.
- README 기준 현재 106개 exercise와 18개 playlist가 연결되어 있다.
- P0/P1 blocker는 없다.
- 현재 분량은 foundation build로는 충분하지만, "그럴싸한 게임" 관점에서는 episode arc와 world ramp가 더 명확해야 한다.
- 콘텐츠 확장은 구현보다 기획이 먼저 필요하다. 특히 다음 milestone이 새 Vim command인지, 기존 command의 applied incident arc인지 분리해야 한다.

기획 backlog:

- Content Arc A: "First Tour to First Dispatch" — tutorial 0~3에서 incident 001로 넘어가는 전환 경험 정리.
- Content Arc B: "Applied Recovery Set" — incident 001~003을 하나의 relay station arc처럼 묶는 제목/briefing/review plan.
- Content Arc C: "Judgment Drills" — incident 005 command-choice를 기준으로 같은 engine 안에서 판단 훈련 beat 후보를 기획.
- Content Arc D: "Next Engine Candidate" — `ci(`/`ci{` pair hardening, search `?`, visual advanced 중 실제 콘텐츠 병목을 만드는 후보만 고른다.
- Done: `docs/roadmap/CONTENT_QUALITY_PLAN_001.md`에 first tour, core toolbelt, first dispatch arc, judgment drill, review loop의 release-quality content shape를 구체화했다.

## Axis 4. Engine / Modules

Release-quality 기준:

- Vim engine은 `State + Key -> State + Events` 순수 전이로 유지한다.
- Runtime은 exercise session, constraints, retry/hint를 담당한다.
- Content loader는 YAML, replay gate, coverage validation을 담당한다.
- Scenario layer는 exercise의 목표와 정답을 바꾸지 않는다.
- Playable model과 renderer는 UI 상태를 `FocusPanel`, app_state evidence로 설명할 수 있어야 한다.
- Progress 저장 포맷은 사용자 승인 없는 품질 작업과 섞지 않는다.

현재 판정:

- 현재 모듈 분리는 release-quality baseline을 진행하기에 충분하다.
- 새 content implementation 없이도 world/UX/content planning audit을 진행할 수 있다.
- 새 engine을 열 필요는 아직 증명되지 않았다.
- P2: renderer style pass를 열면 `internal/playableview` 중심으로 좁게 열어야 한다.
- P2: pre-start modal은 Bubble Tea input routing과 runtime key trace 분리가 필요하므로 별도 ExecPlan이 맞다.
- Done: `docs/roadmap/MODULE_QUALITY_REVIEW_2026-06-02.md`에 현재 package boundary, import flow, 새 engine 필요성, 구조 변경 후보를 문서화했다.
- Done: 큰 구조 변경은 막지 않고, 즉시 진행/checkpoint/사용자 승인 경로로 나누는 운영 원칙을 세웠다.

## Evidence Reviewed

문서:

- `docs/roadmap/PRE_RC_HARDENING_2026-06-02.md`
- `docs/gameplay/world-frame.md`
- `docs/gameplay/scenario-tone.md`
- `docs/gameplay/tui-ux-direction.md`
- `docs/roadmap/UX_BACKLOG_001.md`
- `docs/roadmap/ENGINE_TODO.md`
- `docs/roadmap/PLAYABLE_QUALITY_COMPLETION_AUDIT_2026-06-02.md`

대표 E2E evidence bundle:

- `artifacts/e2e/playable_ftue_first_five_route/screen_final.txt`
- `artifacts/e2e/playable_ftue_first_five_route/screen_timeline.txt`
- `artifacts/e2e/playable_ftue_first_five_route/app_state.json`
- `artifacts/e2e/playable_incident_001_full/screen_final.txt`
- `artifacts/e2e/playable_incident_001_full/screen_timeline.txt`
- `artifacts/e2e/playable_incident_001_full/app_state.json`
- `artifacts/e2e/playable_incident_hint_affordance/screen_final.txt`
- `artifacts/e2e/playable_incident_hint_affordance/screen_timeline.txt`
- `artifacts/e2e/playable_incident_hint_affordance/app_state.json`
- `artifacts/e2e/playable_review_queue/screen_final.txt`
- `artifacts/e2e/playable_review_queue/screen_timeline.txt`
- `artifacts/e2e/playable_review_queue/app_state.json`

Evidence 판정:

- FTUE와 review queue app_state는 success `FocusPanel`, residual risk, next action을 유지한다.
- Incident 001 full route app_state는 `Next runbook: enter`와 다음 residual risk를 유지한다.
- Incident hint affordance는 80x30 viewport에서 hint/action text를 보존한다.
- P0/P1 blocker는 없다.

## Decision

`RELEASE-CANDIDATE-001`은 지금 열지 않는다. 다음 active direction은 `PLAYABLE-QUALITY-BASELINE-001`이다.

첫 구현 후보는 새 content/engine이 아니라 아래 순서로 고른다.

1. Roadmap과 기준 문서 정합성 정리
2. World/UX evidence spot review에서 P2 이상으로 확인된 작은 개선
3. Content arc 기획 문서
4. 필요한 경우 renderer/style 또는 pre-start modal 같은 별도 ExecPlan
