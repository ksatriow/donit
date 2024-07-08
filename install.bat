@echo off

rem Build the binary
echo Building the Go project...
go build -o donit.exe

rem Check if build was successful
if %errorlevel% neq 0 (
    echo Error: Build failed
    exit /b 1
)

rem Move the binary to a directory in PATH
echo Installing the binary...
move donit.exe [PATH-OF-DONIT]\donit.exe

rem Check if move was successful
if %errorlevel% neq 0 (
    echo Error: Installation failed
    exit /b 1
)

echo Installation complete.
exit /b 0
