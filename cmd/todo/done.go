package todo

import (
	"fmt"
	"strconv"

	"github.com/bryantaolong/combo/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "标记任务为已完成",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		data, _ := storage.Load()
		tasks, exists := data.Groups[groupName]
		if !exists {
			fmt.Println(color.RedString("分组 %s 不存在", groupName))
			return nil
		}
		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				if tasks[i].Done {
					fmt.Println(color.YellowString("任务 %d 已经完成", id))
					return nil
				}
				tasks[i].Done = true
				found = true
				break
			}
		}
		if !found {
			fmt.Println(color.RedString("任务 %d 不存在", id))
			return nil
		}
		data.Groups[groupName] = tasks
		_ = storage.Save(data)
		fmt.Println(color.GreenString("任务 %d 已标记完成 (组: %s)", id, groupName))
		return nil
	},
}

func init() {
	TodoCmd.AddCommand(doneCmd)
	doneCmd.Flags().StringVarP(&groupName, "group", "g", "default", "指定分组")
}
