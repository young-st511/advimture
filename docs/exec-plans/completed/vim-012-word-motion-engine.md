# ExecPlan: Word motion engine

Slice-ID: VIM-012
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/vimengine/**
- internal/vimoracle/**
- internal/runtime/**
- content/command_clusters/**
- docs/gameplay/command-catalog.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/vim-012-word-motion-engine.md

## 목표

`word-motion-basic` cluster의 엔진 지원을 구현한다. `w`, `b`, `e`는 Vim의 word motion 학습 목표에 맞게 word/symbol/whitespace 경계를 처리하고, line boundary를 넘어 이동할 수 있어야 한다.

## 영향 도메인

- Gameplay: 새 command는 command catalog에서 출발하며, exercise/scenario 제작은 후속 EXERCISE 루프로 분리한다.
- Vim engine: `State + Key -> State + Events` 순수 전이 계약을 유지한다.

## 수용 기준

- `w`는 현재 word/symbol에서 다음 word/symbol 시작으로 이동한다.
- `b`는 이전 word/symbol 시작으로 이동한다.
- `e`는 현재 또는 다음 word/symbol 끝으로 이동한다.
- word motion은 공백, 문장부호, 줄 경계를 테스트로 고정한다.
- 지원하지 않는 mode에서는 기존 unsupported key 동작을 유지한다.
- `word-motion-basic` command cluster의 engine support를 implemented로 승격한다.

## 범위

- 포함: `internal/vimengine` key constants와 word motion 전이
- 포함: 단위 테스트와 oracle-style expected fixture
- 포함: docs/content command cluster 지원 상태 갱신
- 제외: word motion exercise/scenario playable 승격
- 제외: operator 조합(`dw`, `cw`, `yw`)
- 제외: 실제 Neovim CLI executor

## 검증 계획

- `go test ./internal/vimengine/...`
- `go test ./internal/vimoracle/...`
- `go test ./...`
- `git diff --check`

## 작업 항목

- [x] `w/b/e` Red tests를 작성한다.
- [x] word/symbol/whitespace classifier와 이동 전이를 구현한다.
- [x] oracle-style comparison test를 추가한다.
- [x] command cluster engine support를 implemented로 승격한다.
- [x] 문서를 completed 상태로 동기화한다.

## 의사결정 로그

- 2026-05-18: word motion은 keyword(`letter/digit/_`)와 symbol을 별도 word class로 다룬다.
- 2026-05-18: 줄 사이에는 암묵적 whitespace 경계를 둔다. 빈 줄은 건너뛰되 다음 word/symbol로 이동한다.
- 2026-05-18: 실제 Neovim CLI executor는 아직 없으므로 oracle-style fixture로 비교 계약만 잠근다.
- 2026-05-18: `w`가 지원되면서 runtime unsupported-key 테스트 입력을 `x`로 갱신했다.

## 완료 결과

- `internal/vimengine`에 `w`, `b`, `e` 상수와 word motion 전이를 추가했다.
- 공백, 문장부호, 줄 경계, unsupported mode, DesiredCol 회귀 테스트를 추가했다.
- `word-motion-basic` command cluster를 `engine_support: implemented`로 승격했다.
- exercise/scenario/playable 승격은 후속 EXERCISE 루프로 남겼다.
