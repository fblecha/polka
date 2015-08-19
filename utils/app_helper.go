package utils

import (
	"os"
	"fmt"
)

//AreWeInAppRootDir returns nil if we're in the app root directory; otherwise it retuns the error.
func AreWeInAppRootDir() (string, error) {
  currentDir, err := os.Getwd();
  if err != nil {
    return "", err
  }
  //we are in the app root dir if we have both a ./app and
  //a ./config in the current working dir
  checkDirs := [...]string {
    fmt.Sprintf("%v/app", currentDir),
    fmt.Sprintf("%v/config", currentDir), //expand if needed
  }
  for _, checkDir := range checkDirs {
    if _, err := os.Stat(checkDir); err != nil {
      return "", err
    }
  }
  //if we made it here, all the checkDirs exist, which means we should be good
  return currentDir, nil
}
