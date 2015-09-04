package config

import (
	"github.com/PolkaBand/polka/fileutils"
	"fmt"
)

type ProjectConfig struct {
	ProjectDir string
  S3 string
}

func (p *ProjectConfig) Save() {
	if err := fileutils.Save(p.ProjectDir, "app.json", p, true  ); err != nil {
		fmt.Println(err)
	}
}

func (p *ProjectConfig) Exists() bool {
  return false
}

func CreateProjectConfigAsNeeded(appDir string) (ProjectConfig, error) {
	var config ProjectConfig
  config.ProjectDir = fmt.Sprintf("%s/config", appDir)
	fmt.Printf("app.json stored in %s \n", config.ProjectDir)
  config.S3 = ""
	if !config.Exists() {
		config.Save()
	} else {
		//load the old file  //TODO
	}
	return config, nil
}
