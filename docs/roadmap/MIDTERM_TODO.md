# Midterm Todo

> 새 엔진 모듈을 실제 게임으로 연결하기 위한 중기 순서다. 각 항목은 독립 ExecPlan으로 열고, 테스트와 커밋까지 닫은 뒤 다음 항목으로 넘어간다.

## 실행 보드

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | E2E-003 | completed | app state summary 기반 buffer/cursor/mode/status/score/progress assertion |
| 2 | PLAY-001 | completed | 첫 playable vertical slice: `hjkl` 목표 이동, 성공, 점수, progress 반영 |
| 3 | LEGACY-001 | completed | 새 playable path가 canonical이 된 뒤 기존 editor/game 격리 |
| 4 | SCENARIO-001 | completed | 첫 5분 시나리오 워크플로우로 content loader 요구사항 발견 |
| 5 | VIM-CURRICULUM-001 | completed | Vim curriculum map과 scenario production harness |
| 6 | CONTENT-001 | planned | 새 content schema용 YAML/JSON loader |
| 7 | VIM-012 | planned | 다음 command cluster: `w/b/e` word motion |
| 8 | SURVIVAL-001 | planned | `esc`, `:q!`, `:wq` 생존/종료 루프 |
| 9 | NAV-001 | planned | 후반 navigation 확장: `gg`, `G`, line/file motion |
| 10 | EXCMD-001 | planned | `:` 명령어, substitute, range command 기반 |
| 11 | GAMELOOP-001 | planned | 반복 학습 루프, unlock, 자동 저장, 피드백 구조 |

## 루프별 출구 조건

| ID | 입구 조건 | 필수 산출물 | 검증 | 품질 저하 방지 |
|----|----------|-------------|------|---------------|
| VIM-001 | curriculum defaults 결정 완료 | 첫 5분 command cluster approval packet | `rg`, `go test ./...` | completed: 사용자 승인 후 3개 cluster를 `approved`로 승격했다. |
| CONTENT-001 | VIM-001 승인 또는 approval packet 준비 | root `content/` YAML fixture, loader, validator | `go test ./internal/content/...`, `go test ./...` | draft/planned 콘텐츠는 로드하되 playable 후보에서 제외한다. |
| PLAY-002 | CONTENT-001 통과 | hardcoded playable 제거, file-backed playable | `go test ./internal/playable/...`, `make e2e-smoke` | screen assertion만 믿지 않고 app state/progress를 함께 본다. |
| CONTENT-002 | file-backed playable 통과 | replay/coverage validator | content replay tests, coverage report | `coverage_required`와 `replay_status` 없이는 approved/playable 승격 금지. |
| VIM-012 | word-motion cluster approved | `w/b/e` engine, oracle fixtures | `go test ./internal/vimengine/...`, oracle comparison | word boundary, 공백, 문장부호 edge case를 먼저 고정한다. |
| EXERCISE-001 | VIM-012 통과 | word motion exercise set | replay validator, verifier OK | `w`, `b`, `e`가 각각 optimal trace에 등장해야 한다. |
| SURVIVAL-001 | command-line scope 승인 | `esc`, `:q!`, `:wq` runtime/app semantics | unit tests, E2E smoke | app exit와 mission success를 혼동하지 않는다. |
| NAV-001 | movement fundamentals 통과 | `gg`, `G`, line/file motion clusters | vimengine/oracle tests | 후반 범용 이동은 `coverage_required` 중심으로 범위를 좁힌다. |
| EXCMD-001 | command-line engine 기반 | `:` 명령어, substitute, range command | parser/runtime tests | Ex command는 편집 engine과 scenario success를 분리한다. |
| GAMELOOP-001 | file-backed multi exercise 가능 | playlist/unlock/retry/hint/autosave loop | E2E first-5-minute | 자동 저장은 기존 `.advimture` progress 경계를 유지한다. |

## 진행 원칙

- E2E assertion이 부족하면 app wiring을 멈추고 E2E를 먼저 보강한다.
- playable slice는 한 문제만 end-to-end로 완성한다.
- 기존 구현 archive는 새 path가 동작하고 import graph에서 빠진 뒤 진행한다.
- 새 Vim command는 `command catalog -> vimengine -> oracle comparison -> exercise` 순서로 추가한다.
- 각 루프는 ExecPlan을 열고, 해당 루프의 검증과 커밋까지 닫은 뒤 다음 항목으로 넘어간다.
- 후반부 루프일수록 `scenario-production-harness.md`의 Verifier OK, replay gate, coverage gate를 생략하지 않는다.

## 현재 판단

현재 active slice는 없다. 다음 작업은 **CONTENT-001 새 content schema용 YAML loader**이며, VIM-001에서 승인된 `normal-motion-basic`, `survival-save-quit`, `word-motion-basic`을 입력으로 사용한다.
