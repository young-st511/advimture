# Post Second Dispatch Playtest Review — 2026-07-02

## Summary

`CONTENT-BREADTH-003` 이후 fresh evidence 기준으로 P0 진행 불가 이슈는 보이지 않는다. Second Dispatch Pack은 검색 후 inline 수리, 검증값/줄 재사용, 치환 범위 판단을 각각 짧은 Runbook Dispatch로 읽히게 한다.

다만 완성도를 더 올리려면 신규 콘텐츠 추가보다 아래 세 축이 우선이다.

1. `UI-HINT-LADDER-001`: hint가 exact command spoiler로 너무 빨리 작동하지 않게 한다.
2. `REDTEAM-QA-MATRIX-001`: 우회 가능성 검증을 incident별 matrix로 확장한다.
3. `UI-REPORT-001`: success/failure report의 정보 밀도와 다음 행동 위계를 더 다듬는다.

## Evidence

- `artifacts/e2e/playable_incident_009_full/screen_final.txt`
- `artifacts/e2e/playable_incident_010_full/screen_final.txt`
- `artifacts/e2e/playable_incident_011_full/screen_final.txt`
- `artifacts/e2e/playable_incident_011_redteam_scope_guard/screen_final.txt`
- `artifacts/e2e/playable_incident_011_redteam_scope_guard/app_state.json`
- Commit baseline: `fe2a5e6 Add second dispatch content pack`

## Findings

### P0

없음. 009~011 full route와 011 RedTeam guard는 진행 불가나 progress 오염 없이 완료/실패 상태를 구분한다.

### P1

1. Hint가 학습 ladder라기보다 즉시 도움말에 가깝다.
   - 현재 running hint는 `힌트 내용`으로 잘 분리되지만, 플레이어가 막히는 순간 exact command family와 정답형 문구가 빠르게 보일 수 있다.
   - incident의 기본 running 화면은 command memory를 숨기지만, hint/failure 후에는 `참고 명령`이 즉시 드러난다.
   - 개선 방향: 첫 hint는 판단 관점, 둘째 hint는 command family, 셋째 이후 exact form으로 공개한다.

2. RedTeam 검증 범위가 아직 대표 1개에 가깝다.
   - `playable_incident_011_redteam_scope_guard`는 `%` over-scope route를 잘 막는다.
   - 하지만 직접 재입력 우회, delimiter 포함/삭제 실수, required key 생략, under-scope 치환 같은 유형은 matrix로 체계화되어 있지 않다.
   - 개선 방향: 대표 우회 유형을 문서화하고 최소 2개 이상의 E2E를 추가한다.

3. Success/failure report는 안정적이지만 report density가 여전히 높다.
   - 009/010 success report는 `배운 점`, `기록`, `Runbook`, `잔류 리스크`, `다음 출격 후보`, `다음 행동`을 모두 담아 의미는 충분하다.
   - 011 final success는 가장 깔끔하지만, 일반 success report에서는 “배운 점”과 “다음 행동” 사이에 정보가 많아질 수 있다.
   - 개선 방향: report grammar를 유지하되 80x24에서 배운 점/다음 행동을 더 우선 배치한다.

### P2

1. 009~011은 서로 독립된 dispatch로는 명확하지만, “Second Dispatch Pack”이라는 상위 arc 이름은 아직 화면에 직접 드러나지 않는다.
2. 실패 report는 좋지만 긴 실패 문구와 runtime failure message가 한 줄 그룹에 합쳐질 때 읽기 밀도가 올라간다.

## Recommended Order

1. `UI-HINT-LADDER-001`
2. `REDTEAM-QA-MATRIX-001`
3. `UI-REPORT-001`
4. 이후 fresh evidence를 보고 세 번째 dispatch pack 또는 higher-level arc presentation을 결정한다.

## Out of Scope

- release/tag/push
- progress schema 변경
- content schema 변경
- 새 Vim engine capability
- 새 dependency
