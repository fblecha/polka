package main

import (
	"log"
	"os"
	"fmt"
	"errors"
	//"path/filepath"
	"github.com/kardianos/osext"
	"github.com/mitchellh/cli"
)

func main() {


	templateDir, err := FindTemplateDir()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("template dir = %v\n", templateDir)

	args := os.Args[1:]

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("polka"),
	}

	log.Printf("topflags = %v \n", cli.Args)

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func TestTemplateDir(dirName string) bool  {
	log.Printf("looking in %v\n", dirName)
	_, err := os.Stat(dirName);
	return err == nil  //if err == true, we found the dir; it's false otherwise
}

func FindTemplateDir() (string, error) {
	if folderPath, err := osext.ExecutableFolder(); err == nil {
		log.Println(folderPath)

		devTemplateDirName := fmt.Sprintf("%v/templates", folderPath)
		brewTemplateDirName := fmt.Sprintf("%v/../share/templates", folderPath)

		//maybe we're running this in dev mode, so templates could be at folderPath/templates
		if TestTemplateDir(devTemplateDirName) {
			return devTemplateDirName, nil
		} else if TestTemplateDir(brewTemplateDirName) {
			return brewTemplateDirName, nil
		}
	}
	//we only reach this part if err != nil
	return "", errors.New("unable to find templates dir")

}
