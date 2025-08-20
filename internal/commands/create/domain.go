package create

import (
	"path"

	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/go-liquor/liquor/v3/pkg/lqstring"
	"github.com/spf13/cobra"
)

func createDomain() *cobra.Command {
	var (
		name string
	)
	cm := &cobra.Command{
		Use:   "domain",
		Short: "create a new domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				name = args[0]
			}
			dst := path.Join(rootPath, "internal", "domain", lqstring.ToSnakeCase(name)+".go")

			domainName := lqstring.ToPascalCase(name)
			if err := templates.ParseTemplate(boilerplate.DomainFile, dst, map[string]string{
				"domainName": domainName,
			}); err != nil {
				return err
			}
			stdout.Success("created %s", dst)
			return nil
		},
	}
	cm.Flags().StringVarP(&name, "name", "n", "", "Domain name")
	return cm
}
