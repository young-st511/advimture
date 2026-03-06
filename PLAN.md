# Advimture — Phase 1~2 구현 Plan

## Context
Advimture는 실전 Vim 연습 CLI 게임으로, Go + Bubbletea로 개발한다.
이 Plan은 **Phase 1(코어 엔진 MVP)**과 **Phase 2(학습 시스템)**를 다룬다.
Phase 1에서 Vim 에뮬레이터의 핵심 동작을 구현하고,
Phase 2에서 메인 메뉴, 진행 저장, 튜토리얼 프레임워크, FTUE까지 완성한다.

기획서: `GAME_DESIGN.md`

---

## Step 1: 프로젝트 초기 세팅

- 목표: Go 프로젝트 초기화 + Bubbletea 스켈레톤 + 빈 에디터 화면 띄우기
- 변경 파일:
  - `go.mod`, `go.sum`
  - `main.go`
  - `internal/app/app.go`
  - `internal/editor/editor.go`
  - `internal/ui/styles.go`
  - `.gitignore`
- 상세 작업:
  - [ ] `go mod init github.com/young-st511/advimture`
  - [ ] Bubbletea, Lipgloss, Bubbles 의존성 추가
  - [ ] `.gitignore` 생성 (Go 바이너리, IDE 파일 등)
  - [ ] `main.go` — Bubbletea 프로그램 시작점
  - [ ] `internal/app/app.go` — 최상위 Model (화면 전환 관리 뼈대)
  - [ ] `internal/editor/editor.go` — 에디터 Model 뼈대 (빈 화면 + "Hello Advimture" 텍스트)
  - [ ] `internal/ui/styles.go` — Lipgloss 기본 스타일 정의 (모드별 색상, 테두리)
  - [ ] `go build && ./advimture`로 실행 확인
- 검증: 바이너리 실행 시 터미널에 Bubbletea 화면이 뜨고 `q`로 종료 가능

---

## Step 2: 텍스트 버퍼 + 커서

- 목표: 줄 단위 텍스트 버퍼와 커서 이동 시스템 구현 (로직만, 렌더링은 Step 7)
- 변경 파일:
  - `internal/editor/buffer.go`
  - `internal/editor/buffer_test.go`
  - `internal/editor/cursor.go`
  - `internal/editor/cursor_test.go`
- 상세 작업:
  - [ ] `Buffer` struct: `lines []string` 기반 텍스트 저장
  - [ ] 버퍼 조작 메서드: `InsertChar`, `DeleteChar`, `InsertLine`, `DeleteLine`, `GetLine`, `LineCount`
  - [ ] `Cursor` struct: `Row`, `Col`, `DesiredCol` (긴→짧→긴 줄 이동 시 col 복원용)
  - [ ] 커서 경계 처리: 줄 끝 넘어가기 방지, 빈 줄 처리, 파일 끝 처리
  - [ ] Normal 모드 커서 규칙: 마지막 문자까지만 이동 (줄 끝 '\n' 위 불가)
  - [ ] 단위 테스트: 버퍼 CRUD, 커서 경계 엣지 케이스 (빈 줄, 한 글자 줄, 마지막 줄)
- 검증: `go test ./internal/editor/...` 전체 통과

---

## Step 3: 모드 시스템 + 명령어 파서

- 목표: Normal/Insert/Command/Operator-Pending 모드 전환 + `{count}{operator}{motion}` 파싱
- 변경 파일:
  - `internal/editor/mode.go`
  - `internal/editor/parser.go`
  - `internal/editor/parser_test.go`
- 상세 작업:
  - [ ] `Mode` 타입 정의: `ModeNormal`, `ModeInsert`, `ModeCommand`, `ModeOperatorPending`, `ModeVisual`
  - [ ] 모드 전환 규칙 구현 (기획서 4.1.1 모드 시스템 다이어그램 기반)
  - [ ] `Parser` struct: count 버퍼, pending operator, 키 시퀀스 누적
  - [ ] 파싱 흐름 구현:
    - 숫자 → count 누적
    - `d`/`y`/`c` → Operator-Pending 진입
    - 같은 operator 반복 (`dd`, `yy`, `cc`) → 줄 단위 동작 플래그
    - motion (`w`, `b`, `$`, ...) 또는 text object (`iw`, `i"`, ...) → 범위 계산 후 실행
  - [ ] `ParseResult` struct: `Count`, `Operator`, `Motion`, `TextObject`, `IsLinewise`
  - [ ] 단위 테스트: `3dw`, `dd`, `ciw`, `5j`, `d$`, `2yy` 등 파싱 결과 검증
- 검증: `go test ./internal/editor/...` 전체 통과

---

## Step 4: 기본 Motion 구현

- 목표: `hjkl`, `w`, `b`, `e`, `0`, `$`, `gg`, `G` motion이 커서를 올바르게 이동
- 변경 파일:
  - `internal/editor/motion.go`
  - `internal/editor/motion_test.go`
- 상세 작업:
  - [ ] `Motion` 인터페이스 또는 함수 맵: motion 이름 → 커서 이동 범위 반환
  - [ ] `h`, `j`, `k`, `l` — 기본 이동 + 경계 처리
  - [ ] `j`/`k` — DesiredCol 유지 (긴→짧→긴 줄 복원)
  - [ ] `w`, `b`, `e` — 단어 경계 정의 (알파숫자+언더스코어 / 특수문자 / 공백)
  - [ ] `w`가 줄 끝에서 다음 줄로 넘어가는 동작
  - [ ] `0`, `$` — 줄 처음/끝 (Normal 모드에서 `$`는 마지막 문자 위)
  - [ ] `gg`, `G` — 파일 처음/끝
  - [ ] `{count}` 적용: `3w`, `5j` 등
  - [ ] 단위 테스트: 각 motion별 일반 케이스 + 엣지 케이스 (빈 줄, 파일 끝, 한 글자 줄)
- 검증: `go test ./internal/editor/...` 전체 통과

---

## Step 5: 기본 편집 + Undo/Redo + Register

- 목표: Insert 모드 진입/편집, 삭제, 복사/붙여넣기, Undo/Redo 동작
- 변경 파일:
  - `internal/editor/operator.go`
  - `internal/editor/operator_test.go`
  - `internal/editor/register.go`
  - `internal/editor/undo.go`
  - `internal/editor/undo_test.go`
- 상세 작업:
  - [ ] Insert 모드 진입: `i`, `a`, `o`, `I`, `A`, `O` — 각각 커서 위치/모드 전환 규칙
  - [ ] Insert 모드 편집: 문자 입력, Backspace, Enter(줄 분할)
  - [ ] `Esc`로 Normal 복귀 시 커서 한 칸 왼쪽 이동 (Vim 동작)
  - [ ] Operator 실행: `d` + motion → 범위 삭제, `y` + motion → 범위 복사, `c` + motion → 삭제 + Insert
  - [ ] `x` — 커서 위 문자 삭제
  - [ ] `dd`, `yy`, `cc` — 줄 단위 동작
  - [ ] `Register` struct: `Content string`, `Linewise bool`
  - [ ] `p`/`P` — Linewise 여부에 따른 붙여넣기 동작 차이
  - [ ] `Operation` struct + `UndoStack`/`RedoStack` (Operation Log 방식)
  - [ ] `u` → undo, `Ctrl+r` → redo
  - [ ] 새 편집 발생 시 RedoStack 초기화
  - [ ] 단위 테스트: 삽입/삭제/복사/붙여넣기/undo/redo 시나리오
- 검증: `go test ./internal/editor/...` 전체 통과

---

## Step 6: Command 모드

- 목표: `:` 입력으로 Command 모드 진입, 명령어 실행
- 변경 파일:
  - `internal/editor/command.go`
  - `internal/editor/command_test.go`
- 상세 작업:
  - [ ] `:` 입력 시 Command 모드 진입 — 하단 입력란 활성화
  - [ ] 명령어 파싱: `:w`, `:q`, `:q!`, `:wq`, `:{number}` (줄 이동)
  - [ ] `:s/old/new/g` — 현재 줄 치환
  - [ ] `:%s/old/new/g` — 전체 치환
  - [ ] `:{from},{to}d` — 범위 줄 삭제
  - [ ] `Enter`로 실행, `Esc`로 취소 (Normal 복귀)
  - [ ] 알 수 없는 명령어 → 에러 메시지 표시
  - [ ] `:wq`, `:q` 실행 결과를 에디터 상위로 전달 (저장/종료 이벤트)
  - [ ] 단위 테스트: 각 명령어 파싱 + 실행 결과 검증
- 검증: `go test ./internal/editor/...` 전체 통과

---

## Step 7: 에디터 View 렌더링

- 목표: 텍스트 버퍼를 터미널에 렌더링 — 줄 번호, 커서, 상태바, 모드 표시
- 변경 파일:
  - `internal/editor/editor.go` (Update/View 완성)
  - `internal/editor/view.go`
  - `internal/ui/styles.go` (스타일 확장)
- 상세 작업:
  - [ ] `editor.go`의 `Update()` 완성: 키 입력 → 파서 → 실행 → 상태 반영
  - [ ] `view.go` — View 렌더링:
    - 줄 번호 (회색)
    - 텍스트 영역 + 커서 위치 하이라이트 (블록 커서/바 커서)
    - 빈 줄 `~` 표시 (어두운 회색)
    - 하단 상태바: 모드명, 커서 위치 (Ln, Col), 키스트로크 수
    - Command 모드 시 하단에 `:` 입력란
  - [ ] 모드별 상태바 색상 (Normal=파랑, Insert=초록, Command=노랑, OpPending=주황)
  - [ ] 터미널 크기 대응: `tea.WindowSizeMsg` 처리
  - [ ] `Ctrl+c` × 2 → 종료 (게임 메뉴 복귀용, 현재는 프로그램 종료)
  - [ ] 통합 테스트: 실행하여 텍스트 편집, 모드 전환, 상태바 업데이트 수동 확인
- 검증: `go build && ./advimture`로 실행, 텍스트 편집/이동/모드 전환이 정상 동작

---

## Step 8: 메인 메뉴 + 화면 전환

- 목표: 메인 메뉴 화면과 에디터 사이 전환 구현
- 변경 파일:
  - `internal/app/app.go` (화면 전환 로직 완성)
  - `internal/ui/menu.go`
- 상세 작업:
  - [ ] `menu.go` — 메인 메뉴 Model/Update/View
    - Vim 스타일 조작: `j`/`k` 이동, `Enter` 선택
    - 메뉴 항목: Tutorial, Mission(잠김), Time Attack(잠김), Free Mode, Progress, Cheatsheet, Quit
    - 현재 직급 표시 (기본 Intern)
    - Lipgloss 스타일 박스 레이아웃
  - [ ] `app.go` — 화면 상태 관리:
    - `ScreenMenu`, `ScreenEditor`, `ScreenTutorial`, `ScreenResult` 등 enum
    - 메뉴에서 항목 선택 → 해당 화면으로 전환
    - Free Mode 선택 → 에디터 화면으로 전환 (샘플 텍스트)
  - [ ] Free Mode: 에디터 + 하단 명령어 해설 패널 (키 누를 때마다 한국어 설명 표시)
  - [ ] `Ctrl+c` × 2 → 에디터에서 메인 메뉴 복귀
- 검증: 메뉴 ↔ Free Mode 에디터 전환이 정상 동작

---

## Step 9: 진행 시스템 (저장/로드)

- 목표: `~/.advimture/progress.json`에 진행 데이터 저장/로드
- 변경 파일:
  - `internal/progress/progress.go`
  - `internal/progress/storage.go`
  - `internal/progress/storage_test.go`
  - `internal/progress/rank.go`
- 상세 작업:
  - [ ] `Progress` struct: Player 정보, Tutorial 완료 현황, Mission 기록, TimeAttack 기록
  - [ ] `Rank` 시스템: Intern → Junior → Senior → Staff → Principal → Vim Master
  - [ ] `storage.go`:
    - 경로: `~/.advimture/progress.json`
    - 로드: 파일 없음 → 신규 생성, JSON 파싱 실패 → `.bak` 복구 시도 → 초기화
    - 저장: 임시 파일 → rename (원자적 쓰기), 이전 파일 → `.bak`
  - [ ] 앱 시작 시 자동 로드, 주요 이벤트(튜토리얼 완료 등) 시 자동 저장
  - [ ] 단위 테스트: 저장/로드/손상 복구/신규 생성
- 검증: `go test ./internal/progress/...` 전체 통과 + 앱 재시작 시 데이터 유지 확인

---

## Step 10: 튜토리얼 프레임워크 + 힌트 시스템

- 목표: 5단계 학습 루프 엔진 + 힌트 시스템 + 목표 검증 시스템
- 변경 파일:
  - `internal/game/tutorial.go`
  - `internal/data/loader.go`
  - `internal/data/validator.go`
  - `internal/data/tutorials/` (YAML 스키마 구조)
- 상세 작업:
  - [ ] `Tutorial` Model: substep 순회, 현재 substep의 goal 검증, 진행 관리
  - [ ] 학습 루프: 컨텍스트 표시 → 격리 연습 → 조합 연습 → 미니 챌린지
  - [ ] `Goal` 검증 시스템 구현:
    - `cursor_position` — 커서 위치 확인
    - `cursor_on_char` — 특정 문자 위에 커서
    - `text_match` — 버퍼 텍스트 일치
    - `save_quit` — `:wq` + 텍스트 일치
    - `mode_is` — 현재 모드 확인
    - `command_used` — 특정 명령어 사용 여부
  - [ ] 힌트 시스템:
    - 10초 무입력 → Vi 선배 안내 메시지
    - `?` 1회/2회/3회 → 단계적 힌트 (방향 → 구체적 → 정답 시연)
  - [ ] `allowed_keys` — 스테이지별 허용 키 제한, 차단 시 메시지 표시
  - [ ] `loader.go` — Go embed로 YAML 로딩 + 스키마 검증
  - [ ] Vi 선배 메시지 영역 UI (시안 색상, 상단 패널)
  - [ ] substep 완료 → 다음 substep, 전체 완료 → completion 메시지 + progress 저장
- 검증: 더미 튜토리얼 YAML로 학습 루프 전체 사이클 동작 확인

---

## Step 11: FTUE + Tutorial 1-1 ~ 1-5

- 목표: 첫 실행 인트로 시퀀스 + 실제 튜토리얼 콘텐츠 5개 완성
- 변경 파일:
  - `internal/game/ftue.go`
  - `internal/data/tutorials/t01.yaml` ~ `t05.yaml`
  - `internal/app/app.go` (FTUE 분기 추가)
- 상세 작업:
  - [ ] `ftue.go` — FTUE Model:
    - SSH 접속 타이핑 애니메이션
    - Vim 화면 전환 (당혹감 체험)
    - Vi 선배 등장 + `:q!` 안내
    - 성공 후 Tutorial 1-1로 전환
  - [ ] `app.go` — 첫 실행 감지 (progress에 완료 기록 없음) → FTUE 시작
  - [ ] Tutorial 1-1: 생존 기초 (`i`, `Esc`, `:wq`, `:q!`)
    - substep 4개: `:q!` 연습 → `i`/`Esc` 사이클 → 오타 수정+`:wq` → 미니 챌린지
  - [ ] Tutorial 1-2: 커서 이동 (`h`, `j`, `k`, `l`)
    - substep 4개: 세로 → 가로 → 그리드 → 미로
    - 화살표 키 차단 + 안내 메시지
  - [ ] Tutorial 1-3: 빠른 이동 (`w`, `b`, `e`, `0`, `$`)
    - Aha Moment: hjkl로 20칸 vs `w`로 3번 → 속도 비교
  - [ ] Tutorial 1-4: 삽입의 기술 (`i`, `a`, `o`, `I`, `A`, `O`)
    - Aha Moment: `$a` vs `A` 키스트로크 비교
  - [ ] Tutorial 1-5: 삭제의 기술 (`x`, `dd`, `dw`, `d$`)
    - Aha Moment: `x` 10번 vs `dw` 1번 비교 체험
  - [ ] Tutorial 1-1 ~ 1-3 완료 시 메인 메뉴 해금 (progress 업데이트)
  - [ ] 명령어 언락 연출 (NEW COMMAND UNLOCKED!) 간단 버전
- 검증: 첫 실행 → FTUE → Tutorial 1-1~1-5 순차 플레이 가능, 재실행 시 메인 메뉴에서 시작

---

## 실행 규칙

각 Step은 아래 사이클을 순서대로 수행한다. 어떤 단계도 건너뛰지 않는다.

### 1. 구현
해당 Step의 작업을 수행한다.

### 2. 검증 (SubAgent 병렬 실행)
구현 완료 후, 아래 두 검증을 SubAgent(Agent 도구)로 **동시에 병렬 실행**한다.
두 SubAgent 모두 완료 후, 문제가 발견되면 즉시 수정하고 재검증한다.

**공통 - SubAgent 프롬프트에 포함할 정보:**
- 현재 Step의 Plan (목표, 변경 파일, 상세 작업)
- 기획서(`GAME_DESIGN.md`)의 관련 섹션
- Go 컨벤션: `go vet`, `go test`, 에러는 `fmt.Errorf`로 래핑, 패키지는 `internal/` 하위

**공통 - SubAgent가 직접 수행할 작업:**
- `git diff`로 변경사항 확인
- 필요시 관련 파일을 직접 읽기

#### SubAgent A: Plan 정합성 검증
- Plan의 "목표" 결과물이 모두 구현되었는가
- "상세 작업" 항목이 누락 없이 완료되었는가
- "변경 파일"과 실제 변경 파일이 일치하는가
- 범위 밖 변경이 포함되지 않았는가
- 기획서(`GAME_DESIGN.md`) 요구사항과 일치하는가

#### SubAgent B: 코드 리뷰
- Go 관례: 네이밍(camelCase export, 소문자 unexported), 에러 처리, 패키지 구조
- Bubbletea 패턴: Model/Update/View 분리, `tea.Cmd` 올바른 사용
- 타입 안전성, 중복 코드, 불필요한 export
- 테스트 커버리지: 핵심 로직에 테스트가 있는가

두 SubAgent의 결과를 종합하여 보고한다.

### 3. 커밋
1. AskUserQuestion으로 Step 결과 요약 및 리뷰 요청
2. 승인 후 커밋
   - 커밋 메시지 형식: `feat(scope): 설명` (conventional commits)
   - 예: `feat(editor): implement text buffer and cursor system`

1~3 완료 후 다음 Step으로 이동한다.

### Go 프로젝트 검증 체크리스트 (매 Step)
- [ ] `go build ./...` — 컴파일 성공
- [ ] `go vet ./...` — 정적 분석 통과
- [ ] `go test ./...` — 테스트 전체 통과 (테스트가 있는 Step)
- [ ] 실행 확인이 필요한 Step은 `go build && ./advimture`로 수동 테스트

---

# Phase 3 — 실전 미션 시스템

## Step 12: 검색 기능 (`/`, `n`, `N`) ✅
- `internal/editor/search.go` — SearchState, MatchPos, FindMatches
- `internal/editor/search_test.go` — 단위 테스트 (7개)
- `internal/editor/editor.go` — searchMode/searchInput/searchState 필드, handleSearchKey, n/N 네비게이션
- `internal/editor/view.go` — 검색 하이라이트 렌더링 (SearchHighlightStyle)
- `internal/editor/command.go` — CommandResult에 BufferModified 추가
- `internal/ui/styles.go` — SearchHighlightBg, SearchHighlightStyle 추가

## Step 13: 미션 데이터 시스템 ✅
- `internal/data/loader.go` — MissionData, LoadMission/LoadAllMissions, ValidateMission, CompareText, DiffLine
- `internal/data/missions/m00_test.yaml` — 검증용 더미 미션
- `internal/data/loader_test.go` — 미션 로드/검증/비교 테스트 추가

## Step 14: 키스트로크 추적 + 등급 산출 ✅
- `internal/editor/editor.go` — effectiveKeystrokes 필드, GetEffectiveKeystrokes/GetTotalKeystrokes, SavedAndQuit
- `internal/game/grader.go` — GradeResult, CalcGrade, CalcAccuracy, MentorMessage
- `internal/game/grader_test.go` — 등급 경계값, 정확도 테스트

## Step 15: 미션 게임 모드 + 결과 화면 ✅
- `internal/game/mission.go` — MissionModel (Preview→Challenge→Result), cursor_on_line 특수 처리
- `internal/ui/missionlist.go` — MissionListModel (미션 목록 화면)
- `internal/ui/menu.go` — MenuMission 추가, NewMenu(p) 시그니처 변경 (progress 기반 해금)
- `internal/progress/progress.go` — MissionProgress 재설계(BestGrade/BestKeystrokes/BestTimeMs/Attempts), CompleteMission, IsMissionUnlocked
- `internal/app/app.go` — ScreenMissionList, ScreenMission 추가, startMission

## Step 16: Mission M-01~M-05 + 해금 시스템 ✅
- `internal/data/missions/m01.yaml` — nginx.conf 도메인 변경 (★)
- `internal/data/missions/m02.yaml` — .env DB_HOST 수정 (★)
- `internal/data/missions/m03.yaml` — SSH config Host 블록 추가 (★★)
- `internal/data/missions/m04.yaml` — crontab 백업 Job 추가 (★★)
- `internal/data/missions/m05.yaml` — 로그에서 ERROR 찾기 cursor_on_line (★★)
- Tutorial 1-1~1-3 완료 시 Mission 메뉴 해금, 각 미션은 RequiredTutorials 기반 개별 해금
