@echo off
setlocal enabledelayedexpansion
rem cd %~dp0

set exercise=a b c d e f

(for %%a in (%exercise%) do (
    echo %%a
    mkdir %%a
    copy ..\boilerplate\a.go %%a\%%a.go
    copy ..\boilerplate\a_test.go %%a\%%a_test.go
))
