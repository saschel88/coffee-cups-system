#!/bin/bash

# Coffee Cups System - Code Generation Script
# This script generates Go code from the OpenAPI specification

set -e

echo "🚀 Generating Go code from OpenAPI specification..."

# Check if swagger-codegen is installed
if ! command -v swagger-codegen &> /dev/null; then
    echo "❌ swagger-codegen not found. Installing..."
    
    # Install swagger-codegen
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew install swagger-codegen
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.19/swagger-codegen-cli-2.4.19.jar -O swagger-codegen-cli.jar
        chmod +x swagger-codegen-cli.jar
    else
        echo "❌ Unsupported OS. Please install swagger-codegen manually."
        exit 1
    fi
fi

# Create output directory
mkdir -p generated

# Generate Go client
echo "📦 Generating Go client..."
swagger-codegen generate \
    -i api/swagger.yaml \
    -l go \
    -o generated/client \
    --additional-properties packageName=client,packageVersion=1.0.0

# Generate Go server
echo "🖥️ Generating Go server..."
swagger-codegen generate \
    -i api/swagger.yaml \
    -l go-server \
    -o generated/server \
    --additional-properties packageName=server,packageVersion=1.0.0

# Generate TypeScript client
echo "📱 Generating TypeScript client..."
swagger-codegen generate \
    -i api/swagger.yaml \
    -l typescript-axios \
    -o generated/typescript-client \
    --additional-properties packageName=coffee-cups-client,packageVersion=1.0.0

# Generate documentation
echo "📚 Generating documentation..."
swagger-codegen generate \
    -i api/swagger.yaml \
    -l html2 \
    -o generated/docs

echo "✅ Code generation completed!"
echo ""
echo "Generated files:"
echo "  📁 generated/client/     - Go client library"
echo "  📁 generated/server/     - Go server stub"
echo "  📁 generated/typescript-client/ - TypeScript client"
echo "  📁 generated/docs/       - HTML documentation"
echo ""
echo "Next steps:"
echo "  1. Review generated code"
echo "  2. Integrate with your existing Go application"
echo "  3. Update handlers to match the generated interfaces"
