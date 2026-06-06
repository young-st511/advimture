# World Frame — Remote Recovery Bureau

> Advimture의 중반 이후 scenario skinning 기준이다. 이 문서는 command catalog, exercise, target state를 대체하지 않는다. 모든 콘텐츠는 계속 `Vim command -> Exercise -> Scenario` 순서로 만든다.

## 결론

세계관 기본 프레임은 **원격 시설 복구국 / Runbook Dispatch**로 간다.

플레이어는 전투원이나 탐험가가 아니라, 낡은 원격 시설의 장애 runbook을 Vim으로 복구하는 **원격 복구 오퍼레이터**다. 각 incident는 “시설 하나, 콘솔 하나, 장애 하나”를 닫는 짧은 복구 작전으로 다룬다.

개별 사건에는 **침묵한 릴레이 기지** 감각을 얇게 섞는다. 즉 “릴레이”, “신호”, “폭풍권”, “잔류 리스크” 같은 명사는 쓸 수 있지만, 장대한 SF lore는 쓰지 않는다.

## 왜 이 프레임인가

- Vim command 학습을 가리지 않는다.
- 현재 incident의 로그, 설정, route, secret, mirror, temp/live 재료를 거의 그대로 살린다.
- review queue, best record, key count를 “잔류 리스크”, “최단 복구 기록”, “재점검 대상”으로 자연스럽게 포장할 수 있다.
- visual mode 이후에도 “오염 구역 지정”, “정상 신호 복사”, “범위 정리” 같은 조작 동사로 확장하기 좋다.

## 운영 원칙

- 세계관은 항상 `Vim command -> Exercise -> Scenario` 뒤에 붙인다.
- 브리핑은 “상황 1문장 + Vim 조작 목표 1문장”을 넘기지 않는다.
- 플레이어 역할은 원격 복구 오퍼레이터다.
- 사건 단위는 “시설 하나, 콘솔 하나, 장애 하나”로 제한한다.
- SF 감각은 명사 한두 개로만 쓴다.
- 성공 피드백은 세계관 감탄보다 어떤 Vim 조작이 좋았는지 먼저 말한다.
- 실패 피드백은 복구 힌트와 학습 의도를 우선한다.
- replay/review UI는 저장 포맷 변경 없이 “잔류 리스크”, “재점검”, “최단 복구 기록” 언어를 사용할 수 있다.

## Release-Quality 기준

출시 가능한 품질의 세계관은 lore 양이 아니라 일관성으로 판단한다.

- 처음 보는 플레이어가 1분 안에 "나는 원격 복구 오퍼레이터이고, Vim으로 runbook을 복구한다"를 이해해야 한다.
- tutorial은 첫 투어로 두되, incident로 넘어갈 때 `Runbook Dispatch`에 들어간다는 감각을 짧게 만든다.
- incident는 "문제 1, 2, 3"이 아니라 "복구 작전의 단계"처럼 읽혀야 한다.
- 성공 피드백은 세계관 감탄보다 어떤 Vim 조작이 복구에 기여했는지 먼저 말한다.
- 실패 피드백은 플레이어를 벌주기보다 다음 복구 선택을 알려준다.
- 장대한 배경 설명, 세력 관계, NPC drama는 현재 scope 밖이다.

현재 판정은 `docs/roadmap/PLAYABLE_QUALITY_BASELINE_2026-06-02.md`와 `docs/roadmap/CONTENT_QUALITY_PLAN_001.md`를 따른다.

## Command Mapping

| Vim cluster | 복구 동사 | 짧은 scenario 재료 |
|-------------|-----------|--------------------|
| `search-basic` | 신호 추적 | error, timeout, secret 위치 찾기 |
| `change-with-motion` | 상태값 복구 | down -> up, stale -> fresh |
| `text-object-quote-pair` | 구조 값 복구 | quoted secret, mirror, token |
| `yank-put-basic` | 정상값 재사용 | 검증된 route, 정상 신호, mirror 동기화 |
| `open-line-edit` | guard 삽입 | 기존 줄 아래 보호 설정 추가 |
| `vim-ex-command-substitute` | 범위 정리 | temp -> live, error -> ok |
| `repeat-last-change` | 반복 패치 | 같은 구조의 quoted value 승격 |
| `visual-char-line` | 오염 구역 지정 | 선택한 구간 삭제/복사/정리 |

## Incident Reframe

### Incident 001

권장 제목: **릴레이 기지 001: 원인 신호 추적**

흐름:

1. `/error`로 원인 신호 위치를 잡는다.
2. `n`으로 다음 timeout 신호를 따라간다.
3. `ciw`로 상태값을 복구한다.
4. `o`로 guard 줄을 추가한다.
5. `yy/p`로 정상 route를 재사용한다.
6. `:2,3s`로 문제 구간만 정리한다.

핵심 감각은 “여섯 문제”가 아니라 “원인 신호를 고정하고, 확인한 relay route를 단계적으로 복구하는 6단계 runbook”이다. 성공/실패 피드백은 `/`, `n`, `ciw`, `o`, `yy/p`, range substitute가 왜 맞는 선택인지 먼저 말한다.

### Incident 002

권장 제목: **릴레이 기지 002: 구조 재동기화**

흐름:

1. `/secret`으로 대상 설정을 찾는다.
2. `ci"`로 stale 값을 fresh로 고친다.
3. `yi"` + `P`로 mirror 값을 동기화한다.
4. `:%s`로 temp 상태를 live로 승격한다.
5. `.`으로 같은 quote 패치를 반복한다.

핵심 감각은 “quote 구조를 보존하면서 secret, mirror, temp/live, 반복 quote 값을 순서대로 재동기화하는 복구 작전”이다.

### Incident 003

권장 제목: **릴레이 기지 003: 오염 구간 격리**

흐름:

1. `/contam`으로 오염 표식 위치를 잡는다.
2. visual `d`로 route 안의 bad 구간만 제거한다.
3. visual `y`와 `p`로 정상 ok 신호를 mirror에 전송한다.
4. backward visual `d`로 뒤쪽 stale tail만 제거한다.
5. `:%s`로 잔류 off 상태를 on으로 승격한다.

핵심 감각은 “contam 표식을 검색으로 고정하고, 보이는 오염 구간을 지정하고, 정상 신호를 재사용한 뒤, 전체 상태를 마감하는 격리 runbook”이다.

## First Dispatch Arc Evidence

`incident 001 -> 002 -> 003`은 첫 대표 복구 arc로 읽힌다.

| 순서 | 역할 | Evidence |
|------|------|----------|
| 001 | 원인 신호 추적 | `playable_incident_001_full` final/timeline이 검색, 상태 복구, guard 추가, route 재사용, 문제 구간 range 정리를 보여준다. |
| 002 | 구조 재동기화 | `playable_incident_002_full` final/timeline이 quote 내부 값 복구, mirror 재사용, 전체 상태 승격, 반복 quote 패치를 보여준다. |
| 003 | 오염 구간 격리 | `playable_incident_003_full` final/timeline이 visual 범위 지정, 정상 신호 재사용, 잔류 상태 전체 승격을 보여준다. |

## 금지

- 장대한 릴레이 기지 역사, 세력, 인물 관계 설명
- 체력, 전투, 인벤토리 중심 진행
- 티켓 접수/처리/종결 같은 관료적 표현 반복
- command를 숨기는 미스터리
- 정답을 알아도 실패하는 숨은 조건
- scenario가 target state, optimal keys, constraints를 바꾸는 것
