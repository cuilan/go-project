@echo off
REM --- 此脚本用于将应用程序安装为 Windows 服务。 ---
REM --- 必须以管理员权限运行。 ---

SET "SERVICE_NAME=YourGoProject"
SET "SERVICE_DISPLAY_NAME=Your Go Project Service"
SET "SERVICE_DESCRIPTION=这是一个 Go 模板项目的服务。"

REM --- 智能查找可执行文件的完整路径 ---
REM %~dp0 是批处理文件所在的目录，我们假定它在 'release\init\windows'
SET "RELEASE_ROOT=%~dp0..\\"
SET "EXE_PATH="
FOR /f "delims=" %%F IN ('dir /b /a-d "%RELEASE_ROOT%%SERVICE_NAME%*.exe"') DO SET "EXE_PATH=%RELEASE_ROOT%%%F"

REM 检查可执行文件是否存在
IF "%EXE_PATH%"=="" (
    echo 错误: 在 %RELEASE_ROOT% 未找到 %SERVICE_NAME%*.exe 可执行文件。
    echo 请确保压缩包已正确解压。
    pause
    exit /b 1
)

ECHO ===============================================
ECHO 正在安装服务: %SERVICE_DISPLAY_NAME%
ECHO 服务名称: %SERVICE_NAME%
ECHO 可执行文件路径: %EXE_PATH%
ECHO ===============================================

sc create "%SERVICE_NAME%" binPath= "\"%EXE_PATH%\"" start= auto DisplayName= "%SERVICE_DISPLAY_NAME%"
IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo 错误: 创建服务失败。请确保以管理员身份运行。
    pause
    exit /b %ERRORLEVEL%
)

sc description "%SERVICE_NAME%" "%SERVICE_DESCRIPTION%"

sc start "%SERVICE_NAME%"
IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo 警告: 服务已安装但无法启动。
    echo 请检查 Windows 事件查看器中的错误日志。
    pause
) ELSE (
    echo.
    echo 服务 '%SERVICE_DISPLAY_NAME%' 已成功安装并启动。
)

pause 