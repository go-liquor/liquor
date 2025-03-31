package create

import (
	"github.com/go-liquor/liquor/v2/internal/boilerplate"
	"github.com/go-liquor/liquor/v2/internal/stdout"
	"github.com/go-liquor/liquor/v2/internal/templates"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/spf13/cobra"
	"path"
)

var createPortCmd = &cobra.Command{
	Use:   "port",
	Short: "Create a Port",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		files := map[string]string{
			path.Join("app/ports", lqstring.ToSnakeCase(name)+".go"): boilerplate.Ports,
		}

		if err := templates.ParseTemplates(files, map[string]any{
			"PascalCaseName": lqstring.ToPascalCase(name),
		}); err != nil {
			return err
		}

		stdout.Success("port created successfully")
		return nil
	},
}

func init() {
	createPortCmd.Flags().StringP("name", "n", "", "Name of the Port")
	createPortCmd.MarkFlagRequired("name")
	CreateCmd.AddCommand(createPortCmd)
}
