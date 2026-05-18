# ExecPlan: Isolate Go cache for E2E runner child process

Slice-ID: E2E-004
Created: 2026-05-18
Status: completed
Scope-Mode: narrow
Allowed-Paths:
- cmd/e2e-runner/**
- test/e2e/playable_hjkl_success.yaml
- Makefile
- docs/exec-plans/completed/e2e-004-go-cache-isolation.md

## 목표

TUI E2E runner가 자식 프로세스로 `go run .`을 실행할 때 사용자 Go build cache에 쓰지 않도록 격리한다. sandbox나 CI에서 `/Users/.../Library/Caches/go-build` 쓰기 권한 때문에 화면 assertion이 permission error로 오염되는 flake를 없앤다.

## 범위

- 포함: E2E runner child process의 `GOCACHE` 기본값을 테스트 HOME 아래로 지정
- 포함: child process의 HOME 변경 때문에 module cache가 temp HOME으로 바뀌지 않도록 parent `GOPATH`/`GOMODCACHE`를 고정
- 포함: Makefile에서 runner 자체의 `go run`도 repo-local artifact cache를 사용
- 포함: 격리 cache의 첫 빌드 시간을 반영한 playable scenario timeout 조정
- 포함: 기존 사용자가 명시한 `GOCACHE`는 존중
- 제외: dependency download 정책 변경
- 제외: scenario DSL 변경

## 검증 계획

- `go test ./cmd/e2e-runner`
- `make e2e-smoke`
- `make e2e-playable`

## 승인 체크

- [x] `GOCACHE`가 없으면 child process가 temp HOME 아래 cache를 사용한다.
- [x] temp HOME에서도 child `go run .`이 dependency를 새로 다운로드하려 하지 않는다.
- [x] `make e2e-*`는 runner 자체를 실행할 때 사용자 Go build cache에 쓰지 않는다.
- [x] 기존 `GOCACHE` 환경변수가 있으면 runner가 덮어쓰지 않는다.
- [x] playable E2E가 반복 실행되어도 화면 대기에서 permission error로 실패하지 않는다.

## 완료 결과

- runner child process는 `HOME=temp`에서도 parent `GOPATH`/`GOMODCACHE`를 고정한다.
- runner child process는 `GOCACHE`가 없을 때 temp HOME 아래 Go build cache를 쓴다.
- Makefile의 `e2e-smoke`, `e2e-playable`은 runner 자체도 `artifacts/go-build-cache`로 실행한다.
- playable E2E timeout을 격리 cache 첫 빌드에 맞춰 20초로 조정했다.
