//Package homeplate provides utilities for dealing with ~/.dirname (hidden configuration directories).
package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
  //"os/user"
	//"strings"
)

//Exists returns true or false based on the existance of pathname
func Exists(pathname string) bool {
	_, err := os.Stat(pathname)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return true
}


//Save will save the json marshal output of config into the file named dirname/filename.json
//If overwrite is false, and the file exists, an error will be returned
func Save(pathname string, filename, config interface{}, overwrite bool) error {
	fullFilename := fmt.Sprintf("%s/%s", pathname, filename)

	if overwrite {
		//overwrite == true then we write it either way
    return createAndWrite(fullFilename, config)
	} else {
		//overwrite == false
		if Exists(fullFilename) {
			//overwrite == false && Exists == true
      return fmt.Errorf("Unable to create %s as it exists and did not set overwrite to true", fullFilename)
		} else {
      return createAndWrite(fullFilename, config)
		}
	}
}


func createAndWrite(absfilename string, config interface{} ) error  {
	fmt.Println("blah")
  configFile, err := os.Open(absfilename)
  if err != nil {
    configFile, err = os.Create(absfilename)
		if err != nil {
			return err
		}
  }
  //save json to file
	fmt.Printf("writing file %v", configFile )
  if b, err := json.MarshalIndent(config, "", "  "); err == nil {
    configFile.Write(b)
    return nil
  } else {
    return err
  }
}
