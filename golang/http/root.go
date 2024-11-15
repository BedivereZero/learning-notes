package main

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short:        "http",
		SilenceUsage: true,
	}

	cmd.AddCommand(NewServerCommand())
	cmd.AddCommand(NewClientCommand())

	return cmd
}
