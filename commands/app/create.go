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

func CreateApp() *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new app",
		RunE:  createAppRun,
	}
	cmd.Flags().StringP("name", "n", "", "App name")
	cmd.Flags().StringP("pkg", "p", "", "Go module package")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("pkg")
	return &cmd
}

func createAppRun(cmd *cobra.Command, args []string) error {
	var name, _ = cmd.Flags().GetString("name")
	var pkg, _ = cmd.Flags().GetString("pkg")

	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("%v exists", name)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	liquorHomeDir := path.Join(home, ".liquor")
	frameworkDir := path.Join(liquorHomeDir, "framework")
	if _, err := os.Stat(liquorHomeDir); os.IsNotExist(err) {
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

	message.Success("finish")

	return nil
}
