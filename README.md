# go-boilerplate

Go 언어와 Hexagonal Architecture를 적용한 보일러플레이트 프로젝트입니다.  
새로운 Go 백엔드 프로젝트를 시작할 때 참고할 수 있는 구조와 패턴을 제공합니다.

## 🚀 Quick Start

### **새로운 프로젝트 생성**
```bash
# 1. 이 보일러플레이트를 사용하여 새 프로젝트 생성
make create-project

# 2. 생성된 프로젝트로 이동
cd ../your-project-name

# 3. 개발 서버 실행
make run

# 4. API 문서 확인
open http://localhost:8080/swagger/index.html
```

### **개발 환경에서 직접 실행**
```bash
# 현재 보일러플레이트를 직접 실행
make run

# 또는 전통적인 방법
go run ./cmd/main.go
```

## 🛠️ Make 명령어

모든 주요 작업은 `make` 명령어로 수행할 수 있습니다:

```bash
make help                # 📋 사용 가능한 명령어 목록 표시
make create-project      # 🎯 새 프로젝트 생성
make run                 # 🏃 개발 서버 실행
make build               # 🔨 프로젝트 빌드
make test                # 🧪 테스트 실행
make docs                # 📚 Swagger 문서 생성
make clean               # 🧹 빌드 아티팩트 정리
make install             # 📦 의존성 설치
make lint                # 🔍 코드 린팅
make format              # 🎨 코드 포맷팅
make dev                 # 🔧 개발 환경 설정 + 실행
make quick-start         # ⚡ 빠른 시작 가이드
```

## 아키텍처

이 프로젝트는 **Hexagonal Architecture (포트와 어댑터 패턴)**를 적용하여 구현되었습니다.

### Hexagonal Architecture의 특징
- **도메인 중심**: 핵심 비즈니스 로직이 외부 시스템에 의존하지 않음
- **테스트 용이성**: 각 계층을 독립적으로 테스트 가능
- **유연성**: 외부 시스템 변경 시 어댑터만 수정하면 됨
- **관심사 분리**: 도메인 로직과 인프라 로직의 명확한 분리

### 계층 구조
```
        Inbound Adapters (들어오는 요청)
              ↓
            Ports (포트)
              ↓
          Domain (도메인)
              ↓
            Ports (포트)
              ↓
       Outbound Adapters (나가는 요청)
```

## 프로젝트 구조

```
.
├── cmd/                              # 메인 애플리케이션 진입점
│   └── main.go
├── internal/                        # 비공개 애플리케이션 코드
│   ├── domain/                      # 도메인 계층 (핵심 비즈니스 로직)
│   │   ├── model/                   # 도메인 모델
│   │   │   ├── todo.go
│   │   │   └── user.go
│   │   ├── port/                    # 포트 (인터페이스)
│   │   │   ├── todo_port.go
│   │   │   └── user_port.go
│   │   ├── service/                 # 도메인 서비스
│   │   │   ├── todo_service.go
│   │   │   └── user_service.go
│   │   └── errors.go                # 도메인 에러
│   ├── adapter/                     # 어댑터 계층 (외부 시스템과의 통신)
│   │   ├── inbound/                 # 인바운드 어댑터 (들어오는 요청)
│   │   │   └── http/                # HTTP 핸들러
│   │   │       ├── todo_handler.go
│   │   │       └── user_handler.go
│   │   └── outbound/                # 아웃바운드 어댑터 (나가는 요청)
│   │       └── persistence/         # 데이터 저장소
│   │           ├── todo_repository.go
│   │           └── user_repository.go
│   └── config/                      # 애플리케이션 설정
│       └── config.go
├── scripts/                         # 빌드, 실행, 테스트 스크립트
│   ├── build.sh                     # 빌드 + Swagger 생성 스크립트
│   ├── run.sh                       # 실행 스크립트
│   ├── test.sh                      # 테스트 스크립트
│   └── create-project.sh            # 새 프로젝트 생성 스크립트
├── configs/                         # 설정 파일
├── docs/                            # Swagger 문서 (자동 생성)
├── Makefile                         # Make 명령어 정의
├── go.mod                           # Go 모듈 정의
└── README.md                        # 프로젝트 문서
```

## 포함된 기능들

### **도메인 예시**
- **Todo 관리**: 할 일 생성, 조회, 수정, 삭제
- **User 관리**: 사용자 생성, 조회, 수정, 삭제

### **기술 스택**
- **웹 프레임워크**: Gin
- **API 문서**: Swagger/OpenAPI
- **아키텍처**: Hexagonal Architecture
- **패턴**: 의존성 주입, 포트와 어댑터

### **개발 도구**
- **실시간 리로드**: Air
- **빌드 자동화**: Make + Shell Scripts
- **코드 생성**: Swagger Codegen
- **프로젝트 생성**: 자동화된 템플릿 시스템

## 개발 환경 설정

### 1. Go 설치
```sh
# macOS
brew install go

# Linux
sudo apt-get install golang-go
```

### 2. Air 설치 (실시간 리로드)
```sh
go install github.com/air-verse/air@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

### 3. Swagger 설치
```sh
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

## 실행 방법

### 개발 모드 실행
```sh
# Make를 사용한 실행 (권장)
make run

# Air를 사용한 실시간 리로드
air

# 또는 직접 실행
go run ./cmd/main.go

# 또는 스크립트 사용
./scripts/run.sh
```

### 빌드 (Swagger 문서 포함)
```sh
# Make를 사용한 빌드 (권장)
make build

# 또는 스크립트 사용
./scripts/build.sh
```

### 테스트
```sh
# Make를 사용한 테스트
make test

# 또는 스크립트 사용
./scripts/test.sh
```

## API 문서

### Swagger UI
- URL: http://localhost:8080/swagger/index.html
- API 문서를 실시간으로 확인하고 테스트할 수 있습니다.

### API 문서 갱신
```sh
# Make를 사용한 문서 생성
make docs

# 또는 직접 실행
swag init -g cmd/main.go
```

## API 엔드포인트

### Todo API
- `POST /todos` - 새로운 Todo 생성
- `GET /todos` - 모든 Todo 조회
- `GET /todos/:id` - 특정 Todo 조회
- `PUT /todos/:id` - Todo 수정
- `DELETE /todos/:id` - Todo 삭제

### User API
- `POST /users` - 새로운 User 생성
- `GET /users` - 모든 User 조회
- `GET /users/:id` - 특정 User 조회
- `PUT /users/:id` - User 수정
- `DELETE /users/:id` - User 삭제

## Hexagonal Architecture 개발 가이드

### 1. 새로운 기능 추가 절차

#### 1.1 도메인 모델 정의
```go
// internal/domain/model/your_model.go
type YourModel struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

#### 1.2 포트(인터페이스) 정의
```go
// internal/domain/port/your_port.go
type YourRepositoryPort interface {
    Create(ctx context.Context, model *model.YourModel) error
    GetByID(ctx context.Context, id int) (*model.YourModel, error)
}

type YourServicePort interface {
    CreateYourModel(ctx context.Context, req *model.CreateRequest) (*model.YourModel, error)
}
```

#### 1.3 도메인 서비스 구현
```go
// internal/domain/service/your_service.go
type YourService struct {
    repo port.YourRepositoryPort
}

func NewYourService(repo port.YourRepositoryPort) *YourService {
    return &YourService{repo: repo}
}
```

#### 1.4 어댑터 구현
```go
// Outbound Adapter (Repository)
// internal/adapter/outbound/persistence/your_repository.go

// Inbound Adapter (HTTP Handler)
// internal/adapter/inbound/http/your_handler.go
```

#### 1.5 의존성 주입 (main.go)
```go
// main.go에서 의존성 연결
repo := persistence.NewYourRepository()
service := service.NewYourService(repo)
handler := http.NewYourHandler(service)
```

### 2. 계층별 역할

#### Domain Layer (도메인 계층)
- **Model**: 핵심 비즈니스 엔티티
- **Port**: 외부 시스템과의 인터페이스 정의
- **Service**: 핵심 비즈니스 로직 구현

#### Adapter Layer (어댑터 계층)
- **Inbound Adapter**: 외부에서 들어오는 요청 처리 (HTTP, CLI, gRPC 등)
- **Outbound Adapter**: 외부 시스템 호출 (DB, 외부 API, 파일 시스템 등)

### 3. 코드 컨벤션
- 패키지명은 소문자
- 파일명은 snake_case
- 구조체와 메서드는 CamelCase
- 인터페이스는 `Port` 접미사 사용 (예: `TodoServicePort`)
- 구현체는 구체적인 이름 사용 (예: `TodoService`, `TodoRepository`)

### 4. 테스트 전략
- **도메인 서비스 테스트**: Mock Repository를 사용하여 비즈니스 로직 테스트
- **어댑터 테스트**: 각 어댑터를 독립적으로 테스트
- **통합 테스트**: 전체 플로우 테스트

```sh
# 테스트 실행
make test

# 또는 Go 직접 실행
go test ./...

# 특정 패키지 테스트
go test ./internal/domain/service/...
```

### 5. 의존성 방향
```
Inbound Adapter → Port → Domain Service → Port → Outbound Adapter
```
- 도메인은 어댑터에 의존하지 않음
- 어댑터가 도메인에 의존함
- 포트를 통한 의존성 역전 원칙 적용

## 이 보일러플레이트를 사용하는 방법

### 1. 새 프로젝트 생성 (권장)
```bash
# Make CLI를 사용한 새 프로젝트 생성
make create-project

# 대화형 프롬프트에서 프로젝트 이름 입력
# 예: my-awesome-api
```

### 2. 수동으로 프로젝트 생성
```sh
# 프로젝트 복제
git clone <this-repo> my-new-project
cd my-new-project

# go.mod에서 모듈명 변경
# go-boilerplate → my-new-project 로 변경

# 모든 import 경로 변경
find . -name "*.go" -exec sed -i 's/go-boilerplate/my-new-project/g' {} \;

# 의존성 업데이트
make install
```

### 3. 도메인 모델 교체
- `internal/domain/model/` 에서 Todo, User 를 본인의 도메인으로 교체
- 비즈니스 로직에 맞게 서비스와 리포지토리 구현

### 4. 외부 시스템 연동
- `internal/adapter/outbound/` 에서 실제 데이터베이스, 외부 API 연동
- 현재는 메모리 저장소로 구현되어 있음

## 프로젝트 생성 도구의 특징

### **자동화된 설정**
- ✅ 모든 import 경로 자동 변경
- ✅ 설정 파일 자동 업데이트
- ✅ Swagger 문서 자동 생성
- ✅ 의존성 자동 해결

### **검증 및 테스트**
- ✅ 프로젝트 이름 유효성 검사
- ✅ 중복 디렉토리 확인
- ✅ 초기 빌드 테스트
- ✅ 단계별 성공 확인

### **사용자 친화적**
- ✅ 컬러풀한 출력
- ✅ 진행 상황 표시
- ✅ 에러 처리 및 안내
- ✅ 다음 단계 가이드

## 장점

1. **유지보수성**: 각 계층이 명확히 분리되어 변경 영향 최소화
2. **테스트 용이성**: Mock을 사용한 독립적인 테스트 가능
3. **확장성**: 새로운 어댑터 추가가 쉬움 (예: gRPC, GraphQL 추가)
4. **기술 독립성**: 데이터베이스나 프레임워크 변경이 도메인에 영향 없음
5. **학습 자료**: Clean Architecture 및 Hexagonal Architecture 학습에 적합
6. **생산성**: 새 프로젝트 생성이 자동화되어 빠른 시작 가능

## 라이선스

MIT License - 자유롭게 사용, 수정, 배포하실 수 있습니다.
