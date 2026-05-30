# PLATFORM-RFC-001 — Long-Run Review Platform

Status: accepted planning baseline
Date: 2026-05-21
Last reviewed: 2026-05-30

## 목적

Advimture를 한 번 깨고 끝나는 튜토리얼 모음이 아니라, Vim 명령을 장기적으로 반복 학습하는 플랫폼으로 확장하기 위한 진행/복습/일일 런 구조를 정의한다.

이 RFC는 구현 승인이 아니다. `internal/progress/` 저장 JSON 구조 변경은 별도 ExecPlan과 사용자 승인을 받은 뒤에만 진행한다.

2026-05-30 현재 상태:

- 저장 변경 없는 review queue, daily motivation, success debrief, final dispatch review handoff는 구현됐다.
- progress schema v2는 여전히 보류다.
- 다음 platform 작업은 새 저장 필드가 실제 병목으로 확인될 때 다시 연다. 현재 다음 제품 작업은 `CONTENT-BREADTH-002`다.
- schema v2는 실패 attempt 지속 저장, mastery level, review due date, daily history가 실제 제품 병목으로 확인될 때만 다시 연다.

## 현재 가능한 범위

현재 progress 저장 모델은 `Missions` map에 exercise ID별 완료 여부, best grade, best keystrokes, best time, attempts를 저장한다.

저장 포맷 변경 없이 가능한 기능:

- 성공 화면과 playlist 완료 화면에서 best grade, best key count를 비교한다.
- 완료한 exercise 중 `best_grade != S` 또는 best keystrokes가 optimal보다 큰 항목을 복습 후보로 계산한다.
- command cluster별 완료 수와 S 등급 비율을 runtime에서 계산한다.
- 오늘의 추천을 저장하지 않고, 앱 실행 시 content library와 기존 progress를 조합해 표시한다.

저장 포맷 변경 없이 어려운 기능:

- 실패했지만 아직 성공하지 못한 exercise를 다음 세션에 복구 후보로 유지한다.
- command별 숙련도, 마지막 복습일, 다음 복습 예정일을 추적한다.
- daily streak, daily run history, spaced review interval을 안정적으로 저장한다.
- exercise별 여러 번의 시도 이력과 실수 유형을 분석한다.

## 핵심 개념

### Mastery

목표: 플레이어가 특정 Vim command cluster를 얼마나 자연스럽게 사용할 수 있는지 추정한다.

첫 기준:

- coverage: cluster의 required command를 모두 한 번 이상 성공했는가
- quality: 해당 cluster exercise의 best grade가 S/A 중심인가
- efficiency: best keystrokes가 optimal에 가까운가
- retention: 일정 시간이 지난 뒤에도 성공하는가

### Recovery

목표: 실패, 낮은 grade, 과한 key count를 다음 플레이 이유로 바꾼다.

첫 기준:

- 현재 세션 실패는 즉시 retry UX로 복구한다.
- 저장 포맷 변경 전에는 완료된 exercise의 낮은 best 기록만 복습 후보로 삼는다.
- 저장 포맷 변경 후에는 failed attempt와 mistake focus까지 복습 후보로 저장할 수 있다.

### Spaced Review

목표: 이미 배운 command를 시간이 지난 뒤 다시 꺼내게 한다.

첫 기준:

- S grade + optimal에 가까운 입력은 review interval을 늘린다.
- 실패, forbidden input, required key missing은 review interval을 줄인다.
- review due 계산은 command cluster 단위와 exercise 단위를 모두 지원해야 한다.

### Daily Run

목표: 매일 짧게 플레이할 수 있는 반복 루프를 만든다.

첫 기준:

- 5~8문항 이하로 유지한다.
- 1개 신규 command, 2~3개 복습, 1개 efficiency challenge를 섞는다.
- 실패해도 daily run 자체를 망치지 않고 recovery queue로 보낸다.

## 저장 포맷 후보

아래 구조는 후보이며, 아직 구현하지 않는다.

```json
{
  "schema_version": 2,
  "missions": {},
  "review": {
    "exercise-id": {
      "last_attempted_at": "2026-05-21T00:00:00Z",
      "last_succeeded_at": "2026-05-21T00:00:00Z",
      "last_grade": "A",
      "last_failure_reason": "required_keys_missing",
      "review_due_at": "2026-05-24T00:00:00Z",
      "review_interval_days": 3
    }
  },
  "mastery": {
    "command-cluster-id": {
      "level": "introduced|practicing|comfortable|review",
      "s_success_count": 3,
      "recent_failure_count": 0,
      "last_practiced_at": "2026-05-21T00:00:00Z"
    }
  },
  "daily_runs": {
    "2026-05-21": {
      "completed": true,
      "exercise_ids": ["search-basic-001"],
      "grade": "A"
    }
  }
}
```

## 추천 구현 순서

### PLATFORM-REVIEW-001: 저장 변경 없는 review 후보 표시

Status: completed

- `Missions`와 content library만 사용한다.
- 낮은 best grade, 높은 key count, incomplete exercise를 추천한다.
- 새 progress 필드를 만들지 않는다.
- E2E는 temp HOME progress fixture로 검증한다.

후속 no-schema 후보였던 `PLATFORM-REVIEW-003`은 완료됐다. 성공 debrief와 마지막 dispatch review handoff가 progress v1 위에서 동작한다.

### PROGRESS-SCHEMA-001: progress schema v2 승인/마이그레이션

- 사용자 승인 후에만 진행한다.
- `schema_version`, `review`, `mastery`, `daily_runs` 중 실제 필요한 필드만 최소 추가한다.
- v1 fixture를 v2로 읽고 저장하는 migration test를 만든다.
- 실제 `~/.advimture`가 아닌 테스트 HOME만 사용한다.

### REVIEW-ENGINE-001: spaced review 계산 엔진

- 입력: content library, progress snapshot, 현재 날짜
- 출력: due exercise 후보와 이유
- UI와 저장소에 의존하지 않는 순수 패키지로 둔다.

### DAILY-RUN-001: daily run playlist 생성

- 신규 학습, 복습, efficiency challenge를 5~8문항으로 섞는다.
- 같은 cluster가 과도하게 반복되지 않도록 제한한다.
- 실패한 항목은 recovery queue로 넘기는 정책을 둔다.

## 승인 게이트

다음 변경은 사용자 승인 전까지 금지한다.

- `internal/progress/` JSON field 추가, 삭제, 이름 변경
- 기존 progress 파일 자동 migration
- 실제 `~/.advimture`를 읽거나 쓰는 E2E
- streak, daily history처럼 사용자에게 손실감이 큰 게임성 수치 도입

다음 변경은 저장 포맷 변경 없이 먼저 가능하다.

- 완료 화면의 best record 비교 개선
- content library 기반 command cluster progress summary
- 현재 세션 안의 실패/재시도 debrief
- 저장하지 않는 임시 review recommendation 화면

## 2026-05-25 결정: schema v2 보류

FTUE route, command-choice playable, no-schema daily route를 검증한 결과, 현재 foundation 단계에서는 progress v1과 content library runtime 계산만으로 충분하다.

따라서 progress schema v2는 지금 구현하지 않는다. 다음 조건이 실제 제품 문제로 드러날 때 별도 ExecPlan과 사용자 승인으로 다시 연다.

- 실패했지만 아직 성공하지 못한 attempt를 다음 세션 복구 후보로 유지해야 한다.
- command cluster별 mastery level, last practiced, review due date가 필요하다.
- daily streak/history가 반복 플레이의 핵심 동기가 된다.
- exercise별 실수 유형이나 attempt history를 분석해야 한다.

다시 열 때의 최소 후보는 `review` 필드다. `mastery`, `daily_runs`는 `review` 필드 필요성이 확인된 뒤 후순위로 판단한다.

## 다음 결정 필요 사항

1. Foundation playtest에서 no-schema daily route가 실제 재방문 동기를 만드는지 확인한다.
2. progress schema v2가 필요해지는 시점에 `review` 최소 필드를 먼저 승인할지 결정한다.
3. daily run을 메인 메뉴의 기본 진입점으로 둘지, 별도 메뉴로 둘지 결정한다.
