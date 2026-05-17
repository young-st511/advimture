# Product — 영구 컨텍스트

> 드물게 변경한다. Advimture가 무엇인지, 무엇이 아닌지, 어떤 표면과 기둥을 가질지 기록한다.

## 현재 가정

- Go 기반 TUI adventure game은 유지한다.
- 기존 구현과 기획서는 참고 자료다.
- 새 제품 방향은 이 문서와 `PROGRAM.md`, 관련 spec에서 다시 정의한다.
- 주된 목적은 Vim 학습이다. 어드벤처와 시나리오는 학습 반복을 재미있게 만드는 장치다.

## 목표

Advimture는 플레이어가 Vim을 실제로 유용하게 쓰는 데 필요한 단축어와 명령어를, 반복 가능하고 재미있는 TUI adventure 문항을 통해 익히게 만드는 게임이다.

제품 철학:

> 유용하고 재미있는 Vim 훈련 게임을 먼저 만들고, 거기에 맞는 스토리를 붙인다.

설계 순서는 항상 다음을 따른다.

1. **Vim command**: Vim을 쓸 때 필요한 단축어와 명령어를 고른다.
2. **Exercise**: 그 명령어를 학습하기 위한 문항, 초기 상태, 목표 상태, 정답 키 입력, 채점 기준을 만든다.
3. **Scenario**: 검증된 문항을 어드벤처 사건, 캐릭터 대사, 미션 브리핑으로 감싼다.

스토리나 세계관에서 시작해 문항을 끼워 맞추지 않는다.

## 표면

- TUI 앱
- Vim-like keyboard interaction
- 로컬 progress 저장

## 기둥

- Vim command catalog: 실전에서 자주 쓰이고 조합 가능한 명령어 묶음
- Exercise bank: 각 command cluster를 반복 훈련시키는 명확한 문항과 정답
- Scenario layer: 문항의 의도를 방해하지 않는 어드벤처 포장
- 반복 플레이/성장 구조: 같은 명령을 더 적은 키, 더 정확한 입력, 더 자연스러운 조합으로 풀게 하는 진행
- 검증 가능한 TUI QA Loop: `make e2e-smoke`와 `test/e2e/` scenario를 시작점으로 실행 검증을 축적

## 워크스트림

- Command Catalog: Vim 단축어/명령어 선정, 난이도, 선행 관계 정리
- Exercise Design: 문항, 정답, 힌트, 채점 기준 설계
- Scenario Skinning: 검증된 문항에 어울리는 어드벤처 맥락 부여
- Gameplay Systems: 미션/튜토리얼/실패/보상 구조
- TUI Experience: 화면, 입력, 피드백, 접근성
- Verification: Go tests + TUI E2E QA Loop

## 참고 자료 승격 규칙

`docs/archived/PLAN.md`와 `docs/archived/GAME_DESIGN.md`에서 재사용할 아이디어는 그대로 구현하지 않는다. 먼저 이 문서, `PROGRAM.md`, domain spec, ExecPlan 중 적절한 위치로 옮기고 승인한다.
