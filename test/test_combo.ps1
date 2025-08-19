# Windows PowerShell 测试脚本: test_combo.ps1
# 确保 combo.exe 已经构建好

$combo = ".\combo.exe"

Write-Host "====== Combo CLI 测试开始 ======" -ForegroundColor Cyan

# 1. 查看帮助
Write-Host "`n--- 根命令帮助 ---"
& $combo --help

# 2. 创建分组
Write-Host "`n--- 创建分组 work ---"
& $combo group add work

# 3. 创建默认分组任务
Write-Host "`n--- 添加任务到默认分组 ---"
& $combo todo add "默认任务1"
& $combo todo add "默认任务2"

# 4. 创建工作分组任务
Write-Host "`n--- 添加任务到 work 分组 ---"
& $combo todo add "写日报" -g work
& $combo todo add "开会准备" -g work

# 5. 列出所有分组
Write-Host "`n--- 列出所有分组 ---"
& $combo group list

# 6. 列出默认分组任务
Write-Host "`n--- 列出默认分组任务 ---"
& $combo todo list

# 7. 列出 work 分组任务
Write-Host "`n--- 列出 work 分组任务 ---"
& $combo todo list -g work

# 8. 标记任务完成
Write-Host "`n--- 标记默认分组任务 1 为已完成 ---"
& $combo todo done 1

Write-Host "`n--- 标记 work 分组任务 2 为已完成 ---"
& $combo todo done 2 -g work

# 9. 列出任务查看状态
Write-Host "`n--- 查看任务完成状态 ---"
& $combo todo list
& $combo todo list -g work

# 10. 删除任务
Write-Host "`n--- 删除默认分组任务 2 ---"
& $combo todo delete 2

Write-Host "`n--- 删除 work 分组任务 1 ---"
& $combo todo delete 1 -g work

# 11. 清空分组任务
Write-Host "`n--- 清空 work 分组任务 ---"
& $combo todo clear -g work

# 12. 删除分组
Write-Host "`n--- 删除 work 分组 ---"
& $combo group delete work

Write-Host "`n====== Combo CLI 测试结束 ======" -ForegroundColor Cyan
