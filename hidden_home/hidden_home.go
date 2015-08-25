//Package hidden_home provides utilities for dealing with ~/.dirname hidden configuration directories.
package hidden_home

import (
	"encoding/json"
	"fmt"
	"os"
  "os/user"
	"strings"
)

//Exists returns true or false based on the existance of the ~/.dirname directory
func Exists(dirname string) bool {
	_, err := os.Stat(dirname)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return true
}
//ExpandTilde assumes that it was passed in basename and needs to expand it to /Users/fb3/.basename
func ExpandTilde(basename string) string {
  usr, _ := user.Current()
  homedir := usr.HomeDir
  absname := fmt.Sprintf("%v/.%v", homedir, basename)
  return absname
}

//Create creates ~/.dirname
func Create(dirname string) error {
  //if dirname starts with ~/. then pull that out
  basename := strings.Replace(dirname, "~/.", "", 1)
  absname := ExpandTilde(basename)

  if !Exists(absname) {
    return os.Mkdir(absname, 0777)
  }
  return nil
}

//Save will save the json marshal output of config into the file named ~/.dirname/configFilename.json
func Save(dirname string, configFilename, config interface{}) error {
  //create dirname if it doesn't exist
	if err := Create(dirname); err != nil {
		return fmt.Errorf("Unable to create ~/.%v ", dirname)
	}

  homedir := ExpandTilde(dirname)
	fullConfigFilename := fmt.Sprintf("%v/%v.json", homedir, configFilename)

	configFile, err := os.Create(fullConfigFilename)
	if err != nil {
		return err
	}

	if b, err := json.MarshalIndent(config, "", "  "); err == nil {
		configFile.Write(b)
		return nil
	} else {
		return err
	}
}
