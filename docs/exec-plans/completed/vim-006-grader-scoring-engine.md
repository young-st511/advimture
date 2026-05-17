# ExecPlan: Grader / scoring engine

Slice-ID: VIM-006
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/vim-006-grader-scoring-engine.md
- internal/scoring/**

## 목표

Exercise Runtime의 성공/진행 상태와 Content Schema의 expected key metadata를 받아, 순수한 평가 결과를 산출하는 scoring 엔진을 만든다.

## 범위

- 포함: `internal/scoring` 패키지 생성
- 포함: pass/fail 판정
- 포함: expected keys와 실제 key trace 비교
- 포함: 효율성 ratio와 grade 산정
- 포함: attempt/hint penalty
- 제외: runtime 상태 변경
- 제외: progress 저장
- 제외: UI 표시 문구
- 제외: YAML loader

## 연결된 엔진

- 입력: runtime session state, content expected keys, hint/attempt metadata
- 출력: score result
- 다음 연결 후보: Scenario Runtime, Progress Adapter

## 구현 계획

1. `Input`, `Result`, `Grade` 타입을 만든다.
2. `Evaluate(Input)`은 순수 함수로 유지한다.
3. `StatusSucceeded`가 아니면 pass=false로 평가한다.
4. expected keys가 있으면 key trace와 exact match를 계산한다.
5. expected key 수 대비 실제 key 수로 efficiency ratio를 계산한다.
6. hint와 retry는 감점 요소로 반영한다.

## 검증 계획

- `go test ./internal/scoring`
- `go test ./...`
- 변경 종료 전 `git diff -- docs/roadmap/ENGINE_TODO.md docs/roadmap/PROGRAM.md docs/exec-plans/active/vim-006-grader-scoring-engine.md internal/scoring`

## E2E Evidence

이번 slice는 TUI runtime에 연결하지 않으므로 TUI E2E는 필수 evidence가 아니다.

## 승인 체크

- [x] succeeded runtime state만 pass 처리된다.
- [x] exact expected keys match를 계산한다.
- [x] extra key 입력은 efficiency와 grade에 반영된다.
- [x] hint와 retry penalty가 grade에 반영된다.
- [x] scoring 엔진은 runtime 상태를 변경하지 않는다.
- [x] 전체 테스트가 통과한다.

## 의사결정 로그

- 2026-05-18: scoring은 runtime을 조작하지 않는 순수 함수로 설계했다. progress 저장과 UI 표시 정책은 이후 adapter에서 연결한다.
- 2026-05-18: expected keys가 없으면 exact match는 true로 간주한다. content가 아직 최적 키를 제공하지 않는 exercise도 성공 상태를 평가할 수 있게 하기 위함이다.
