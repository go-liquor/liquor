package create

import (
	"fmt"
	"os"
	"path"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

var createEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Create a new entity",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		return createEntityRun(name)
	},
}

func createEntityRun(name string) error {

	var filename string = commons.ToFilename(name)
	var entityPath string = "internal/domain/entity"

	if commons.IsNotExist(entityPath) {
		os.MkdirAll(entityPath, 0755)
	}

	var outputFile string = path.Join(entityPath, filename)

	if commons.IsExist(outputFile) {
		return fmt.Errorf("file %v already exists", outputFile)
	}

	if err := templates.ParseTemplate(templates.Entity, outputFile, map[string]string{
		"PascalCaseName": textcase.PascalCase(name),
	}); err != nil {
		return err
	}
	message.Success("created %v", outputFile)

	return nil
}

func init() {
	createEntityCmd.Flags().StringP("name", "n", "", "Model name")
	createEntityCmd.MarkFlagRequired("name")
}
