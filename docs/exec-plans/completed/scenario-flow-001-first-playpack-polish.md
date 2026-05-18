# ExecPlan: Reusable scenario flow and first playpack polish

Slice-ID: SCENARIO-FLOW-001
Created: 2026-05-18
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/workflows/scenario-flow-workbench.md
- docs/workflows/scenario-production-harness.md
- docs/gameplay/spec.md
- docs/gameplay/exercise-bank.md
- docs/gameplay/scenario-bank.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/scenario-flow-001-first-playpack-polish.md
- docs/exec-plans/completed/scenario-flow-001-first-playpack-polish.md
- content/exercises/first-five-minutes.yaml
- content/playlists/first-five-minutes.yaml
- content/scenarios/first-five-minutes.yaml
- internal/content/loader_test.go
- internal/playable/model_test.go
- test/e2e/**

## 목표

SubAgent 토론 기반 시나리오 제작 워크플로우를 재사용 가능한 운영 문서로 만들고, 그 워크플로우를 첫 5분 플레이팩에 적용해 콘텐츠 품질을 한 번 끌어올린다.

## 영향 도메인

- Gameplay: scenario layer의 briefing, feedback, learning reinforcement 품질을 높인다.
- Workflow: command-first 콘텐츠 제작을 반복 가능한 SubAgent workbench로 정리한다.
- Verification: 기존 playable E2E가 깨지지 않는지 확인한다.

## 수용 기준

- `docs/workflows/scenario-flow-workbench.md`는 입력, 역할, 토론 순서, 산출물, OK 기준, 검증 명령을 포함한다.
- 워크플로우는 `Vim command -> Exercise -> Scenario` 순서를 보존한다.
- SubAgent 토론 결과는 must-fix / should-fix / defer로 통합한다.
- 첫 플레이팩의 scenario copy는 exercise의 target state, optimal keys, allowed keys를 바꾸지 않는다.
- 플레이리스트에 포함된 beat는 approved + replay pass exercise와 approved scenario를 참조한다.
- 변경된 문구가 E2E assertion에 쓰이면 해당 E2E YAML도 함께 동기화한다.
- full playable E2E가 통과한다.

## 범위

- 포함: 재사용 가능한 scenario flow workbench 문서
- 포함: 첫 플레이팩 scenario 문구/문서 polish
- 포함: playlist에 이미 포함된 draft/pending beat의 lifecycle 정합성 보정
- 포함: 필요 시 E2E screen assertion 문구 동기화
- 제외: command cluster 추가
- 제외: exercise target/optimal/allowed key 변경
- 제외: 새 exercise 추가
- 제외: content schema 변경
- 제외: progress 저장 포맷 변경
- 제외: 새 Go 의존성 추가

## 검증 계획

- `go test ./internal/content/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## SubAgent Review 결과

### must-fix

- `normal-motion-basic-002`가 playlist에 포함되어 있지만 exercise/scenario가 `draft`, `replay_status: pending`이었다.
- full E2E는 14개 exercise만 기대하고 있어 playlist와 playable readiness gate가 어긋났다.

### should-fix

- survival, word motion scenario copy가 명확하지만 반복적이고 일부 failure가 정답을 너무 직접 제시했다.
- `:%s/TODO/DONE/g`는 `g` flag 필요성이 데이터로 강하게 드러나지 않는다.

### defer

- `normal-motion-basic`의 `h/k` optimal trace coverage가 아직 없다. 새 exercise 추가가 필요하므로 별도 content/exercise slice로 분리한다.
- `normal-motion-basic` prerequisite과 playlist 시작 순서의 해석은 과거 사용자 결정(`normal-motion-basic -> survival-save-quit -> word-motion-basic`)을 유지하되, 향후 onboarding pacing slice에서 재검토한다.
- 첫 5분 pack의 범위가 넓은 문제는 플레이 가능한 vertical slice 이후 pacing 재설계로 다룬다.

## 완료 결과

- `docs/workflows/scenario-flow-workbench.md`를 추가해 SubAgent 토론 기반 scenario flow를 재사용 가능하게 만들었다.
- `normal-motion-basic-002`를 approved + replay pass + E2E assertion 상태로 승격했다.
- 첫 플레이팩 full E2E를 15개 exercise 경로로 동기화했다.
- survival/word motion scenario copy를 command 학습이 더 잘 드러나도록 다듬었다.
- `h/k` coverage 보강과 첫 pack pacing 재설계는 후속 slice로 남겼다.

## 검증 결과

- `go test ./internal/content/...` 통과
- `go test ./internal/playable/...` 통과
- `go test ./...` 통과
- `make e2e-smoke` 통과
- `make e2e-playable` 통과
- `git diff --check` 통과

## 작업 항목

- [x] current docs/content 기준을 확인한다.
- [x] 재사용 가능한 scenario flow workbench를 문서화한다.
- [x] SubAgent 역할별 독립 평가를 수행한다.
- [x] must-fix / should-fix / defer 결론을 통합한다.
- [x] 첫 플레이팩 scenario copy를 필요한 만큼 개선한다.
- [x] `normal-motion-basic-002` lifecycle 불일치를 해소한다.
- [x] E2E assertion과 문서를 동기화한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
