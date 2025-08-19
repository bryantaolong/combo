package todo

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupName string

var addCmd = &cobra.Command{
	Use:   "add [content]",
	Short: "添加一个任务",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := args[0]
		data, _ := storage.Load()
		tasks := data.Groups[groupName]
		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}
		tasks = append(tasks, storage.Task{ID: newID, Content: content, Done: false})
		data.Groups[groupName] = tasks
		_ = storage.Save(data)
		fmt.Println(color.GreenString("任务添加成功: %s (组: %s)", content, groupName))
		return nil
	},
}

func init() {
	TodoCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&groupName, "group", "g", "default", "指定分组")
}
