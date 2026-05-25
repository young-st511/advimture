# CHOICE-JUDGMENT-001 — Command Choice Judgment Cue

Slice-ID: CHOICE-JUDGMENT-001
Created: 2026-05-26
Status: completed
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/gameplay/spec.md
- docs/gameplay/command-choice-drills.md
- docs/exec-plans/active/choice-judgment-001-command-choice-cue.md
- content/exercises/command-choice.yaml
- content/scenarios/command-choice.yaml
- internal/playable/
- test/e2e/playable_command_choice_scope.yaml

## 목표

`command-choice-scope-001`이 단순히 `V j d` 순서를 암기하는 문항처럼 보이지 않도록, “복구 범위 판단 → 적절한 Vim 도구 선택”의 학습 의도를 콘텐츠와 검증에 더 분명하게 만든다.

## 범위

- 포함:
  - briefing/hint/success/failure 문구를 선택 판단 중심으로 조정
  - `scope-choice` 설계 문서와 playable content 정합성 확인
  - command-choice E2E가 선택 이유 문구를 검증하도록 보강
- 제외:
  - 새 YAML schema 추가
  - 새 Vim engine 기능 추가
  - command-choice exercise 대량 추가
  - progress 저장 포맷 변경

## 수용 기준

- 첫 문장은 정답 key sequence를 직접 말하지 않고, 범위 판단을 요구한다.
- hint/failure/success는 “단어/quote 값이 아니라 줄 묶음”이라는 선택 이유를 강조한다.
- E2E는 성공 buffer뿐 아니라 command-choice 판단 copy를 화면에서 검증한다.
- `go test ./internal/content ./internal/playable`과 focused E2E를 통과한다.

## Step 1: Gap Review

- [x] 현재 command-choice content와 E2E 문구 확인
- [x] 암기처럼 보일 수 있는 표현 식별

## Step 2: Content and Test

- [x] exercise/scenario 문구 개선
- [x] E2E assertion 보강
- [x] 필요한 경우 playable test 보강

## Step 3: Verification

- [x] focused Go tests
- [x] focused E2E
- [x] `go test ./...`
- [x] `make e2e-playable`
- [x] `git diff --check`

## 결과

- `command-choice-scope-001`의 목표/힌트/시나리오 문구를 “값/단어가 아니라 줄 묶음을 복구 범위로 판단한다”는 선택 이유 중심으로 조정했다.
- `playable_command_choice_scope`는 초기 판단 copy와 성공 copy를 함께 검증한다.
- `playable_incident_004_full`의 다음 잔류 리스크 기대 문구와 command-line 입력 간격을 갱신해 full suite 안정성을 높였다.
