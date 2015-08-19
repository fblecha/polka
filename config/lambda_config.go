package config


import (
//	"log"
	"os"
	"fmt"
  "github.com/PolkaBand/polka/utils"
	"encoding/json"
	//"io/ioutil"
	//"path/filepath"
)

type LambdaConfig struct {
	Name 			string
	RAM	 			int //in MB
	Timeout		int //in seconds
	AWSRegion string
}

//FindAll finds all the lambda configs that exist under the ./config directory.
//func FindAll() ([]LambdaConfig, error) {
//then we're in the app root dir, so config directory tree is under ./config
//find all lambda_config.json files -- these should marshall directly into a

//}

//Create a lambda_config.json file that can be used as part of the input into AWS.
//conceptName should be something like "todo" or anything else anything else that might exist under app/endpoint/{conceptName}
//dir should be ??
func Create(conceptName string, dir string) error {
  //I believe we should be in the app root directory
  appRootDir, err := utils.AreWeInAppRootDir()
  if err != nil {
    return err
  }
  //so we can create the config for conceptName as
  //./config/{dir}/{conceptName}_lambdaconfig.json
  configDir := fmt.Sprintf("%v/config/%v", appRootDir, dir)
  filename := fmt.Sprintf("%v/%v_lambdaconfig.json", configDir, conceptName )
  if err := os.MkdirAll(configDir, 0777); err != nil {
    return err
  }
  //we made configDir and all parents as needed
  configFile, err := os.Create(filename)
  if err != nil {
    return err
  }
  //TODO move default config ?
  defaultConfig := LambdaConfig {
    Name: "blah",  //TODO this needs to be a better name
    RAM: 1024, //TODO make this configurable
    Timeout: 10,
    AWSRegion: "us-east-1", //TODO make this configurable
  }

  b, err := json.MarshalIndent(defaultConfig, "", "  ")
  if err != nil {
    return err
  }
  //else we're good
  configFile.Write(b)

  return nil
}
