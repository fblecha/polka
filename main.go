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

	fmt.Printf("template dir = %v\n", templateDir)

	args := os.Args[1:]

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("polka"),
	}

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func FindTemplateDir() (string, error) {
	fp, e := osext.ExecutableFolder()
	fmt.Printf("exec folder path = %v & err = %v\n", fp, e )

	if folderPath, err := osext.ExecutableFolder(); err == nil {
		fmt.Println(folderPath)
		//maybe we're running this in dev mode, so templates could be at folderPath/templates
		devTemplateDirName := fmt.Sprintf("%v/templates", folderPath)
		fmt.Printf("(dev)looking in %v\n", devTemplateDirName)
		if _, err := os.Stat(devTemplateDirName); err == nil {
			return devTemplateDirName, err
		} else {
			//maybe we're running this from a brew install, so the templates would be at folderPath/../share/templates
			brewTemplateDirName := fmt.Sprintf("%v/../share/templates", folderPath)
					fmt.Printf("(brew)looking in %v\n", brewTemplateDirName)
			if _, err := os.Stat(brewTemplateDirName); err == nil {
				return brewTemplateDirName, err
			} else {
				return "", err
			}
		}

	}

	//we only reach this part if err != nil
	return "", errors.New("unable to find templates dir")
}
