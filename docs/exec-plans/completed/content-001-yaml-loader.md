# ExecPlan: YAML content loader

Slice-ID: CONTENT-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- content/**
- internal/content/**
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/content-001-yaml-loader.md

## 목표

repo root `content/` 아래 YAML 콘텐츠를 읽어 command cluster, exercise, scenario, playlist를 검증하고, 실행 가능한 exercise를 기존 `internal/content.CompileExercise` 경로로 변환한다.

## 범위

- 포함: `content/command_clusters`, `content/exercises`, `content/scenarios`, `content/playlists` fixture
- 포함: YAML loader와 cross-reference validator
- 포함: `engine_support: planned` 콘텐츠를 로드하되 playable candidate에서 제외하는 API
- 포함: `normal-motion-basic-001`을 compiled exercise로 변환
- 포함: key trace 정규화, `optimal_key_count`, allowed/forbidden 충돌 검증
- 제외: playable app wiring
- 제외: replay validator 자동화
- 제외: JSON import/export
- 제외: runtime에서 forbidden key 차단

## 수용 기준

- loader는 repo root `content/` 아래 YAML 파일을 우선 읽는다.
- loader는 command cluster, exercise, scenario, playlist 파일을 각각 읽는다.
- loader는 `engine_support: planned` 콘텐츠를 로드하되 playable 후보에서는 제외한다.
- loader는 `normal-motion-basic-001`을 현재 hardcoded playable exercise와 같은 compiled exercise로 변환한다.
- loader는 exercise가 참조하는 command cluster가 없으면 오류를 반환한다.
- loader는 scenario가 참조하는 exercise가 없으면 오류를 반환한다.
- loader는 cursor target이 buffer 범위를 벗어나면 오류를 반환한다.
- loader는 optimal key trace를 정규화하고 `optimal_key_count`와 길이가 다르면 오류를 반환한다.
- loader는 optimal key가 allowed keys에 없거나 forbidden keys와 충돌하면 오류를 반환한다.
- loader는 approved exercise가 approved 또는 implemented command cluster만 참조하도록 검증한다.
- loader는 approved scenario가 approved 또는 implemented exercise만 참조하도록 검증한다.

## 검증 계획

- `go test ./internal/content/...`
- `go test ./...`

## 승인 체크

- [x] YAML fixture가 root `content/` 아래에 있다.
- [x] loader/validator 단위 테스트가 실패 케이스를 포함한다.
- [x] planned 콘텐츠가 playable candidate에서 제외된다.
- [x] `normal-motion-basic-001` compile 결과가 기존 playable target과 일치한다.

## 의사결정 로그

- 2026-05-18: 사람이 쓰기 쉬운 YAML의 `buffer: |`를 loader에서 `[]string`으로 정규화해 기존 `CompileExercise`에 연결했다.
- 2026-05-18: `coverage_required`는 CONTENT-001에서 실패 처리하지 않고 `CoverageReports()`로 보고한다. `normal-motion-basic`의 `h/k` gap은 후속 CONTENT-002/EXERCISE 루프에서 보강한다.
- 2026-05-18: `replay_status` 자동 검증은 CONTENT-002로 남기고, 이번 루프는 YAML load/cross-reference/schema validation에 집중했다.

## 완료 결과

- root `content/` 아래 command cluster, exercise, scenario, playlist YAML fixture를 추가했다.
- `internal/content.LoadLibrary`로 YAML을 로드하고 cross-reference를 검증한다.
- `Library.PlayableExercises`는 approved/implemented + `engine_support: implemented` exercise만 반환한다.
- `Library.CompileExercise`는 loaded exercise를 기존 runtime compiled exercise로 변환한다.
- 실패 케이스 테스트를 추가했다: missing command cluster, missing scenario exercise, out-of-range cursor, optimal key count mismatch, invalid key policy.
