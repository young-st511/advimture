# TUI UX Direction

> UI/UX 본격 개발을 위한 통합 결론이다. 세계관은 화면의 분위기를 만들지만, Vim 학습 목표보다 앞서지 않는다.

## 현재 판단

현재 TUI는 Vim 훈련 도구로는 기능한다. mode, cursor, input limit, required key coaching, success grade, key count, progress 저장이 모두 노출되어 있어 Agent와 플레이어가 같은 상태를 볼 수 있다.

하지만 화면 정보 계층은 아직 제품 UI라기보다 테스트 가능한 개발 UI에 가깝다.

- 현재 조작 목표보다 `재점검 대상`, `오늘의 복구 루트`가 먼저 보인다.
- tutorial과 incident가 같은 action panel 언어를 사용한다.
- `A D V I M T U R E - Playable Slice`, `Mode`, `Status`, `Cursor`, `ACTION`은 세계관보다 디버그 표면에 가깝다.
- buffer/cursor/selection은 색상 없이도 검증 가능하지만, `[]`, `{}` 기호 삽입이 긴 줄과 visual selection에서 과밀해진다.
- E2E screen artifact는 누적 ANSI log 성격이 강해 현재 frame 검증을 보조하기 어렵다.

## UI 원칙

- Vim 학습 목표가 1순위다.
- 초반 tutorial은 정답 key를 명시해도 된다. 첫 학습에서는 `이번에 익힐 키`, `왜 쓰는지`, `남은 입력`을 우선한다.
- incident는 정답 key보다 선택 판단을 우선한다. 기본 화면에서는 `상황`, `판단 목표`, `복구 단계`를 먼저 보여주고 key sequence는 hint/failure에서 점진 공개한다.
- review/daily 언어는 세계관에 맞지만 보조 정보다. 현재 exercise 목표보다 높은 계층에 두지 않는다.
- Runbook Dispatch 톤은 “상황 1문장 + Vim 조작 목표 1문장”을 넘기지 않는다.
- UI 검증은 화면 문자열뿐 아니라 app_state의 의미 필드를 우선한다.

## 권장 정보 구조

```text
Header
  ADVIMTURE / current playlist or incident / exercise index / status

Briefing
  situation sentence
  Vim operation or judgment goal

Console
  runbook buffer with cursor and selection

Status Line
  mode / status / cursor / inputs left / command or search line

Action or Debrief Panel
  tutorial: trained key, hint, retry, next
  incident: judgment cue, hint, retry, next
  success: recovery record, best record, residual risk, next
```

넓은 터미널에서는 `Console + Side Rail`을 고려할 수 있다. 좁은 터미널에서는 위 구조를 세로로 쌓는다.

## 개발 순서

1. **UX-QA-001**: review/daily route를 typed app_state assertion과 evidence로 검증한다.
2. **UI-CONTRACT-001**: header/briefing/console/status/action panel의 영역 contract를 고정한다.
3. **UI-RENDER-001**: 렌더링을 `internal/playableview` 같은 순수 renderer로 분리한다.
4. **UI-HIERARCHY-001**: 현재 exercise 목표를 상위에 두고 review/daily를 보조 정보로 낮춘다.
5. **UI-MODE-001**: tutorial과 incident action panel 언어를 분리한다.
6. **UI-EVIDENCE-001**: 현재 frame 또는 frame timeline evidence를 추가한다.

## 첫 구현 범위

첫 개발은 화면 리디자인이 아니라 QA 기반부터 시작한다. `UX-QA-001`은 기존 화면을 크게 바꾸지 않고, review/daily route를 문자열 `contains` 대신 typed `app_state.review` assertion으로 검증하며, E2E evidence에 app_state/progress snapshot을 남긴다.

