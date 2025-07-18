@echo off
REM --- This script queries the status and configuration of the Windows service. ---
REM --- Run as administrator if needed. ---

SET "SERVICE_NAME=your-app"

ECHO ===============================================
ECHO Querying service status for: %SERVICE_NAME%
ECHO ===============================================
sc query "%SERVICE_NAME%"

ECHO.
ECHO ===============================================
ECHO Querying service configuration for: %SERVICE_NAME%
ECHO ===============================================
sc qc "%SERVICE_NAME%"

pause
