# Foundation Exit Review

Date: 2026-05-30
Slice: `FOUNDATION-EXIT-001`
Status: completed review

## 결론

Foundation은 다음 단계로 넘어가도 된다. 다만 판정은 "출시 가능"이 아니라 **출시 가능한 게임 루프를 만들 수 있는 foundation 통과**다.

현재 Advimture는 Vim engine, content replay gate, playable E2E, progress v1 경계, Runbook Dispatch UI frame을 갖췄다. 다음 병목은 새 Vim command 수가 아니라, 이미 학습한 command를 계속 다시 플레이하게 만드는 mission/review/game loop다.

권장 다음 순서는 `PLATFORM-REVIEW-003 -> CONTENT-BREADTH-002 -> QUOTE-PAIR-HARDEN-001 -> release readiness`다.

## 확인한 Evidence

- `go test ./...`: pass
- `make e2e-playable`: pass
- playable target은 smoke/constraint/retry/review/resume/tutorial/incident/command-choice를 포함한 35개 scenario를 통과했다.
- 대표 long route summary:
  - `playable_command_choice_scope`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
  - `playable_incident_006_full`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
  - `playable_incident_007_full`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
- 대표 final screen spot review:
  - `incident-006`: target 검색 후 delimiter 보존 수리 성공 modal이 읽힌다.
  - `incident-007`: 5-beat mixed recovery가 `RUNBOOK SEALED`, grade/key count, dispatch complete까지 표시한다.
  - command-choice scope: 선택 이유와 quote value reuse 결과가 성공 modal에 남는다.

## 영역별 판정

| 영역 | 판정 | 근거 | 다음 리스크 |
|------|------|------|-------------|
| Vim engine | pass | movement, edit, operator, search, visual, char-find까지 tutorial/E2E 완료 | 새 command보다 quote/pair hardening이 작은 후속 후보 |
| Runtime/content | pass | YAML loader, replay gate, constraints, required/forbidden key, app_state assertion 동작 | applied content 증가 시 fixture/evidence 누락 주의 |
| TUI/E2E loop | pass | final/timeline evidence, focus panel, modal, review queue app_state 검증 | terminal size별 release smoke는 아직 부족 |
| Game loop | conditional | review queue/daily route/best record는 있으나 아직 "다시 들어올 이유"가 화면 단위에 흩어져 있음 | 다음 slice에서 mission/review loop를 제품 루프로 묶어야 함 |
| Content breadth | conditional | tutorial과 incident 001~007이 있으나 applied run 다양성은 아직 첫 확장 단계 | 새 engine 없이 기존 command 조합 incident를 더 늘려야 함 |
| Release readiness | not ready | 설치/첫 실행/터미널 크기/known limitations/release build gate가 아직 별도 문서와 검증으로 묶이지 않음 | 공개 전 `RELEASE-READINESS-001` 필요 |

## Blockers

### P0

없음.

현재 known bug였던 arrow key 우회, `ctrl+c` 즉시 종료, `?` hint affordance, command mismatch feedback은 E2E 대상에 들어와 있다.

### P1

1. `PLATFORM-REVIEW-003`: mission/review/daily/best record가 아직 하나의 게임 루프로 충분히 통합되지 않았다.
2. `CONTENT-BREADTH-002`: 기존 command를 섞는 applied incident와 command-choice가 더 필요하다.
3. `RELEASE-READINESS-001`: 첫 공개 전 설치/실행/터미널 크기/known limitations/release build 검증이 필요하다.
4. `UI-POLISH-002`: 현재 화면은 기능적으로 읽히지만, 출시 직전에는 색/강조/command memory/briefing polish가 더 필요하다.

## 다음 권장 순서

### 1. PLATFORM-REVIEW-003 — Mission/Review Game Loop

가장 먼저 연다.

목표는 저장 포맷 변경 없이 지금 있는 `재점검`, `잔류 리스크`, `오늘의 복구 루트`, `best record`, playlist 완료 화면을 하나의 반복 플레이 루프로 묶는 것이다.

포함:

- 첫 진입에서 오늘 무엇을 복구해야 하는지 더 명확히 보여준다.
- 성공/실패/완료 화면이 다음 행동을 게임 세계관 안에서 제안한다.
- 낮은 grade, 높은 key count, incomplete exercise가 다음 플레이 이유가 되게 한다.

제외:

- progress schema v2
- spaced review due date 저장
- daily streak/history 저장

### 2. CONTENT-BREADTH-002 — Applied Content Expansion

`PLATFORM-REVIEW-003`로 반복 루프를 정리한 뒤, 기존 engine만 사용해 적용 run을 늘린다.

후보:

- search-then-act incident
- repeat-change choice
- linewise reuse choice
- mixed incident 008

### 3. QUOTE-PAIR-HARDEN-001 — Quote/Pair Text Object Hardening

새 engine을 연다면 가장 작은 후보는 `ci'`, `di'`, `yi'`, `ci(`, `ci{` 계열이다. 단, release loop가 먼저 필요하므로 세 번째로 둔다.

### 4. RELEASE-READINESS-001

첫 외부 공개 전에는 별도 release gate가 필요하다.

포함 후보:

- `README.md` 설치/실행/테스트 안내
- terminal size별 smoke
- progress 파일 안전성 재확인
- release build command
- known limitations

## 사용자 결정사항

지금 당장 필요한 결정은 하나다.

- 다음 active slice를 `PLATFORM-REVIEW-003`으로 열지 승인한다.

아래는 지금 막지 않고 후속에서 결정한다.

- 첫 공개 기준을 internal alpha로 둘지 public beta로 둘지
- progress schema v2를 언제 다시 열지
- quote/pair hardening에서 single quote와 parenthesis/brace 중 어디까지 첫 scope로 넣을지

## Exit 판정

Foundation phase는 **조건부 통과**다.

조건:

- 다음 작업은 새 engine 직행이 아니라 `PLATFORM-REVIEW-003`으로 game loop를 먼저 다듬는다.
- content breadth를 늘릴 때 long route에는 final/timeline evidence를 계속 남긴다.
- progress 저장 포맷은 사용자 승인 전까지 변경하지 않는다.
