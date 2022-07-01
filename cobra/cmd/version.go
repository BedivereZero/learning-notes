package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	format string
)

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().StringVarP(&format, "format", "f", "", "format of output")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sample v0.1.0")
	},
	Args: cobra.NoArgs,
}
