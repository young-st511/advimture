# ExecPlan: File-backed playable

Slice-ID: PLAY-002
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- internal/playable/**
- internal/content/**
- internal/app/**
- test/e2e/**
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/exec-plans/completed/play-002-file-backed-playable.md

## 목표

첫 playable slice의 hardcoded exercise/scenario 정의를 제거하고, CONTENT-001의 root `content/` YAML loader에서 읽은 playable exercise/scenario를 실행한다.

## 범위

- 포함: `internal/playable`이 `content.LoadLibrary`를 사용해 첫 playable exercise를 선택
- 포함: app 기본 content root를 `content/`로 연결
- 포함: playable unit test와 E2E scenario를 file-backed copy에 맞게 갱신
- 제외: multi-exercise playlist 진행
- 제외: replay/coverage validator 자동화
- 제외: mission/progress schema 변경

## 수용 기준

- playable model은 hardcoded `content.ExerciseSpec`를 직접 만들지 않는다.
- playable model은 root `content/` YAML에서 approved + implemented exercise만 실행한다.
- scenario title/briefing/success text는 YAML scenario에서 온다.
- `normal-motion-basic-001`을 `l`, `l`로 성공시켜 score/progress/e2e_state가 유지된다.
- `make e2e-smoke`가 file-backed content 기준으로 통과한다.

## 검증 계획

- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`

## 승인 체크

- [x] hardcoded exercise definition이 제거됐다.
- [x] file-backed playable unit test가 통과한다.
- [x] TUI E2E가 content fixture copy를 기준으로 통과한다.

## 의사결정 로그

- 2026-05-18: playable은 `content.LoadLibrary`의 첫 playable exercise와 matching scenario를 사용한다. multi-exercise playlist 진행은 GAMELOOP-001로 남긴다.
- 2026-05-18: 앱 기본 content root는 repo root `content/`로 둔다. 테스트는 `../../content`를 명시해 working directory 차이를 피한다.
- 2026-05-18: 리뷰에서 content load 실패 화면의 quit 처리와 빈 playable scenario copy를 발견해 PLAY-002 완료 전 같은 루프에서 보강했다.
- 2026-05-18: 배포형 바이너리의 cwd-independent content root 정책은 패키징/실행 정책 루프로 분리한다.

## 완료 결과

- `internal/playable`의 hardcoded `content.ExerciseSpec`를 제거했다.
- `internal/app`이 `ContentRoot: "content"`를 전달한다.
- E2E scenario를 YAML scenario copy 기준으로 갱신했다.
- content load 실패 화면에서도 `q`로 종료할 수 있다.
- approved/implemented scenario는 title/briefing/success copy가 비어 있으면 load 단계에서 실패한다.
- `l`, `l` 성공, score/progress/e2e_state 흐름은 유지했다.
