@echo off
REM --- This script is used to install the application as a Windows service. ---
REM --- Must be run with administrator privileges. ---

SET "SERVICE_NAME=your-app"
SET "SERVICE_DISPLAY_NAME=Your App Service"
SET "SERVICE_DESCRIPTION=This is a your app service."

REM --- Intelligently find the full path of the executable file ---
SET "RELEASE_ROOT=%~dp0"
SET "BIN_PATH=%RELEASE_ROOT%bin"
IF NOT EXIST "%BIN_PATH%" (
    echo Error: bin directory not found at %BIN_PATH%
    echo Please ensure the package has been extracted correctly.
    pause
    exit /b 1
)

SET "EXE_PATH=%BIN_PATH%\%SERVICE_NAME%.exe"
IF NOT EXIST "%EXE_PATH%" (
    echo Error: %SERVICE_NAME%.exe executable file not found at %EXE_PATH%
    echo Please ensure the package has been extracted correctly and the executable is in the bin/ directory.
    pause
    exit /b 1
)

REM Check if the executable file exists
IF "%EXE_PATH%"=="" (
    echo Error: %SERVICE_NAME%.exe executable file not found in %RELEASE_ROOT%\bin\
    echo Please ensure the package has been extracted correctly and the executable is in the bin/ directory.
    pause
    exit /b 1
)

REM Check if configs directory exists
SET "CONFIG_DIR=%RELEASE_ROOT%configs"
IF NOT EXIST "%CONFIG_DIR%" (
    echo Error: configs directory not found at %CONFIG_DIR%
    echo Please ensure the package has been extracted correctly.
    pause
    exit /b 1
)

ECHO ===============================================
ECHO Installing service: %SERVICE_DISPLAY_NAME%
ECHO Service name: %SERVICE_NAME%
ECHO Executable path: %EXE_PATH%
ECHO Configs path: %CONFIG_DIR%
ECHO ===============================================

sc create "%SERVICE_NAME%" binPath= "\"%EXE_PATH%\" \"--config-dir\" \"%CONFIG_DIR%\"" start= auto DisplayName= "%SERVICE_DISPLAY_NAME%"
IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Error: Failed to create service. Please ensure you are running as administrator.
    pause
    exit /b %ERRORLEVEL%
)

sc description "%SERVICE_NAME%" "%SERVICE_DESCRIPTION%"

sc start "%SERVICE_NAME%"
IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Warning: Service has been installed but cannot be started.
    echo Please check the error logs in Windows Event Viewer.
    pause
) ELSE (
    echo.
    echo Service '%SERVICE_DISPLAY_NAME%' has been successfully installed and started.
)

pause 