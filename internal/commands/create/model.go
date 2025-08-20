package create

import (
	"path"

	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"github.com/spf13/cobra"
)

func createModel() *cobra.Command {
	var (
		name string
	)
	cm := &cobra.Command{
		Use:   "model",
		Short: "create a new model",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				name = args[0]
			}
			return createModelRun(name)
		},
	}
	cm.Flags().StringVarP(&name, "name", "n", "", "Model name")
	return cm
}

func createModelRun(name string) error {
	dst := path.Join(rootPath, "internal", "adapters", "db", lqstring.ToSnakeCase(name)+"_model.go")

	modelName := lqstring.ToPascalCase(name)
	if err := templates.ParseTemplate(boilerplate.ModelFile, dst, map[string]string{
		"modelName": modelName,
	}); err != nil {
		return err
	}
	stdout.Success("created %s", dst)
	return nil
}
