# Playtest Review — 2026-05-28 Applied Mastery Runs

## 결론

Applied Mastery Runs는 다음 단계로 넘어갈 수 있다. `INCIDENT-006`, `CHOICE-PLAY-003`, `INCIDENT-007` 모두 content replay, focused E2E, 전체 Go test, 전체 playable E2E를 통과했다.

다만 다음 작업은 새 content 확장보다 E2E evidence 품질 보강을 먼저 권장한다. 긴 incident의 key trace와 app_state 검증은 충분하지만, summary의 `screen_timeline_evidence`와 `screen_final_evidence`가 false로 남아 UI/UX 회귀를 사람이 훑기에는 증거가 약하다.

## 완료 범위

- `PLAN-REFRESH-008`: Applied Mastery Runs 중기 플랜 고정
- `INCIDENT-006`: `/target`, `n`, `ct,` applied incident
- `REUSE-CHOICE-001`: quote value reuse-choice 설계
- `CHOICE-PLAY-003`: `yi"` + `P` quote value reuse-choice playable
- `INCIDENT-007`: `/breach`, `ci"`, `Vj d`, `ct,`, `:%s` mixed recovery run

## 검증 결과

- `go test ./internal/content`: pass
- `go test ./...`: pass
- focused E2E:
  - `playable_incident_006_full`: pass
  - `playable_command_choice_scope`: pass
  - `playable_incident_007_full`: pass
- `make e2e-playable`: pass
- app_state/key_trace/progress evidence: pass

## RedTeam Notes

- `incident-007`은 5-beat mixed run이라 학습 부하가 높다. 현재는 각 beat가 하나의 판단만 요구하므로 허용 가능하지만, 다음 mixed run에서 beat 하나에 여러 새 판단을 섞으면 품질이 빠르게 떨어질 수 있다.
- command-choice incident가 4-beat가 되면서 `incident-005`의 역할이 넓어졌다. 앞으로 command-choice beat를 더 붙이기보다 새 command-choice incident로 분리하는 편이 낫다.
- loader test가 총 exercise/scenario 개수에 의존한다. content가 늘수록 유지 비용이 커지지만, 현재는 content drift를 빨리 잡는 장점이 더 크다.

## UX/Evidence Notes

- 긴 incident의 final app_state와 key trace는 검증된다.
- cleaned `screen.txt`는 저장되지만 summary의 `screen_timeline_evidence`와 `screen_final_evidence`는 false다.
- UI/UX 회귀를 점검하려면 다음 QA 루프에서 긴 incident의 final screen 또는 timeline evidence를 명시적으로 저장/요약하는 편이 좋다.

## 다음 후보

1. `E2E-EVIDENCE-008`: long incident final/timeline evidence를 summary에 명시하고, representative route에서 final screen을 빠르게 리뷰 가능하게 만든다.
2. `INCIDENT-008` 또는 새 engine gap planning은 evidence 보강 이후로 둔다.

추천: `E2E-EVIDENCE-008`을 먼저 수행한다.
