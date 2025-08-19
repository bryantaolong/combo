package todo

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "清空指定分组的任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, _ := storage.Load()
		if _, exists := data.Groups[groupName]; !exists {
			fmt.Println(color.RedString("分组 %s 不存在", groupName))
			return nil
		}
		data.Groups[groupName] = []storage.Task{}
		_ = storage.Save(data)
		fmt.Println(color.GreenString("分组 %s 已清空任务", groupName))
		return nil
	},
}

func init() {
	TodoCmd.AddCommand(clearCmd)
	clearCmd.Flags().StringVarP(&groupName, "group", "g", "default", "指定分组")
}
