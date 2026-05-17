# Legacy Inventory

> 새 엔진 모듈을 기준으로 기존 구현을 어떻게 다룰지 정리한다. 실제 archive는 충돌이 발생하는 slice에서 작게 수행한다.

## 원칙

- 기존 구현이 새 엔진 경계를 방해하지 않으면 즉시 이동하지 않는다.
- 새 runtime, scenario, adapter와 책임이 겹치는 순간 legacy/archive 후보로 격리한다.
- `.go` 파일을 `docs/archived`에 그대로 두지 않는다. 필요하면 `.go.txt` 또는 압축 스냅샷으로 보관한다.
- archive 후에는 `go test ./...`가 통과해야 한다.

## 패키지별 상태

| 패키지 | 현재 판단 | 이유 | Archive Trigger |
|--------|-----------|------|-----------------|
| `internal/editor` | archive 후보 | 새 `internal/vimengine`과 Vim semantics 책임이 겹친다. | 새 app/runtime이 `vimengine`을 사용하기 시작하고 기존 editor import가 제거될 때 |
| `internal/game` | archive 후보 | 기존 tutorial/mission/grader가 새 `runtime/scenario/scoring`과 책임이 겹친다. | 첫 새 scenario runtime 기반 mission이 app에 연결될 때 |
| `internal/app` | review | 기존 Bubble Tea app entry와 화면 전환을 포함한다. | 새 TUI adapter를 Bubble Tea model에 연결할 때 |
| `internal/ui` | review | 기존 menu/result view는 재사용 가능성이 있다. | 새 view model과 맞지 않는 UI state를 강제할 때 |
| `internal/data` | review | 기존 YAML loader는 참고 가능하지만 새 content schema와 다르다. | VIM-005 spec용 loader를 만들 때 기존 schema와 충돌하면 archive |
| `internal/progress` | keep with caution | 기존 저장 포맷과 테스트가 있고, 새 progress adapter가 이를 감싼다. | save schema 변경이 필요해지는 별도 승인 slice |

## 새 모듈 기준

| 새 패키지 | 책임 |
|----------|------|
| `internal/vimengine` | Vim-like 상태 전이 |
| `internal/runtime` | exercise session 진행 |
| `internal/content` | exercise spec validation/compile |
| `internal/scoring` | pass/grade/efficiency 평가 |
| `internal/scenario` | adventure scenario orchestration |
| `internal/tuiadapter` | TUI input/action, state/view model 변환 |
| `internal/progressadapter` | scenario result를 progress model로 변환 |
| `internal/vimoracle` | optional oracle comparison |

## 다음 Archive 후보

1. `internal/editor`
2. `internal/game`
3. 기존 `internal/data` YAML schema

단, 실제 이동은 새 app wiring 또는 content loader slice에서 import graph를 확인한 뒤 진행한다.
