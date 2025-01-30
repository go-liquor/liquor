package commands

import (
	"github.com/go-liquor/liquor/commands/app"
	"github.com/spf13/cobra"
)

func AppCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "app",
		Short: "app commands",
	}

	cmd.AddCommand(
		app.CreateApp(),
		app.EnableModuleCmd(),
	)

	return &cmd
}
