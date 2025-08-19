package group

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "创建一个分组",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		data, _ := storage.Load()
		if _, exists := data.Groups[name]; exists {
			fmt.Println(color.RedString("分组 %s 已存在", name))
			return nil
		}
		data.Groups[name] = []storage.Task{}
		_ = storage.Save(data)
		fmt.Println(color.GreenString("分组 %s 创建成功", name))
		return nil
	},
}

func init() {
	GroupCmd.AddCommand(groupAddCmd)
}
