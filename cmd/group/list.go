package group

import (
	"fmt"

	"github.com/bryantaolong/combo/storage"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有分组",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := storage.Load()
		if err != nil {
			return err
		}
		if len(data.Groups) == 0 {
			fmt.Println(color.YellowString("暂无分组"))
			return nil
		}
		fmt.Println(color.CyanString("现有分组："))
		for name := range data.Groups {
			fmt.Println(" -", color.BlueString(name))
		}
		return nil
	},
}

func init() {
	GroupCmd.AddCommand(groupListCmd)
}
