@echo off
setlocal enabledelayedexpansion

:: 脚本说明：
:: 此脚本用于在当前系统环境下编译 Go 项目，生成适用于当前平台的二进制文件。
:: 输出文件名将包含操作系统和架构信息，与 cross_compile 脚本保持一致。

:: 切换到项目根目录
cd /d "%~dp0\.."

:: --- 脚本执行 ---

:: 创建 bin 目录 (如果不存在)
if not exist "bin" (
    echo "Creating directory: bin"
    mkdir "bin"
)

:: 获取操作系统和架构
set "GOOS=windows"
set "GOARCH="
if /i "%PROCESSOR_ARCHITECTURE%"=="AMD64" set "GOARCH=amd64"
if /i "%PROCESSOR_ARCHITECTURE%"=="x86" set "GOARCH=386"
if /i "%PROCESSOR_ARCHITECTURE%"=="ARM64" set "GOARCH=arm64"

if not defined GOARCH (
    echo "Unsupported architecture: %PROCESSOR_ARCHITECTURE%"
    exit /b 1
)

echo "OS: %GOOS%, Arch: %GOARCH%"

:: 整理 Go 模块依赖
echo "Running go mod tidy..."
go mod tidy

:: 遍历 cmd 目录并编译所有应用
echo "Building commands for %GOOS%/%GOARCH%..."
for /d %%c in (cmd\*) do (
    set "TARGET_NAME=%%~nc"
    set "PACKAGE_NAME=.\cmd\!TARGET_NAME!"

    set "TARGET_FILE=bin\!TARGET_NAME!_%GOOS%_%GOARCH%.exe"
    echo "Building executable !PACKAGE_NAME! => !TARGET_FILE!"
    go build -o "!TARGET_FILE!" "!PACKAGE_NAME!"
)

echo.
echo "Build completed successfully."
endlocal
