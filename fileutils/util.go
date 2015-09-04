//Package homeplate provides utilities for dealing with ~/.dirname (hidden configuration directories).
package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
  "os/user"
	"strings"
)

//Exists returns true or false based on the existance of the ~/.dirname directory
func Exists(path string) bool {
	_, err := os.Stat(path)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return true
}
//ExpandTilde assumes that it was passed in basename (~/.basename) and needs to expand
//it (e.g. /Users/fb3/.basename )
func ExpandTilde(basename string) string {
  usr, _ := user.Current()
  absname := fmt.Sprintf("%v/.%v", usr.HomeDir, basename)
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
//If overwrite is false, and the file exists, an error will be returned
func Save(dirname string, configFilename, config interface{}, overwrite bool) error {
  //create dirname if it doesn't exist
	if err := Create(dirname); err != nil {
		return fmt.Errorf("Unable to create ~/.%v ", dirname)
	}
  homedir := ExpandTilde(dirname)
	fullConfigFilename := fmt.Sprintf("%v/%v.json", homedir, configFilename)

	if overwrite {
		//overwrite == true then we write it either way
    return createAndWrite(fullConfigFilename, config)
	} else {
		//overwrite == false
		if Exists(fullConfigFilename) {
			//overwrite == false && Exists == true
      return fmt.Errorf("Unable to create %s as it exists and did not set overwrite to true", fullConfigFilename)
		} else {
      return createAndWrite(fullConfigFilename, config)
		}
	}
}


func createAndWrite(absfilename string, config interface{} ) error  {
	fmt.Println("blah")
  configFile, err := os.Create(absfilename)
  if err != nil {
    return err
  }
  //save json to file
	fmt.Printf("writing file %v", configFile)
  if b, err := json.MarshalIndent(config, "", "  "); err == nil {
    configFile.Write(b)
    return nil
  } else {
    return err
  }
}
