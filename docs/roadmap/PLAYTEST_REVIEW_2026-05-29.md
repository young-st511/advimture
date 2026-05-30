# Playtest Review — Long Incident Evidence

Date: 2026-05-29
Scope: `E2E-EVIDENCE-008`
Superseded-by: `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`

## 결론

긴 incident route의 UI/UX 재검토를 막던 evidence 부족은 해소됐다. runner 기능을 새로 만들지 않고, long route fixture가 `screen_timeline.txt`와 `screen_final.txt`를 항상 남기도록 고정했다.

이 review의 다음 후보 판단은 `FOUNDATION-EXIT-001`에서 후속 검토됐다. 현재 다음 권장은 `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`를 따른다.

## 확인한 Evidence

- `playable_command_choice_scope`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
- `playable_incident_006_full`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
- `playable_incident_007_full`: pass, `screen_timeline_evidence: true`, `screen_final_evidence: true`
- full playable E2E 이후 `playable_incident_001_full`, `002`, `003`, `004`, command-choice scope, `006`, `007` summary flag 모두 true
- `make e2e-playable`: pass

`playable_incident_007_full/screen_final.txt`는 마지막 `ADVIMTURE |` frame, runbook console, success modal, status line, command line을 사람이 읽을 수 있는 형태로 보존한다.

## RedTeam Notes

- `incident-007`은 5-beat mixed run이라 학습 부하가 높다. 다만 각 beat가 한 가지 판단을 요구하고 final screen evidence도 읽히므로 현재는 허용 가능하다.
- 다음 long incident를 추가할 때 evidence bundle이 빠지면 UI/UX 회귀를 다시 사람이 추측하게 된다. long route fixture에는 timeline/final evidence를 기본으로 요구한다.
- full E2E는 로컬/Agent 루프 기준으로 충분하다. CI 포함 여부는 아직 별도 결정으로 남긴다.

## 다음 후보

아래 후보 판단은 2026-05-29 당시의 review 결과다. 2026-05-30 현재는 `FOUNDATION-EXIT-001`이 완료됐고, 다음 권장은 `PLATFORM-REVIEW-003`이다.

1. `FOUNDATION-EXIT-001`: Foundation exit review와 다음 중기 플랜 수립. 지금까지의 engine/content/UI가 첫 제품 기둥을 충분히 세웠는지 점검한다.
2. `QUOTE-PAIR-HARDEN-001`: 새 engine을 연다면 `ci'`, `ci(`, `ci{` 같은 pair text object 확장이 가장 작고 실무 효용이 높다.
3. `PLATFORM-REVIEW-003`: 저장 포맷 변경 없이 mission map/review entry를 더 게임답게 묶을 수 있는지 검토한다.

최신 권장은 `docs/roadmap/FOUNDATION_EXIT_REVIEW_2026-05-30.md`와 `docs/roadmap/FORWARD_PLAN.md`를 따른다.
