# Health Check — 2026-05-22

## 결론

판정: **Conditional Go**

`Structure Editing and Applied Survival`은 완료로 본다. quote text object는 엔진, tutorial, E2E까지 연결됐고, `incident-002`는 search/substitute/quote/dot repeat 조합을 실제 playable 적용 런으로 검증했다.

다만 다음 단계는 바로 visual mode 구현이 아니다. Visual mode는 다음 큰 후보가 맞지만, selection state와 TUI 표시, app_state/E2E assertion 계약이 먼저 필요하다. 동시에 incident run은 기능적으로 통과했으나 아직 “하나의 복구 작전”보다 “종합시험”에 가깝다. 다음 중기 플랜은 incident UX 보강과 visual readiness를 함께 다루는 것이 좋다.

## 점검 방식

- Architecture Review는 별도 SubAgent에게 위임했다.
- Learning Design / Game UX Review는 별도 SubAgent에게 위임했다.
- 본 세션에서는 두 리뷰 결과를 현재 문서와 코드 구조에 대조했다.
- 파일 수정 전 기준 문서와 domain contract를 확인했다.

## Architecture 판단

좋은 점:

- `vimengine -> runtime -> content/scenario -> playable/progressadapter` 경계는 여전히 작동한다.
- progress 저장 포맷과 dependency boundary는 지켜졌다.
- E2E runner는 temp HOME, key trace, app_state assertion을 갖고 있어 다음 검증 확장의 기반이 있다.

위험:

| 심각도 | 항목 | 판단 |
|--------|------|------|
| High | visual selection state contract 부재 | `vimengine.State`, `tuiadapter.ViewModel`, `e2estate.State`가 selection을 표현하지 못한다. |
| Medium | content/E2E assertion schema 한계 | 현재 assertion은 buffer/cursor/mode/status/command 중심이라 선택 중간 상태를 검증하기 어렵다. |
| Medium | runtime constraint의 path 검증 약함 | `required_keys`는 포함 여부 중심이라 `v -> motion -> d/y`의 순서와 범위를 충분히 고정하지 못한다. |
| Low | TUI rendering은 cursor 단일 표시 | visual mode는 화면 표시 자체가 학습 피드백이므로 최소 highlight 계약이 필요하다. |

## Learning / Game UX 판단

좋은 점:

- quote tutorial은 4문항으로 짧고 실무형이다.
- `incident-002`의 `secret 복구 -> mirror 동기화 -> temp 승격 -> 반복 패치` 흐름은 좋은 뼈대다.
- search/substitute/quote/dot repeat 조합은 중급 적용 런으로 적절하다.

위험:

| 심각도 | 항목 | 판단 |
|--------|------|------|
| High | incident game-feel 부족 | 각 exercise가 독립 buffer로 시작해 하나의 사건을 누적 해결하는 감각이 약하다. |
| High | visual mode UX 선행 필요 | selection이 보이고 잡히고 적용되는 피드백 없이 콘텐츠부터 만들면 혼란스럽다. |
| Medium | 최근 학습 조합 편중 | quote/dot/search/substitute가 연달아 등장해 일부 플레이어에게 정답 암기처럼 느껴질 수 있다. |
| Medium | incident hint 깊이 부족 | harness는 최소 2단계 힌트를 요구하지만 incident beat 일부는 1단계 힌트에 머문다. |

## 다음 중기 플랜 추천

1. **PLAN-REFRESH-003 — Applied Learning / Visual Readiness Plan**
   - 중기 플랜을 incident UX, 반복 동기, visual readiness 순서로 재정렬한다.

2. **INCIDENT-UX-003 — Incident Run Game-Feel 강화**
   - 새 엔진 기능 없이 briefing, beat 연결감, 성공 debrief, 2단계 힌트를 보강한다.
   - 목표는 incident가 “문제 세트”가 아니라 “하나의 복구 작전”처럼 느껴지게 하는 것이다.

3. **E2E-FIXTURE-001 — Progress Fixture 유지보수 완화**
   - full playlist E2E의 긴 progress fixture가 계속 늘고 있다.
   - visual/playpack 추가 전에 최소 fixture 또는 fixture builder 전략을 검토한다.

4. **VISUAL-GAP-002 — Visual Selection Contract**
   - selection kind, anchor/head, mode 이름, reset 규칙, charwise/linewise 우선순위를 결정한다.
   - content assertion과 e2e app_state의 selection 표현을 함께 정의한다.

5. **E2E-007 — Selection App State Assertion**
   - `internal/e2estate`, `cmd/e2e-runner`, content `e2e_assertions`에 같은 의미의 selection assertion을 추가한다.

6. **VIM-027 / TUI-003 — Minimal Visual Foundation**
   - `v` charwise selection, motion 갱신, `esc` reset, 최소 화면 표시까지만 닫는다.
   - `d/y` operator application과 playpack은 그 다음 루프로 분리한다.

## 보류

- visual block `<C-v>`
- indentation `>`, `<`, `=`
- count prefix, register prefix
- mouse/terminal selection 연동
- progress schema v2, mastery/daily run 저장
- incident 신규 엔진 기능 추가
- visual mode와 incident run을 한 playpack 안에 섞기

## 사용자 결정 필요 사항

1. 다음 중기 플랜을 `INCIDENT-UX-003`으로 시작할지, `VISUAL-GAP-002`로 시작할지 결정한다.
2. visual 첫 구현을 charwise `v`로 갈지, linewise `V`로 갈지 결정한다.
3. E2E fixture builder를 다음 플랜에 포함할지 결정한다.

