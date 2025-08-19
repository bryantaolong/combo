package todo

import "github.com/spf13/cobra"

var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "任务管理相关命令",
	Long:  "添加、列出、完成、删除任务",
}
