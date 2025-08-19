package group

import "github.com/spf13/cobra"

var GroupCmd = &cobra.Command{
	Use:   "group",
	Short: "分组管理相关命令",
	Long:  "创建、列出、删除分组",
}
