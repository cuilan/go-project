@echo off
REM Windows 批处理文件，用于简化 Makefile.windows 的使用
REM 使用方法: make-windows.bat [target]

setlocal enabledelayedexpansion

REM 检查是否提供了目标参数
if "%1"=="" (
    echo Usage: make-windows.bat [target]
    echo.
    echo Examples:
    echo   make-windows.bat help
    echo   make-windows.bat build
    echo   make-windows.bat install-service
    echo.
    echo Available targets:
    echo   all              Build all Windows platform binaries
    echo   build            Build current Windows platform binary
    echo   build-all        Build all Windows platform binaries
    echo   dist             Build and package as ZIP distribution
    echo   lint             Run all code checkers
    echo   test             Run all unit tests
    echo   test-cover       Run tests and generate HTML coverage report
    echo   mod-vendor       Update vendor directory
    echo   mod-tidy         Tidy go.mod file
    echo   mod-download     Download modules to local cache
    echo   version          Show version information
    echo   run              Run application
    echo   run-bin          Run executable file
    echo   install-service  Install Windows service
    echo   uninstall-service Uninstall Windows service
    echo   start-service    Start Windows service
    echo   stop-service     Stop Windows service
    echo   restart-service  Restart Windows service
    echo   clean            Clean build artifacts and temporary files
    echo   install-tools    Install code checking development tools
    echo   help             Show this help information
    exit /b 1
)

REM 检查 make 命令是否可用
where make >nul 2>&1
if %errorlevel% neq 0 (
    echo Error: 'make' command not found.
    echo Please install GNU Make for Windows from:
    echo http://gnuwin32.sourceforge.net/packages/make.htm
    echo.
    echo After installation, make sure 'make' is in your PATH.
    exit /b 1
)

REM 检查 Makefile.windows 是否存在
if not exist "Makefile.windows" (
    echo Error: Makefile.windows not found in current directory.
    echo Please run this script from the project root directory.
    exit /b 1
)

REM 执行 make 命令
echo Running: make -f Makefile.windows %*
make -f Makefile.windows %*

REM 检查执行结果
if %errorlevel% neq 0 (
    echo.
    echo Make command failed with exit code %errorlevel%
    echo Please check the error messages above.
    exit /b %errorlevel%
)

echo.
echo Make command completed successfully. 