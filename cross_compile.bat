@echo off
setlocal enabledelayedexpansion

:: 设置目标目录和文件名
set "TARGET_DIR=release"
set "TARGET_NAME=your-go-project-name"
:: 获取时间
set "TIMESTAMP=%date:~0,4%%date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%%time:~6,2%"
:: 设置平台列表
:: set "PLATFORMS=darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/amd64"
set "PLATFORMS=windows/amd64"

:: 清理并创建目标目录
if exist "%TARGET_DIR%" rmdir /s /q "%TARGET_DIR%"
mkdir "%TARGET_DIR%"

:: 遍历平台列表
for %%p in (%PLATFORMS%) do (
    :: 解析平台和架构
    for /f "tokens=1,2 delims=/" %%a in ("%%p") do (
        set "GOOS=%%a"
        set "GOARCH=%%b"
    )

    :: 设置目标文件名
    set "TARGET=%TARGET_DIR%\%TARGET_NAME%"
    if "!GOOS!"=="windows" (
        set "TARGET=%TARGET_DIR%\%TARGET_NAME%_!GOOS!_!GOARCH!_%TIMESTAMP%.exe"
    ) else (
        set "TARGET=%TARGET_DIR%/%TARGET_NAME%_!GOOS!_!GOARCH!"
    )

    :: 打印构建信息
    echo "build => !TARGET!"

    :: 设置环境变量并构建
    set CGO_ENABLED=1
    set "GOOS=!GOOS!"
    set "GOARCH=!GOARCH!"
    ::go build -mod=vendor -trimpath -o "!TARGET!"
    go build -mod=vendor -o "!TARGET!"
)
