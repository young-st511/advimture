# Module Quality Review — 2026-06-02

## 목적

`PLAYABLE-QUALITY-BASELINE-001`의 엔진/모듈화 축을 현재 코드 기준으로 점검한다.

질문은 "새 엔진을 열어야 하는가?"가 아니라, 세계관/UX/UI/콘텐츠 품질을 계속 올릴 때 현재 모듈 경계가 충분히 좁은 변경 단위를 제공하는가다.

## 설계 변경 운영 원칙

큰 구조 변경은 금지 대상이 아니다. release-quality 목표에 더 가까워지려면 `playable` orchestration, input routing, renderer DTO, engine capability를 열 수 있다.

다만 변경은 blast radius에 따라 세 경로로 나눈다.

| 경로 | 예시 | 처리 |
|------|------|------|
| 즉시 진행 | renderer label/copy, 순수 함수 분리, test-only fixture, 문서 보강 | 이 ExecPlan 안에서 구현하고 focused test로 검증한다. |
| checkpoint 필요 | `internal/playable` flow 분리, pre-start modal input boundary, app_state projection 변경 | 별도 ExecPlan 또는 현재 ExecPlan의 명시 step으로 열고 수용 기준과 E2E evidence를 먼저 둔다. |
| 사용자 승인 필요 | progress 저장 포맷, content schema/ID, 새 의존성, 신규 대형 콘텐츠 구현 | 목표와 무관하게 바로 진행하지 않는다. 승인 후 별도 slice로 연다. |

따라서 막아야 하는 것은 "큰 수정" 자체가 아니라, evidence 없이 제품 계약을 흔드는 변경이다.

## 현재 패키지 경계

`go list ./internal/...` 기준 현재 내부 패키지는 아래 역할로 나뉜다.

| 패키지 | 역할 | release-quality 판정 |
|--------|------|----------------------|
| `internal/vimengine` | Vim-like `State + Key -> State + Events` 전이 | 충분 |
| `internal/runtime` | exercise session, constraints, retry/hint, goal matching | 충분 |
| `internal/scoring` | grade/key count/intent scoring | 충분 |
| `internal/content` | YAML loader, replay gate, coverage validation | 충분 |
| `internal/scenario` | exercise를 briefing/success/failure/hint runbook layer로 감싸기 | 충분 |
| `internal/tuiadapter` | scenario/runtime state를 view model로 변환 | 충분 |
| `internal/playableview` | 순수 renderer, viewport wrapping, modal rendering | 충분 |
| `internal/playable` | Bubble Tea model, entry progression, progress/review/E2E state orchestration | 충분하나 허브 리스크 있음 |
| `internal/review` | progress v1 기반 review queue 계산 | 충분 |
| `internal/progressadapter` | scenario result를 progress completion으로 변환 | 충분 |
| `internal/progress` | progress v1 저장 경계 | 충분, 변경 주의 |
| `internal/e2estate` | typed E2E app_state snapshot | 충분 |
| `internal/vimoracle` | optional Neovim oracle comparison | 충분 |
| `internal/app` | app entrypoint wiring | 충분 |

## Import Flow

현재 import flow는 release-quality 작업을 진행하기에 충분히 단방향이다.

```text
vimengine
  <- runtime
  <- content, tuiadapter, vimoracle

runtime
  <- scoring
  <- scenario
  <- progressadapter

content
  <- scenario
  <- review
  <- playable

scenario + tuiadapter + playableview + review + progressadapter + progress + e2estate
  <- playable
  <- app
```

관찰:

- `vimengine`은 UI, content, progress를 모른다.
- `runtime`은 content YAML이나 Bubble Tea를 모른다.
- `scenario`는 content/runtime/scoring만 알고, progress 저장을 모른다.
- `playableview`는 Bubble Tea model을 모르고 `Screen` DTO를 렌더링한다.
- `e2estate`는 runtime/playable에 의존하지 않는 JSON snapshot shape다.
- `playable`은 의도적으로 orchestration hub다. 품질 개선이 계속 `playable.Model`에 쌓이면 후속 분리 후보가 된다.

## 충분한 점

현재 구조는 아래 작업을 좁은 slice로 처리할 수 있다.

- world/UX copy polish: content scenario copy 또는 renderer label만 변경
- viewport/hint/modal 안정화: `internal/playableview` 테스트와 focused E2E로 검증
- command-choice 기획/검증: existing content constraints와 E2E로 처리
- review/debrief copy 조정: `internal/playable`, `internal/review`, renderer surface로 제한 가능
- engine hardening: command cluster 단위로 `vimengine -> runtime -> content -> E2E` 순서 개방 가능

## 리스크와 분리 후보

### P1. Playable orchestration hub

`internal/playable`은 content loading, Bubble Tea update, progress save, review queue, focus panel, E2E state를 모두 묶는다.

현재는 release-quality baseline을 막지 않는다. 다만 아래 조건이 생기면 분리한다.

- pre-start modal처럼 runtime key trace와 UI-only input을 분리해야 한다.
- review/daily route가 더 복잡해져 playable model 안의 상태 조합이 늘어난다.
- focus panel line generation이 tutorial/incident/review/world rules를 과하게 많이 알게 된다.

후보 분리:

- `internal/playable/focus` 또는 별도 package: FocusPanel identity/lines 생성
- `internal/playable/flow`: next entry, review dispatch, retry/success transition
- `internal/playable/state`: E2E state projection

### P1. Pre-start modal input boundary

`UI-PRESTART-001`은 단순 renderer 문제가 아니다. modal open 상태의 `enter`는 Vim runtime key trace로 들어가면 안 된다.

열기 조건:

- episode/runbook 시작 전 설명이 Mission HUD 2줄로 부족하다는 evidence가 반복된다.
- first dispatch 진입 전 player role 설명이 현재 success/debrief copy만으로 부족하다.

필요한 계약:

- UI-only input state
- key trace exclusion
- E2E app_state에 modal state 추가 여부

### P2. Playableview DTO coupling

`internal/playableview`는 현재 selection rendering 때문에 `internal/tuiadapter.SelectionView`를 참조한다.

현재는 문제 없다. 다만 renderer를 더 독립적인 package로 만들거나 screenshot-like evidence를 강화하려면 `playableview.Selection` DTO를 내부에 둘 수 있다.

열기 조건:

- renderer 테스트가 tuiadapter shape 변화에 자주 흔들린다.
- visual rendering style pass에서 selection model이 커진다.

### P2. Content schema expansion

현재 content schema는 release-quality content planning을 감당한다. command-choice의 `choice_focus`, `why_intended` 같은 기획 필드는 문서에만 있다.

열기 조건:

- 선택 이유를 E2E/app_state로 typed 검증해야 한다.
- command-choice authoring이 scenario copy만으로 유지하기 어려워진다.

주의:

- content schema 변경은 사용자 확인과 별도 ExecPlan이 필요하다.

## 새 Engine 필요성 판정

현재 release-quality baseline에는 새 Vim engine이 필요하지 않다.

이유:

- first tour, core toolbelt, first dispatch arc, judgment drill, review loop가 모두 현재 implemented clusters로 구성된다.
- incident 001~003과 incident 005는 기존 engine만으로 제품 skeleton을 설명한다.
- `ci(`/`ci{`, search `?`, visual advanced는 유용하지만 "현재 품질 기준을 막는 blocker"는 아니다.

다음 engine 후보를 열려면 아래 중 하나가 evidence로 확인되어야 한다.

- 콘텐츠 arc가 기존 quote/visual/search 기능만으로 반복감에 빠진다.
- 특정 실무 편집 감각이 현재 engine 범위로 설명되지 않는다.
- E2E evidence에서 플레이어가 "왜 이 Vim 기능이 빠졌는지"를 체감할 정도의 빈틈이 보인다.

## 결론

현재 모듈화와 엔진 범위는 release-quality baseline을 계속 진행하기에 충분하다.

다음 큰 구조 변경 후보는 새 engine이 아니라 `playable` orchestration 분리 또는 pre-start modal input boundary다. 둘 다 지금 즉시 열 필요는 없고, first dispatch 전환 UX evidence에서 실제 병목이 확인될 때 별도 ExecPlan으로 연다.
