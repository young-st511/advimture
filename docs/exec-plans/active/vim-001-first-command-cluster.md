# ExecPlan: 첫 command cluster 설계

Slice-ID: VIM-001
Created: 2026-05-18
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/gameplay/command-catalog.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/gameplay/spec.md
- docs/workflows/vim-learning-loops.md

## 목표

Advimture의 첫 5분 플레이 루프에 들어갈 Vim command cluster를 선정한다. 이 slice의 핵심은 스토리나 화면 구현이 아니라, 플레이어가 가장 먼저 배워야 할 Vim 조작의 순서와 학습 목표를 명확히 하는 것이다.

## 범위

- 포함: `survival-save-quit`, `normal-motion-basic`, `word-motion-basic` 후보 검토와 승인 기준 정리
- 포함: command cluster별 선행 관계, 실무 유용성, common mistakes 보강
- 제외: 실제 게임 데이터 구현
- 제외: full exercise/scenario 구현
- 제외: E2E schema 확장

## Loop 1 — Vim Command

### 입력

- 후보 command: `esc`, `:q!`, `:wq`, `h`, `j`, `k`, `l`, `w`, `b`, `e`
- 선행 command: 없음
- 플레이어 숙련도: Vim 입문자

### 산출물

`docs/gameplay/command-catalog.md`의 draft command cluster를 승인 가능한 수준으로 다듬는다.

### 승인 체크

- [ ] command cluster가 하나의 학습 목표로 묶인다.
- [ ] 실무 유용성이 명확하다.
- [ ] 선행 관계가 명시되어 있다.
- [ ] 첫 5분 플레이 루프에 포함할 cluster와 뒤로 미룰 cluster를 구분한다.
- [ ] 사람이 승인하여 `status: approved`로 바꿨다.

## Loop 2 — Exercise

이번 slice에서는 exercise를 구현하지 않는다. 단, command cluster 판단에 필요한 최소 예시는 `docs/gameplay/exercise-bank.md`의 draft 항목으로 참고한다.

## Loop 3 — Scenario

이번 slice에서는 scenario를 구현하지 않는다. 단, command cluster 판단에 필요한 최소 예시는 `docs/gameplay/scenario-bank.md`의 draft 항목으로 참고한다.

## 구현 계획

- 코드 변경 없음
- 문서 변경 중심
- 승인 후 다음 slice에서 exercise 설계로 이동

## 검증 계획

- `rg -n "status: draft|status: approved" docs/gameplay`
- `go test ./...`

## E2E Evidence

- 이번 slice는 문항 설계 slice라 E2E evidence를 요구하지 않는다.

## 의사결정 로그

(작업 중 추가)

## 미해결 질문

- 첫 5분에서 `word-motion-basic`까지 포함할지, `survival-save-quit` + `normal-motion-basic`까지만 포함할지 결정 필요.
