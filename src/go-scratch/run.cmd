@echo off
setlocal

:: Synchronize the root directory before deferring control back to gclient.py.
:: call "%~dp0\update_depot_tools.bat" %*

set GOPATH=%~dp0src;%GOPATH%
:: Defer control.
go run "%~dp0src\main.go"
pause