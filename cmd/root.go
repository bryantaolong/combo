package cmd

import (
	"fmt"

	"github.com/bryantaolong/combo/cmd/group"
	"github.com/bryantaolong/combo/cmd/todo"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "combo",
	Short: "Combo 是一个命令行 To Do 应用，支持分组和颜色输出",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 combo! 输入 --help 查看命令")
	},
}

func Execute() {
	rootCmd.AddCommand(todo.TodoCmd)
	rootCmd.AddCommand(group.GroupCmd)
	rootCmd.AddCommand(versionCmd)
	cobra.CheckErr(rootCmd.Execute())
}
