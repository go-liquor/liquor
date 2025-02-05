package commands

import (
	"github.com/go-liquor/liquor/internal/commons"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run app with go run",
	RunE: func(cmd *cobra.Command, args []string) error {
		return commons.Command(".", "go", "run", "cmd/app/main.go")
	},
}
