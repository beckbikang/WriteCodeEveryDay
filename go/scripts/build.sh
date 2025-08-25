#!/bin/bash

# 构建脚本
echo "Building WriteCodeEveryDay..."

# 清理之前的构建
rm -rf bin/

# 创建输出目录
mkdir -p bin/

# 构建应用程序
go build -o bin/writecode cmd/main.go

echo "Build completed successfully!"
