package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	//"log"
	//"os"
	"strings"
	//"text/template"
)

type ConfigCommand struct {
	Name string
	Ui   cli.Ui
}


func (c *ConfigCommand) Help() string {
	helpText := `
Usage: polka config [item]

Configure a global polka item that applies to the entire app.

Items:

	s3root	the root bucket for this app \n

`
	return strings.TrimSpace(helpText)
}


func (c *ConfigCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}
	//assume that we want to create app with name args[0] in the current working directory

	// if currentDir, err := os.Getwd(); err == nil {
	// 	CreateNewApp(currentDir, args[0])
	// 	return 0
	// } else {
	// 	log.Fatal(err)
	// 	return 1
	// }
	return 1
}

func (c *ConfigCommand) Synopsis() string {
	return "configure a global polka app item"
}
