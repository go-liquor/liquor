package create

import (
	"fmt"
	"path"

	"github.com/charmbracelet/huh"
	"github.com/go-liquor/liquor/v3/app/adapters/database"
	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/gomod"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"github.com/go-liquor/liquor/v3/pkg/maps"
	"github.com/spf13/cobra"
)

func createRepository() *cobra.Command {
	var (
		name   string
		driver string
	)
	cm := &cobra.Command{
		Use:     "repository",
		Aliases: []string{"repo"},
		Short:   "Create a repository files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				name = args[0]
			}
			if name == "" || driver == "" {
				fields := []huh.Field{}
				if name == "" {
					fields = append(fields, huh.NewInput().Title("Set repository name").
						Value(&name))
				}
				if driver == "" {

					values := []huh.Option[string]{}
					for driverName, _ := range database.Drivers {
						values = append(values, huh.NewOption(driverName, driverName))
					}

					fields = append(fields, huh.NewSelect[string]().
						Options(values...).
						Title("Choice the database driver").
						Value(&driver))
				}
				if err := huh.NewForm(
					huh.NewGroup(
						fields...,
					),
				).Run(); err != nil {
					return err
				}
			}

			if _, ok := database.Drivers[driver]; !ok {
				return fmt.Errorf("driver needs is %v", maps.GetKeys(database.Drivers))
			}

			module, err := gomod.GetModPath(rootPath)
			if err != nil {
				return err
			}

			var values = map[string]string{
				"repositoryName": lqstring.ToPascalCase(name),
				"module":         module,
			}

			filename := lqstring.ToSnakeCase(name) + "_repository.go"
			portFileDst := path.Join(rootPath, "internal", "ports", filename)
			dbFileDst := path.Join(rootPath, "internal", "adapters", "db", filename)

			if err := templates.ParseTemplate(boilerplate.RepositoryPortFile, portFileDst, values); err != nil {
				return fmt.Errorf("failed to create %s: %v", portFileDst, err)
			}
			stdout.Success("Created %s", portFileDst)

			content, err := boilerplate.RepositoryImplFiles.ReadFile(fmt.Sprintf("repository/%v.go.tpl", driver))
			if err != nil {
				return fmt.Errorf("file driver repository not found: %v", err)
			}
			if err := templates.ParseTemplate(string(content), dbFileDst, values); err != nil {
				return fmt.Errorf("failed to create %s: %v", dbFileDst, err)
			}
			stdout.Success("Created %s", dbFileDst)

			if err := createModelRun(name); err != nil {
				return fmt.Errorf("failed to create model: %v", err)
			}

			return nil
		},
	}
	cm.Flags().StringVarP(&name, "name", "n", "", "Repository name")
	cm.Flags().StringVarP(&driver, "driver", "d", "", "Database driver")
	return cm
}
