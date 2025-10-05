#!/bin/bash

# Coffee Cups System - Rules Enforcement Script
# This script checks if the project follows the defined rules

set -e

echo "🔍 Checking Coffee Cups System rules compliance..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
PASSED=0
FAILED=0
WARNINGS=0

# Function to print results
print_result() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
        ((PASSED++))
    else
        echo -e "${RED}❌ $2${NC}"
        ((FAILED++))
    fi
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
    ((WARNINGS++))
}

echo ""
echo "📋 Running rule checks..."
echo ""

# 1. Check project structure
echo "🏗️  Checking project structure..."
if [ -d "cmd" ] && [ -d "internal" ] && [ -d "api" ] && [ -d "docs" ]; then
    print_result 0 "Project structure follows Clean Architecture"
else
    print_result 1 "Project structure does not follow Clean Architecture"
fi

# 2. Check Go code quality
echo ""
echo "💻 Checking Go code quality..."
if command -v gofmt &> /dev/null; then
    if [ "$(gofmt -s -l . | wc -l)" -eq 0 ]; then
        print_result 0 "Go code is properly formatted"
    else
        print_result 1 "Go code is not properly formatted"
        echo "Run: gofmt -s -w ."
    fi
else
    print_warning "gofmt not found, skipping format check"
fi

# 3. Check for go vet issues
if command -v go &> /dev/null; then
    if go vet ./... 2>/dev/null; then
        print_result 0 "Go vet passed"
    else
        print_result 1 "Go vet found issues"
    fi
else
    print_warning "Go not found, skipping vet check"
fi

# 4. Check test coverage
echo ""
echo "🧪 Checking test coverage..."
if [ -f "coverage.out" ]; then
    COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
    if (( $(echo "$COVERAGE >= 80" | bc -l) )); then
        print_result 0 "Test coverage is $COVERAGE% (>= 80%)"
    else
        print_result 1 "Test coverage is $COVERAGE% (< 80%)"
    fi
else
    print_warning "No coverage file found, run tests first"
fi

# 5. Check API documentation
echo ""
echo "📚 Checking API documentation..."
if [ -f "api/swagger.yaml" ]; then
    print_result 0 "API documentation exists"
    
    # Check if swagger file is valid
    if command -v swagger &> /dev/null; then
        if swagger validate api/swagger.yaml 2>/dev/null; then
            print_result 0 "API specification is valid"
        else
            print_result 1 "API specification is invalid"
        fi
    else
        print_warning "Swagger CLI not found, skipping validation"
    fi
else
    print_result 1 "API documentation missing"
fi

# 6. Check for security issues
echo ""
echo "🔒 Checking security..."
if command -v gosec &> /dev/null; then
    if gosec ./... 2>/dev/null; then
        print_result 0 "No security issues found"
    else
        print_result 1 "Security issues found"
    fi
else
    print_warning "gosec not found, skipping security check"
fi

# 7. Check Docker configuration
echo ""
echo "🐳 Checking Docker configuration..."
if [ -f "Dockerfile" ]; then
    print_result 0 "Dockerfile exists"
    
    # Check if Dockerfile follows best practices
    if grep -q "FROM.*alpine" Dockerfile; then
        print_result 0 "Using Alpine Linux base image"
    else
        print_warning "Consider using Alpine Linux for smaller image size"
    fi
    
    if grep -q "USER" Dockerfile; then
        print_result 0 "Running as non-root user"
    else
        print_result 1 "Not running as non-root user"
    fi
else
    print_result 1 "Dockerfile missing"
fi

# 8. Check environment configuration
echo ""
echo "⚙️  Checking environment configuration..."
if [ -f "configs/config.yaml" ]; then
    print_result 0 "Configuration file exists"
else
    print_result 1 "Configuration file missing"
fi

# 9. Check for secrets
echo ""
echo "🔐 Checking for secrets..."
if grep -r -i "password\|secret\|key\|token" --exclude-dir=.git --exclude="*.md" . | grep -v "config.yaml" | grep -v "swagger.yaml" | grep -v "RULES.md" | grep -v "check-rules.sh"; then
    print_result 1 "Potential secrets found in code"
else
    print_result 0 "No secrets found in code"
fi

# 10. Check documentation
echo ""
echo "📖 Checking documentation..."
if [ -f "README.md" ]; then
    print_result 0 "README.md exists"
else
    print_result 1 "README.md missing"
fi

if [ -f "docs/API.md" ]; then
    print_result 0 "API documentation exists"
else
    print_warning "API documentation missing"
fi

# 11. Check for proper error handling
echo ""
echo "🚨 Checking error handling..."
if grep -r "_ = " --include="*.go" . | grep -v "test" | grep -v "vendor"; then
    print_result 1 "Found ignored errors (using _ = )"
else
    print_result 0 "No ignored errors found"
fi

# 12. Check for proper logging
echo ""
echo "📝 Checking logging..."
if grep -r "fmt.Print" --include="*.go" . | grep -v "test" | grep -v "vendor"; then
    print_result 1 "Found fmt.Print usage (use structured logging)"
else
    print_result 0 "No fmt.Print usage found"
fi

# 13. Check for magic numbers
echo ""
echo "🔢 Checking for magic numbers..."
if grep -r "[^a-zA-Z_]\([0-9]\{3,\}\)" --include="*.go" . | grep -v "test" | grep -v "vendor" | grep -v "time" | grep -v "http"; then
    print_warning "Found potential magic numbers"
else
    print_result 0 "No magic numbers found"
fi

# 14. Check for proper imports
echo ""
echo "📦 Checking imports..."
if grep -r "import.*\"" --include="*.go" . | grep -v "test" | grep -v "vendor"; then
    print_result 1 "Found string imports (use proper import paths)"
else
    print_result 0 "No string imports found"
fi

# 15. Check for proper database usage
echo ""
echo "🗄️  Checking database usage..."
if grep -r "Raw(" --include="*.go" . | grep -v "test" | grep -v "vendor"; then
    print_result 1 "Found Raw() queries (use parameterized queries)"
else
    print_result 0 "No Raw() queries found"
fi

# Summary
echo ""
echo "📊 Summary:"
echo "==========="
echo -e "${GREEN}✅ Passed: $PASSED${NC}"
echo -e "${RED}❌ Failed: $FAILED${NC}"
echo -e "${YELLOW}⚠️  Warnings: $WARNINGS${NC}"

if [ $FAILED -eq 0 ]; then
    echo ""
    echo -e "${GREEN}🎉 All critical rules passed!${NC}"
    exit 0
else
    echo ""
    echo -e "${RED}❌ Some rules failed. Please fix the issues above.${NC}"
    exit 1
fi
