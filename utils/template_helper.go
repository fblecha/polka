package utils

import (
  "log"
  "fmt"
  "os"
  "errors"
  "github.com/kardianos/osext"
)
func TemplateDirExists(dirName string) bool  {
	log.Printf("looking in %v\n", dirName)
	_, err := os.Stat(dirName);
	return err == nil  //if err == true, we found the dir; it's false otherwise
}

// FindTemplateDir returns the absolute path to the templates directory as a string.
func FindTemplateDir() (string, error) {
	if folderPath, err := osext.ExecutableFolder(); err == nil {
		log.Println(folderPath)

		devTemplateDirName := fmt.Sprintf("%v/templates", folderPath)
		brewTemplateDirName := fmt.Sprintf("%v/../share/templates", folderPath)

		//maybe we're running this in dev mode, so templates could be at folderPath/templates
		if TemplateDirExists(devTemplateDirName) {
			return devTemplateDirName, nil
		} else if TemplateDirExists(brewTemplateDirName) {
			return brewTemplateDirName, nil
		}
	}
	//we only reach this part if err != nil
	return "", errors.New("unable to find templates dir")

}
