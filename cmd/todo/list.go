package todo

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := storage.Load()
		if err != nil {
			return err
		}
		tasks, exists := data.Groups[groupName]
		if !exists {
			fmt.Println(color.RedString("分组 %s 不存在", groupName))
			return nil
		}
		if len(tasks) == 0 {
			fmt.Println(color.YellowString("分组 %s 暂无任务", groupName))
			return nil
		}
		fmt.Println(color.CyanString("分组 %s 的任务：", groupName))
		for _, t := range tasks {
			if t.Done {
				fmt.Printf("[%d] %s\n", t.ID, color.GreenString(t.Content))
			} else {
				fmt.Printf("[%d] %s\n", t.ID, t.Content)
			}
		}
		return nil
	},
}

func init() {
	TodoCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&groupName, "group", "g", "default", "指定分组")
}
