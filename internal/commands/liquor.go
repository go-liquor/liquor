package commands

import (
	"github.com/go-liquor/liquor/v2/internal/commands/create"
	"github.com/spf13/cobra"
)

var LiquorRootCmd = &cobra.Command{
	Use:   "liquor",
	Short: "Liquor CLI",
}

func Execute() error {
	LiquorRootCmd.AddCommand(
		create.CreateCmd,
	)
	if err := LiquorRootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
