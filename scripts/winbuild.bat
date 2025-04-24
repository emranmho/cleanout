@echo off
set APP_NAME=cleanout
cd ..
go build -o %APP_NAME%.exe
echo Build complete: %APP_NAME%.exe
