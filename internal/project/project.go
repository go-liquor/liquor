// Package project provides functionality for managing Liquor project configuration.
package project

import (
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// ProjectFile is the name of the configuration file used to store project settings.
const (
	ProjectFile = ".liquor.yaml"
)

// Project represents the configuration structure for a Liquor project.
// It contains settings that define how the project should behave.
type Project struct {
	// DatabaseDriver specifies the type of database driver to be used in the project
	DatabaseDriver string `yaml:"databaseDriver"`
}

// GetProject reads the project configuration from the ProjectFile and returns
// a new Project instance. If the configuration file doesn't exist, it creates
// a new one with default values.
//
// Returns:
//   - *Project: A pointer to the project configuration
func GetProject() *Project {
	var project Project
	vp := viper.New()
	vp.SetConfigFile(ProjectFile)
	if _, err := os.Stat(ProjectFile); os.IsNotExist(err) {
		vp.Set("databaseDriver", "")
		vp.WriteConfig()
		return &Project{}
	}
	vp.ReadInConfig()
	vp.Unmarshal(&project)
	return &project
}

// UpdateProject writes the provided project configuration to the ProjectFile.
// It serializes the project struct to YAML format.
//
// Parameters:
//   - p: A pointer to the Project struct containing the configuration to be saved
func UpdateProject(p *Project) {
	content, _ := yaml.Marshal(&p)
	os.WriteFile(ProjectFile, content, 0755)
}
