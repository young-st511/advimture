# Content Quality Plan 001 — Release-Quality Content Target

Status: updated
Created: 2026-06-02
Updated: 2026-06-06

## 목적

이 문서는 신규 콘텐츠 구현 계획이 아니다. Advimture가 "출시 가능한 수준"으로 느껴지려면 어느 정도의 콘텐츠 구조가 필요한지 기준을 정하고, 현재 콘텐츠가 그 기준에 얼마나 가까운지 점검한다.

새 YAML 콘텐츠, 새 engine, 새 schema는 이 문서에서 구현하지 않는다.

## Release-Quality Content Target

첫 release-quality baseline의 콘텐츠 목표는 많은 분량이 아니라 아래 구조다.

1. **First Tour**
   - Vim을 처음 만난 플레이어가 모드, 이동, 생존, 단어 이동, 작은 수정을 짧은 episode로 익힌다.
   - 기준: tutorial 0~3 초반까지 막힘 없이 진행되고, 각 episode가 8문항 이하를 유지한다.

2. **Core Toolbelt**
   - operator grammar, yank/put, text object, open line, repeat, search, visual, inline target을 짧은 tutorial로 익힌다.
   - 기준: 각 command cluster는 replay gate와 full route E2E를 통과하고, 첫 화면에서 학습 command가 보인다.

3. **First Dispatch Arc**
   - 이미 배운 command를 조합해 incident 001~003을 하나의 릴레이 기지 복구 arc처럼 느끼게 한다.
   - 기준: 새 command를 가르치지 않고, search/change/yank/substitute/visual을 복구 판단으로 조합한다.

4. **Judgment Drill**
   - command-choice incident로 "아는 명령 중 무엇을 고를지"를 훈련한다.
   - 기준: 성공/실패 피드백이 정답 key보다 선택 이유를 강화한다.

5. **Review Loop**
   - 성공 후 잔류 리스크, 다음 출격, review queue가 반복 동기를 만든다.
   - 기준: progress v1만으로 debrief와 재출격이 설명된다.

## 현재 콘텐츠 판정

현재 콘텐츠는 foundation build 기준으로 release-quality target의 skeleton을 이미 갖고 있다.

| 영역 | 현재 상태 | 판정 |
|------|-----------|------|
| First Tour | tutorial 0~3 초반 canonical route 존재 | 충분 |
| Core Toolbelt | tutorial 4~9, 90~94 구현 및 full E2E 존재 | 충분 |
| First Dispatch Arc | incident 001~003 mixed/structure/visual recovery 존재, copy polish 완료 | 충분 |
| Judgment Drill | incident 005 command-choice 7 beat 존재 | 충분. search-then-scope와 bracket scope 후보는 playable route로 승격 완료 |
| Review Loop | success debrief, residual risk, next dispatch 존재 | 충분 |

현재 부족분은 새 콘텐츠 수가 아니라 **arc framing과 pacing 설명**이다. 즉 다음 콘텐츠 작업은 구현보다 "어떤 순서와 감각으로 플레이어에게 보일지"를 정리하는 것이 우선이다.

## Minimum Release-Quality Content Shape

바로 출시하지 않더라도, "출시 가능한 수준으로 개발 중"이라고 말하려면 첫 playable 콘텐츠는 아래 형태를 갖춰야 한다.

| 단계 | 플레이어 감각 | 현재 구현 | release-quality 판정 |
|------|---------------|-----------|----------------------|
| 1. 첫 조작 | Vim에서 당황하지 않고 움직이고 빠져나간다. | `tutorial-0-movement`, `tutorial-1-survival` | 충분 |
| 2. 첫 효율 | 한 글자씩 움직이는 대신 Vim motion을 고른다. | `tutorial-2-fast-navigation` | 충분 |
| 3. 첫 편집 | Normal/insert/undo/redo로 작은 수정을 닫는다. | `tutorial-3-small-edits` | 충분 |
| 4. 핵심 문법 | operator, yank/put, text object, search, visual, char find를 짧게 익힌다. | tutorial 4~9, 90~94 | 충분 |
| 5. 첫 적용 run | 배운 command를 조합해 하나의 복구 작전을 닫는다. | incident 001~003 | 충분 |
| 6. 선택 판단 | 같은 목표라도 어떤 Vim 도구가 맞는지 고른다. | incident 005 command-choice | 충분, 선택 이유 review 필요 |
| 7. 반복 동기 | 성공 후 잔류 리스크와 다음 출격이 보인다. | review queue/debrief | 충분 |

이 기준에서 현재 콘텐츠 수는 부족하지 않다. 2026-06-06 `CONTENT-ARC-POLISH-001`에서는 새 YAML을 추가하지 않고 incident 001~003의 title/briefing/feedback copy만 다듬어 **원인 신호 추적 -> 구조 재동기화 -> 오염 구간 격리** 흐름을 첫 dispatch arc로 고정했다.

## Evidence 기준

콘텐츠 release-quality 판정은 아래 evidence를 우선한다.

| Evidence | 확인하는 질문 |
|----------|---------------|
| `playable_ftue_first_five_route` final/timeline/app_state | 첫 투어가 tutorial 0~3 초반까지 막힘 없이 이어지는가? |
| `playable_full_first_five_minute` | legacy full route가 현재 episode 구조와 충돌하지 않는가? |
| `playable_incident_001_full` | 첫 incident가 검색 -> 수정 -> guard -> 재사용 -> 범위 정리의 runbook처럼 읽히는가? |
| `playable_incident_002_full` | quote 구조 복구와 mirror 동기화가 한 사건으로 묶이는가? |
| `playable_incident_003_full` | visual selection이 "오염 구간 격리"로 이해되는가? |
| `playable_command_choice_scope` | command-choice가 정답 암기가 아니라 선택 판단을 강화하는가? |
| `playable_review_queue` | 잔류 리스크와 다음 출격이 반복 동기로 이어지는가? |

## Planned Content Arcs

### Arc A. First Tour to First Dispatch

목표: tutorial 0~3에서 incident 001로 넘어갈 때 플레이어가 "연습 문제에서 실제 복구 작전으로 들어간다"는 전환을 이해한다.

완료된 polish:

- Tutorial/incident 전환은 header track의 `Runbook Dispatch`와 다음 action copy에서 첫 dispatch 진입으로 읽히게 한다.
- Incident 001 첫 화면은 새 command 소개가 아니라 "검색으로 원인 위치를 고정하는 첫 복구 작전"으로 시작한다.
- 새 command를 노출하지 않고, hint/failure에서만 command memory를 점진 공개한다.

완료 기준:

- 첫 dispatch 진입 직전 success modal이 `다음 런북`을 단순 다음 문제 이동이 아니라 첫 복구 작전 진입으로 읽히게 한다.
- Incident 001 첫 화면은 새 command 소개가 아니라 "배운 command 조합"이라는 태도를 유지한다.
- 이 전환은 `playable_ftue_first_five_route`와 `playable_incident_001_full` evidence로 spot review한다.

### Arc B. Relay Station 001-003

목표: incident 001~003을 서로 독립된 문제 묶음이 아니라 "릴레이 기지 복구가 단계적으로 깊어진다"는 감각으로 묶는다.

현재 구현된 arc:

| Incident | 역할 | command 조합 | 플레이어 감각 |
|----------|------|--------------|---------------|
| 001: 원인 신호 추적 | 첫 mixed run | `/`, `n`, `ciw`, `o`, `yy/p`, `:2,3s` | 원인 신호를 고정하고, 상태값/guard/정상 route/문제 구간 정리로 첫 야간 핫픽스 runbook을 닫는다. |
| 002: 구조 재동기화 | 구조 값 복구 run | `/`, `ci"`, `yi"` + `P`, `:%s`, `.` | quote 구조를 보존하면서 secret/mirror/temp/반복 quote 값을 재동기화한다. |
| 003: 오염 구간 격리 | 보이는 범위 지정 run | `/`, visual `d`, visual `y/p`, backward visual `d`, `:%s` | contam 위치를 찾고 bad/stale/off 구간을 visual/range 판단으로 격리한다. |

구현 주의:

- 각 incident의 exercise target/optimal keys/constraints는 유지한다.
- 제목/briefing/feedback만 world-frame 기준으로 다듬는다.
- 세 incident 모두 "검색으로 진입 -> 구조/범위 조작 -> 마감 치환" 리듬을 공유한다. 이 반복 구조가 지루한 반복이 아니라 숙련도 상승처럼 느껴지는지 evidence로 확인한다.
- 001~003 이후에는 바로 새 command를 추가하기보다 command-choice로 넘어가 "어떤 도구를 고를지"를 훈련시키는 편이 자연스럽다.

추가 구현 후보는 이 goal에서 열지 않는다. `CONTENT-ARC-POLISH-001`은 기존 scenario copy만 다뤘고, exercise target/optimal keys/constraints는 유지했다.

### Arc C. Judgment Drill Set

목표: incident 005를 기준으로 "Vim을 아는 것"에서 "어떤 Vim 동작을 고를지 판단하는 것"으로 넘어간다.

기획 후보:

- Scope triage: 단어 하나가 아니라 줄 묶음을 선택해야 하는 상황
- Repeat or substitute: 같은 literal이 반복될 때 전체 범위 치환을 선택하는 상황
- Inline target range: delimiter를 보존해야 하는 상황
- Quote value reuse: 검증된 값을 직접 재입력하지 않는 상황
- Repeat change reuse: 같은 변경을 두 번 입력하지 않는 상황

구현 주의:

- 새 engine 없이 기존 constraints로 의도 선택을 고정한다.
- 실패 copy는 "왜 이 command가 상황에 맞지 않는지"를 설명한다.

현재 구현된 judgment set:

| Beat | 판단 질문 | 의도 선택 |
|------|-----------|-----------|
| scope triage | 값/단어가 아니라 줄 묶음인가? | linewise visual delete |
| repeat or substitute | 반복 변경보다 전체 치환이 맞는가? | `:%s` |
| inline target range | delimiter를 보존해야 하는가? | `ct,` |
| quote value reuse | 검증된 quote 값을 다시 치지 않고 재사용할 수 있는가? | `yi"` + `P` |
| repeat change reuse | 같은 변경을 다시 입력하지 않아도 되는가? | `.` |

완료 기준:

- 성공/실패 copy가 "왜 이 선택이 맞는지"를 설명한다.
- E2E는 final buffer뿐 아니라 key trace/app_state evidence를 남긴다.
- 새 command나 schema 없이 기존 constraints로 선택 의도를 고정한다.

### Arc D. Next Engine Candidate

목표: 콘텐츠가 실제로 막히는 지점이 확인될 때만 새 engine을 연다.

후보:

- `quote-pair-hardening`: `ci(`, `ci{`
- `search-hardening`: `?`, regex, highlight, history 중 하나
- `visual-advanced`: visual block, multi-line charwise, count/register prefix

판정 기준:

- 현재 콘텐츠 arc가 기존 engine으로 충분하면 열지 않는다.
- 새 engine을 열면 command catalog -> engine test -> runtime/content -> E2E 순서를 따른다.
- `?` backward search는 현재 hint key와 충돌하므로 UX 결정이 선행되어야 한다.

현재 판정:

- release-quality baseline을 위해 지금 새 engine은 필요하지 않다.
- `ci(`/`ci{`는 실용적이지만, 현재 quote-pair와 visual/char-find 적용 run만으로 first release-quality skeleton은 충분하다.
- search `?`는 hint key와 충돌하므로 단순 engine 문제가 아니라 UX/input contract 문제다.
- visual advanced는 blast radius가 커서 content arc가 실제로 막힐 때만 연다.

## Exit Criteria For Content Planning

- First Tour, Core Toolbelt, First Dispatch, Judgment Drill, Review Loop의 역할이 문서화되어 있다.
- 현재 구현이 어디까지 충족하는지 명시되어 있다.
- 구현하지 않을 미래 콘텐츠 후보가 command/arc 단위로 분리되어 있다.
- 새 콘텐츠 구현이 필요하면 별도 ExecPlan을 연다.

## Next Content Decision

다음 콘텐츠 관련 작업은 **구현**이 아니라 아래 중 하나를 고른다.

1. `CONTENT-ARC-POLISH-001`: incident 001~003의 기존 scenario title/briefing/feedback을 world-frame 기준으로 spot polish한다.
2. `JUDGMENT-DRILL-REVIEW-001`: incident 005의 성공/실패 copy가 선택 이유를 충분히 설명하는지 evidence 기준으로 점검한다.
3. `CONTENT-EVIDENCE-BUNDLE-001`: first tour, incident 001~003, command-choice, review queue evidence를 한 번에 읽는 review bundle을 정리한다.

새 tutorial/incident YAML 구현은 이 goal의 현재 slice에서는 열지 않는다.
