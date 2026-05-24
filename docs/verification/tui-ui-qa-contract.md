# TUI UI QA Contract

> UI 개선이 진행되어도 Agent가 화면 의미를 안정적으로 검증하기 위한 계약이다.

## 원칙

- 화면 문구는 UX 개선 과정에서 바뀔 수 있다.
- 핵심 상태는 `app_state` typed assertion으로 검증한다.
- E2E evidence는 임시 HOME 삭제 후에도 Agent가 다시 읽을 수 있어야 한다.
- 실제 사용자 HOME과 progress file은 사용하지 않는다.

## Review State Assertion

review/daily route 영역은 다음 shape로 검증한다.

```yaml
assert:
  app_state:
    review:
      queue_count: 3
      primary_exercise_id: normal-motion-basic-002
      primary_reason: incomplete
      daily_route: "오늘의 복구 루트: 경고 지점으로 이동하기(미복구) 외 2건 대기"
```

문자열 `contains`는 fallback으로만 사용한다. 새 UI가 review 문구를 다른 위치로 옮겨도 `review` state가 유지되면 테스트는 안정적으로 통과해야 한다.

## Evidence Snapshot

E2E runner는 기존 `summary.json`, `raw.log`, `screen.txt`, `key_trace.txt`에 더해 다음 snapshot을 남길 수 있어야 한다.

- `app_state.json`: test HOME의 `.advimture/e2e_state.json` 복사본
- `progress.json`: test HOME의 `.advimture/progress.json` 복사본
- `screen_timeline.txt`: cleaned terminal text의 누적 흐름 evidence. UI 위계 변경 후 Agent가 최종 문구뿐 아니라 지나온 화면 흐름을 확인할 때 사용한다.

이 snapshot은 temp HOME이 삭제된 뒤에도 실패 원인을 확인하기 위한 evidence다.

## 다음 후보

- `residual_risk` typed state: 성공 후 다음 재점검 대상의 exercise id/reason 검증
- `screen_final.txt`: 누적 ANSI stream이 아니라 마지막 viewport 중심 evidence
- `frames/*.txt`: wait/send 단위 세분화 frame timeline
