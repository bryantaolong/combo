package todo

import (
	"fmt"
	"strconv"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "删除指定任务",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("任务 ID 必须是数字")
		}
		data, _ := storage.Load()
		tasks, exists := data.Groups[groupName]
		if !exists {
			fmt.Println(color.RedString("分组 %s 不存在", groupName))
			return nil
		}
		newTasks := []storage.Task{}
		found := false
		for _, t := range tasks {
			if t.ID == id {
				found = true
				continue
			}
			newTasks = append(newTasks, t)
		}
		if !found {
			fmt.Println(color.RedString("任务 %d 不存在", id))
			return nil
		}
		data.Groups[groupName] = newTasks
		_ = storage.Save(data)
		fmt.Println(color.GreenString("任务 %d 已删除 (组: %s)", id, groupName))
		return nil
	},
}

func init() {
	TodoCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&groupName, "group", "g", "default", "指定分组")
}
