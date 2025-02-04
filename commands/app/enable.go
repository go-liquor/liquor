package app

import (
	"fmt"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/spf13/cobra"
)

func enableModuleCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "enable <module>",
		Short: "Enable modules",
		RunE:  enableModuleRun,
	}

	return &cmd
}

func enableModuleRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need specify a module")
	}
	for _, module := range args {
		if err := commons.Command(".", "go", "get", "github.com/go-liquor/liquor-sdk/modules/"+module); err != nil {
			return err
		}
	}
	return nil
}
