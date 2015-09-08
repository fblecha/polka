//Package homeplate provides utilities for dealing with ~/.dirname (hidden configuration directories).
package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
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
func Save(pathname string, filename, config interface{} ) error {
	absfilename := fmt.Sprintf("%s/%s", pathname, filename)
  configFile, err := os.Open(absfilename)
	defer configFile.Close()

  if err != nil {
		//if the file doesn't exist, try to open it for create
    configFile, err = os.Create(absfilename)
		log.Printf("createAndWrite writing %v \n", config)
		if err != nil {
			return err
		}
  }
  //save json to file
  if b, err := json.MarshalIndent(config, "", "  "); err == nil {
    configFile.Write(b)
		configFile.Sync()
		//log.Printf("wrote %v to \n file %s \n", config, absfilename )

    return nil
  } else {
    return err
  }
}
