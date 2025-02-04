package app

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/constants"
	"github.com/go-liquor/liquor/internal/message"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var createApp = &cobra.Command{
	Use:   "create",
	Short: "Create a new app",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		var pkg, _ = cmd.Flags().GetString("pkg")

		if commons.IsExist(name) {
			return fmt.Errorf("%v exists", name)
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		liquorHomeDir := path.Join(home, ".liquor")
		frameworkDir := path.Join(liquorHomeDir, "framework")
		if commons.IsNotExist(frameworkDir) {
			os.MkdirAll(liquorHomeDir, 0755)
			if _, err := git.PlainClone(frameworkDir, false, &git.CloneOptions{
				URL:      constants.FrameworkRepo,
				Progress: os.Stdout,
			}); err != nil {
				return err
			}
		}

		if err := cp.Copy(frameworkDir, name); err != nil {
			return err
		}
		message.Success("create %v", name)
		os.RemoveAll(path.Join(name, ".git"))

		if err := commons.ReplacePackage(name, "github.com/go-liquor/framework", pkg); err != nil {
			return err
		}

		mod, err := commons.GetModFile(name)
		if err != nil {
			return fmt.Errorf("failed to get mod file: %v", err)
		}

		mod.DropRequire("github.com/go-liquor/liquor-sdk")
		mod.AddRequire("github.com/go-liquor/liquor-sdk", constants.SdkVersion)
		mod.DropReplace("github.com/go-liquor/liquor-sdk", "")
		content, err := mod.Format()
		if err != nil {
			return fmt.Errorf("failed in format mod file: %v", err)
		}
		if err := os.WriteFile(path.Join(name, "go.mod"), content, 0755); err != nil {
			return fmt.Errorf("failed to recreate go.mod: %v", err)
		}
		commons.Command(name, "go", "mod", "tidy")
		message.Success("finish")
		return nil
	},
}

func init() {
	createApp.Flags().StringP("name", "n", "", "App name")
	createApp.Flags().StringP("pkg", "p", "", "Go module package")
	createApp.MarkFlagRequired("name")
	createApp.MarkFlagRequired("pkg")
}
