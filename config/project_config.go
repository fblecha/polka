package config

import (
	"github.com/PolkaBand/polka/fileutils"
	"fmt"
	"log"
	"os"
	"encoding/json"
)

type ProjectConfig struct {
	ProjectDir string
  S3 string
}

func (p *ProjectConfig) Save(overwrite bool) {
	if err := fileutils.Save(p.ProjectDir, "app.json", p  ); err != nil {
		fmt.Println(err)
	}
}

func (p *ProjectConfig) Exists() bool {
  return false
}

func CreateProjectConfigAsNeeded(appDir string) (ProjectConfig, error) {
	var config ProjectConfig
  config.ProjectDir = fmt.Sprintf("%s/config", appDir)
  config.S3 = ""
	if !config.Exists() {
		log.Println("created new app.json")
		config.Save(true)
	} else {
		//load the old file
		log.Println("loaded existing app.json")
		abspath := fmt.Sprintf("%s/app.json", config.ProjectDir)
		configFile, err := os.Open(abspath)
		if err != nil {
				return config, err
		}
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&config); err != nil {
			return config, err
		}

	}
	return config, nil
}
