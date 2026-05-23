# UI-EVIDENCE-002 — Final Screen Evidence

Slice-ID: UI-EVIDENCE-002
Created: 2026-05-23
Status: completed
Completed: 2026-05-23
Scope-Mode: normal
Allowed-Paths:
- docs/roadmap/PROGRAM.md
- docs/verification/spec.md
- docs/verification/tui-e2e-loop.md
- cmd/e2e-runner/
- test/e2e/

## 영향 도메인

- Verification: TUI E2E는 테스트 전용 HOME을 사용하고, 실패/회귀 분석용 evidence는 `artifacts/` 아래에 격리한다.
- Gameplay: 이번 slice는 UI 동작을 바꾸지 않고 evidence만 보강한다.

## 중기 플랜

1. `UI-EVIDENCE-002`: 마지막 화면 evidence를 `screen_final.txt`로 분리해 다음 UI 개편의 관찰 기반을 만든다.
2. `UI-LAYOUT-001`: `tea.WindowSizeMsg`와 renderer width/height를 도입해 FocusPanel을 responsive modal-like layout으로 배치한다.
3. `UI-MODAL-002`: tutorial intro/failure/success modal의 입력 레이어를 분리해 modal open 중 입력이 Vim key trace로 섞이지 않게 한다.
4. `UI-SKIN-001`: Runbook Dispatch 세계관에 맞춰 header/status/focus panel 문구와 색을 다듬되 Vim 학습 목표를 앞에 둔다.
5. `UI-QA-002`: final screen evidence와 app_state UI assertion을 묶은 regression fixture를 확장한다.

## 목표

현재 `screen.txt`와 `screen_timeline.txt`는 cleaned terminal stream 전체를 담아 최종 프레임만 빠르게 보기 어렵다. 이 slice는 E2E evidence에 `screen_final.txt`를 추가해 Agent가 UI 변경 후 최종 화면을 직접 읽고 비교할 수 있게 한다.

## 범위

- 포함:
  - E2E runner evidence option `save_screen_final`
  - final frame 추출 함수와 단위 테스트
  - summary JSON에 final screen evidence 여부 기록
  - 대표 fixture의 final screen evidence 저장
  - verification docs 동기화
- 제외:
  - 실제 TUI layout 변경
  - ANSI full terminal emulator/parser 도입
  - screenshot/image evidence
  - 새 dependency 추가

## 수용 기준

- `save_screen_final: true`이면 `artifacts/e2e/{scenario_id}/screen_final.txt`를 쓴다.
- `screen_final.txt`는 누적 stream 전체가 아니라 마지막 `ADVIMTURE |` frame 중심이어야 한다.
- final frame 추출은 playable error frame도 보존한다.
- summary JSON은 `screen_final_evidence` boolean을 기록한다.
- 기존 `screen.txt`, `screen_timeline.txt`, raw ANSI, key trace, app_state evidence 동작은 유지된다.

## Step 1: Final Screen Evidence

- 목표: E2E runner에 final screen evidence 저장 기능을 추가한다.
- 변경 파일:
  - `cmd/e2e-runner/main.go`
  - `cmd/e2e-runner/main_test.go`
  - `test/e2e/playable_hjkl_success.yaml`
- 테스트 파일:
  - `cmd/e2e-runner/main_test.go`
- 충족 기준: `screen_final.txt` 생성, summary flag, 기존 evidence 유지
- Boundaries 주의: evidence는 `artifacts/` 아래에만 생성한다.
- 상세 작업:
  - [x] `evidence.save_screen_final` 옵션 추가
  - [x] final frame 추출 함수와 테스트 추가
  - [x] `writeEvidence`에서 `screen_final.txt` 저장
  - [x] 대표 smoke fixture에 옵션 추가

## Step 2: 문서 동기화와 완료 처리

- 목표: verification docs와 Program 상태를 현재 동작에 맞춘다.
- 변경 파일:
  - `docs/verification/spec.md`
  - `docs/verification/tui-e2e-loop.md`
  - `docs/roadmap/PROGRAM.md`
- 테스트 파일: n/a
- 충족 기준: final screen evidence의 목적과 한계가 문서에서 명확하다.
- Boundaries 주의: full terminal parser는 후속으로 남긴다.
- 상세 작업:
  - [x] verification spec에 final screen evidence 추가
  - [x] TUI E2E loop evidence 목록 갱신
  - [x] ExecPlan 완료 이동

## 검증 결과

- `go test ./cmd/e2e-runner`
- `go run ./cmd/e2e-runner --scenario test/e2e/playable_hjkl_success.yaml`

---

## 실행 규칙

각 Step은 Spec 기반 테스트 작성 → 구현 → 변경 범위 테스트 → 필요한 E2E → 문서 동기화 → 커밋/푸시 순서로 진행한다.
