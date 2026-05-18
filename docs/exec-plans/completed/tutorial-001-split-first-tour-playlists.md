# ExecPlan: Split first tour into tutorial episodes

Slice-ID: TUTORIAL-001
Created: 2026-05-18
Status: completed
Scope-Mode: normal
Allowed-Paths:
- content/playlists/first-five-minutes.yaml
- docs/gameplay/spec.md
- docs/gameplay/content-requirements.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/PROGRAM.md
- docs/exec-plans/active/tutorial-001-split-first-tour-playlists.md
- docs/exec-plans/completed/tutorial-001-split-first-tour-playlists.md
- internal/content/loader.go
- internal/content/loader_test.go
- internal/playable/model.go
- internal/playable/model_test.go
- test/e2e/**

## 목표

현재 17개 exercise를 한 번에 이어 실행하는 `first-5-minute` vertical slice를 실제 학습 UX에 맞게 8문항 이하의 tutorial episode playlist들로 분리한다.

## 영향 도메인

- Gameplay: first tour를 짧은 tutorial episode 단위로 진행한다.
- Content: playlist를 movement/survival/navigation/ex-command episode로 나눈다.
- Playable: 화면에 tutorial title과 episode-local exercise count를 보여준다.
- Verification: E2E assertion을 episode-aware 진행 기준으로 갱신한다.

## 수용 기준

- approved/implemented tutorial playlist는 각각 8문항 이하이다.
- `first-5-minute` legacy vertical slice는 default playable path에서 실행되지 않는다.
- default playable path는 tutorial episode 순서로 진행한다.
- 화면은 현재 tutorial title과 episode-local exercise count를 보여준다.
- 한 tutorial 마지막 성공에서 다음 tutorial이 있으면 `Next tutorial: enter`를 보여주고 `enter`로 다음 tutorial에 진입한다.
- 기존 progress `Missions` map만 사용하고 저장 포맷은 바꾸지 않는다.
- `make e2e-playable`은 tutorial episode split 이후에도 통과한다.

## 범위

- 포함: playlist YAML 분리
- 포함: playlist status 필터링과 episode metadata
- 포함: playable view copy와 E2E 갱신
- 제외: 새 content schema 의존성 추가
- 제외: progress 저장 포맷 변경
- 제외: episode 선택 메뉴
- 제외: exercise constraints/failure UX 구현

## 중기 플랜

1. `TUTORIAL-001`: 17개 first tour를 8문항 이하 tutorial episodes로 분리한다.
2. `CONSTRAINT-001`: exercise constraint schema와 runtime fail-on-max-input/forbidden-input을 구현한다.
3. `SCORING-002`: required command 사용 여부를 grade/coaching에 반영한다.
4. `FAILURE-001`: 실패 화면, `r`/`enter` retry, 남은 입력 수 UI를 완성한다.
5. `QA-001`: forbidden input, max input 초과, non-intended route, retry/hint E2E를 보강한다.
6. `CONTENT-003`: Ex command를 중반 고급 튜토리얼로 위치시키고 첫 투어 콘텐츠 pacing을 재검증한다.

## 검증 계획

- `go test ./internal/content/...`
- `go test ./internal/playable/...`
- `go test ./...`
- `make e2e-smoke`
- `make e2e-playable`
- `git diff --check`

## 검증 결과

- `go test ./internal/content/...`: pass
- `go test ./internal/playable/...`: pass
- `go test ./...`: pass
- `make e2e-smoke`: pass
- `make e2e-playable`: pass
- `git diff --check`: pass

## 작업 항목

- [x] 중기 플랜을 `MIDTERM_TODO.md`에 반영한다.
- [x] playlist YAML을 tutorial episode로 분리한다.
- [x] loader/playable이 approved tutorial playlist들을 순서대로 읽게 한다.
- [x] playable view가 tutorial title과 episode-local count를 보여준다.
- [x] E2E fixtures를 episode-aware UI에 맞게 갱신한다.
- [x] 검증 명령을 실행한다.
- [x] ExecPlan을 completed로 이동하고 PROGRAM을 갱신한다.
