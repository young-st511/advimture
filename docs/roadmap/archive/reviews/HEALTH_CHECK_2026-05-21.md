# Health Check — 2026-05-21

## 결론

판정: **Conditional Go**

Advimture는 Vim command를 추가하고 playable tutorial로 승격하는 엔진/콘텐츠/E2E 루프는 충분히 작동한다. 다음 단계로 갈 수 있다. 다만 다음 2~4주는 새 Vim 기능을 더 늘리기보다, 저장 포맷 변경 없는 review loop와 mixed incident run으로 반복 플레이 동기를 먼저 만드는 편이 좋다.

progress schema v2, spaced review 저장, daily streak, calendar 기반 daily run은 아직 No-Go다. 별도 승인 전까지 `internal/progress/` 저장 JSON 구조는 변경하지 않는다.

## 현재 완성도

- `vimengine -> runtime -> content/scenario -> playable/progressadapter` 경계는 다음 command/playpack 확장에 충분하다.
- `open-line-edit`, `repeat-last-change`, `search-basic`은 gap planning, engine, content, E2E까지 닫혔다.
- `make e2e-playable`은 중간점검 중 stale expectation을 복구했고, 최신 full playpack E2E를 포함한다.
- Learning design은 Go다. `required_keys`, `forbidden_keys`, `max_inputs`, `mistake_focus`가 “의도 command로 풀기”를 잘 보호한다.
- Gameplay UX는 Conditional Go다. 첫 클리어까지는 학습 게임이지만, 반복 플레이 동기는 아직 best record 비교 수준이다.

## 구조상 강점

- Vim engine은 `State + Key -> State + Events` 순수 전이로 유지된다.
- content loader는 approved/implemented/replay/e2e assertion gate를 강하게 검증한다.
- runtime은 forbidden/max/required key를 engine 진행과 분리해 처리한다.
- E2E runner는 temp HOME, unsafe HOME guard, key trace, app state assertion을 갖췄다.
- progress v1 경계와 `PLATFORM_RFC_001.md`의 승인 게이트가 일치한다.

## 위험과 부채

| 심각도 | 항목 | 판단 |
|--------|------|------|
| High | 반복 동기가 약함 | 완료 후 best record는 있으나 “오늘 무엇을 다시 할지”가 없다. |
| High | progress loader error policy | schema v2 전에 손상/마이그레이션 실패 정책을 명확히 해야 한다. |
| Medium-High | strict constraints 체감 | `max_inputs == optimal` 문항은 초보자에게 주문 입력처럼 느껴질 수 있다. |
| Medium | `playable.Model` 비대화 가능성 | UI, content load, scenario, progress, E2E state가 한 모델에 모인다. |
| Medium | playlist ID 정렬 의존 | `tutorial-90-search-basic`처럼 순서 회피용 ID 압력이 생겼다. |
| Medium | 일부 문항 질감 반복 | text-object/open-line은 비슷한 key spell이 이어져 변주가 약하다. |

## 중간점검 중 즉시 수정한 것

Architecture Review에서 full E2E 일부가 stale expectation을 가진다는 신호가 나왔다. 기존 기억 규칙대로 다음 기획으로 넘어가지 않고 `QA-PLATFORM-001`을 먼저 수행했다.

수정:

- 중간 tutorial full E2E의 마지막 action을 `Playlist complete`에서 `Next tutorial: enter`로 갱신했다.
- `make e2e-playable`에 operator/yank/text-object/open-line/repeat/search full E2E를 추가했다.
- verification docs의 E2E suite 설명을 현재 흐름에 맞췄다.

검증:

- `make e2e-playable` passed
- `GOCACHE=... go test ./...`는 최종 검증에서 재확인한다.

## Learning / Gameplay 판단

좋은 점:

- 커리큘럼은 실무형이다. 이동, 생존, 작은 수정, operator, yank/put, text object, open-line, repeat, search가 플랫폼 엔지니어 기본기에 맞다.
- playpack 단위가 3~7문항으로 작아 과부하가 적다.
- 시나리오는 command를 가리지 않고 “상황 + 조작 목표”를 짧게 유지한다.

위험:

- search는 실무 체감이 큰 기능인데 현재 순서상 늦게 나온다.
- Ex command는 search 경험 후 bridge run에서 다시 다루면 더 자연스럽다.
- strict constraints를 tutorial 기본 경험으로 계속 쓰면 탐색보다 실패가 먼저 올 수 있다.

토론 결론:

- tutorial 순서 자체는 지금 바꾸지 않는다. progress fixture와 E2E가 이미 안정화된 자산이므로 대규모 재배치는 보류한다.
- search 체감은 별도 `Search + Substitute Bridge Run`과 `Incident Run`으로 보완한다.
- Practice/Challenge schema 분리는 보류한다. 우선 기존 schema에서 proactive coaching과 Incident Run의 완화된 constraints로 해결한다.

## 다음 중기 플랜 추천

1. **PLATFORM-REVIEW-001**
   - 저장 변경 없는 review 후보 계산 엔진과 read-only UI.
   - 입력: content library + progress v1 snapshot.
   - 출력: 낮은 grade, 높은 key count, incomplete exercise 추천.
   - `playable` 내부에 바로 넣지 말고 순수 review 패키지로 시작한다.

2. **INCIDENT-RUN-001**
   - 새 Vim 기능 없이 기존 command를 섞는 5~8문항 생존 어드벤처형 mixed playlist.
   - 추천 주제: 장애 로그에서 원인 찾고 설정 핫픽스하기.
   - 후보 흐름: `/error` -> `n` -> `ciw` -> `o` -> `yy/p` -> `:2,3s`.

3. **COACHING-001**
   - schema 변경 없이 첫 실패 전 coaching 노출을 개선한다.
   - strict tutorial은 유지하되 Incident Run은 `max_inputs = optimal + 2~4` 정도로 완화한다.

4. **PROGRESS-SCHEMA-001**
   - 아직 구현 금지.
   - schema v2 승인 packet, 손상 파일/마이그레이션 실패/backup 정책을 먼저 설계한다.

## 사용자 결정 필요 사항

1. 다음 구현을 `PLATFORM-REVIEW-001`로 시작할지, `INCIDENT-RUN-001`로 시작할지 결정한다.
2. Review queue를 메인 첫 화면에 둘지, playlist 완료 후 추천 행동으로 둘지 결정한다.
3. Incident Run을 `tutorial-91` 같은 후속 tutorial로 둘지, 별도 `incident-*` 카테고리로 분리할지 결정한다.

## 보류

- progress schema v2 실제 구현
- daily streak/calendar 저장
- Practice/Challenge schema 분리
- tutorial 순서 재배치
- quote/pair text object, visual selection, macro/register
- search highlight, `?` backward search, regex
