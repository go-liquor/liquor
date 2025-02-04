package app

import (
	"github.com/spf13/cobra"
)

var AppCommand = &cobra.Command{
	Use:   "app",
	Short: "app commands",
}

func init() {
	AppCommand.AddCommand(
		createApp,
		enableModuleCmd,
	)

}
