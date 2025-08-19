@echo off
REM ==========================
REM Combo CLI 测试脚本 (CMD)
REM ==========================

set COMBO=combo.exe

echo ======================= Combo CLI 测试开始 =======================

REM 1. 查看帮助
echo.
echo --- 根命令帮助 ---
%COMBO% --help

REM 2. 创建分组
echo.
echo --- 创建分组 work ---
%COMBO% group add work

REM 3. 创建默认分组任务
echo.
echo --- 添加任务到默认分组 ---
%COMBO% todo add "默认任务1"
%COMBO% todo add "默认任务2"

REM 4. 创建工作分组任务
echo.
echo --- 添加任务到 work 分组 ---
%COMBO% todo add "写日报" -g work
%COMBO% todo add "开会准备" -g work

REM 5. 列出所有分组
echo.
echo --- 列出所有分组 ---
%COMBO% group list

REM 6. 列出默认分组任务
echo.
echo --- 列出默认分组任务 ---
%COMBO% todo list

REM 7. 列出 work 分组任务
echo.
echo --- 列出 work 分组任务 ---
%COMBO% todo list -g work

REM 8. 标记任务完成
echo.
echo --- 标记默认分组任务 1 为已完成 ---
%COMBO% todo done 1

echo.
echo --- 标记 work 分组任务 2 为已完成 ---
%COMBO% todo done 2 -g work

REM 9. 列出任务查看状态
echo.
echo --- 查看任务完成状态 ---
%COMBO% todo list
%COMBO% todo list -g work

REM 10. 删除任务
echo.
echo --- 删除默认分组任务 2 ---
%COMBO% todo delete 2

echo.
echo --- 删除 work 分组任务 1 ---
%COMBO% todo delete 1 -g work

REM 11. 清空分组任务
echo.
echo --- 清空 work 分组任务 ---
%COMBO% todo clear -g work

REM 12. 删除分组
echo.
echo --- 删除 work 分组 ---
%COMBO% group delete work

echo.
echo ======================= Combo CLI 测试结束 =======================
pause
