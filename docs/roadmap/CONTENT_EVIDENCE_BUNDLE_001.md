# Content Evidence Bundle 001 — First Dispatch + Judgment Drill

Status: current
Created: 2026-06-06

## 목적

이 문서는 새 콘텐츠 구현 계획이 아니다. 현재 playable evidence를 사람이 빠르게 읽고 Advimture의 중기 방향을 판단할 수 있도록 묶는 review bundle이다.

이번 bundle의 질문은 하나다.

> Advimture가 "Vim 키 암기"를 넘어, Runbook Dispatch 상황에서 적절한 Vim 도구를 고르는 게임으로 읽히는가?

## 대표 Evidence

| Evidence | 확인하는 흐름 | 판정 |
|----------|---------------|------|
| `playable_ftue_first_five_route` | First Tour에서 tutorial 0~3 초반까지 막힘 없이 이어지고, incident 전환 직전 `Runbook Dispatch` 감각이 보이는가? | required |
| `playable_incident_001_full` | 원인 신호 추적, 상태값 복구, guard 추가, route 재사용, 문제 구간 range 정리로 첫 hotfix runbook이 닫히는가? | required |
| `playable_incident_002_full` | quote 구조를 보존하며 secret/mirror/temp/반복 quote 값을 재동기화하는가? | required |
| `playable_incident_003_full` | contam 표식 검색, visual 격리, 정상 신호 재사용, 잔류 off 전체 승격이 오염 구간 격리로 읽히는가? | required |
| `playable_command_choice_scope` | scope/range/inline/reuse/repeat-change 판단이 command 이름보다 선택 이유로 설명되는가? | required |
| `playable_review_queue` | 성공 후 residual risk와 next dispatch가 반복 동기로 이어지는가? | supporting |

## First Dispatch Reading

`incident 001 -> 002 -> 003`은 아래 세 단계로 읽는다.

1. 원인 신호 추적: 검색과 range 판단으로 첫 relay hotfix를 닫는다.
2. 구조 재동기화: quote 내부 값과 mirror 값을 보존/재사용하며 구조를 맞춘다.
3. 오염 구간 격리: visual selection과 전체 범위 치환으로 보이는 오염과 잔류 상태를 정리한다.

이 bundle은 새 command, 새 schema, 새 progress 저장 포맷 없이 현재 playable loop가 충분한 release-quality skeleton을 갖췄는지 확인한다.

## Judgment Drill Reading

`incident-005-command-choice`는 아래 다섯 판단 질문을 한 route에서 검증한다.

| Beat | 질문 |
|------|------|
| scope choice | 고칠 대상이 글자/단어가 아니라 줄 묶음인가? |
| range choice | 반복 수정보다 전체 범위 치환이 맞는가? |
| inline target choice | delimiter를 보존해야 하는가? |
| reuse choice | 검증된 값을 다시 입력하지 않고 재사용할 수 있는가? |
| repeat-change reuse | 같은 변경을 다시 입력하지 않고 반복할 수 있는가? |

성공/실패 copy는 command 이름보다 선택 이유를 먼저 설명해야 한다. 이 원칙은 `docs/gameplay/command-choice-drills.md`의 Current Playable Mapping과 `content/scenarios/command-choice.yaml`의 scenario copy를 기준으로 검토한다.

## Bundle Exit Criteria

- first tour, first dispatch, judgment drill, review loop evidence가 서로 같은 제품 방향을 가리킨다.
- evidence가 부족할 때 새 콘텐츠를 추가하기보다 먼저 copy, screen contract, action language를 보강한다.
- 새 engine, content schema, progress schema, dependency 변경은 이 bundle 밖이다.
