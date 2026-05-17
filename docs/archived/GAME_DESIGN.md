# Advimture — 실전 Vim 연습 CLI 게임 기획서

## 1. 개요

### 1.1 컨셉

SSH 접속, 서버 설정 파일 편집 등 **실전 상황을 시뮬레이션**하여 Vim을 자연스럽게 익히는 CLI 기반 연습 게임.
단순 명령어 암기가 아니라, 실제 시나리오 속에서 미션을 해결하며 Vim 근육 기억을 만드는 것이 목표.

### 1.2 타겟 사용자

- Vim을 처음 접하거나 기초만 아는 개발자
- SSH 환경에서 설정 파일을 편집해야 하는 상황이 잦은 사람
- GUI 에디터 없이 터미널만으로 작업해야 하는 경우가 있는 사람

### 1.3 기술 스택

- **언어**: Go
- **TUI 프레임워크**: Bubbletea + Lipgloss + Bubbles
- **빌드**: 단일 바이너리 배포

---

## 2. 내러티브

### 2.1 세계관

플레이어는 **신입 DevOps 엔지니어**로, 첫 출근 날부터 서버 장애에 투입된다.
회사의 서버에는 GUI가 없고 Vim만 사용할 수 있다.
선배 엔지니어 **"Vi 선배"**가 터미널 메시지로 가이드를 제공한다.

### 2.2 스토리 흐름

```
Day 1: 첫 출근 — Vim 기초 (Tutorial 1-1 ~ 1-5)
  "일단 서버에 접속해봐. vim으로 파일 열어서 수정하고 저장하는 것부터."

Day 2: 첫 장애 대응 — 이동과 편집 심화 (Tutorial 1-6 ~ 1-10)
  "어젯밤에 서버 설정이 꼬였어. 빠르게 찾아서 고쳐야 해."

Day 3~: 실전 투입 — Mission 시작
  "이제 혼자 해볼 때가 됐어. 긴급 티켓이 들어왔다."
  각 미션은 "긴급 티켓"으로 도착하며, 시나리오에 긴박감 부여.

승급 시스템:
  Intern → Junior → Senior → Staff → Principal
  미션 클리어 수 + 평균 등급에 따라 직급 상승
```

### 2.3 Vi 선배의 역할

- 튜토리얼에서 단계별 안내 메시지 제공
- 미션 시작 시 상황 브리핑
- 실패 시 격려 + 힌트 ("그 키 대신 이런 방법도 있어")
- 성공 시 칭찬 + 프로 팁 ("잘했어. 참고로 `.`을 쓰면 반복할 수 있어")

---

## 3. 게임 구조

### 3.1 첫 실행 경험 (FTUE: First Time User Experience)

게임을 처음 실행하면 메인 메뉴가 아니라 **인트로 시퀀스**가 시작된다.
목표: **첫 30초 안에 "이 게임은 재밌겠다"를 느끼게 하는 것.**

```
$ advimture

  [0초] 빈 터미널, 커서 깜빡임
  [1초] 타이핑 애니메이션 시작
  > ssh production-server-01
  [2초] Connecting... ██████████ 100%
  [3초] Connected.
  [4초]
  > vim /etc/nginx/nginx.conf
  [5초] 화면이 갑자기 Vim 인터페이스로 전환 (깜짝 효과)

  [에디터 화면에 nginx 설정이 펼쳐짐]
  [7초] 플레이어가 아무 키나 누르면...]

  Case 1: 'q' 입력 → 화면에 아무 변화 없음
  Case 2: ':q' 입력 → "E37: 변경 사항이 저장되지 않았습니다" 에러
  Case 3: 아무 키나 막 입력 → 이상한 동작 발생

  [10초] Vi 선배 등장 (화면 하단에 메시지):
  "당황하지 마. 다 그래. 처음엔 나가는 것도 어렵지.
   내가 도와줄게. Esc를 한 번 누르고, :q! 를 입력해봐."

  [플레이어가 :q! 입력 성공]
  "좋아. 이제 나가는 법은 알았네. 나머지도 알려줄게.
   준비됐어?"

  [Enter] 시작하기 → Tutorial 1-1로 진입
```

**설계 의도:**

- 처음 5초: 실제 SSH 접속하는 듯한 몰입감 (타이핑 애니메이션)
- 5~10초: 진짜 Vim처럼 조작이 안 되는 당혹감 체험 → "나도 이랬지"라는 공감
- 10초~: Vi 선배 등장으로 안도감 + 첫 성공 경험 (:q!)
- 이 30초가 "Vim은 어렵지만 이 게임이 도와줄 수 있다"는 메시지 전달

**인터랙션 피드백 원칙:**

- 모든 키 입력에 즉각적 시각 반응 (0.016초 이내, 60fps 기준)
- 유효한 명령 실행 시: 상태바에 명령어 이름 잠깐 표시 ("dw: 단어 삭제")
- 무효한 키 입력 시: 상태바에 "알 수 없는 명령" 표시 (비프음 없음, 시각적으로만)
- 모드 전환 시: 상태바 색상이 부드럽게 변경 (instant이 아닌 100ms 트랜지션)
- 튜토리얼 목표 달성 시: 체크마크 애니메이션 + Vi 선배 한 줄 칭찬

- 첫 실행 시 자동으로 Tutorial 1-1로 진입
- Tutorial 1-1 ~ 1-3 완료 후 메인 메뉴 해금
- 이후 실행부터는 메인 메뉴에서 시작

### 3.2 메인 메뉴

```
╔══════════════════════════════════════════╗
║         ADVIMTURE v1.0              ║
║         Rank: Junior DevOps             ║
║                                         ║
║  [t] Tutorial    - Vim 기초   8/10      ║
║  [m] Mission     - 실전 미션  3/10      ║
║  [a] Time Attack - 속도 도전            ║
║  [f] Free Mode   - 자유 연습            ║
║  [p] Progress    - 진행 현황            ║
║  [c] Cheatsheet  - 명령어 모음          ║
║                                         ║
║  [q] Quit                               ║
╚══════════════════════════════════════════╝
```

Vim 키 스타일로 메뉴 조작: `j`/`k`로 이동, `Enter`로 선택 (숫자 키도 가능).

### 3.3 게임 모드 상세

#### Mode 1: Tutorial (튜토리얼)

단계별로 Vim의 핵심 기능을 하나씩 배우는 가이드 모드.

**학습 루프 패턴 (모든 스테이지 공통):**

```
Step 1: 컨텍스트 — "왜 이걸 배우는가" (Vi 선배의 상황 설명)
Step 2: 데모    — 화면에서 명령어 동작을 시각적으로 보여줌
Step 3: 격리 연습 — 새 명령어 하나만 반복 사용하는 단순 과제 (3~5회)
Step 4: 조합 연습 — 이전에 배운 명령어와 섞어서 사용하는 과제
Step 5: 미니 챌린지 — 실전풍 텍스트에서 제한 키스트로크 내로 목표 달성
```

**실패/막힘 처리:**

- 10초간 입력 없음 → Vi 선배: "혹시 막혔어? `?`를 누르면 힌트를 볼 수 있어"
- `?` 1회 → 방향 힌트 ("이동 계열 명령어를 써봐")
- `?` 2회 → 구체적 힌트 ("w를 누르면 다음 단어로 이동해")
- `?` 3회 → 정답 시연 (키 입력을 화면에서 재생)
- 힌트를 사용하면 별도 표시 (등급에는 영향 없음, 학습이 목적이므로)

**스테이지 총괄:**

| 스테이지 | 주제          | 학습 내용                       | Aha Moment 설계                                               |
| -------- | ------------- | ------------------------------- | ------------------------------------------------------------- |
| 1-1      | 생존 기초     | `i`, `Esc`, `:wq`, `:q!`        | "Vim에서 나가는 법을 아는 것만으로 상위 10%입니다"            |
| 1-2      | 커서 이동     | `h`, `j`, `k`, `l`              | 화살표 키를 막고 hjkl로만 미로 탈출 → 속도 체감               |
| 1-3      | 빠른 이동     | `w`, `b`, `e`, `0`, `$`         | hjkl로 20칸 이동 vs `w`로 3번 이동 → 속도 비교 표시           |
| 1-4      | 삽입의 기술   | `i`, `a`, `o`, `I`, `A`, `O`    | 줄 끝에 텍스트 추가: `$` → `a` vs `A` 한 번 → 키스트로크 비교 |
| 1-5      | 삭제의 기술   | `x`, `dd`, `dw`, `d$`           | `x`를 10번 vs `dw` 1번으로 동일 결과 → "이게 Vim의 힘"        |
| 1-6      | 실수는 괜찮아 | `u`, `Ctrl+r`                   | 의도적으로 잘못 편집시킨 후 `u`로 복구하는 체험               |
| 1-7      | 복사의 기술   | `yy`, `p`, `P`, `yw`            | 같은 줄을 5번 타이핑 vs `yy` + `4p` → 효율 체감               |
| 1-8      | 검색          | `/pattern`, `n`, `N`            | 500줄 로그에서 `j`로 찾기 vs `/ERROR`로 즉시 찾기             |
| 1-9      | 치환          | `:s`, `:%s`                     | 10군데의 오타를 하나씩 vs `:%s`로 한번에 → 압도적 차이        |
| 1-10     | Vim의 문법    | `ciw`, `ci"`, `di(`, `ct.`, `.` | "Vim은 동사+명사 언어다" — 조합의 원리를 깨닫는 순간          |

**서브스텝 상세 (대표 스테이지 3개)**

**Tutorial 1-1: 생존 기초**

```
컨텍스트:
  Vi 선배: "서버에 접속했어. 앞에 보이는 파일을 수정해야 해.
           근데 먼저... vim에서 나가는 법부터 알려줄게."

서브스텝:
  1) ':q!'로 나가기
     - 화면에 파일이 열려 있음. "일단 나가보자. :q! 를 입력해봐."
     - ':q!' 입력 성공 → "좋아. 이제 어떤 상황에서든 나갈 수 있어."
     - 반복 3회 (매번 다른 파일 내용)

  2) 'i'로 Insert 모드 진입 + 'Esc'로 복귀
     - "이번엔 글자를 입력해볼게. i를 눌러봐."
     - 하단 상태바가 -- INSERT -- 로 변경되는 것을 확인
     - "좋아. 다시 Normal 모드로 돌아가려면 Esc를 눌러."
     - i → (타이핑) → Esc 사이클 3회 반복

  3) 텍스트 수정 후 ':wq'로 저장
     - 파일 내용: "Hello Wrold" (오타)
     - 목표: "Hello World"로 수정 후 저장
     - i로 진입 → 오타 수정 → Esc → :wq
     - 이 시점에서는 커서 이동은 화살표 키 허용 (아직 hjkl을 안 배움)

  4) 미니 챌린지: 실전 연습
     - Vi 선배: "설정 파일에 오타가 있어. 고치고 저장해봐."
     - 간단한 설정 파일에서 한 글자 수정 후 :wq

  완료 시: "축하해. 이제 넌 'Vim에서 나갈 수 있는 사람'이야. 이것만으로 대단한 거야."
```

**Tutorial 1-2: 커서 이동**

````
컨텍스트:
  Vi 선배: "이제부터 화살표 키는 잊어. 진짜 Vim 유저는 hjkl을 써."

서브스텝:
  1) 'j'/'k' 세로 이동
     - 5줄짜리 텍스트. "j를 눌러서 아래로 이동해봐."
     - 목표 줄에 도달하면 성공.
     - 화살표 키 입력 시: Vi 선배 "화살표 키 대신 j를 써봐. 손이 홈 포지션에서 안 떠나."
     - j/k만으로 3회 반복

  2) 'h'/'l' 가로 이동
     - 한 줄 텍스트에서 특정 문자 위로 커서 이동
     - h/l만으로 3회 반복

  3) 조합 연습: hjkl 네 방향
     - 그리드 형태 텍스트에서 지정된 위치로 이동
     ```
     . . . . .
     . . . . .
     . . X . .     ← X 위치로 커서를 이동하세요
     . . . . .
     . . . . .
     ```
     - 화살표 키는 차단됨, hjkl만 사용 가능
     - 3개 위치 연속 이동

  4) 미니 챌린지: 미로
     - '#'은 벽, '.'은 길, 'G'는 골
     ```
     . . # # #
     . # . . #
     . . . # .
     # # . . .
     # # # . G
     ```
     - hjkl로 S(시작)에서 G(골)까지 이동
     - '#' 위로는 이동 불가 (이동 시도 시 경고음)
     - 최소 키스트로크로 도달하면 보너스 메시지

  완료 시:
     Vi 선배: "어때? 손이 키보드 중앙에서 안 움직이지?
              이게 hjkl의 장점이야. 속도가 붙으면 화살표보다 빨라."
     → 키스트로크 수 표시 + "화살표로 했다면 N키 더 필요했을 거야" (비교)
````

**Tutorial 1-5: 삭제의 기술**

```
컨텍스트:
  Vi 선배: "삭제도 하나씩 할 필요 없어. Vim에선 '범위'를 지정해서 삭제할 수 있어."

서브스텝:
  1) 'x'로 문자 삭제
     - 텍스트: "Helllo World" (l이 하나 더 있음)
     - "커서를 l 위에 놓고 x를 눌러봐."
     - x로 단일 문자 삭제 3회 반복

  2) 'dd'로 줄 삭제
     - 5줄짜리 텍스트 중 3번째 줄이 불필요
     - "그 줄 전체를 지우고 싶으면 dd를 눌러."
     - dd로 줄 삭제 3회 반복

  3) 'dw'로 단어 삭제
     - 텍스트: "delete this unnecessary word here"
     - "unnecessary를 지워볼게. 단어 위에서 dw를 눌러."
     - dw, d$, d0 각각 1회씩 연습

  4) Aha Moment: 비교 체험
     - 동일한 편집 작업을 두 번 수행:
       Round 1: x만 사용 가능 (10글자 삭제 → x를 10번)
       Round 2: dw 사용 가능 (10글자 삭제 → dw 1번)
     - 완료 후 키스트로크 비교 화면:
       "Round 1: 10 keystrokes"
       "Round 2: 2 keystrokes  (dw + motion)"
       Vi 선배: "이게 Vim의 힘이야. 덩어리로 생각해."

  5) 미니 챌린지
     - 설정 파일에서 주석 3줄 삭제 + 불필요한 단어 2개 삭제
     - 제한: 15 키스트로크 이내
     - x, dd, dw, d$ 자유롭게 조합

  완료 시: Vi 선배: "삭제는 문자 단위가 아니라 '의미 단위'로 하는 거야. 기억해."
```

**Tutorial 1-10: Vim의 문법 (핵심 Aha Moment 스테이지)**

```
컨텍스트:
  Vi 선배: "Vim의 진짜 힘은 '문법'에 있어.
           영어처럼 동사 + 명사로 명령을 조합할 수 있지."

서브스텝:
  1) 문법 소개 — 동사 + 명사
     - 화면에 문법 구조 표시:
       ┌──────────────────────────────────┐
       │  Vim의 문법:  [동사] + [명사]     │
       │                                   │
       │  동사: d(삭제) c(변경) y(복사)     │
       │  명사: w(단어) $(줄끝) iw(안쪽단어) │
       │                                   │
       │  d + w = 단어 삭제                 │
       │  c + iw = 단어 변경               │
       │  y + $ = 줄끝까지 복사            │
       └──────────────────────────────────┘
     - Vi 선배: "이 조합을 알면 암기량이 확 줄어. 3개 동사 × 10개 명사 = 30개 명령."

  2) ciw (Change Inner Word) 격리 연습
     - 텍스트: name: "old_value"
     - 목표: name: "new_value" — old_value 위에서 ciw → new_value 입력
     - 3회 반복 (다른 변수명/값)

  3) ci" (Change Inner Quote) 격리 연습
     - 텍스트: title: "Hello World"
     - 목표: title: "Goodbye World" — 따옴표 안 어디서든 ci" → 전체 교체
     - Vi 선배: "커서가 따옴표 안 어디에 있든 상관없어. ci\"는 안쪽 전체를 바꿔."
     - 3회 반복

  4) di( (Delete Inner Paren) + ct. (Change To dot)
     - 텍스트: call(unnecessary_arg)
     - 목표: call() — di(로 괄호 안 삭제
     - 텍스트: first.second.third
     - 목표: replaced.second.third — ct.로 점 직전까지 변경
     - 각 2회

  5) . (Dot Repeat) — Aha Moment
     - 텍스트: 5줄에 걸쳐 "foo"라는 단어가 있음
     - Round 1: ciw로 하나씩 변경 (ciw → bar → Esc 를 5번)
     - Round 2: 첫 번째만 ciw → bar → Esc, 나머지는 n으로 이동 후 . 반복
     - 키스트로크 비교:
       "Round 1: 25 keystrokes"
       "Round 2: 11 keystrokes"
     - Vi 선배: "이게 dot repeat의 힘이야. 한 번 한 작업을 .으로 반복해."

  6) 미니 챌린지: 조합 실전
     - JSON 설정 파일에서:
       - 키 이름 변경 (ciw)
       - 값 변경 (ci")
       - 불필요한 필드 삭제 (dd)
       - 같은 변경을 다른 위치에 반복 (.)
     - 제한: 20 키스트로크

  완료 시:
     Vi 선배: "동사 + 명사. 이 문법만 기억하면 새로운 명사를 배울 때마다
              이미 아는 동사와 조합할 수 있어. 이게 Vim의 진짜 힘이야."
     → 특별 연출: "VIM GRAMMAR MASTERED!" + 전체 조합표 표시
```

**선행 조건 시스템:**

- Tutorial 1-1 ~ 1-5 완료 → Mission M-01 ~ M-03 해금
- Tutorial 1-6 ~ 1-8 완료 → Mission M-04 ~ M-07 해금
- Tutorial 1-9 ~ 1-10 완료 → Mission M-08 ~ M-10 해금
- Mission 5개 클리어 → Time Attack 해금

#### Mode 2: Mission (실전 미션)

실제 서버 작업 시나리오를 재현한 미션 모드.
각 미션은 "긴급 티켓" 형태로 제시된다.

```
╔═══════════════════════════════════════════════╗
║  TICKET #001                    Priority: LOW ║
║──────────────────────────────────────────────║
║  Title: nginx 도메인 변경                      ║
║                                               ║
║  "example.com 도메인을 mysite.com으로          ║
║   변경해야 합니다. nginx.conf를 수정해주세요."  ║
║                                               ║
║  Required Skills: 검색, 치환                   ║
║  Optimal Keystrokes: 8                        ║
║                                               ║
║  [Enter] 시작  [b] 돌아가기                    ║
╚═══════════════════════════════════════════════╝
```

**미션 목록:**

| 미션 | 시나리오                           | 난이도 | 필요 스킬          | 최적해                                 |
| ---- | ---------------------------------- | ------ | ------------------ | -------------------------------------- |
| M-01 | nginx.conf: server_name 변경       | ★      | 검색, 치환         | `:%s/example.com/mysite.com/g` + `:wq` |
| M-02 | .env: DB_HOST 값 수정              | ★      | 검색, `ciw`        | `/DB_HOST` → `wciwvalue`               |
| M-03 | SSH config: 호스트 블록 추가       | ★★     | `G`, `o`, 삽입     | `Go` + 텍스트 입력                     |
| M-04 | crontab: cron job 추가             | ★★     | `G`, `o`, 삽입     | `Go` + 텍스트 입력                     |
| M-05 | 로그 파일: ERROR 라인 찾기         | ★★     | `/`, `n`, `N`      | `/ERROR` + `n` 반복                    |
| M-06 | Docker Compose: 포트/환경변수 수정 | ★★     | 검색, `ci"`, `o`   | 복합 편집                              |
| M-07 | /etc/hosts: 도메인 매핑 추가       | ★      | `G`, `o`           | `Go` + 입력                            |
| M-08 | k8s YAML: replicas 수 변경         | ★★     | `/`, `ciw`         | `/replicas` → `wciw3`                  |
| M-09 | 설정 파일: 주석 블록 제거          | ★★★    | `V`, `d`, 줄 번호  | `:{from},{to}d`                        |
| M-10 | 다중 IP 치환                       | ★★★    | `:%s`, 정규식 기초 | `:%s/10\.0\.0\.1/10.0.0.50/g`          |

**등급 산출 기준:**

```
S등급: 키스트로크 <= 최적해 * 1.0  (최적해와 동일하거나 더 적음)
A등급: 키스트로크 <= 최적해 * 1.5
B등급: 키스트로크 <= 최적해 * 2.5
C등급: 완료하면 부여
```

**복수 최적해 처리:**
하나의 미션에 여러 정답 경로가 존재할 수 있다.
최적해는 "알려진 최단 경로들의 최소 키스트로크"를 기준으로 하되,
등급 산출 시에는 키스트로크 수만 비교하므로 경로 자체는 검증하지 않는다.
결과 텍스트가 목표와 일치하면 어떤 방법이든 인정.

#### Mode 3: Time Attack (속도 도전)

**Sprint:**

- 10개의 마이크로 편집 문제를 연속 출제
- 각 문제는 1~3줄의 짧은 텍스트 + 단일 편집 목표
- 문제 풀 (30개 이상)에서 랜덤 10개 출제
- 각 문제 제한 시간: 30초 (초과 시 자동 스킵)
- 총 소요 시간으로 순위 결정

```
문제 풀 예시:
  - "Hello Wrold" → "Hello World"로 수정       (난이도: 쉬움)
  - port: 8080 → port: 3000 으로 변경            (난이도: 쉬움)
  - 줄 "# 주석" 삭제                             (난이도: 쉬움)
  - name: "old" → name: "new" (ci" 활용)         (난이도: 보통)
  - 3줄의 중복 블록 삭제                           (난이도: 보통)
  - 5군데의 "foo"를 "bar"로 치환                   (난이도: 어려움)
```

**Endurance:**

- HP: 3칸 (하트 표시)
- 문제 난이도가 점진적으로 상승 (Easy 5개 → Medium 5개 → Hard → ...)
- 정답: 콤보 +1, 3콤보마다 HP +1 (최대 5)
- 오답 (`:wq` 시 expected와 불일치): HP -1, 콤보 초기화
- 시간 초과 (문제당 60초): HP -1
- HP 0 → 게임 오버, 총 클리어 수가 기록

```
  ♥ ♥ ♥          Combo: 5x          Score: 47
  ┌─────────────────────────────────────┐
  │  DB_HOST=localhost                  │
  │                                     │
  │  → DB_HOST=10.0.0.5 로 변경하세요  │
  └─────────────────────────────────────┘
```

**Daily Challenge:**

- 날짜(YYYY-MM-DD)를 시드로 사용하여 문제 풀에서 10개 고정 출제
- 같은 날은 항상 같은 문제 → 비교 가능
- 완료 시 결과를 **클립보드 복사 가능한 텍스트**로 제공:

```
ADVIMTURE - Daily Challenge 2026-03-06
Time: 32.5s | Keys: 47 | Accuracy: 94%
Rank: A | Combo: 8x
```

측정 지표: 총 소요 시간, 키스트로크 수, 정확도, 콤보 (연속 정답)

#### Mode 4: Free Mode (자유 연습)

빈 버퍼 또는 샘플 파일을 열어 자유롭게 Vim 조작을 연습.
화면 하단에 현재 모드, 입력한 키 시퀀스, 실행된 명령어를 실시간 표시.

```
┌─────────────────────────────────────┐
│ server {                            │
│     listen 80;                      │
│     server_name █example.com;       │
│ }                                   │
│                                     │
├─────────────────────────────────────┤
│ MODE: NORMAL | Keys: dw | Cmd: 단어 삭제 │
│ Cursor: Ln 3, Col 17                │
└─────────────────────────────────────┘
```

**명령어 해설 기능**: 키를 누를 때마다 하단에 해당 명령어의 동작을 한국어로 표시.
학습과 실험을 동시에 할 수 있는 샌드박스 역할.

---

## 4. 핵심 시스템 설계

### 4.1 Vim 에뮬레이터

게임 목적에 맞는 최소한의 Vim 동작만 구현한다. 완전한 Vim 클론이 목표가 아님.

#### 4.1.1 모드 시스템

```
Normal Mode ──→ Insert Mode    (i, a, o, I, A, O)
Normal Mode ──→ Command Mode   (:)
Normal Mode ──→ Visual Mode    (v, V)
Normal Mode ──→ Operator-Pending Mode  (d, y, c 입력 후)
Insert Mode ──→ Normal Mode    (Esc)
Command Mode ──→ Normal Mode   (Enter, Esc)
Visual Mode ──→ Normal Mode    (Esc, 명령 실행 후)
Operator-Pending Mode ──→ Normal Mode  (motion 입력 후, Esc)
```

**Operator-Pending Mode:**
`d`, `y`, `c` 등 operator를 입력하면 이 모드에 진입.
다음 입력(motion 또는 text object)을 기다린 후 조합하여 실행.

```
예: d → (Operator-Pending) → w → "dw" 실행 (단어 삭제)
예: c → (Operator-Pending) → i → w → "ciw" 실행 (단어 변경)
```

#### 4.1.2 명령어 파싱 구조

Vim의 명령어 문법: `{count}{operator}{count}{motion/text-object}`

```
파싱 흐름:
1. 숫자 입력 → count 버퍼에 누적
2. Operator 입력 (d, y, c) → Operator-Pending 진입
   - 같은 operator 반복 (dd, yy, cc) → 줄 단위 동작
3. Motion 입력 (w, b, $, 0, ...) → count * motion 범위 계산
4. 또는 Text Object 입력 (iw, i", a(, ...) → 오브젝트 범위 계산
5. operator + 범위 → 실행
```

**구현할 Operator:**
| Operator | 동작 | 예시 |
|----------|------|------|
| `d` | 삭제 | `dw`, `dd`, `d$`, `diw` |
| `y` | 복사 (yank) | `yw`, `yy`, `y$` |
| `c` | 변경 (삭제 후 Insert) | `cw`, `cc`, `ciw`, `ci"` |

**구현할 Motion:**
| Motion | 동작 |
|--------|------|
| `h`, `j`, `k`, `l` | 기본 이동 |
| `w`, `b`, `e` | 단어 단위 이동 |
| `0`, `$` | 줄 처음/끝 |
| `gg`, `G` | 파일 처음/끝 |
| `{number}G` | 특정 줄 이동 |
| `f{char}`, `t{char}` | 문자까지 이동 |
| `Ctrl+d`, `Ctrl+u` | 반 페이지 스크롤 |

**구현할 Text Object:**
| Text Object | 범위 |
|-------------|------|
| `iw` / `aw` | 단어 안/밖 |
| `i"` / `a"` | 따옴표 안/밖 |
| `i'` / `a'` | 작은따옴표 안/밖 |
| `i(` / `a(` | 괄호 안/밖 |
| `i{` / `a{` | 중괄호 안/밖 |

**추가 명령어:**
| 명령 | 동작 |
|------|------|
| `x` | 커서 위 문자 삭제 |
| `r{char}` | 커서 위 문자 교체 |
| `J` | 줄 합치기 |
| `.` | 마지막 변경 명령 반복 (dot repeat) |
| `u` | 되돌리기 |
| `Ctrl+r` | 다시하기 |
| `p`, `P` | 붙여넣기 (커서 뒤/앞) |
| `/pattern` | 검색 |
| `n`, `N` | 다음/이전 검색 결과 |
| `v` | Visual Mode (문자 단위) |
| `V` | Visual Line Mode (줄 단위) |

**Command Mode:**
| 명령 | 동작 |
|------|------|
| `:w` | 저장 |
| `:q`, `:q!` | 종료 |
| `:wq` | 저장 후 종료 |
| `:{number}` | 줄 이동 |
| `:s/old/new/g` | 현재 줄 치환 |
| `:%s/old/new/g` | 전체 치환 |
| `:{from},{to}d` | 범위 줄 삭제 |

#### 4.1.3 Unnamed Register

삭제(`d`, `x`), 복사(`y`) 명령 실행 시 대상 텍스트를 unnamed register에 저장.
`p`/`P`는 항상 unnamed register의 내용을 붙여넣기.

```
레지스터 상태:
- content: string     — 저장된 텍스트
- linewise: bool      — 줄 단위 여부 (dd, yy → true, dw, yw → false)

붙여넣기 동작:
- linewise=true: p → 현재 줄 아래에 새 줄로 삽입, P → 위에 삽입
- linewise=false: p → 커서 뒤에 삽입, P → 커서 앞에 삽입
```

#### 4.1.4 Dot Repeat (`.`)

마지막으로 실행한 "변경 명령"을 기록하고, `.` 입력 시 동일하게 재실행.

```
기록 대상: Insert/Change/Delete 계열 명령의 전체 시퀀스
예: ciw → "hello" → Esc → 이 전체가 하나의 변경으로 기록
   . 입력 시 → 커서 위치의 단어를 "hello"로 변경
```

#### 4.1.5 커서 규칙

- **Normal Mode**: 커서는 문자 위에 위치 (블록 커서). 빈 줄이면 col=0.
  줄의 마지막 문자까지만 이동 가능 (줄 끝 '\n' 위에 올 수 없음).
- **Insert Mode**: 커서는 문자 사이에 위치 (바 커서). 줄 끝 뒤에도 위치 가능.
- `$`로 줄 끝 이동 시 마지막 문자 위에 위치.
- `A`는 줄 끝 뒤로 커서 이동 + Insert 진입 (실질적으로 `$a`와 동일하지만 한 키).

#### 4.1.6 엣지 케이스 동작 정의

모든 엣지 케이스에서 Vim 실제 동작과 동일하게 처리한다.

**커서 경계:**
| 상황 | 입력 | 동작 |
|------|------|------|
| 첫 번째 줄에서 | `k` | 아무 일도 안 일어남 (경고음 없음) |
| 마지막 줄에서 | `j` | 아무 일도 안 일어남 |
| 줄 첫 번째 문자에서 | `h` | 아무 일도 안 일어남 |
| 줄 마지막 문자에서 | `l` | 아무 일도 안 일어남 (Normal), 줄 끝 뒤로 이동 가능 (Insert) |
| 빈 줄에서 | `w`, `b`, `e` | 다음/이전 줄의 첫 단어로 이동 |
| 빈 줄에서 | `dw` | 빈 줄 삭제 (다음 줄과 합침) |
| 빈 줄에서 | `x` | 아무 일도 안 일어남 |
| 한 글자 줄에서 | `x` | 문자 삭제 → 빈 줄이 됨 |
| 파일에 줄이 1개만 있을 때 | `dd` | 줄 내용 삭제 → 빈 줄 1개 남음 (줄 자체는 유지) |
| 파일 마지막 줄에서 | `dd` | 줄 삭제 → 커서가 위 줄로 이동 |
| 파일 끝에서 | `G` | 마지막 줄 첫 번째 비공백 문자로 이동 |
| 검색 결과 없을 때 | `n` | "Pattern not found" 메시지 (상태바) |
| Normal 모드에서 | 일반 문자 키 (q, z 등) | 미구현 명령 → 아무 일도 안 일어남 |

**긴 줄 ↔ 짧은 줄 이동:**

- `j`/`k`로 이동 시, 대상 줄이 현재 col보다 짧으면 줄 끝으로 이동
- "원래 col" 기억 (desired column): 다시 긴 줄로 이동하면 원래 위치로 복귀
- 예: col=20 → 짧은 줄(10자) → col=9 → 다시 긴 줄 → col=20 복귀

**`w` 단어 경계 정의:**

- 단어: 연속된 알파벳/숫자/언더스코어 OR 연속된 특수문자
- `w`는 다음 단어의 첫 글자로 이동
- 줄 끝에서 `w`는 다음 줄의 첫 단어로 이동
- 예: `server_name` → 하나의 단어, `server.name` → `server`, `.`, `name` 3개

### 4.2 텍스트 버퍼

- 줄 단위 배열(`[]string`)로 텍스트 관리
- 커서 위치: `(row, col)` 좌표

**Undo/Redo: Operation Log 방식**

```
Operation {
  Type:      "insert" | "delete" | "replace"
  StartPos:  (row, col)
  EndPos:    (row, col)
  Text:      string      // 삭제된 또는 삽입된 텍스트
  PrevText:  string      // replace 시 이전 텍스트
}

UndoStack: []Operation   // u 실행 시 pop → 역연산 실행 → RedoStack에 push
RedoStack: []Operation   // Ctrl+r 시 pop → 재실행 → UndoStack에 push
새 편집 발생 시 RedoStack 초기화
```

### 4.3 미션 검증 시스템

- 미션마다 **목표 상태**(expected text)를 정의
- 플레이어가 `:wq`로 저장하면 현재 버퍼와 목표 상태를 **줄 단위 비교**
- `:q!`로 나가면 미완료 처리 (진행에 영향 없음)

**비교 알고리즘:**

```
1. 양쪽 텍스트를 줄 단위로 분할
2. 줄 수가 다르면 → 불일치
3. 각 줄을 정확 비교 (trailing whitespace 무시)
4. 모든 줄 일치 → 성공
5. 하나라도 불일치 → 실패
```

**실패 시 diff 화면:**

```
┌─ DIFF ──────────────────────────────┐
│                                      │
│  Ln 3:                               │
│  ✗ yours:    server_name exampl.com; │  ← 빨간색
│  ✓ expected: server_name mysite.com; │  ← 초록색
│                                      │
│  Vi 선배: "거의 다 맞았어! 3번째 줄을  │
│           다시 확인해봐."              │
│                                      │
│  [r] 재도전  [b] 메뉴                 │
└──────────────────────────────────────┘
```

- 최대 3줄까지 diff 표시 (그 이상 틀리면 "N줄이 다릅니다" 요약)
- 틀린 부분이 어디인지 문자 수준까지 밑줄로 표시 (가능한 경우)

**M-05 같은 "검색만 하는" 미션 처리:**

- `expected_text` 대신 `goal` 필드 사용 가능
- `goal.type: "cursor_on_line"` — 커서가 특정 패턴을 포함하는 줄에 위치
- `goal.type: "yank_contains"` — 레지스터에 특정 텍스트가 복사됨
- 검증은 `:wq` 대신 별도 트리거 (예: 특정 줄에서 `Enter`)

### 4.4 키스트로크 추적

```
KeyLog {
  Key:       string    // 입력한 키
  Timestamp: int64     // 밀리초 타임스탬프
  Mode:      string    // 입력 시점의 모드
  Command:   string    // 해석된 명령어 (예: "dw", "insert 'hello'")
  Effective: bool      // 실제로 버퍼를 변경했는지
}
```

- 모든 키 입력을 기록
- `Effective=false`인 키 (잘못 누른 키, 취소된 명령)를 세어 정확도 산출
- 미션 완료 시 키스트로크 수 = Effective=true인 키만 카운트
- 최적해 대비 비교하여 효율성 피드백 제공
- "이렇게 하면 더 빠릅니다" 팁: 미션별로 2~3개의 대체 풀이를 미리 정의

### 4.5 튜토리얼 데이터 스키마 (YAML)

```yaml
id: '1-2'
title: '커서 이동'
day: 1 # 스토리 Day
story: |
  Vi 선배: "이제부터 화살표 키는 잊어. 진짜 Vim 유저는 hjkl을 써."

allowed_keys: # 이 스테이지에서 허용하는 키 (나머지는 차단)
  - 'h'
  - 'j'
  - 'k'
  - 'l'
  - 'Esc'

blocked_keys_message: '이 스테이지에서는 hjkl만 사용할 수 있습니다.'

substeps:
  - id: 'vertical'
    instruction: 'j를 눌러서 4번째 줄로 이동해보세요.'
    initial_text: |
      line 1
      line 2
      line 3
      line 4
      line 5
    cursor_start: { row: 0, col: 0 }
    goal:
      type: 'cursor_position' # 커서 위치 확인
      row: 3
      col: 0
    repeat: 3 # 다른 텍스트로 3회 반복
    on_success: '좋아! j는 아래, k는 위로 이동이야.'

  - id: 'horizontal'
    instruction: "l을 눌러서 'X' 문자 위로 이동하세요."
    initial_text: 'find the X here'
    cursor_start: { row: 0, col: 0 }
    goal:
      type: 'cursor_on_char' # 특정 문자 위에 커서
      char: 'X'
    repeat: 3

  - id: 'grid'
    instruction: 'X 위치로 커서를 이동하세요.'
    initial_text: |
      . . . . .
      . . . . .
      . . X . .
      . . . . .
    cursor_start: { row: 0, col: 0 }
    goal:
      type: 'cursor_position'
      row: 2
      col: 4
    repeat: 3 # 매번 X 위치가 달라짐

  - id: 'maze'
    instruction: "미로를 탈출하세요! '#'은 벽입니다."
    type: 'maze' # 특수 모드: 벽 충돌 감지
    initial_text: |
      . . # # #
      . # . . #
      . . . # .
      # # . . .
      # # # . G
    cursor_start: { row: 0, col: 0 }
    goal:
      type: 'cursor_on_char'
      char: 'G'
    walls: '#' # 이동 불가 문자
    track_keystrokes: true # 최소 이동 비교 표시

completion_message: |
  Vi 선배: "어때? 손이 키보드 중앙에서 안 움직이지?
           이게 hjkl의 장점이야."
unlock_commands: ['h', 'j', 'k', 'l']
```

**튜토리얼 Goal 타입:**
| type | 검증 방식 |
|------|----------|
| `cursor_position` | 커서가 지정된 (row, col)에 위치 |
| `cursor_on_char` | 커서가 지정된 문자 위에 위치 |
| `text_match` | 버퍼 텍스트가 expected와 일치 |
| `text_contains` | 버퍼에 특정 문자열이 포함 |
| `mode_is` | 현재 모드가 지정 모드와 일치 |
| `command_used` | 특정 명령어를 사용했는지 확인 |
| `save_quit` | `:wq`로 저장 종료 + 텍스트 일치 |

### 4.6 미션 데이터 스키마 (YAML)

```yaml
id: 'm01'
title: 'nginx 도메인 변경'
difficulty: 1 # 1~3 (★ 수)
story: |
  프로덕션 서버의 도메인이 변경되었습니다.
  nginx.conf의 server_name을 수정해주세요.
required_tutorials:
  - '1-1'
  - '1-9'

initial_text: |
  server {
      listen 80;
      server_name example.com;

      location / {
          proxy_pass http://localhost:3000;
      }
  }

expected_text: |
  server {
      listen 80;
      server_name mysite.com;

      location / {
          proxy_pass http://localhost:3000;
      }
  }

cursor_start:
  row: 0
  col: 0

optimal_keystrokes: 8
optimal_solutions:
  - description: '전체 치환'
    keys: ':%s/example.com/mysite.com/g<Enter>:wq<Enter>'
    count: 8
  - description: '검색 후 ciw'
    keys: '/example<Enter>ciwmysite<Esc>:wq<Enter>'
    count: 12

tips:
  - trigger: 'keystroke_over_20'
    message: ':%s/old/new/g 를 사용하면 파일 전체를 한번에 치환할 수 있습니다.'
  - trigger: 'used_hjkl_over_10'
    message: '/keyword 로 검색하면 원하는 위치로 바로 이동할 수 있습니다.'
```

---

## 5. UI/UX 설계

### 5.1 게임 ↔ Vim 에뮬레이터 조작 충돌 해결

**핵심 문제:** Vim에서 `Esc`는 모드 전환인데, 게임 메뉴로 나가려면?

**해결:**

- `Ctrl+c` × 2회 연속 (0.5초 이내) → 게임 메뉴로 복귀
- 화면 상단에 항상 표시: `Ctrl+c ×2: 메뉴`
- `Esc`는 오직 Vim 모드 전환에만 사용
- Command Mode에서 `:quit` → 미션 포기 (`:q!`와 동일)
- Tutorial 중에는 추가로 `Ctrl+h`로 힌트 표시 (= `?`와 동일)

### 5.2 화면 레이아웃

**최소 터미널 크기: 80 x 24** (미달 시 경고 메시지 표시)

```
┌─ ADVIMTURE ─────────────── Mission M-03 ─┐  ← 헤더 (1줄)
│                                               │
│  Vi 선배: "SSH config에 새 호스트를 추가해."   │  ← 미션 설명 (2줄)
│                                               │
├───────────────────────────────────────────────┤
│  Host production                              │  ← 에디터 영역
│      HostName 10.0.1.100                      │     (가변, 최소 15줄)
│      User deploy                              │
│      Port 22                                  │
│  █                                            │
│  ~                                            │
│  ~                                            │
│  ~                                            │
├───────────────────────────────────────────────┤
│  -- NORMAL --      Ln 5, Col 1    Keys: 12    │  ← 상태바 (1줄)
│  Ctrl+c ×2: 메뉴 | ?: 힌트                    │  ← 도움말 (1줄)
└───────────────────────────────────────────────┘
```

- 터미널 크기 변경 시 에디터 영역만 리사이즈
- 헤더, 미션 설명, 상태바, 도움말은 고정 높이

### 5.3 색상 테마 (Lipgloss)

```
모드별 상태바 색상:
  Normal Mode:           파랑 배경, 흰색 텍스트
  Insert Mode:           초록 배경, 흰색 텍스트
  Visual Mode:           보라 배경, 흰색 텍스트
  Command Mode:          노랑 배경, 검정 텍스트
  Operator-Pending Mode: 주황 배경, 흰색 텍스트

커서:
  Normal:  블록 커서 (반전 색상)
  Insert:  밑줄 커서 (또는 얇은 바)
  Visual:  선택 영역 반전

에디터 영역:
  줄 번호:       회색
  빈 줄 (~):     어두운 회색
  검색 하이라이트: 노란 배경
  diff 틀린 줄:   빨간 배경
  diff 정답 줄:   초록 배경

Vi 선배 메시지:  시안(cyan) 텍스트, 이탤릭
```

**색각 이상 대응:**

- 색상만으로 구분하지 않고 기호도 병행 (예: 틀린 줄에 `✗`, 정답 줄에 `✓`)
- 향후 확장으로 고대비 테마 옵션 고려

### 5.4 피드백 시스템

**미션 완료 시:**

```
╔═══════════════════════════════════════════╗
║          MISSION COMPLETE!                ║
║                                           ║
║  Vi 선배: "깔끔하게 처리했네. 잘했어."     ║
║                                           ║
║  키스트로크:  15  (최적: 8)                ║
║  소요 시간:   12.3초                       ║
║  정확도:      87%                          ║
║  등급:        B                            ║
║                                           ║
║  TIP: :%s 를 사용하면 한번에               ║
║       치환할 수 있습니다                    ║
║                                           ║
║  [Enter] 다음 미션  [r] 재도전  [b] 메뉴   ║
╚═══════════════════════════════════════════╝
```

**등급별 Vi 선배 반응:**

- S: "완벽해. 이 정도면 내가 가르칠 게 없어."
- A: "깔끔하게 처리했네. 잘했어."
- B: "해냈네! 근데 더 빠른 방법이 있어. 힌트를 확인해봐."
- C: "일단 해결한 건 좋아. 다시 도전해서 더 효율적으로 해보자."

---

## 6. 진행 시스템

### 6.1 저장 데이터

JSON 파일(`~/.advimture/progress.json`)에 로컬 저장.

```json
{
  "player": {
    "name": "dawn",
    "rank": "Junior",
    "total_missions_cleared": 3,
    "total_keystrokes": 1247,
    "created_at": "2026-03-06T10:00:00Z"
  },
  "tutorials": {
    "1-1": { "completed": true, "hints_used": 0 },
    "1-2": { "completed": true, "hints_used": 1 }
  },
  "missions": {
    "m01": {
      "best_grade": "A",
      "best_keystrokes": 10,
      "best_time_ms": 8500,
      "attempts": 3
    }
  },
  "time_attack": {
    "sprint_best_ms": 45200,
    "endurance_best_combo": 12,
    "daily": {
      "2026-03-06": { "time_ms": 32500, "grade": "A" }
    }
  }
}
```

### 6.2 승급 시스템

| 직급           | 조건                                  |
| -------------- | ------------------------------------- |
| Intern         | 시작 시 기본                          |
| Junior         | Tutorial 5개 완료                     |
| Senior         | Mission 5개 클리어 (C등급 이상)       |
| Staff          | Mission 10개 클리어 (평균 B등급 이상) |
| Principal      | 전 미션 A등급 이상                    |
| **Vim Master** | 전 미션 S등급                         |

### 6.3 진행 현황 화면

```
╔═════════════════════════════════════════════╗
║  PROGRESS              Rank: Junior DevOps  ║
║                                             ║
║  Tutorial:  ████████░░  8/10                ║
║  Missions:  ███░░░░░░░  3/10                ║
║                                             ║
║  M-01  nginx.conf       ★     S  [0:08]    ║
║  M-02  .env 편집         ★     A  [0:12]    ║
║  M-03  SSH config       ★★    B  [0:25]    ║
║  M-04  crontab          ★★    --           ║
║  M-05  로그 검색         ★★    (잠김)       ║
║  ...                                        ║
║                                             ║
║  Time Attack Best: 45.2s (Sprint)           ║
║  총 키스트로크: 1,247                        ║
║                                             ║
║  [b] 돌아가기                                ║
╚═════════════════════════════════════════════╝
```

---

## 7. 리워드 시스템

### 7.1 명령어 언락 연출

새로운 Vim 명령어를 처음 배우거나 사용할 때 "언락" 연출.

```
  ┌────────────────────────────┐
  │   NEW COMMAND UNLOCKED!    │
  │                            │
  │        ciw                 │
  │   Change Inner Word        │
  │                            │
  │   "커서 위의 단어를 삭제하고  │
  │    바로 입력 모드로 전환"    │
  │                            │
  │   Cheatsheet에 추가됨      │
  └────────────────────────────┘
```

### 7.2 Cheatsheet (수집 요소)

배운 명령어들이 쌓이는 개인 치트시트. 게임 중 `Ctrl+h` 또는 메뉴에서 확인 가능.

```
╔═══════════════════════════════════╗
║  MY CHEATSHEET         32/50     ║
║                                   ║
║  [이동]                           ║
║  h j k l    기본 이동        ✓   ║
║  w b e      단어 이동        ✓   ║
║  0 $        줄 처음/끝       ✓   ║
║  gg G       파일 처음/끝     ✓   ║
║  f{c} t{c}  문자 찾기        ✓   ║
║                                   ║
║  [편집]                           ║
║  x          문자 삭제        ✓   ║
║  dd         줄 삭제          ✓   ║
║  ciw        단어 변경        ✓   ║
║  ci"        따옴표 안 변경   NEW  ║
║  .          반복             ???  ║ ← 아직 안 배운 명령어
║  ...                              ║
╚═══════════════════════════════════╝
```

- `✓`: 배우고 사용해본 명령어
- `NEW`: 방금 배운 명령어
- `???`: 아직 해금되지 않은 명령어 (존재만 보여줘서 궁금증 유발)

---

## 8. 안정성 설계

### 8.1 데이터 파일 에러 처리

**진행 데이터 (`progress.json`):**

- 파일 없음 → 신규 생성 (기본값)
- JSON 파싱 실패 → 백업 파일(`progress.json.bak`) 존재 시 복구 시도, 없으면 초기화 + 경고
- 저장 시 항상 임시 파일에 먼저 쓴 후 rename (원자적 쓰기)
- 저장 성공 시 이전 파일을 `.bak`으로 보관

**미션/튜토리얼 YAML:**

- Go embed로 바이너리에 내장 → 파일 누락 불가
- YAML 스키마 검증: 필수 필드 누락 시 게임 시작 전 에러 메시지 + 종료
- 개발 모드에서만 외부 YAML 로딩 지원 (핫 리로드로 빠른 콘텐츠 테스트)

### 8.2 터미널 호환성

- **지원 대상**: iTerm2, Terminal.app, Windows Terminal, GNOME Terminal, Alacritty
- **색상**: TrueColor(24bit) 우선, 미지원 시 256color fallback, 최소 16color
- **유니코드**: 박스 드로잉 문자(─, │, ╔ 등) 사용. 미지원 터미널용 ASCII fallback (+, -, |)
- **Ctrl+키 조합**: `Ctrl+c`, `Ctrl+r`, `Ctrl+d`, `Ctrl+u`, `Ctrl+h` — OS별 시그널 충돌 확인
  - `Ctrl+c`: Bubbletea에서 기본적으로 SIGINT 처리. 게임 내에서 intercept하여 메뉴 복귀로 사용

---

## 9. 콘텐츠 확장 가이드라인

### 9.1 미션 추가 절차

```
1. missions/ 디렉토리에 YAML 파일 생성
2. 필수 필드 작성 (id, title, difficulty, initial_text, expected_text 등)
3. optimal_solutions에 최소 1개 이상의 풀이 등록
4. tips에 2~3개의 조건부 팁 작성
5. required_tutorials에 선행 조건 명시
6. 테스트: 최적해대로 풀었을 때 S등급이 나오는지 확인
```

### 9.2 튜토리얼 추가 절차

```
1. tutorials/ 디렉토리에 YAML 파일 생성
2. substep 최소 3개 (격리 연습 + 조합 연습 + 미니 챌린지)
3. allowed_keys로 학습 범위 제한
4. unlock_commands로 Cheatsheet 연동
5. 테스트: 힌트 없이 처음 플레이어가 클리어 가능한지 확인
```

### 9.3 Time Attack 문제 추가

```
1. 문제 풀 YAML에 항목 추가
2. 난이도 태그 필수 (easy/medium/hard)
3. 목표 키스트로크 5 이하 (마이크로 편집이므로)
4. 초기 텍스트 3줄 이내
```

---

## 10. 구현 우선순위 (개발 로드맵)

### Phase 1: 코어 엔진 (MVP)

1. 프로젝트 초기 세팅 (Go + Bubbletea)
2. 텍스트 버퍼 + Operation Log 기반 Undo/Redo
3. 모드 시스템 (Normal/Insert/Command/Operator-Pending)
4. 명령어 파서 (`{count}{operator}{motion}` 문법)
5. 기본 이동 (`hjkl`, `w`, `b`, `e`, `0`, `$`, `gg`, `G`)
6. 기본 편집 (`i`, `a`, `o`, `x`, `dd`, `dw`, `d$`, `p`, `u`, `Ctrl+r`)
7. Command Mode (`:wq`, `:q`, `:q!`, `:{number}`)
8. Unnamed Register
9. 에디터 View 렌더링 (줄 번호, 커서, 상태바)

### Phase 2: 학습 시스템

10. FTUE (첫 실행 인트로 시퀀스)
11. 튜토리얼 프레임워크 (5단계 학습 루프, 힌트 시스템)
12. Tutorial 1-1 ~ 1-5 구현
13. 메인 메뉴 + 네비게이션
14. 진행 상태 저장/로드

### Phase 3: 미션 시스템

15. 미션 YAML 로더 + 검증 시스템
16. 키스트로크 추적 + 등급 산출
17. 미션 결과 화면 (Vi 선배 피드백, 팁)
18. Mission M-01 ~ M-05 구현
19. 선행 조건 / 해금 시스템

### Phase 4: 명령어 확장

20. 검색 (`/pattern`, `n`, `N`) + 치환 (`:s`, `:%s`)
21. 텍스트 오브젝트 (`ciw`, `ci"`, `di(` 등)
22. Visual Mode (v, V)
23. Dot Repeat (`.`)
24. `f`, `t`, `r`, `J`
25. Tutorial 1-6 ~ 1-10 + Mission M-06 ~ M-10

### Phase 5: 게임성 확장

26. Time Attack (Sprint, Endurance)
27. 승급 시스템 + 명령어 언락 연출
28. Cheatsheet (수집 요소)
29. Free Mode (명령어 해설 기능)
30. Daily Challenge

### Phase 6: 폴리싱

31. 색상 테마 완성 + 색각 이상 대응
32. 터미널 리사이즈 대응
33. Vi 선배 메시지 전체 작성
34. 애니메이션/전환 효과
35. 결과 복사 기능 (Time Attack)

---

## 11. 프로젝트 구조

```
advimture/
├── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── app/
│   │   └── app.go              # 메인 Bubbletea 앱, 화면 전환 관리
│   ├── editor/
│   │   ├── buffer.go           # 텍스트 버퍼 ([]string + 편집 연산)
│   │   ├── cursor.go           # 커서 관리 + 이동 규칙
│   │   ├── mode.go             # 모드 정의 (Normal/Insert/Command/Visual/OpPending)
│   │   ├── parser.go           # 명령어 파서 (count + operator + motion)
│   │   ├── operator.go         # d, y, c 등 operator 실행
│   │   ├── motion.go           # w, b, $, 0, gg, G 등 motion 계산
│   │   ├── textobject.go       # iw, i", a( 등 text object 범위 계산
│   │   ├── register.go         # Unnamed register
│   │   ├── undo.go             # Operation Log 기반 Undo/Redo
│   │   ├── search.go           # /, n, N 검색
│   │   ├── dotrepeat.go        # . (dot repeat) 기록/재실행
│   │   └── editor.go           # 에디터 통합 (Model/Update/View)
│   ├── game/
│   │   ├── ftue.go             # First Time User Experience
│   │   ├── tutorial.go         # 튜토리얼 모드 (학습 루프 엔진)
│   │   ├── mission.go          # 미션 모드 (시나리오 + 검증)
│   │   ├── timeattack.go       # 타임어택 모드
│   │   ├── freemode.go         # 자유 연습 모드
│   │   └── cheatsheet.go       # 치트시트 (수집 현황)
│   ├── data/
│   │   ├── loader.go           # 미션/튜토리얼 데이터 로딩
│   │   ├── validator.go        # 결과 검증 (diff 비교)
│   │   ├── tutorials/          # 튜토리얼 데이터 (임베디드)
│   │   │   ├── t01.yaml
│   │   │   └── ...
│   │   └── missions/           # 미션 데이터 (임베디드)
│   │       ├── m01.yaml
│   │       └── ...
│   ├── progress/
│   │   ├── progress.go         # 진행 상태 관리
│   │   ├── rank.go             # 승급 시스템
│   │   └── storage.go          # JSON 파일 저장/로드
│   └── ui/
│       ├── menu.go             # 메인 메뉴
│       ├── result.go           # 미션 결과 화면
│       ├── progress_view.go    # 진행 현황 화면
│       ├── unlock.go           # 명령어 언락 연출
│       └── styles.go           # Lipgloss 스타일 정의
└── README.md
```
