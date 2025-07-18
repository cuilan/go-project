@echo off
REM --- This script is used to uninstall the application's Windows service. ---
REM --- Must be run with administrator privileges. ---

SET "SERVICE_NAME=your-app"
SET "SERVICE_DISPLAY_NAME=Your App Service"

ECHO ===============================================
ECHO Uninstalling service: %SERVICE_DISPLAY_NAME%
ECHO ===============================================

sc stop "%SERVICE_NAME%"
sc delete "%SERVICE_NAME%"

IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Error: Failed to uninstall service.
    echo The service may not be installed, or you need to run as administrator.
    pause
    exit /b %ERRORLEVEL%
)

echo.
echo Service '%SERVICE_DISPLAY_NAME%' has been successfully uninstalled.
pause 