# CONTENT-BREADTH-001 — Command Choice Extension

Slice-ID: CONTENT-BREADTH-001
Created: 2026-05-26
Status: active
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/content-breadth-001-command-choice-extension.md
- content/
- test/e2e/

## 목표

새 engine 기능 없이 기존 Vim command를 재조합해 command-choice/applied learning 콘텐츠 폭을 한 단계 늘린다.

## 선택한 범위

`choice-002-repeat-or-substitute`를 playable로 승격한다. 같은 literal이 파일 전체에 반복될 때 `.` 반복보다 `:%s`가 더 적합하다는 range-choice 판단을 훈련한다.

## 포함

- command-choice 두 번째 exercise/scenario 추가
- incident command-choice playlist에 두 번째 beat 추가
- focused E2E 추가 또는 기존 command-choice E2E 확장
- content replay gate와 full playable E2E 통과

## 제외

- 새 Vim engine 기능
- 새 YAML schema
- progress 저장 포맷 변경
- command-choice 대량 확장

## 수용 기준

- 새 exercise는 approved + implemented + replay_status pass다.
- `required_keys`는 `:`, `%`, `s`, `enter` 기반 substitute route를 고정한다.
- dot repeat/manual edit/visual route가 학습 목표를 무너뜨리면 forbidden route로 막는다.
- scenario copy는 `.`가 아니라 `:%s`를 선택해야 하는 이유를 설명한다.
- focused E2E와 `make e2e-playable`을 통과한다.

## Step 1: Content Design

- [ ] existing command-choice/Ex command content 확인
- [ ] `choice-002-repeat-or-substitute` target/optimal/constraints 설계

## Step 2: Implementation

- [ ] exercise/scenario/playlist 추가
- [ ] focused E2E 추가/갱신
- [ ] docs/spec 갱신

## Step 3: Verification

- [ ] content tests
- [ ] focused E2E
- [ ] `go test ./...`
- [ ] `make e2e-playable`
- [ ] `git diff --check`
