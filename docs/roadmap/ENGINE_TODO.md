# Engine / Module Todo

> Advimture는 엔진, 미들웨어, 어댑터를 분리해서 만든다. 각 항목은 독립 ExecPlan으로 열고, 테스트와 커밋까지 닫은 뒤 다음 항목으로 넘어간다.

## 진행 원칙

- 한 루프는 하나의 엔진 또는 모듈 계약만 다룬다.
- 새 Vim command 추가와 runtime/UI 연결을 같은 루프에서 섞지 않는다.
- 기존 구현이 새 경계를 방해하면 먼저 legacy/archive 격리를 검토한다.
- TUI 연결 전까지는 unit test와 model-level test를 우선한다.
- TUI로 연결되는 순간 E2E assertion이 부족하면 구현을 멈추고 E2E를 보강한다.

## 순서

| 순서 | ID | 엔진/모듈 | 상태 | 핵심 책임 |
|------|----|-----------|------|-----------|
| 1 | VIM-002 | Vim Engine Foundation | completed | `State + Key -> State + Events`, `h/j/k/l` |
| 2 | VIM-003 | Vim Engine Contract Hardening | completed | 상태 주입, key trace replay, copy/normalize 경계 |
| 3 | VIM-004 | Exercise Runtime Foundation | completed | 문항 세션, 목표 판정, key trace, retry, hint |
| 4 | VIM-005 | Content Schema Foundation | completed | command/exercise/scenario 데이터를 runtime spec으로 변환 |
| 5 | VIM-006 | Grader / Scoring Engine | completed | 정답, 최적 키 대비 평가, 실수/힌트 평가 |
| 6 | VIM-007 | Scenario Runtime | completed | 검증된 exercise를 어드벤처 사건으로 감싸기 |
| 7 | VIM-008 | TUI Adapter | completed | Bubble Tea input/output과 runtime 연결 |
| 8 | VIM-009 | Progress Adapter | completed | runtime 결과를 로컬 저장 포맷으로 변환 |
| 9 | VIM-010 | Neovim Oracle Runner | completed | `exact` tier command를 Neovim 결과와 비교 |
| 10 | VIM-011 | Legacy Archive Pass | completed | 기존 `internal/editor/game/app/ui/data` 유지/격리 결정 |

## 다음 루프 후보

현재 진행 중: 없음

다음 후보는 새 엔진 모듈 완료 상태에 따라 결정한다.

- 실제 Bubble Tea app wiring을 시작하려면 먼저 E2E assertion을 보강한다.
- 첫 playable slice는 `content -> scenario -> tuiadapter -> progressadapter`를 실제 app에 연결한다.
- 새 Vim command를 늘릴 때는 `command catalog -> vimengine -> oracle comparison` 순서로 진행한다.
