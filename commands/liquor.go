package commands

import "github.com/spf13/cobra"

var LiquorCmd = &cobra.Command{
	Use:   "liquor",
	Short: "Liquor CLI commands",
	Long:  "Liquor is a framework to web development backend with golang",
}

func ExecuteLiquor() error {
	LiquorCmd.AddCommand(AppCommand())
	if err := LiquorCmd.Execute(); err != nil {
		return err
	}
	return nil
}
