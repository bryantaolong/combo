package group

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "删除一个分组及其所有任务",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if name == "default" {
			fmt.Println(color.RedString("默认分组不能删除"))
			return nil
		}
		data, err := storage.Load()
		if err != nil {
			return err
		}
		if _, exists := data.Groups[name]; !exists {
			fmt.Println(color.RedString("分组 %s 不存在", name))
			return nil
		}
		delete(data.Groups, name)
		if err := storage.Save(data); err != nil {
			return err
		}
		fmt.Println(color.GreenString("分组 %s 已删除（含任务）", name))
		return nil
	},
}

func init() {
	GroupCmd.AddCommand(groupDeleteCmd)
}
