package commands

import (
	"fmt"
	"os"

	"github.com/go-liquor/liquor/commands/app"
	"github.com/go-liquor/liquor/commands/create"
	"github.com/go-liquor/liquor/internal/constants"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

var LiquorCmd = &cobra.Command{
	Use:   "liquor",
	Short: "Liquor CLI commands",
	Long:  "Liquor is a framework to web development backend with golang",
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Printf("😍 CLI Version: %v\n🚀 SDK Version: %v\n", aurora.Cyan(constants.CliVersion), aurora.Cyan(constants.SdkVersion))
			os.Exit(0)
		}
	},
}

func init() {
	LiquorCmd.Flags().Bool("version", false, "Show version")
	LiquorCmd.AddCommand(
		app.AppCommand,
		create.CreateCmd,
		runCmd,
	)
}

func ExecuteLiquor() error {

	if err := LiquorCmd.Execute(); err != nil {
		return err
	}
	return nil
}
