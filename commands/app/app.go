package app

import (
	"github.com/spf13/cobra"
)

func AppCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "app",
		Short: "app commands",
	}

	cmd.AddCommand(
		createApp(),
		enableModuleCmd(),
	)

	return &cmd
}
