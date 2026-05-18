# 0004. Tutorial pacing and exercise constraints

Status: accepted
Date: 2026-05-18

## Context

첫 playable pack은 17개 exercise를 한 번에 완주할 수 있다. 이는 vertical slice와 E2E 검증에는 유용하지만, 플레이어가 실제로 학습할 때는 한 playlist 안에서 다루는 command 범위가 넓다.

또한 현재 exercise는 목표 cursor/buffer state에 도달하면 성공하므로, 의도한 Vim command를 쓰지 않고 우회해도 성공할 수 있다. Vim 학습 게임으로 완성하려면 문항별 입력 제약, 최대 입력 수, 실패 후 재시도 UX가 필요하다.

## Decision

- 초반은 본격 어드벤처가 아니라 짧은 튜토리얼 에피소드 묶음으로 운영한다.
- 각 튜토리얼 playlist는 기본적으로 8문항 이하로 유지한다.
- 초반 튜토리얼은 “첫 투어” 느낌으로 Vim command를 넓게 맛보게 한다.
- 중후반부터는 생존 어드벤처, 탐험, 위기 해결 비중을 높인다.
- `:s`, `:%s`, range substitute는 초반 튜토리얼에서 빼고, 중반에 등장하는 별도 고급 튜토리얼로 분리한다.
- 최대 입력 수를 넘으면 즉시 실패하고 재시도를 유도한다.
- 금지 입력이나 금지 우회 전략을 쓰면 즉시 실패한다.
- 새 command를 처음 배우는 문항에서는 의도 command 외 우회 입력을 불가능하게 설계한다.
- 실패 횟수는 기본 무제한으로 둔다. 단, 후반 콘텐츠를 위해 `attempt_limit` 설정 여지는 남긴다.
- 실패 후 재시도는 `r`과 `enter`를 모두 허용한다.
- 초반 튜토리얼 코칭은 개념 힌트 중심으로 한다. 단, 새 command 첫 소개 시에는 command 의미를 명시한다.
- Scenario 제작 워크플로우에는 문항별 허용/금지 입력과 우회 전략을 설계하는 Constraint Designer 역할을 추가한다.

## Consequences

- 현재 17개 exercise first pack은 구현 검증용 vertical slice로 유지할 수 있지만, 실제 학습 UX는 여러 짧은 tutorial playlist로 재구성해야 한다.
- 다음 구현 slice는 playlist split, exercise constraint schema, failure/retry UI, command-intent scoring 중 하나를 명시적으로 선택해야 한다.
- Scenario copy만 좋아지는 것으로는 문항 완성으로 보지 않는다. Exercise constraint와 failure coaching까지 함께 설계해야 한다.
- Ex substitute 계열은 초반 맛보기에서 빠지고, 중반 고급 튜토리얼의 독립된 학습 단위가 된다.
