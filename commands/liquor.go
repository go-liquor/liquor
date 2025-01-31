package commands

import (
	"github.com/go-liquor/liquor/commands/app"
	"github.com/go-liquor/liquor/commands/create"
	"github.com/spf13/cobra"
)

var LiquorCmd = &cobra.Command{
	Use:   "liquor",
	Short: "Liquor CLI commands",
	Long:  "Liquor is a framework to web development backend with golang",
}

func ExecuteLiquor() error {
	LiquorCmd.AddCommand(
		app.AppCommand(),
		create.CreateCommand())
	if err := LiquorCmd.Execute(); err != nil {
		return err
	}
	return nil
}
