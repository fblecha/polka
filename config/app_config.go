package config

import (
	"fmt"
)

type AppConfig struct {
	AppDir string
	S3     string
}

func (p *AppConfig) Save() {

}

func (p *AppConfig) Exists() bool {
	return false
}

func CreateAppConfigAsNeeded(appDir string) (AppConfig, error) {
	var config AppConfig
	config.AppDir = fmt.Sprintf("%s/config/%s", appDir, "app.json")
	config.S3 = ""
	if !config.Exists() {
		config.Save()
	} else {
		//load the old file  //TODO
	}
	return config, nil
}
