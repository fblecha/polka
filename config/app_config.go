package config

import (
	"fmt"
)

type ProjectConfig struct {
	ProjectDir string
  S3 string
}

func (p *ProjectConfig) Save() {

}

func (p *ProjectConfig) Exists() bool {
  return false
}

func CreateProjectConfigAsNeeded(appDir string) (ProjectConfig, error) {
	var config ProjectConfig
  config.ProjectDir = fmt.Sprintf("%s/config/%s", appDir, "app.json")
	fmt.Printf("app.json stored in %s \n", config.ProjectDir)
  config.S3 = ""
	if !config.Exists() {
		config.Save()
	} else {
		//load the old file  //TODO
	}
	return config, nil
}
