# ExecPlan: Vim curriculum and scenario production harness

Slice-ID: VIM-CURRICULUM-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/README.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/workflows/vim-learning-loops.md
- docs/workflows/scenario-production-harness.md
- docs/gameplay/vim-curriculum-map.md
- docs/gameplay/command-catalog.md
- docs/exec-plans/completed/vim-curriculum-001-scenario-production-harness.md

## 목표

`Vim command -> Exercise -> Scenario` 제작 워크플로우가 얕은 시나리오 생성이나 기능 누락으로 흐르지 않도록, SubAgent 생성/검증 루프까지 포함한 고성능 하네스를 세팅한다.

## 범위

- 포함: 전체 Vim 커리큘럼 맵 초안
- 포함: command/exercise/scenario 각 루프의 입력, 산출물, OK 기준, 실패 조건
- 포함: SubAgent 생성자/검증자 반복 프로토콜
- 포함: 다음 콘텐츠 제작 루프에서 사용할 체크리스트
- 제외: 실제 command cluster 승인
- 제외: Go 구현
- 제외: content loader 구현

## 검증 계획

- SubAgent 초안과 별도 검증 결과를 비교한다.
- 검증자가 지적한 결함을 문서에 반영한다.
- `rg -n "Scenario Production Harness|VIM-CURRICULUM-001|OK Gate|SubAgent" docs`
- `go test ./...`

## 승인 체크

- [x] Vim 기능 범위가 첫 5분보다 넓게 chapter/cluster로 정리되어 있다.
- [x] command-first 원칙을 위반하는 산출물을 거부하는 기준이 있다.
- [x] exercise가 기계 검증 가능하지 않을 때 거부하는 기준이 있다.
- [x] scenario가 exercise 목표를 바꿀 때 거부하는 기준이 있다.
- [x] SubAgent 생성자/검증자 반복 방식이 문서화되어 있다.
- [x] 다음 CONTENT-001 또는 command catalog 확장 루프가 이 하네스를 입력으로 사용할 수 있다.

## 의사결정 로그

- 2026-05-18: Producer SubAgent는 chapter/cluster 커리큘럼과 Producer/Verifier 반복 프로토콜을 제안했다.
- 2026-05-18: Verifier SubAgent는 story-first 역류, draft/playable 누수, `h/k` coverage 누락, replay gate 부재, exact tier oracle 기준 누락을 지적했다.
- 2026-05-18: 지적 사항을 반영해 `coverage_required`, `trained_commands`, `replay_status`, `learning_reinforcement`, `does_not_change`, lifecycle/replay/coverage gates를 문서화했다.
- 2026-05-18: 최종 Verifier SubAgent가 OK 판정을 냈고 blocking must-fix는 없었다.

## 완료 결과

- `docs/workflows/scenario-production-harness.md` 추가
- `docs/gameplay/vim-curriculum-map.md` 추가
- command/exercise/scenario bank schema에 coverage/replay/alignment 필드 보강
- CONTENT-001 acceptance draft에 lifecycle/replay/coverage gate 추가
- SubAgent Producer/Verifier 반복 프로토콜과 OK Gate 문서화
