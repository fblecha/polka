package command

/*
 Cribbed **heavily** from https://github.com/hashicorp/consul
*/

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"log"
	"strings"
	//"text/template"
)

type NewCommand struct {
	Name string
	Ui   cli.Ui
}

type PolkaDir struct {
	RootDir string 	//root dir absolute path e.g. /Users/fb3/code/todo_app
	DirName string  //relative name to RootDir, e.g. app, test, config
}

func (d *PolkaDir) Create() error {
	log.Printf("creating %v/%v \n", d.RootDir, d.DirName)
	return nil
}



func (c *NewCommand) Help() string {
	helpText := `
Usage: polka new app_name

Generate a new polka application

`
	return strings.TrimSpace(helpText)
}

//Create a new app in the absolute location specified by name
//eg it should be ""/Users/fb3/code/todo" not "todo"
func CreateAppRootDir(name string) error {
	log.Printf("app root dir = %v\n", name)

	//mode := int(0777)
	if err := os.MkdirAll(name, 0777); err != nil {
		panic(err)
	}

	return nil
}

func CreateNewApp(rootDir string, name string) {
	appDir := fmt.Sprintf("%v/%v", rootDir, name)

	CreateAppRootDir(appDir)

	childDirs := [...]PolkaDir{
		{appDir, "app"}, //the primary app location -- most new code will go in here
		{appDir, "bin"}, //app specific commands -- note that "generic" commmands will be installed as part of polka
		{appDir, "config"}, //central source for the app config
		{appDir, "test"},
	}
	for _, child := range childDirs {
		if err := child.Create(); err != nil {
			panic(err)
		}
	}
}

func (c *NewCommand) Run(args []string) int {

	log.Printf("new command # of args = %v\n", len(args))
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}
	//assume that we want to create app with name args[0] in the current working directory

	if currentDir, err := os.Getwd(); err == nil  {
		CreateNewApp(currentDir, args[0])
		return 0
	} else {
		log.Fatal(err)
		return 1
	}
}

func (c *NewCommand) Synopsis() string {
	return "creates a new polka app"
}
