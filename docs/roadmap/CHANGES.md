# Changes Log — 시퀀싱/가정 변경 기록

> append-only. 새 항목을 위에 추가하고 기존 항목은 수정하지 않는다.

## 2026-05-30 — Command-choice에 repeat-change 판단 추가

이전 가정: `incident-005-command-choice`는 linewise scope, global substitute, inline target range, quote value reuse까지 다루며, 다음 content breadth 후보는 line reuse 또는 repeat-change reuse였다.

새 가정: `incident-005-command-choice`는 fifth beat로 repeat-change choice를 포함한다. 같은 단어 교체가 이어질 때 두 번째 변경은 직접 재입력하지 않고 `.`으로 마지막 변경을 반복하는 판단을 검증한다.

이유: line reuse는 이미 linewise tutorial/incident에서 여러 번 다뤘고, `.` 반복 변경은 Vim 효율 체감과 실무 설정 편집 재사용성이 더 크다. 기존 engine/schema만으로 replay/E2E 검증도 가능하다.

영향: `CONTENT-BREADTH-002`는 완료됐다. 다음 권장 slice는 `QUOTE-PAIR-HARDEN-001`이다.

## 2026-05-30 — Mission/review loop를 next dispatch로 연결

이전 가정: review queue, daily route, best record는 화면에 표시되지만 성공 후 다음 행동과는 느슨하게 연결돼 있었다.

새 가정: 저장 포맷 변경 없이도 성공 debrief가 `이번 복구`, `최단 복구`, `목표 입력`, `잔류 리스크`, `다음 출격`을 보여주고, 마지막 dispatch에서 review 후보가 있으면 `enter`로 primary review exercise에 재진입한다.

이유: Foundation exit review에서 다음 병목은 새 Vim command 수가 아니라 반복 플레이 루프라고 판정했다. 이미 있는 progress v1과 review queue만으로도 “성공 -> 기록 확인 -> 잔류 리스크 -> 재출격” 행동을 만들 수 있다.

영향: `PLATFORM-REVIEW-003`은 완료됐다. 다음 권장 slice는 기존 engine만 사용하는 `CONTENT-BREADTH-002`다.

## 2026-05-30 — Foundation exit review 완료

이전 가정: `FOUNDATION-EXIT-001`에서 platform loop, content breadth, engine hardening 중 다음 우선순위를 evidence로 결정해야 했다.

새 가정: Foundation은 조건부 통과했으며, 다음 병목은 새 Vim command 수가 아니라 mission/review/game loop다. 다음 권장 slice는 `PLATFORM-REVIEW-003`이다.

이유: `go test ./...`와 `make e2e-playable`이 통과했고, long incident final/timeline evidence도 사람이 읽을 수 있는 상태다. 반면 review queue, daily route, best record, playlist 완료 흐름은 아직 하나의 반복 플레이 루프로 충분히 묶이지 않았다.

영향: 다음 active slice는 사용자 승인 후 `PLATFORM-REVIEW-003`으로 연다. `CONTENT-BREADTH-002`, `QUOTE-PAIR-HARDEN-001`, release readiness는 후속 후보로 둔다.

## 2026-05-30 — Foundation forward plan 도입

이전 가정: 다음 권장 후보는 `PROGRAM.md`와 `MIDTERM_TODO.md`에만 두면 충분했다.

새 가정: `PROGRAM.md`는 현재 상태, `MIDTERM_TODO.md`는 현재 중기 보드, `FORWARD_PLAN.md`는 앞으로 2~8주 방향과 decision gate를 맡는다.

이유: 문서 cleanup 이후에도 다음 작업 순서를 지속적으로 참고하려면 current board와 rolling plan을 분리해야 한다.

영향: 다음 실제 작업은 `FOUNDATION-EXIT-001`이다. 이후 platform loop, content breadth, quote-pair hardening 중 하나를 evidence 기반으로 선택한다.

## 2026-05-30 — Roadmap 문서 신선도 규칙 도입

이전 가정: `PROGRAM.md`, `MIDTERM_TODO.md`, `ENGINE_TODO.md`, health/review 문서가 모두 roadmap root에 남아 있어도 최신 항목을 사람이 구분할 수 있었다.

새 가정: root roadmap 문서는 현재 판단만 담고, 오래된 health/review와 긴 중기 플랜 이력은 `docs/roadmap/archive/`로 이동한다. `PROGRAM.md`는 active slice와 최근 완료만, `MIDTERM_TODO.md`는 현재 중기 보드만 유지한다.

이유: Agent가 stale "next 후보"를 active 계획처럼 읽지 않게 하려면 canonical/current/history 경계를 문서 구조로 강제해야 한다.

영향: 다음 루프는 여전히 `PLAN-REFRESH-009`다. 새 engine/content를 열기 전에 slimmed `PROGRAM.md`, current `MIDTERM_TODO.md`, 최신 review를 기준으로 판단한다.

## 2026-05-29 — Long incident evidence bundle 고정

이전 가정: Applied Mastery Runs는 기능적으로 통과했지만 긴 incident의 `summary.json`에서 `screen_timeline_evidence`와 `screen_final_evidence`가 false라 UI/UX 회귀를 사람이 훑기 어려웠다.

새 가정: 긴 incident full route와 command-choice applied route는 `screen_timeline.txt`와 `screen_final.txt`를 기본 evidence bundle로 남긴다. runner 기능은 그대로 두고 fixture와 verification contract를 고정한다.

이유: 새 content나 engine을 늘리기 전에, 긴 플레이 루트를 Agent와 사람이 같은 증거로 검토할 수 있어야 한다.

영향: 다음 루프는 `PLAN-REFRESH-009`로 Foundation exit review와 다음 중기 플랜을 먼저 고정한다. 새 engine 후보를 연다면 quote-pair hardening이 가장 작은 후보지만, evidence 기반 review를 먼저 둔다.

## 2026-05-28 — Applied mastery run 시작

이전 가정: `Inline Target Motions` 완료 후 다음 후보는 `incident-006-inline-target-repair` 또는 reuse-choice drill이었다.

새 가정: 다음 중기 플랜은 `Applied Mastery Runs`로 고정한다. `incident-006-inline-target-repair`, quote value reuse-choice, `incident-007-mixed-recovery`를 완료했고, 검색 후 delimiter 보존 수리와 mixed command 적용을 playable path에 연결했다.

이유: 새 engine을 늘리는 것보다 이미 배운 command를 실제 incident에서 조합하는 판단이 현재 학습/게임성 병목에 더 가깝다.

영향: 다음 루프는 새 content가 아니라 `E2E-EVIDENCE-008`로 긴 incident의 final/timeline evidence를 보강하는 것을 권장한다.

## 2026-05-26 — Inline target motion을 tutorial까지 승격

이전 가정: `f/t`는 다음 engine 후보였고, operator 결합과 playable tutorial 범위는 아직 분리 계획 단계였다.

새 가정: `char-find-line`은 forward same-line `f/t`와 `df/dt/cf/ct`를 구현하고, `tutorial-94-char-find-line` 6문항으로 replay/coverage/E2E까지 연결한다. 이어서 `incident-005-command-choice`에 `choice-005-inline-target-range`를 third beat로 추가해 `ct,`와 `cf,`의 delimiter 보존 판단까지 적용한다.

이유: `f/t`는 단순 이동보다 delimiter/quote/comma를 기준으로 정확한 범위를 잡을 때 학습 가치가 크다. 따라서 tutorial 직후 적용 루프로 넘어가야 장기 반복 학습 플랫폼의 “유용하고 재미있는 command 선택” 철학과 맞는다.

영향: `Inline Target Motions` 중기 플랜은 scope planning, engine, tutorial, applied command-choice까지 완료됐다. 다음 후보는 `/target` 검색 후 `ct,`로 comma 앞 값만 교체하는 `incident-006-inline-target-repair` 또는 reuse-choice drill이다.

## 2026-05-25 — Daily route를 primary 복구 대상으로 구체화

이전 가정: 저장 포맷 변경 전에는 review queue 개수만 읽어 `오늘의 복구 루트: N건 대기`로 표시한다.

새 가정: 저장 포맷은 그대로 두되, review queue의 primary 대상과 이유를 함께 읽어 `오늘의 복구 루트: 목표 문자까지 이동하기(미복구) 외 2건 대기`처럼 표시한다.

이유: 단순 개수는 반복 학습 동기를 만들기 약하므로, 플레이어가 오늘 무엇부터 복구해야 하는지 즉시 알아야 한다.

영향: progress v2 없이도 daily loop는 더 선명해졌고, mastery/due date 저장 여부는 `PROGRESS-V2-DECISION-001`에서 별도 판단한다.

## 2026-05-23 — BUGFIX-REQUIRED-KEY-ROUTE 완료

이전 가정: 목표 상태에 도달했을 때 required key가 없으면 즉시 실패하는 방식으로 우회를 막을 수 있었다.

새 가정: redo 학습처럼 중간 상태가 목표와 같아지는 문항은 입력 여지가 남아 있으면 즉시 실패하지 않고 계속 진행할 수 있다. 단, 목표 상태에 머문 채 required key를 나중에 덧붙이는 우회는 실패한다.

이유: `undo-redo-basic-002`는 `ctrl+r` 자체가 학습 목표인데 숨은 `h` 입력을 요구했고, word motion 문항은 `h/l` 우회를 막는 constraints가 부족했다.

영향: `undo-redo-basic-002`는 `x`, `u`, `ctrl+r` 3입력으로 clear된다. `word-motion-basic-*`는 `w/b/e` required key와 h/l forbidden route를 갖는다.

## 2026-05-23 — UI-EVIDENCE-001 완료

이전 가정: E2E evidence의 `screen.txt`는 cleaned terminal text였지만, UI QA에서 이것이 최종 화면인지 누적 흐름인지 이름만으로 분명하지 않았다.

새 가정: UI QA용 누적 화면 흐름은 `screen_timeline.txt`로 명시 저장할 수 있다. `summary.json`도 `screen_timeline_evidence` 저장 여부를 기록한다.

이유: TUI 리디자인 반복 중 Agent가 화면 변화의 흐름을 재검토할 수 있어야 하고, 당장 불안정한 viewport parser를 도입하지 않아도 evidence 이름을 명확히 분리할 수 있다.

영향: `TUI Experience Foundation` 중기 플랜은 완료됐다. 다음 UI 루프에서 필요하면 `screen_final.txt` 또는 `frames/*.txt`를 별도 구현한다.

## 2026-05-23 — UI-MODE-001 완료

이전 가정: tutorial과 incident 모두 action panel에서 `Coach: 훈련 키 ...`를 보여줬고, incident도 첫 화면에서 정답 key sequence를 과하게 드러낼 수 있었다.

새 가정: tutorial은 훈련 키를 직접 노출하고, incident는 running 상태에서 판단 cue를 우선한다. incident 실패 후에는 retry를 돕기 위해 필요한 키 힌트를 `복구 힌트`로 표시할 수 있다.

이유: tutorial은 command 습득이 목적이지만 incident는 이미 배운 command를 고르는 적용 run이므로, UI 언어도 학습 단계와 적용 단계를 구분해야 한다.

영향: 다음 활성 slice는 화면 변경 evidence를 보강하는 `UI-EVIDENCE-001`이다.

## 2026-05-23 — UI-HIERARCHY-001 완료

이전 가정: 화면 상단은 playlist, review queue, daily route, 현재 문항을 순서대로 보여줬고, 현재 목표보다 메타 정보가 먼저 읽힐 수 있었다.

새 가정: 화면은 current task first다. Header는 위치와 진행 상태를 압축하고, title/briefing을 바로 보여준 뒤 review/daily route는 `복구 현황` 보조 영역으로 낮춘다.

이유: Vim 학습 게임에서 플레이어가 가장 먼저 알아야 하는 정보는 "지금 무엇을 조작해야 하는가"이며, 반복 학습 동기는 그 다음이어야 한다.

영향: 다음 활성 slice는 tutorial과 incident action panel 언어를 분리하는 `UI-MODE-001`이다.

## 2026-05-23 — UI-RENDER-001 완료

이전 가정: playable 화면 출력은 `internal/playable.Model.View()` 안에서 game state, progress 표시, action panel 문자열과 함께 구성됐다.

새 가정: `internal/playableview`가 순수 renderer 책임을 가진다. `internal/playable`은 상태와 action line을 구성하고, 화면 문자열 조립과 selection/cursor 표시는 renderer에 위임한다.

이유: 실제 UI 정보 위계와 모드별 문구를 개선하기 전에 game loop와 rendering 책임을 분리해야 한다.

영향: 다음 활성 slice는 현재 exercise 목표를 상위로 올리는 `UI-HIERARCHY-001`이다.

## 2026-05-23 — UI-CONTRACT-001 완료

이전 가정: TUI 개선 방향은 `header + briefing + console + status + action/debrief` 수준의 권장 구조였지만, tutorial과 incident의 화면 책임 차이는 아직 상세하지 않았다.

새 가정: TUI 화면은 Header, Briefing, Console, Status Line, Action/Debrief Panel 5개 영역으로 나눈다. Tutorial은 훈련 키 노출을 허용하고, Incident는 판단 cue와 점진 힌트를 우선한다.

이유: 이후 renderer 분리와 정보 위계 개편이 흔들리지 않도록, 코드 변경 전에 화면 영역의 책임과 우선순위를 고정해야 한다.

영향: 다음 활성 slice는 동작 변경 없는 `UI-RENDER-001` renderer 분리다.

## 2026-05-23 — TUI Experience Foundation 시작

이전 가정: TUI는 기능과 E2E 검증을 중심으로 성장했고, 제품 UI/UX는 action panel 같은 개별 문제 해결 수준에서만 다뤄졌다.

새 가정: 다음 중기 플랜은 `TUI Experience Foundation`이다. 화면을 크게 바꾸기 전에 typed app_state assertion, evidence snapshot, 화면 영역 contract, renderer 분리를 먼저 진행한다.

이유: SubAgent 리뷰 결과, 현재 화면은 Vim 훈련에는 충분하지만 tutorial/incident 정보 위계와 Runbook Dispatch 콘솔 감각이 아직 약하다. UI를 안정적으로 개선하려면 LLM이 의미 상태를 문자열이 아니라 typed state와 evidence로 확인할 수 있어야 한다.

영향: 첫 slice `UX-QA-001`은 review/daily route를 typed `app_state.review`와 `app_state.json`/`progress.json` evidence로 검증한다. 다음 활성 slice는 `UI-CONTRACT-001`이다.

## 2026-05-23 — PLATFORM-REVIEW-002 완료

이전 가정: review queue는 메인 화면과 성공 debrief에 표시됐지만, 매일 짧게 다시 들어올 이유를 주는 daily framing은 아직 progress schema v2 후보에 머물러 있었다.

새 가정: 저장 포맷 변경 전에도 앱 실행 시 review queue를 읽어 `오늘의 복구 루트: N건 대기`로 표시할 수 있다. E2E state summary도 review count와 primary candidate를 노출한다.

이유: daily streak/calendar를 저장하지 않아도 플레이어에게 "오늘 처리할 잔류 리스크" 감각을 줄 수 있고, progress schema v2 승인 전까지 안전하게 반복 플레이 동기를 강화할 수 있다.

영향: `Linewise Visual and Choice Training` 중기 플랜은 완료됐다. 다음 루프는 새 중기 플랜 수립으로 시작한다.

## 2026-05-23 — COMMAND-CHOICE-001 완료

이전 가정: incident 004까지 완료되어 배운 command는 늘었지만, 플레이어가 어느 상황에서 어떤 command를 고를지 훈련하는 별도 기준은 아직 없었다.

새 가정: command choice drill은 새 command cluster가 아니라 기존 implemented cluster를 섞어 범위, 재사용, 검색 후 편집, range command 선택을 훈련하는 적용 레이어다.

이유: Vim 학습이 key sequence 암기에서 멈추지 않으려면 같은 목표를 여러 방식으로 풀 수 있는 상황에서 가장 적합한 도구를 고르는 판단을 훈련해야 한다.

영향: 다음 활성 slice는 저장 포맷 변경 없이 review/daily 동기를 강화하는 `PLATFORM-REVIEW-002`다.

## 2026-05-23 — INCIDENT-RUN-004 완료

이전 가정: linewise visual은 tutorial까지 연결됐지만, incident 적용 run에서는 아직 실제 config block 복구 도구로 쓰이지 않았다.

새 가정: `incident-004-linewise-block-recovery`는 `/block`, linewise `Vd`, linewise `Vy+p`, linewise `VGd`, `:%s`를 하나의 Runbook Dispatch 적용 런으로 검증한다.

이유: linewise visual은 줄 묶음 삭제/복사에 실무 체감이 크므로, tutorial 직후 config block 복구 run으로 적용력을 강화한다.

영향: 다음 활성 slice는 배운 command 중 적절한 도구를 고르는 `COMMAND-CHOICE-001`이다.

## 2026-05-23 — PLAYPACK-011 완료

이전 가정: linewise `V`는 engine/runtime/TUI에서 구현됐지만 playable tutorial에는 아직 연결되지 않았다.

새 가정: `tutorial-93-visual-line`은 `Vd`, `Vy` + `p`, `VGd`를 다루는 3문항 linewise visual tutorial이며 full E2E를 통과한다.

이유: linewise visual을 incident에 적용하기 전에 v와 V의 단위 차이, linewise register, 파일 끝 tail 삭제를 짧게 학습해야 한다.

영향: 다음 활성 slice는 linewise visual을 적용하는 `INCIDENT-RUN-004`다.

## 2026-05-23 — VIM-029 완료

이전 가정: visual mode는 같은 줄 charwise `v` + `d/y`만 구현했고, linewise `V`는 문서 후보였다.

새 가정: linewise `V`는 engine/runtime/TUI에서 구현됐다. 첫 범위는 `V` 진입/해제, `j/k/G/gg` row motion, linewise `d/y`, unnamed linewise register, app_state/TUI `selection.kind: linewise`다.

이유: 설정 블록 단위 삭제/복사는 실무 학습 체감이 크고, 기존 linewise register/put 모델과 안정적으로 연결된다.

영향: 다음 활성 slice는 3문항 이하 tutorial과 full E2E를 추가하는 `PLAYPACK-011`이다.

## 2026-05-23 — VISUAL-LINE-001 완료

이전 가정: linewise `V`를 먼저 구현하기로 했지만, row motion, normalized range, register kind, cursor landing 기준이 구현 전 문서로 충분히 고정되지 않았다.

새 가정: 첫 linewise visual 구현 범위는 `V` 진입/해제, `j/k/gg/G` row motion, linewise `d/y`, unnamed linewise register, TUI/app_state `selection.kind: linewise`다.

이유: 엔진 구현이 selection 정규화, register, undo, TUI/E2E 표면을 동시에 건드리므로 Red 테스트 기준을 먼저 좁혀야 한다.

영향: 다음 활성 slice는 `VIM-029`다.

## 2026-05-23 — PLAN-REFRESH-005 완료

이전 가정: `Post-Visual Applied Mastery and Hardening` 완료 후 다음 우선순위가 linewise visual, command choice drill, platform review loop 사이에 열려 있었다.

새 가정: 다음 중기 플랜은 `Linewise Visual and Choice Training`이다. 우선 linewise `V`를 테스트 우선으로 구현하고, tutorial과 incident 적용 run으로 승격한 뒤 command choice와 review 동기를 보강한다.

이유: linewise `V`는 설정 블록 삭제/복사 학습 가치가 크고, 이미 있는 linewise register/put 모델과 잘 맞는다. 엔진 복잡도가 높으므로 scope approval과 Red 테스트를 먼저 둔다.

영향: 다음 활성 slice는 `VISUAL-LINE-001`이다.

## 2026-05-22 — VISUAL-LINE-GAP-001 완료

이전 가정: charwise visual 이후 linewise `V`와 multi-line charwise visual 중 무엇을 먼저 열지 결정되지 않았다.

새 가정: 다음 visual 후보는 linewise `V` + row motion + `d/y`다. multi-line charwise `v` operator와 visual block은 linewise 구현 이후로 미룬다.

이유: linewise `V`는 설정 블록 삭제/복사 학습 가치가 크고, 기존 linewise register/put 모델과 연결되어 검증 표면이 더 안정적이다.

영향: `Post-Visual Applied Mastery and Hardening` 중기 플랜은 완료됐다. 다음 루프는 새 중기 플랜 수립으로 시작한다.

## 2026-05-22 — INCIDENT-FLOW-001 완료

이전 가정: incident 001~003은 기능적으로 완주 가능했지만, 일부 beat 문구가 독립 문제처럼 읽힐 수 있었다.

새 가정: incident briefing/success/failure는 `진입 조치 -> 후속 조치 -> 마감 조치` 흐름을 화면 문구에서 직접 드러낸다. target state, optimal keys, constraints는 변경하지 않는다.

이유: Runbook Dispatch 프레임을 얹되 Vim 학습 목표를 가리지 않으려면 별도 narrative screen보다 각 beat 문구의 연결감을 높이는 편이 안전하다.

영향: 다음 활성 slice는 linewise `V`와 multi-line visual 범위를 구현 전 결정하는 `VISUAL-LINE-GAP-001`이다.

## 2026-05-22 — INCIDENT-RUN-003 완료

이전 가정: visual selection은 tutorial에서만 검증됐고, incident 적용 런에서는 아직 실제 복구 도구로 쓰이지 않았다.

새 가정: `incident-003-visual-recovery`는 `/contam`, visual delete/yank, backward visual delete, `:%s`를 하나의 Runbook Dispatch 적용 런으로 검증한다.

이유: visual selection을 튜토리얼 전용 기능이 아니라 실제 복구 상황에서 쓰는 도구로 승격하기 위해서다.

영향: 다음 활성 slice는 incident 001~003의 연결 문구와 continuity를 다듬는 `INCIDENT-FLOW-001`이다.

## 2026-05-22 — charwise visual invariant 테스트 보강

이전 가정: charwise visual foundation과 `d/y` engine은 동작했지만, 빈 줄, 줄 경계, undo/register, unsupported multi-line 같은 edge/invariant 검증이 넓지 않았다.

새 가정: 같은 줄 charwise visual의 현재 behavior는 engine/runtime 테스트로 더 촘촘히 고정한다. 새 visual 기능은 추가하지 않았다.

이유: visual을 incident에 적용하기 전에 selection/range/undo/register 회귀를 빠르게 잡을 안전망이 필요하다.

영향: `VISUAL-HARDEN-001`은 완료됐고, 다음 활성 slice는 visual을 실제 적용 런으로 승격하는 `INCIDENT-RUN-003`이다.

## 2026-05-22 — vimengine visual/selection helper 분리

이전 가정: visual selection type, visual key handling, visual `d/y` helper가 `internal/vimengine/engine.go`에 함께 있었다.

새 가정: selection type/normalization은 `selection.go`, visual mode key handling과 charwise visual operator helper는 `visual.go`로 분리한다. behavior는 바꾸지 않는다.

이유: visual 이후 linewise/multi-line 확장을 검토하기 전에 engine 파일의 책임을 낮춰야 한다.

영향: `ENGINE-SPLIT-001`은 완료됐고, 다음 활성 slice는 charwise visual invariant 테스트를 보강하는 `VISUAL-HARDEN-001`이다.

## 2026-05-22 — content replay에서 selection assertion 검증

이전 가정: visual selection은 app_state/E2E assertion으로 검증할 수 있었지만, content replay gate는 `e2e_assertions.selection` mismatch를 직접 비교하지 않았다.

새 가정: `replay_status: pass` exercise가 selection assertion을 선언하면 content loader가 optimal key replay 결과의 selection active/kind/anchor/head/start/end를 비교한다.

이유: visual 관련 content가 화면 문구나 최종 buffer만으로 통과하면 selection anchor/head/range 회귀를 놓칠 수 있다.

영향: `QA-SEL-001`은 완료됐고, 다음 활성 slice는 behavior 변화 없는 `ENGINE-SPLIT-001`이다.

## 2026-05-22 — post-visual 중기 플랜을 적용력 강화로 전환

이전 가정: visual selection tutorial 완료 후 다음 후보는 linewise visual이나 또 다른 Vim command 확장일 수 있었다.

새 가정: 다음 중기 플랜은 `Post-Visual Applied Mastery and Hardening`이다. 우선순위는 selection replay hardening, vimengine 내부 분리, charwise visual invariant 보강, visual 적용 incident, incident continuity, linewise visual gap planning 순서다.

이유: Architecture, Learning/Game UX, Scenario/World 리뷰 모두 visual 이후 바로 기능을 넓히기보다 검증/분리/적용력을 먼저 강화해야 한다고 판단했다.

영향: `QA-SEL-001`이 다음 활성 slice다. progress schema v2, daily run, spaced review, linewise visual 구현은 별도 승인 또는 후속 플랜으로 둔다.

## 2026-05-22 — visual selection tutorial 완성

이전 가정: visual `d/y` engine은 구현됐지만 플레이 가능한 tutorial과 E2E가 아직 없었다.

새 가정: `visual-char-line`은 approved + implemented command cluster이며, `tutorial-92-visual-selection`은 3문항으로 forward deletion, visual yank-put, backward selection normalization을 훈련한다.

이유: visual mode는 선택 상태만 이해하는 것보다 “보이는 범위를 잡고 삭제/복사한다”는 반복 문항이 있어야 실제 학습 단위로 완성된다.

영향: 다음 루프는 visual hardening으로 바로 확장하기보다, 현재 튜토리얼/incident 흐름을 기준으로 다음 command 또는 게임성 보강 후보를 재점검한다.

## 2026-05-22 — charwise visual `d/y` engine 구현

이전 가정: visual mode는 selection을 만들고 표시할 수 있었지만, 선택 범위에 실제 Vim operator를 적용하지 못했다.

새 가정: 같은 줄 charwise visual selection에는 `d`와 `y`를 적용할 수 있다. `d`는 선택 범위를 삭제하고 register에 저장하며 undo 가능하다. `y`는 선택 범위를 unnamed charwise register에 저장한다.

이유: visual mode 학습은 “눈으로 범위를 잡고 삭제/복사한다”는 행동까지 이어져야 실제 Vim command로서 의미가 있다.

영향: 다음 루프는 `PLAYPACK-010`으로 3~4문항 visual tutorial을 만든다. multi-line, linewise `V`, visual block은 후속 hardening이다.

## 2026-05-22 — visual operator 범위를 같은 줄 `d/y`로 고정

이전 가정: visual foundation 이후 operator 적용과 tutorial 중 무엇을 먼저 할지 열려 있었다.

새 가정: 다음 구현은 같은 줄 charwise visual selection에 `d`와 `y`를 적용하는 `VIM-028`이다. content/tutorial은 engine 통과 이후 `PLAYPACK-010`으로 분리한다.

이유: visual mode가 학습 단위로 완성되려면 “선택 후 삭제/복사”가 필요하지만, multi-line/linewise/block까지 한 번에 넣으면 검증 표면이 커진다.

영향: `V`, visual block, count/register prefix, multi-line charwise operator, dot repeat 연계는 후속 hardening으로 둔다.

## 2026-05-22 — charwise visual foundation 구현

이전 가정: visual mode는 selection 계약과 app_state assertion까지만 준비되어 있었다.

새 가정: engine은 `ModeVisual`과 charwise selection state를 가지며, `v` 진입, motion으로 head 이동, `esc`/`v` reset을 지원한다. TUI는 selection range와 selected cell 표시를 제공한다.

이유: visual mode는 operator 적용 전에도 “범위를 눈으로 지정한다”는 학습 감각이 필요하며, 후속 `d/y` 적용과 playpack을 위한 최소 platform state가 필요하다.

영향: 다음 계획은 visual selection에 `d/y`를 적용할지, 먼저 visual tutorial playpack draft를 만들지 결정해야 한다.

## 2026-05-22 — selection app_state assertion 추가

이전 가정: app_state summary는 buffer, cursor, mode, status, score, progress만 검증했다.

새 가정: app_state summary와 runner assertion은 optional `selection` object를 지원한다. content `e2e_assertions.selection`도 같은 shape를 보존한다.

이유: visual mode는 화면 텍스트만으로 검증하면 selection anchor/head/range가 빠질 수 있으므로, 구현 전에 관찰 가능한 상태 계약이 필요하다.

영향: 다음 루프는 `VIM-027-TUI-003`으로 `v` charwise selection, motion update, `esc` reset, 최소 TUI 표시를 구현한다.

## 2026-05-22 — visual selection 계약을 charwise `v`로 축소

이전 가정: visual mode 후보는 `v`, `V`, `d`, `y`를 함께 검토하되 첫 구현 범위가 아직 열려 있었다.

새 가정: 첫 구현은 charwise `v`만 다룬다. selection은 `kind`, `anchor`, `head`, normalized inclusive `start`/`end`로 표현하고, E2E는 app_state `selection` object를 검증한다.

이유: visual mode는 화면 표시와 selection state가 학습 이해에 직접 영향을 주므로, operator 적용보다 state/render/assertion 계약을 먼저 고정하는 편이 안정적이다.

영향: 다음 루프는 `E2E-007`로 `e2estate`, runner assertion, content assertion schema에 selection을 추가한다.

## 2026-05-22 — E2E progress fixture builder 도입

이전 가정: 후반 playlist/incident E2E는 `setup.progress_file`에 긴 JSON을 직접 넣어 현재 exercise 직전까지의 progress를 만들었다.

새 가정: runner는 `setup.complete_before: <exercise-id>`를 지원하며, 현재 playable content 순서를 읽어 지정 exercise 직전까지 completed progress를 생성한다. 긴 inline JSON은 예외적 raw fixture가 필요할 때만 사용한다.

이유: content 순서가 늘어날수록 progress JSON 복붙이 낡기 쉽고, 후속 visual mode E2E 작성 시 fixture 비용이 커지기 때문이다.

영향: 다음 visual mode 검증 루프는 짧은 `complete_before` fixture를 사용해 app_state assertion에 집중한다.

## 2026-05-22 — 복구국 언어를 progress/debrief UI에 적용

이전 가정: review queue와 성공 debrief는 기능적으로 동작했지만 `Debrief`, `Best`, `Playlist`, `재진단 큐`처럼 플랫폼 내부 용어가 섞여 있었다.

새 가정: 저장 포맷은 그대로 두고 화면 언어만 `재점검 대상`, `잔류 리스크`, `복구 기록`, `최단 복구 기록`, `Runbook 복구 완료`로 표현한다.

이유: 반복 학습 UI가 세계관 밖의 점수판처럼 보이지 않고, 원격 시설 복구국의 작업 기록처럼 읽히게 하기 위함이다.

영향: 다음 루프는 `E2E-FIXTURE-001`로 긴 progress fixture 유지보수 부담을 줄인다.

## 2026-05-22 — 원격 시설 복구국 세계관 프레임 채택

이전 가정: 중반 이후 시나리오 톤은 터미널 문제 해결 생존 어드벤처였지만, canonical 세계관 명칭과 반복 플레이 언어는 아직 고정되지 않았다.

새 가정: 중반 이후 세계관 기본 프레임은 `원격 시설 복구국 / Runbook Dispatch`로 한다. 플레이어는 낡은 원격 시설의 장애 runbook을 Vim으로 복구하는 원격 복구 오퍼레이터다. 개별 incident에는 `침묵한 릴레이 기지` 감각을 얇게 섞되, 장대한 lore는 금지한다.

이유: 세 명의 시나리오 작가 SubAgent와 메타 리뷰 결과, 이 혼합안이 Vim command 학습을 가장 덜 가리면서 incident 001/002, review queue, visual mode 확장을 자연스럽게 묶는다고 판단했다.

영향: 다음 루프는 `INCIDENT-UX-003`으로 incident 제목, briefing, feedback, 2단계 hint를 복구 작전 프레임에 맞춘다.

## 2026-05-22 — Structure editing 완료 후 visual readiness와 incident UX 우선

이전 가정: `text-object-quote-pair`와 두 번째 incident run을 닫은 뒤 다음 후보인 `visual-char-line` 구현으로 바로 넘어갈 수 있다.

새 가정: visual mode는 다음 큰 Vim capability 후보가 맞지만, 구현 전에 selection state contract, TUI rendering, content/E2E app_state assertion을 먼저 닫는다. 또한 incident run은 기능적으로 통과했지만 아직 종합시험 감각이 남아 있으므로, 새 command 추가 전에 incident game-feel과 2단계 힌트 보강을 검토한다.

이유: Architecture Review와 Learning/Game UX Review 모두 visual mode의 표시/검증 계약 선행 필요성과 incident 적용 런의 게임성 보강 필요를 지적했다.

영향: 다음 중기 플랜 후보는 `INCIDENT-UX-003`, `E2E-FIXTURE-001`, `VISUAL-GAP-002`, `E2E-007`, `VIM-027/TUI-003` 순서로 검토한다.

## 2026-05-21 — Utility command tutorial 완료와 장기 반복 학습 RFC 분리

이전 가정: `o/O`, `.`, `/ n N`를 단기 utility command 확장으로 다루되, 이후 장기 반복 학습 플랫폼의 저장 구조는 아직 정리되지 않았다.

새 가정: `open-line-edit`, `repeat-last-change`, `search-basic`은 각각 playable tutorial과 E2E까지 연결된 foundation cluster로 승격한다. 장기 반복 학습은 `PLATFORM-RFC-001`에서 저장 변경 없는 review 후보와 progress schema v2 후보를 분리한다.

이유: Vim 학습 플랫폼은 명령어 추가만으로 완성되지 않고, 복습과 숙련도 추적이 필요하다. 다만 사용자 로컬 저장 포맷은 영향이 크므로 schema 변경은 별도 승인 게이트 뒤에 둔다.

영향: 다음 루프는 저장 변경 없는 `PLATFORM-REVIEW-001`로 시작하거나, 사용자 승인 후 `PROGRESS-SCHEMA-001`로 progress v2를 설계한다.

## 2026-05-19 — 중반 생존 어드벤처 톤과 operator grammar 순서 확정

이전 가정: 중반부터 생존 어드벤처 비중을 높인다는 방향은 있었지만, 생존감의 범위와 다음 콘텐츠 시퀀스가 구체적으로 고정되지 않았다.

새 가정: 중반 톤은 `docs/gameplay/scenario-tone.md`의 “터미널 문제 해결 생존 어드벤처”를 따른다. 생존감은 전투, 체력, 인벤토리가 아니라 로그, 설정, 통신, 임시 패치, 저장/폐기 판단 같은 텍스트 조작 압력에서 온다. 다음 중기 플랜은 `delete-with-motion`, `change-with-motion` 중심의 operator grammar adventure intro로 진행한다.

이유: 작은 수정 튜토리얼 이후 Vim이 실무 도구처럼 느껴지려면 `operator + motion` 문법으로 넘어가는 것이 자연스럽다. 시나리오가 command 학습을 가리지 않도록 톤 가드레일을 먼저 고정했다.

영향: `docs/gameplay/scenario-tone.md`, `docs/gameplay/vim-curriculum-map.md`, `docs/roadmap/MIDTERM_TODO.md`, `docs/roadmap/PROGRAM.md`를 기준으로 다음 루프를 진행한다.

## 2026-05-18 — 튜토리얼 페이싱과 제약 설계 원칙 확정

이전 가정: 첫 playable pack은 17개 exercise를 한 번에 완주하는 vertical slice이며, 초반/중반 튜토리얼 분리와 실패/재시도/입력 제약 기준은 아직 열려 있었다.

새 가정: 초반은 8문항 이하의 짧은 튜토리얼 에피소드 묶음으로 운영한다. 초반은 “첫 투어” 느낌으로 command를 넓게 맛보게 하고, 중후반부터 생존 어드벤처와 탐험 비중을 높인다. `:s`, `:%s`, range substitute는 초반에서 빼고 중반 고급 튜토리얼로 분리한다. 최대 입력 수 초과와 금지 입력/금지 우회 전략 사용은 즉시 실패이며, 초반 재시도는 무제한 기본이다. 재시도는 `r`과 `enter`를 모두 허용한다.

이유: 첫 pack이 E2E vertical slice로는 유효하지만 실제 학습 단위로는 넓기 때문에, 플레이어가 부담 없이 익히는 짧은 episode 구조가 필요하다. 또한 Vim 학습 게임으로서 “목표에 도착했는가”뿐 아니라 “의도 command를 썼는가”를 강제해야 한다.

영향: `docs/roadmap/decisions/0004-tutorial-pacing-and-constraints.md`, `docs/gameplay/spec.md`, `docs/gameplay/vim-curriculum-map.md`, `docs/workflows/scenario-flow-workbench.md`를 기준으로 다음 구현 루프를 진행한다.

## 2026-05-18 — 콘텐츠/커리큘럼 기본값 확정

이전 가정: CONTENT-001의 파일 포맷, 위치, draft 파일 정책, 다음 엔진 확장 순서, scenario tone, coverage 예외, 자동 저장 기준이 아직 열려 있었다.

새 가정: 콘텐츠는 YAML 우선이며 repo root `content/`에 둔다. draft 파일은 허용하지만 playable에는 approved/implemented 및 `engine_support: implemented`만 연결한다. 첫 커리큘럼 순서는 `normal-motion-basic -> survival-save-quit -> word-motion-basic`이고, 다음 엔진 확장은 `w/b/e`를 우선한다. 후반에는 `gg`, `G`, `:` command, substitute, range command까지 폭넓게 다룬다. 시나리오 톤은 DevOps/터미널 문제 해결이며 과하지 않은 억까 상황을 허용한다. 진행 사항은 자동 저장한다.

이유: CONTENT-001과 이후 game loop 구현 전에 사용자 결정을 명시적으로 고정해 Agent가 임의 선택하지 않게 하기 위함이다.

영향: `docs/roadmap/decisions/0003-content-and-curriculum-defaults.md`, `docs/gameplay/content-requirements.md`, `docs/workflows/scenario-production-harness.md`를 기준으로 다음 루프를 진행한다.

## 2026-05-17 — Harness-first 재기획으로 전환

이전 가정: 기존 `docs/archived/PLAN.md`와 `docs/archived/GAME_DESIGN.md`를 이어서 구현한다.

새 가정: 기존 구현은 참고 자료로 두고, Go 기반 TUI adventure game이라는 큰 방향만 유지한 채 제품 기획과 검증 워크플로우를 다시 세팅한다.

이유: 현재 Advimture의 제품 방향과 재미가 충분하지 않다고 판단했다.

영향: 새 구현 전 `docs/roadmap/`, `docs/exec-plans/`, `docs/gameplay/`, `docs/verification/`를 기준으로 작업한다.
