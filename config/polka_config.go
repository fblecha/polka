package config

import (
	"fmt"
	"github.com/fblecha/homeplate"
	"os"
	"encoding/json"
)

type PolkaConfig struct {
	Name               string
	NumberOfExecutions int
}

func (p *PolkaConfig) Save() {
	homeDir := "polka"
	homeplate.Save(homeDir, "polka", p, false)
}

func (p *PolkaConfig) Exists() bool {
	return false //TODO
}

//CreatePolkaHomeConfig creates ~/.polka if it doesn't exist already.
//It then places the contents of PolkaConfig into it as ~/.polka/polka_config.json
func CreatePolkaHomeConfigAsNeeded() (PolkaConfig, error) {
	var config PolkaConfig

	abspath := fmt.Sprintf("%s/%s", homeplate.ExpandTilde("polka"), "polka.json")

	if !homeplate.Exists(abspath) {
		config := PolkaConfig{
			Name:               "polka",
			NumberOfExecutions: 0,
		}
		config.Save()
	} else {
		//load the old file
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
