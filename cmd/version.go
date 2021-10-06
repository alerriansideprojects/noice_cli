package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Noice CLI tool",
	Long:  `This command can be used get the version number of noice CLI tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("noice v0.0.1-alpha")
	},
}
