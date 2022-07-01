package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool

	region string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkFlagRequired("region")
}

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "This is a cobra example",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
