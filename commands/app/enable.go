package app

import (
	"fmt"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/spf13/cobra"
)

var enableModuleCmd = &cobra.Command{
	Use:   "enable <module>",
	Short: "Enable modules",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you need specify a module")
		}
		for _, module := range args {
			if err := commons.Command(".", "go", "get", "github.com/go-liquor/liquor-sdk/modules/"+module); err != nil {
				return err
			}
		}
		return nil
	},
}
