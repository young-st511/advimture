# E2E-EVIDENCE-008 — Long Incident Final/Timeline Evidence

Slice-ID: E2E-EVIDENCE-008
Created: 2026-05-29
Completed: 2026-05-29
Status: completed
Scope-Mode: e2e-fixture-and-docs
Allowed-Paths:
- docs/exec-plans/active/e2e-evidence-008-long-incident-evidence.md
- docs/exec-plans/completed/e2e-evidence-008-long-incident-evidence.md
- docs/roadmap/PROGRAM.md
- docs/roadmap/MIDTERM_TODO.md
- docs/roadmap/CHANGES.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- test/e2e/playable_incident_*_full.yaml
- test/e2e/playable_command_choice_scope.yaml
- docs/roadmap/PLAYTEST_REVIEW_2026-05-29.md

## 목표

Applied Mastery Runs 이후 긴 incident route를 사람이 직접 재검토할 때 마지막 화면과 진행 흐름 evidence가 부족했다. runner 기능은 이미 있으므로, 이번 slice는 긴 incident E2E fixture가 `screen_timeline.txt`와 `screen_final.txt`를 일관되게 남기도록 고정한다.

## 범위

- 포함:
  - `playable_incident_*_full.yaml` evidence 옵션 정리
  - `playable_command_choice_scope.yaml` evidence 옵션 정리
  - long route evidence 기준 문서화
  - 대표 long route evidence 재생성과 summary flag 확인
- 제외:
  - E2E runner 새 기능
  - TUI layout 변경
  - game/content schema 변경
  - progress 저장 포맷 변경
  - CI 정책 변경

## 수용 기준

- 모든 `test/e2e/playable_incident_*_full.yaml` fixture는 raw ANSI, cleaned stream, key trace, summary, timeline, final screen evidence를 저장한다.
- `test/e2e/playable_command_choice_scope.yaml`도 같은 evidence bundle을 저장한다.
- 대표 route의 `summary.json`은 `screen_timeline_evidence: true`, `screen_final_evidence: true`를 기록한다.
- 생성된 `screen_final.txt`는 마지막 `ADVIMTURE |` frame 중심으로 사람이 읽을 수 있다.
- verification docs는 long incident/playable route의 evidence bundle 기준을 설명한다.
- 앱 동작, content schema, progress schema는 변경하지 않는다.

## Step 1: Plan and Contract

- [x] 중기 플랜에 `E2E-EVIDENCE-008`을 추가한다.
- [x] Program active slice를 갱신한다.
- [x] verification docs에 long route evidence 기준을 추가한다.

## Step 2: Fixture Evidence

- [x] incident full route fixture에 `save_screen_timeline`과 `save_screen_final`을 추가한다.
- [x] command-choice scope route fixture에 같은 evidence bundle을 추가한다.
- [x] 기존 raw/clean/key/summary evidence는 유지한다.

## Step 3: Verification

- [x] runner unit test를 실행한다.
- [x] 대표 long route focused E2E를 실행한다.
- [x] full playable E2E를 실행한다.
- [x] 전체 Go 테스트와 diff check를 실행한다.

## 검증 계획

- `go test ./cmd/e2e-runner/...`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_006_full.yaml`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_007_full.yaml`
- `make e2e-playable`
- `go test ./...`
- `git diff --check`

## 의사결정 로그

- 2026-05-29: runner 구현은 이미 충분하므로 새 기능보다 fixture contract와 representative evidence를 우선한다.
- 2026-05-29: long incident route는 향후 UI/UX RedTeam과 Agent playtest의 기본 관찰 단위로 본다.
- 2026-05-29: 보강된 evidence 기준으로 다음 후보는 새 engine 직행보다 Foundation exit review를 먼저 권장한다.

## 검증 결과

- `go test ./cmd/e2e-runner/...`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_command_choice_scope.yaml`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_006_full.yaml`: pass
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_incident_007_full.yaml`: pass
- `make e2e-playable`: pass
- `go test ./...`: pass
- `git diff --check`: pass
- all configured long route summaries confirmed `screen_timeline_evidence: true` and `screen_final_evidence: true`:
  - `playable_incident_001_full`
  - `playable_incident_002_full`
  - `playable_incident_003_full`
  - `playable_incident_004_full`
  - `playable_command_choice_scope`
  - `playable_incident_006_full`
  - `playable_incident_007_full`
- representative summary flag:
  - `playable_command_choice_scope`: `screen_timeline_evidence: true`, `screen_final_evidence: true`
  - `playable_incident_006_full`: `screen_timeline_evidence: true`, `screen_final_evidence: true`
  - `playable_incident_007_full`: `screen_timeline_evidence: true`, `screen_final_evidence: true`

## 미해결 질문

- full E2E를 CI에 올릴지는 이번 slice에서 결정하지 않는다.
