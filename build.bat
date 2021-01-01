@echo off

set Task=%1
set CurrentDir=%~dp0
set OutputFilePath="%CurrentDir%build\comm_server.exe"

cd "%CurrentDir%src"
if "%Task%" == "-b" (
	go build -o "%OutputFilePath%"
	exit 0
)

if "%Task%" == "-br" (
	go build -o "%OutputFilePath%"
	cmd /K %OutputFilePath%
	exit 0
)

if "%Task%" == "-r" (
	cmd /K %OutputFilePath%
	exit 0
)