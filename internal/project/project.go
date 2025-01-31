package project

import (
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const (
	ProjectFile = ".liquor.yaml"
)

type Project struct {
	DatabaseDriver string `yaml:"databaseDriver"`
}

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

func UpdateProject(p *Project) {
	content, _ := yaml.Marshal(&p)
	os.WriteFile(ProjectFile, content, 0755)
}
