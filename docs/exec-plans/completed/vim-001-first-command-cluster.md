# ExecPlan: 첫 command cluster 설계

Slice-ID: VIM-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
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

- [x] command cluster가 하나의 학습 목표로 묶인다.
- [x] 실무 유용성이 명확하다.
- [x] 선행 관계가 명시되어 있다.
- [x] 첫 5분 플레이 루프에 포함할 cluster와 뒤로 미룰 cluster를 구분한다.
- [x] 사람이 승인하여 `status: approved`로 바꿨다.

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

- 2026-05-18: 첫 playable 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`로 결정했다.
- 2026-05-18: CONTENT-001은 YAML 우선, repo root `content/`, draft 파일 허용, planned/playable 분리를 기본값으로 삼는다.
- 2026-05-18: coverage는 cluster 목적별 `coverage_required`를 기준으로 판단한다. 초반 기본 이동은 모든 방향 command coverage가 필요하지만, 후반 범용 이동은 억지로 모든 command를 한 루프에 넣지 않는다.

## 미해결 질문

- 없음.

## Approval Packet

추천 판정:

| Cluster | 추천 상태 | 이유 | 구현 주의 |
|---------|-----------|------|-----------|
| `normal-motion-basic` | approve now | 현재 engine/playable에서 구현 가능하고 첫 성공 경험에 적합하다. | 승인 전 `h`, `k` optimal trace exercise를 CONTENT/EXERCISE 루프에서 보강한다. |
| `survival-save-quit` | approve as planned | Vim 입문자의 불안감을 줄이는 필수 생존 cluster다. | command-line mode, app exit, autosave semantics를 SURVIVAL-001에서 분리 구현한다. |
| `word-motion-basic` | approve as planned | `hjkl` 반복 대비 효율 체감을 주며 Vim 문법의 다음 단계로 이어진다. | `w/b/e` word boundary oracle fixture를 VIM-012에서 먼저 고정한다. |

승인하지 않고 보류할 경우:

- `normal-motion-basic`만 approval 후보로 두고 CONTENT-001은 implemented draft fixture만 로드한다.
- `survival-save-quit`, `word-motion-basic`은 curriculum backlog에 남기되 playable playlist에는 넣지 않는다.

## 완료 결과

- `normal-motion-basic`, `survival-save-quit`, `word-motion-basic`을 `approved`로 승격했다.
- `survival-save-quit`, `word-motion-basic`은 `engine_support: planned`를 유지해 playable 연결과 분리했다.
- 후속 CONTENT/EXERCISE 루프에서 `normal-motion-basic`의 `h`, `k` coverage gap을 보강해야 한다.
