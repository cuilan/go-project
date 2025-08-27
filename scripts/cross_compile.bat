@echo off
setlocal enabledelayedexpansion

:: 脚本说明：
:: 此脚本用于交叉编译 Go 项目，生成适用于不同操作系统和架构的二进制文件。

:: 切换到项目根目录
:: %~dp0 会展开为当前批处理文件所在的驱动器和路径
cd /d "%~dp0\.."

:: --- 可配置变量 ---
:: 目标二进制文件输出目录
set "TARGET_DIR=bin"
:: 目标二进制文件名 (不含后缀)
:: set "TARGET_NAME=client"
:: 要编译的 Go 包路径
:: set "PACKAGE_NAME=./cmd/client"
:: 目标平台列表，格式为 "操作系统/架构"，用空格分隔
:: 完整的平台列表: "darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/amd64"
set "PLATFORMS=linux/amd64 windows/amd64"

:: --- 脚本执行 ---

:: 清理并创建目标目录
if exist "%TARGET_DIR%" (
    echo "Cleaning up directory: %TARGET_DIR%"
    rmdir /s /q "%TARGET_DIR%"
)
echo "Creating directory: %TARGET_DIR%"
mkdir "%TARGET_DIR%"

:: 遍历 cmd 目录下的所有子目录并开始编译
for /d %%c in (cmd\*) do (
    set "TARGET_NAME=%%~nc"
    set "PACKAGE_NAME=.\cmd\!TARGET_NAME!"
    echo --- Building command: !TARGET_NAME! ---

    for %%p in (%PLATFORMS%) do (
        :: 解析平台和架构
        for /f "tokens=1,2 delims=/" %%a in ("%%p") do (
            set "GOOS=%%a"
            set "GOARCH=%%b"
        )
    
        :: 禁用 CGO
        set CGO_ENABLED=0
    
        :: 构造输出文件名
        set "TARGET_FILE=%TARGET_DIR%\!TARGET_NAME!_!GOOS!_!GOARCH!"
        if "!GOOS!" == "windows" (
            set "TARGET_FILE=!TARGET_FILE!.exe"
        )
    
        :: 打印构建信息
        echo build =^> !TARGET_FILE!
    
        :: 执行编译命令
        go build -trimpath -o "!TARGET_FILE!" "!PACKAGE_NAME!"
    )
)

echo.
echo All builds completed.
endlocal
