package hidden_home

import(
  "os"
  "encoding/json"
  "fmt"
  "strings"
)

//DoesPolkaHomeDirExist returns true or false based on the existance of the ~/.polka directory
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

//CreatePolkaHomeDir creates ~/.basename
func Create(basename string) error {
  //TODO see if homedir starts with "~/."
  var absoluteDirName string
  if strings.HasPrefix(basename, "~/.") {
    absoluteDirName = basename
  } else {
    absoluteDirName = fmt.Sprintf("~/.%v", basename)
  }
	return os.Mkdir(absoluteDirName, 0777)
}

//Save will save the json output of config into the file named ~/.basename/configleFilename
func Save(basename string, configFilename, config interface{} ) error {
  if !Exists(basename) {
    if err := Create(basename); err != nil {
      return fmt.Errorf("Unable to create ~/.%v ", basename)
    }
  }

  fullConfigFilename := fmt.Sprintf("%v/%v", basename, configFilename)

  configFile, err := os.Create(fullConfigFilename)
  if err != nil {
    return err
  }

  if b, err := json.MarshalIndent(config, "", "  "); err == nil {
    //else we're good
    configFile.Write(b)
    return nil
  } else {
    return err
  }

}
