package create

import (
	"path"

	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/gomod"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"github.com/spf13/cobra"
)

func createUseCase() *cobra.Command {
	var (
		name     string
		rootPath string
	)
	cm := &cobra.Command{
		Use:     "usecase",
		Aliases: []string{"uc"},
		Short:   "create a usecase",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				name = args[0]
			}

			module, err := gomod.GetModFile(rootPath)
			if err != nil {
				return err
			}
			useCaseName := lqstring.ToPascalCase(name)

			values := map[string]string{
				"module":      module.Module.Mod.Path,
				"useCaseName": useCaseName,
			}

			appDst := path.Join(rootPath, "internal/application", lqstring.ToSnakeCase(name)+"_usecase.go")
			if err := templates.ParseTemplate(boilerplate.UsecaseFile, appDst, values); err != nil {
				return err
			}
			stdout.Success("created %s", appDst)
			portDst := path.Join(rootPath, "internal/ports", lqstring.ToSnakeCase(name)+"_usecase.go")
			if err := templates.ParseTemplate(boilerplate.UsecasePortFile, portDst, values); err != nil {
				return err
			}
			stdout.Success("created %s", portDst)
			return nil
		},
	}
	cm.Flags().StringVarP(&name, "name", "n", "", "Usecase name")
	cm.Flags().StringVarP(&rootPath, "root", "r", ".", "Root path")
	return cm
}
