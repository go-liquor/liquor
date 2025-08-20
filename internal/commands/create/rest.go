package create

import (
	"fmt"
	"path"

	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"github.com/spf13/cobra"
)

func createRest() *cobra.Command {
	var (
		name  string
		group string
	)
	cm := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"api"},
		Short:   "Create a rest api",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				name = args[0]
			}

			var values = map[string]string{
				"restName": lqstring.ToPascalCase(name),
				"group":    group,
			}
			restFileName := path.Join(rootPath, "internal", "adapters", "rest", lqstring.ToSnakeCase(name)+".go")
			if err := templates.ParseTemplate(boilerplate.RestApiFile, restFileName, values); err != nil {
				return fmt.Errorf("failed to create %s: %v", restFileName, err)
			}
			stdout.Success("Created %s", restFileName)
			return nil
		},
	}
	cm.Flags().StringVarP(&name, "name", "n", "", "Rest api")
	cm.Flags().StringVarP(&group, "group", "g", "/", "Group Rest")
	return cm
}
