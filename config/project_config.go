package config

import (
	"github.com/PolkaBand/polka/utils"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

type ProjectConfig struct {
  S3 string
}

func LoadProjectConfig() (ProjectConfig, error) {
	var config ProjectConfig
	if appDir, err := utils.AreWeInAppRootDir(); err == nil {
		//attempt to load the old file

		log.Println("loaded existing app.json")
		abspath := fmt.Sprintf("%s/config/app.json", appDir)
		configFile, err := os.Open(abspath)
		defer configFile.Close()
		if err == nil {
			jsonParser := json.NewDecoder(configFile)
			if err = jsonParser.Decode(&config); err == nil {
				return config, err
			} else {
				return config, err
			}
		} else {
			//ok, that failed, let's make a new one and return it.  Not that it is **not saved**.
			return config, nil
		}
	} else {
		//ok, that failed, let's make a new one and return it.  Not that it is **not saved**.
		return config, err
	}
}

//Save this project config in config/app.json.  Always overwrites the previous file.
func (p *ProjectConfig) Save() {
	if appDir, err := utils.AreWeInAppRootDir(); err == nil {
		fileName := fmt.Sprintf("%s/config/app.json", appDir )
		if b, err := json.MarshalIndent(p, "", "  "); err == nil {

			if err := ioutil.WriteFile(fileName, b, 0644); err != nil {
				log.Println(err)
			}
		}
	}
}
