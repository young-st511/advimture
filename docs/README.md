# docs/ 가이드

> 이 디렉토리는 Advimture 재기획, 실행 계획, 검증 루프를 관리한다. 기존 구현의 동작 역추출 문서는 아직 작성하지 않는다.

## 이 프로젝트의 spec 작성 모드

**모드 B — Agent 초안 + 사람 승인**을 사용한다.

- 사람이 의도를 표현하면 Agent가 `[draft]` 접두사로 수용 기준 초안을 작성한다.
- 사람이 검토하고 승인하면 `[draft]` 접두사를 제거한다.
- Agent는 `[draft]`가 없는 항목만 구현 대상으로 인식한다.
- Agent는 자기가 작성한 `[draft]`를 자기가 승인하지 않는다.

## 디렉토리 구조

```text
docs/
├── README.md
├── guardrails.md
├── gameplay/
│   ├── domain-contract.md
│   ├── spec.md
│   ├── command-catalog.md
│   ├── exercise-bank.md
│   ├── scenario-bank.md
│   └── content-requirements.md
├── verification/
│   ├── domain-contract.md
│   ├── spec.md
│   └── tui-e2e-loop.md
├── archived/
│   ├── PLAN.md
│   └── GAME_DESIGN.md
├── workflows/
│   └── vim-learning-loops.md
├── exec-plans/
│   ├── README.md
│   ├── templates/
│   ├── active/
│   └── completed/
└── roadmap/
    ├── PRODUCT.md
    ├── PROGRAM.md
    ├── CHANGES.md
    ├── decisions/
    └── archive/
```

## 문서별 안내

| 문서 | 대상 독자 | 목적 | 업데이트 시점 |
|------|----------|------|-------------|
| `roadmap/PRODUCT.md` | 사람 + AI | 제품의 영구 컨텍스트 | 제품의 본질, 표면, 기둥 변경 시 |
| `roadmap/PROGRAM.md` | 사람 + AI | 현재 phase와 활성 slice | 작업 우선순위나 slice 변경 시 |
| `roadmap/CHANGES.md` | 사람 + AI | 시퀀싱/가정 변경 로그 | phase 가정 변경 시 append-only |
| `exec-plans/*` | 사람 + AI | 비사소한 작업의 실행 계획 | 작업 시작/진행/완료 시 |
| `archived/*` | 사람 + AI | 과거 구현/기획 스냅샷 보관 | 과거 자료 이동 또는 보존 정책 변경 시 |
| `workflows/*` | 사람 + AI | 반복 실행 가능한 설계 loop | 제품 설계 프로세스 변경 시 |
| `{domain}/domain-contract.md` | AI | 도메인별 불변 규칙 | 규칙, 저장 포맷, 검증 방식 변경 시 |
| `{domain}/spec.md` | 사람 + AI | 승인된 수용 기준과 현재 동작 | 기능 구현 전/후 |
| `gameplay/command-catalog.md` | 사람 + AI | Vim command cluster 축적 | 새 command 학습 목표 제안/승인 시 |
| `gameplay/exercise-bank.md` | 사람 + AI | command 기반 문항 축적 | 새 exercise 제안/승인 시 |
| `gameplay/scenario-bank.md` | 사람 + AI | exercise 기반 시나리오 축적 | 새 scenario 제안/승인 시 |
| `gameplay/content-requirements.md` | 사람 + AI | loader가 받아야 할 콘텐츠 구조 요구사항 | scenario workflow로 새 콘텐츠 요구가 발견될 때 |
| `guardrails.md` | 사람 | 자동 검증과 안전장치 현황 | CI, hooks, 테스트 루프 변경 시 |

## spec.md 운영 규칙

현재 단계에서는 기존 구현을 canonical spec으로 역추출하지 않는다. 새로 기획한 기능만 `[draft]` 수용 기준으로 들어오고, 승인 후 구현된다.

### 수용 기준

- 각 항목은 pass/fail 판단이 가능해야 한다.
- "적절히", "자연스럽게", "좋게" 같은 표현은 구체적 관찰 가능 조건으로 바꾼다.
- Agent가 초안을 작성할 수 있지만, 사람 승인 전 구현하지 않는다.

### 현재 동작

- 구현이 완료된 항목만 "현재 동작"으로 이동한다.
- 이동할 때 관련 소스 파일과 검증 명령을 함께 남긴다.
- 기존 구현에서 가져온 아이디어는 "참고"일 뿐이며, 승인된 수용 기준으로 승격되기 전까지 현재 동작으로 기록하지 않는다.

## 하네스 진화

이 하네스의 각 규칙은 "현재 모델이 혼자서는 할 수 없다"는 가정을 인코딩한다. 모델이 업그레이드되면 별도 harness audit으로 각 가정이 여전히 유효한지 검증한다.

- 완전 소멸: 모델 역량이 초월한 규칙은 제거한다.
- 조건부: 단순 작업에는 불필요하지만 복잡한 작업에 필요한 규칙은 조건부로 바꾼다.
- 영구: 자기 평가 편향, 작성자/승인자 분리처럼 모델 역량과 무관한 규칙은 유지한다.
