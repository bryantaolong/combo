package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示 combo 的版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("combo CLI %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
