package create

import (
	"fmt"
	"path"

	"github.com/go-liquor/liquor/v2/internal/boilerplate"
	"github.com/go-liquor/liquor/v2/internal/gomod"
	"github.com/go-liquor/liquor/v2/internal/stdout"
	"github.com/go-liquor/liquor/v2/internal/templates"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/spf13/cobra"
)

var createApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Create a new api",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")

		gm, err := gomod.GetModFile(".")
		if err != nil {
			return fmt.Errorf("failed to get go.mod file: %v", err)
		}

		files := map[string]string{
			path.Join("app/adapters/rest", lqstring.ToSnakeCase(name)+".go"): boilerplate.Api,
		}

		if err := templates.ParseTemplates(files, map[string]any{
			"PascalCaseName": lqstring.ToPascalCase(name),
			"Package":        gm.Module.Mod.Path,
		}); err != nil {
			return err
		}

		stdout.Success("api created successfully")
		return nil
	},
}

func init() {
	createApiCmd.Flags().StringP("name", "n", "", "Api name")
	createApiCmd.MarkFlagRequired("name")
	CreateCmd.AddCommand(createApiCmd)
}
