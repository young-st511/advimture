# Midterm Todo

> 새 엔진 모듈을 실제 게임으로 연결하기 위한 중기 순서다. 각 항목은 독립 ExecPlan으로 열고, 테스트와 커밋까지 닫은 뒤 다음 항목으로 넘어간다.

## 순서

| 순서 | ID | 상태 | 목표 |
|------|----|------|------|
| 1 | E2E-003 | completed | app state summary 기반 buffer/cursor/mode/status/score/progress assertion |
| 2 | PLAY-001 | planned | 첫 playable vertical slice: `hjkl` 목표 이동, 성공, 점수, progress 반영 |
| 3 | LEGACY-001 | planned | 새 playable path가 canonical이 된 뒤 기존 editor/game 격리 |
| 4 | CONTENT-001 | planned | 새 content schema용 YAML/JSON loader |
| 5 | VIM-012 | planned | 다음 command cluster: `w/b/e` word motion |
| 6 | GAMELOOP-001 | planned | 반복 학습 루프, unlock, 피드백 구조 |

## 진행 원칙

- E2E assertion이 부족하면 app wiring을 멈추고 E2E를 먼저 보강한다.
- playable slice는 한 문제만 end-to-end로 완성한다.
- 기존 구현 archive는 새 path가 동작하고 import graph에서 빠진 뒤 진행한다.
- 새 Vim command는 `command catalog -> vimengine -> oracle comparison -> exercise` 순서로 추가한다.

## 현재 판단

다음 작업은 **PLAY-001 첫 playable vertical slice**다. 이 루프에서는 앱이 `ADVIMTURE_E2E=1`일 때 state summary를 쓰고, E2E runner가 그 파일로 성공 상태를 검증한다.
