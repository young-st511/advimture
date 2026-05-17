# ExecPlan 가이드

> ExecPlan은 비사소한 작업을 시작하기 전에 작성하는 실행 계획서다. 코드 변경의 의도, 범위, 검증 방법을 사전에 합의한다.

## 라이프사이클

1. **작성**: `templates/`에서 적합한 템플릿을 복사하여 `active/<plan-id>.md`로 시작한다.
2. **활성**: 작업 진행 중 plan을 갱신한다.
3. **완료**: PR 머지 후 `completed/<plan-id>.md`로 이동한다.

## 언제 ExecPlan을 만드는가

- 다중 파일 변경
- 게임 루프, 화면 전환, 저장 포맷, 미션 데이터 스키마 변경
- Bubble Tea 프로그램 구조 변경
- TUI E2E 러너 또는 QA 자동화 변경
- 새 의존성 추가
- 공개 contract나 roadmap에 영향을 주는 변경

## ExecPlan에 포함되는 것

- 목표
- 변경 범위와 제외 범위
- 검증 계획
- 의사결정 로그
- 위험과 미해결 질문

## 템플릿 선택

| 템플릿 | 사용 시점 |
|--------|----------|
| `general.md` | 일반 기능, 문서, 구현 slice |
| `strict-scope.md` | 저장 포맷, 의존성, 공유 contract처럼 범위 제한이 필요한 작업 |
| `vim-learning-slice.md` | Vim command → exercise → scenario 순서로 새 학습 문항을 설계/구현하는 작업 |

## strict-scope 모드

블래스트 반경이 큰 작업은 `Scope-Mode: strict`와 `Allowed-Paths`를 명시한다. Agent는 허용 경로 밖을 수정하지 않는다.
