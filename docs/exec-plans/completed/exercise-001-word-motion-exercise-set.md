# ExecPlan: Word motion exercise set

Slice-ID: EXERCISE-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- content/exercises/**
- content/scenarios/**
- content/playlists/**
- internal/content/**
- docs/gameplay/spec.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/exercise-001-word-motion-exercise-set.md

## 목표

VIM-012에서 구현한 `w/b/e` word motion을 실제 학습 콘텐츠로 승격한다. 시나리오가 아니라 command coverage에서 출발해, `w`, `b`, `e`가 각각 optimal trace에 등장하는 exercise set을 만든다.

## 영향 도메인

- Gameplay: 설계 순서는 `Vim command -> Exercise -> Scenario`를 유지한다.
- Content: 기존 YAML ID/파일명은 변경하지 않고, 새 exercise/scenario ID만 추가 또는 기존 draft를 승인한다.
- Verification: replay gate와 coverage report가 콘텐츠 품질 게이트 역할을 한다.

## 수용 기준

- `word-motion-basic`의 `coverage_required`인 `w`, `b`, `e`가 approved + implemented exercise optimal trace에 모두 등장한다.
- word motion exercise는 모두 `replay_status: pass`이며 loader replay gate를 통과한다.
- 각 exercise는 player-facing goal, target state, optimal keys, allowed keys, hints, E2E assertions를 가진다.
- 각 approved + implemented exercise는 approved + implemented scenario와 연결된다.
- playlist에는 word motion beats가 남되, app playable 첫 문제는 기존 `normal-motion-basic-001`을 유지한다.

## 범위

- 포함: root `content/` YAML exercise/scenario/playlist 갱신
- 포함: content loader test의 root fixture 기대값 갱신
- 제외: multi-exercise app 진행
- 제외: TUI에서 word motion exercise 선택 UI
- 제외: 새 engine command 구현

## 검증 계획

- `go test ./internal/content/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 작업 항목

- [x] word motion coverage Red test를 작성한다.
- [x] `w`, `b`, `e`를 각각 훈련하는 YAML exercise를 작성한다.
- [x] 각 exercise에 맞는 scenario copy를 작성한다.
- [x] playlist beat를 word motion exercise set에 맞게 갱신한다.
- [x] replay/coverage/content/playable 검증을 통과시킨다.
- [x] 문서를 completed 상태로 동기화한다.

## 의사결정 로그

- 2026-05-18: `word-motion-basic-001`은 기존 draft를 승인하고 `w` 전용 exercise로 유지한다.
- 2026-05-18: `b`, `e`는 별도 exercise로 분리한다. 한 문항에 모두 몰아넣지 않고 각 command의 의미를 반복 가능하게 하기 위해서다.
- 2026-05-18: app playable의 첫 문제는 기존 `normal-motion-basic-001`로 유지한다. multi-exercise 진행은 GAMELOOP-001에서 다룬다.
- 2026-05-18: YAML `after_keys`가 0으로 읽히던 문제를 발견해 `HintSpec` yaml tag와 보존 테스트를 추가했다.

## 완료 결과

- `word-motion-basic` coverage required인 `w`, `b`, `e`가 approved + implemented exercise optimal trace에 모두 포함됐다.
- word motion exercise 3개와 scenario 3개가 replay gate를 통과한다.
- playlist에 word motion beat 3개를 연결했다.
- content loader test가 playable 후보 4개와 word motion coverage를 검증한다.
