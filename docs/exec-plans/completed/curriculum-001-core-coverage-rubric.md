# ExecPlan: Core Vim coverage rubric

Slice-ID: CURRICULUM-001
Created: 2026-05-19
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/gameplay/spec.md
- docs/gameplay/vim-curriculum-map.md
- docs/gameplay/command-catalog.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/ENGINE_TODO.md
- docs/exec-plans/active/curriculum-001-core-coverage-rubric.md
- docs/exec-plans/completed/curriculum-001-core-coverage-rubric.md

## 목표

다음 콘텐츠와 엔진 구현이 흔들리지 않도록 Vim 핵심 학습 범위를 chapter, command cluster, engine 필요도, playpack 우선순위로 정리한다.

## 영향 도메인

- Gameplay: 어떤 Vim 능력을 어떤 순서로 가르칠지 고정한다.
- Content: 다음 playpack 후보 command cluster를 정한다.
- Engine: 다음 구현 gap 분석의 입력을 만든다.
- Scenario: 중반 생존 어드벤처로 넘어가기 전 학습 순서를 보존한다.

## 수용 기준

- `docs/gameplay/vim-curriculum-map.md`가 현재 구현 상태와 맞는다.
- Chapter 0-4의 핵심 command cluster가 `foundation`, `next`, `later` 중 하나로 분류된다.
- 다음 playpack 후보는 6~8문항으로 나눌 수 있는 command cluster 묶음이어야 한다.
- 다음 playpack 후보는 “작은 수정” 중심이어야 하며 이동 복습을 주목표로 삼지 않는다.
- 각 next cluster는 engine support 상태와 oracle 필요 여부를 가진다.
- `docs/roadmap/ENGINE_TODO.md`가 ENGINE-GAP-001 입력으로 쓸 수 있는 후보를 가진다.

## 범위

- 포함: curriculum map 정합화
- 포함: command catalog draft cluster 정리
- 포함: engine gap 후보 문서화
- 제외: Go engine 구현
- 제외: YAML exercise 구현
- 제외: 시나리오 본문 작성

## 검증 계획

- `rg "CURRICULUM-001|ENGINE-GAP-001|PLAYPACK-002|SCENARIO-TONE-001" docs/roadmap docs/exec-plans`
- `rg "Next Playpack Candidate|Coverage Rubric|Known Coverage Gaps" docs/gameplay/vim-curriculum-map.md`
- `go test ./...`
- `git diff --check`

## 검증 결과

- `rg "CURRICULUM-001|ENGINE-GAP-001|PLAYPACK-002|SCENARIO-TONE-001" docs/roadmap docs/exec-plans`: pass
- `rg "Next Playpack Candidate|Coverage Rubric|Known Coverage Gaps" docs/gameplay/vim-curriculum-map.md`: pass
- `go test ./...`: pass
- `git diff --check`: pass

## 작업 항목

- [x] 새 중기 플랜을 TODO/PROGRAM에 반영한다.
- [x] curriculum map의 chapter/coverage/gap 정보를 현재 구현 상태에 맞춘다.
- [x] command catalog에 다음 cluster draft를 추가한다.
- [x] ENGINE_TODO에 gap 분석 후보를 정리한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
