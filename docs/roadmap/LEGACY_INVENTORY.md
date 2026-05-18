# Legacy Inventory

> 새 엔진 모듈을 기준으로 기존 구현을 어떻게 다룰지 정리한다. LEGACY-001에서 active code path와 충돌하는 옛 구현을 archive했다.

## 원칙

- 기존 구현이 새 엔진 경계를 방해하지 않으면 즉시 이동하지 않는다.
- 새 runtime, scenario, adapter와 책임이 겹치는 순간 legacy/archive 후보로 격리한다.
- `.go` 파일을 `docs/archived`에 그대로 두지 않는다. 필요하면 `.go.txt` 또는 압축 스냅샷으로 보관한다.
- archive 후에는 `go test ./...`가 통과해야 한다.

## 패키지별 상태

| 패키지 | 현재 판단 | 이유 | Archive Trigger |
|--------|-----------|------|-----------------|
| `internal/editor` | archived | 새 `internal/vimengine`과 Vim semantics 책임이 겹친다. | `docs/archived/legacy-code/2026-05-18/internal/editor/*.go.txt` |
| `internal/game` | archived | 기존 tutorial/mission/grader가 새 `runtime/scenario/scoring/playable`과 책임이 겹친다. | `docs/archived/legacy-code/2026-05-18/internal/game/*.go.txt` |
| `internal/app` | keep | 새 playable model을 기본 앱 경로로 감싸는 얇은 Bubble Tea entry다. | 새 multi-screen game loop가 필요할 때 별도 ExecPlan에서 확장 |
| `internal/ui` | archived | 기존 menu/result view가 새 playable/view model 경계를 강제하지 않도록 격리했다. | `docs/archived/legacy-code/2026-05-18/internal/ui/*.go.txt` |
| `internal/data` | archived | 기존 YAML loader/schema는 새 `internal/content` schema와 다르다. | `docs/archived/legacy-code/2026-05-18/internal/data/**` |
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
| `internal/playable` | 첫 playable vertical slice의 Bubble Tea model |
| `internal/e2estate` | E2E app state summary schema |

## Archive 결과

- 2026-05-18: `internal/editor`, `internal/game`, `internal/data`, `internal/ui`를 `docs/archived/legacy-code/2026-05-18/`로 이동했다.
- 2026-05-18: obsolete FTUE scenario `test/e2e/ftue_ctrl_c_quit.yaml`을 `docs/archived/legacy-e2e/2026-05-18/`로 이동했다.
- 2026-05-18: archived Go source는 `*.go.txt`로 보관하여 `go test ./...` 대상에서 제외했다.
