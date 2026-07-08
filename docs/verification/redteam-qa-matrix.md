# RedTeam QA Matrix

## 목적

Advimture의 RedTeam QA는 “화면상 성공처럼 보이는가”보다 “요구한 Vim 판단을 만족하지 않고도 clear 가능한가”를 검증한다. 각 fixture는 progress 저장 여부, app_state status, key trace를 함께 확인한다.

## Matrix

| 유형 | 대표 위험 | 현재 대표 E2E | 기대 결과 |
|---|---|---|---|
| range over-scope | 현재 줄만 바꿔야 하는데 `%` 전체 범위를 사용한다. | `playable_incident_011_redteam_scope_guard` | failed, progress 미완료 |
| delimiter 포함 삭제 | delimiter를 보존해야 하는데 `cf,`로 comma까지 포함해 change한다. | `playable_incident_009_redteam_delimiter_guard` | failed, progress 미완료 |
| 직접 재입력 우회 | 검증값을 register로 재사용해야 하는데 직접 typing한다. | `playable_incident_010_redteam_retype_guard` | failed, progress 미완료 |
| under-scope | 전체 파일을 바꿔야 하는데 현재 줄만 치환한다. | 후보 | 성공 불가, 필요 시 failed evidence 추가 |
| required key 생략 | target state만 맞춘 뒤 의도 command를 쓰지 않는다. | 기존 constraint required-key E2E | failed 또는 running 유지 후 failed |

## 운영 규칙

- RedTeam E2E는 반드시 `setup.home: temp` 또는 `setup.complete_before`를 사용한다.
- RedTeam E2E는 최소한 `status: failed`, `progress.completed: false`, key trace를 검증한다.
- 새 content schema, progress schema, engine capability를 RedTeam을 위해 추가하지 않는다.
- RedTeam이 UX blind spot을 발견하면 해당 blind spot을 gameplay/verification spec에 승격할지 검토한다.
