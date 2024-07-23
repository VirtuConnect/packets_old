@echo off
setlocal enabledelayedexpansion

:: Check if the correct number of arguments are provided
if "%~1"=="" (
    echo "Error: No commit message provided."
    echo "Usage: deploy.bat \"commit message\" tag"
    exit /b 1
)
if "%~2"=="" (
    echo "Error: No tag provided."
    echo "Usage: deploy.bat \"commit message\" tag"
    exit /b 1
)

:: Store the arguments in variables
set "commit_message=%~1"
set "tag=%~2"

:: Stage all changes
git add *

:: Commit with the provided message
git commit -m "%commit_message%"

:: Push the changes
git push

:: Tag the commit with the provided tag
git tag "%tag%"

:: Push the tags
git push --tags

:: Navigate to the backend directory and update the Go package
cd ../backend
go get github.com/VirtuConnect/packets@%tag%

:: Navigate to the client directory and update the Go package
cd ../client
go get github.com/VirtuConnect/packets@%tag%

:: Return to the initial directory
cd ../packets