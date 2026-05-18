# Gameplay Domain Contract

> 게임플레이 도메인은 플레이어가 보는 TUI 경험, 미션/튜토리얼 흐름, 저장 진행도를 다룬다.

## 공통 불변 규칙

| 규칙 | 근거 | 검증 방법 |
|------|------|----------|
| 기존 구현은 참고 자료이며 canonical spec이 아니다. | 프로젝트를 처음부터 재기획하기로 했으므로 과거 구현을 기준으로 새 기능을 정당화하면 안 된다. | 새 작업 시작 시 `docs/roadmap/PRODUCT.md`, `PROGRAM.md`, 관련 spec의 승인 기준을 먼저 확인한다. |
| 사용자 노출 문구는 한국어 톤을 우선한다. | 현재 게임 콘셉트와 기존 메시지가 한국어 사용자 경험을 전제로 한다. | UI 문자열 변경 diff를 수동 리뷰한다. |
| 저장 포맷 변경은 사용자 승인 없이 하지 않는다. | `~/.advimture/progress.json`은 사용자 로컬 데이터를 직접 건드린다. | `internal/progress/`, 저장 JSON 필드, migration 로직 변경 여부를 `git diff`로 확인한다. |
| 새 content schema의 ID와 파일명 변경은 기획 승인 없이 하지 않는다. | 진행도와 미션 참조가 ID에 의존할 수 있다. | `internal/content/`, 향후 content fixture/data 디렉터리 diff를 확인한다. |
| 새 의존성은 ExecPlan과 사용자 승인 후 추가한다. | TUI 게임은 단일 바이너리 배포 단순성이 중요하다. | `go.mod`, `go.sum` diff를 확인한다. |
| 설계 순서는 `Vim command → Exercise → Scenario`를 따른다. | 제품의 주 목적은 Vim 학습이며, 시나리오는 문항을 감싸는 장치다. | 신규 미션/튜토리얼/기획 문서에 command cluster와 exercise 정의가 scenario보다 먼저 있는지 확인한다. |

## Gameplay 고유 규칙

| 규칙 | 근거 | 검증 방법 |
|------|------|----------|
| 시나리오에서 출발해 문항을 끼워 맞추지 않는다. | 스토리 우선 설계는 학습 목표를 흐리게 만든다. | `docs/workflows/vim-learning-loops.md`의 3단계 루프 산출물을 순서대로 확인한다. |
| 모든 문항은 하나 이상의 Vim command cluster에 연결된다. | 플레이어가 무엇을 반복 훈련하는지 명확해야 한다. | exercise spec에 command cluster, 선행 조건, 목표 조작을 명시했는지 확인한다. |
| 새 게임 루프는 먼저 수용 기준으로 정의한다. | "재미"와 "학습 효과"는 구현 후 테스트만으로 판정하기 어렵다. | `docs/gameplay/spec.md`에 승인된 항목이 있는지 확인한다. |
| 미션 성공 조건은 사람이 읽을 수 있는 문장과 기계 검증 가능한 조건을 함께 가진다. | Agent와 플레이어가 같은 목표를 봐야 한다. | spec 또는 ExecPlan에 player-facing goal과 assertion을 둘 다 기록한다. |
| TUI 조작은 키 입력 trace로 재현 가능해야 한다. | E2E loop가 실패를 재현하려면 입력 시퀀스가 필요하다. | QA evidence에 key trace가 남는지 확인한다. |
