#!/bin/bash

set -e  # Exit on any error

# 색상 코드
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 로고 출력
echo -e "${BLUE}"
echo "  ________      ____        _ __                 __      __"
echo " / ____/ /___  / __ )____  (_) /_  ___  _____  / /_____/ /_"
echo "/ / __/ / __ \/ __  / __ \/ / / / _ \/ ___/ __ \/ / ___/ __/"
echo "/ /_/ / / /_/ / /_/ / /_/ / / / /  __/ /  / /_/ / / /  / /_"
echo "\____/_/\____/_____/\____/_/_/_/\___/_/  / .___/_/_/   \__/"
echo "                                       /_/"
echo -e "${NC}"
echo -e "${GREEN}🚀 Go Boilerplate Project Generator${NC}"
echo ""

# 프로젝트 이름 입력받기
while true; do
    echo -e "${YELLOW}📝 Enter your new project name:${NC}"
    read -p "> " PROJECT_NAME
    
    # 입력 검증
    if [[ -z "$PROJECT_NAME" ]]; then
        echo -e "${RED}❌ Project name cannot be empty!${NC}"
        continue
    fi
    
    if [[ ! "$PROJECT_NAME" =~ ^[a-zA-Z0-9_-]+$ ]]; then
        echo -e "${RED}❌ Project name can only contain letters, numbers, hyphens, and underscores!${NC}"
        continue
    fi
    
    break
done

# 현재 디렉토리 확인
CURRENT_DIR=$(pwd)
PARENT_DIR=$(dirname "$CURRENT_DIR")
TARGET_DIR="$PARENT_DIR/$PROJECT_NAME"

echo -e "${BLUE}📂 Current directory: $CURRENT_DIR${NC}"
echo -e "${BLUE}🎯 Target directory: $TARGET_DIR${NC}"

# 대상 디렉토리가 이미 존재하는지 확인
if [[ -d "$TARGET_DIR" ]]; then
    echo -e "${YELLOW}⚠️  Directory '$TARGET_DIR' already exists!${NC}"
    echo -e "${YELLOW}🗑️  Removing existing directory...${NC}"
    rm -rf "$TARGET_DIR"
    echo -e "${GREEN}✅ Existing directory removed${NC}"
fi

echo ""
echo -e "${GREEN}🚀 Creating new project '$PROJECT_NAME'...${NC}"
echo ""

# 1. 현재 디렉토리를 상위로 복사
echo -e "${YELLOW}📋 Step 1: Copying project files...${NC}"
cp -r "$CURRENT_DIR" "$TARGET_DIR"
echo -e "${GREEN}✅ Files copied successfully${NC}"

# 2. 불필요한 파일들 삭제
echo -e "${YELLOW}🧹 Step 2: Cleaning up unnecessary files...${NC}"
cd "$TARGET_DIR"

# .git, bin, tmp 디렉토리 삭제
rm -rf .git bin tmp .cursor 2>/dev/null || true

# 로그 파일들 삭제
find . -name "*.log" -delete 2>/dev/null || true

echo -e "${GREEN}✅ Cleanup completed${NC}"

# 3. go.mod에서 모듈명 변경
echo -e "${YELLOW}🔧 Step 3: Updating module name...${NC}"
sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" go.mod && rm go.mod.bak
echo -e "${GREEN}✅ Module name updated${NC}"

# 4. 모든 Go 파일에서 import 경로 변경
echo -e "${YELLOW}🔄 Step 4: Updating import paths...${NC}"
find . -name "*.go" -type f -exec sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" {} \; -exec rm {}.bak \;
echo -e "${GREEN}✅ Import paths updated${NC}"

# 5. 설정 파일 업데이트
echo -e "${YELLOW}⚙️  Step 5: Updating configuration files...${NC}"

# config.yaml 업데이트
if [[ -f "configs/config.yaml" ]]; then
    sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" configs/config.yaml && rm configs/config.yaml.bak
    sed -i.bak "s/go_boilerplate/${PROJECT_NAME//-/_}/g" configs/config.yaml && rm configs/config.yaml.bak
fi

# config.dev.yaml 업데이트  
if [[ -f "configs/config.dev.yaml" ]]; then
    sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" configs/config.dev.yaml && rm configs/config.dev.yaml.bak
    sed -i.bak "s/go_boilerplate_dev/${PROJECT_NAME//-/_}_dev/g" configs/config.dev.yaml && rm configs/config.dev.yaml.bak
fi

echo -e "${GREEN}✅ Configuration files updated${NC}"

# 6. 스크립트 파일 업데이트
echo -e "${YELLOW}📜 Step 6: Updating script files...${NC}"

# build.sh 업데이트
if [[ -f "scripts/build.sh" ]]; then
    sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" scripts/build.sh && rm scripts/build.sh.bak
fi

# run.sh 업데이트
if [[ -f "scripts/run.sh" ]]; then
    sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" scripts/run.sh && rm scripts/run.sh.bak
fi

echo -e "${GREEN}✅ Script files updated${NC}"

# 7. README.md 업데이트
echo -e "${YELLOW}📖 Step 7: Updating README.md...${NC}"
if [[ -f "README.md" ]]; then
    sed -i.bak "s/go-boilerplate/$PROJECT_NAME/g" README.md && rm README.md.bak
    
    # README 제목과 설명 개인화
    PROJECT_TITLE=$(echo "$PROJECT_NAME" | tr '[:lower:]' '[:upper:]' | tr '-' ' ')
    sed -i.bak "1s/.*/# $PROJECT_NAME/" README.md && rm README.md.bak
    sed -i.bak "s/Go 언어와 Hexagonal Architecture를 적용한 보일러플레이트 프로젝트입니다./$PROJECT_TITLE - Go 언어와 Hexagonal Architecture를 적용한 프로젝트입니다./" README.md && rm README.md.bak
fi

echo -e "${GREEN}✅ README.md updated${NC}"

# 8. Swagger 주석 업데이트
echo -e "${YELLOW}📚 Step 8: Updating Swagger documentation...${NC}"
if [[ -f "cmd/main.go" ]]; then
    PROJECT_TITLE_CAPITALIZED=$(echo "$PROJECT_NAME" | sed 's/-/ /g' | sed 's/\b\(.\)/\u\1/g')
    sed -i.bak "s/Go Boilerplate API/$PROJECT_TITLE_CAPITALIZED API/g" cmd/main.go && rm cmd/main.go.bak
    sed -i.bak "s/This is a sample go boilerplate API server./This is a $PROJECT_NAME API server./g" cmd/main.go && rm cmd/main.go.bak
fi

echo -e "${GREEN}✅ Swagger documentation updated${NC}"

# 9. go mod tidy 실행
echo -e "${YELLOW}📦 Step 9: Resolving dependencies...${NC}"
go mod tidy
echo -e "${GREEN}✅ Dependencies resolved${NC}"

# 10. Swagger 문서 생성 (선택사항)
echo -e "${YELLOW}📋 Step 10: Generating Swagger documentation...${NC}"
SWAG_PATH="$HOME/go/bin/swag"
if [[ -f "$SWAG_PATH" ]]; then
    "$SWAG_PATH" init -g cmd/main.go
    echo -e "${GREEN}✅ Swagger documentation generated${NC}"
else
    echo -e "${YELLOW}⚠️  Swagger CLI not found. Run 'go install github.com/swaggo/swag/cmd/swag@latest' to install.${NC}"
fi

# 11. 초기 빌드 테스트
echo -e "${YELLOW}🔨 Step 11: Testing initial build...${NC}"
if go build -o "bin/$PROJECT_NAME" ./cmd/main.go; then
    echo -e "${GREEN}✅ Build successful${NC}"
else
    echo -e "${RED}❌ Build failed. Please check the errors above.${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}🎉 Project '$PROJECT_NAME' created successfully!${NC}"
echo ""
echo -e "${BLUE}📍 Project location: $TARGET_DIR${NC}"
echo ""
echo -e "${YELLOW}🚀 Next steps:${NC}"
echo -e "   1. cd $TARGET_DIR"
echo -e "   2. ./scripts/run.sh                  # Run the development server"
echo -e "   3. open http://localhost:8080/swagger/index.html  # View API docs"
echo ""
echo -e "${BLUE}🛠️  Available commands:${NC}"
echo -e "   make run          # Start development server"
echo -e "   make build        # Build the project"  
echo -e "   make test         # Run tests"
echo -e "   make docs         # Generate Swagger docs"
echo ""
echo -e "${GREEN}Happy coding! 🎯${NC}" 