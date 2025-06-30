@echo off
REM --- 此脚本用于卸载应用的 Windows 服务。 ---
REM --- 必须以管理员权限运行。 ---

SET "SERVICE_NAME=YourGoProject"
SET "SERVICE_DISPLAY_NAME=Your Go Project Service"

ECHO ===============================================
ECHO 正在卸载服务: %SERVICE_DISPLAY_NAME%
ECHO ===============================================

sc stop "%SERVICE_NAME%"
sc delete "%SERVICE_NAME%"

IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo 错误: 卸载服务失败。
    echo 服务可能未安装，或者您需要以管理员身份运行。
    pause
    exit /b %ERRORLEVEL%
)

echo.
echo 服务 '%SERVICE_DISPLAY_NAME%' 已成功卸载。
pause 