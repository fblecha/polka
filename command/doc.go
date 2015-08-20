package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
	//"github.com/PolkaBand/polka/config"
	//"encoding/json"
	//"io/ioutil"
	"path/filepath"
)

type DocCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *DocCommand) Help() string {
	helpText := `
Usage: polka doc [subcommand]

Generates documentation for your polka app.

If no subcommand is specified, polka documents everything.

Subcommands:
	endpoints 	- list out all the endpoints in this application
	integrations	-
	architecture	-
	resources 	- summary table of the resources requested for each lambda function
`
	return strings.TrimSpace(helpText)
}

func (c *DocCommand) Run(args []string) int {

	log.Printf("new command # of args = %v\n", len(args))

	/*
		if the current working directory doesn't contain an app directory  (e.g. ./app/ exists),
		then we fatal out.  Otherwise, let's iterate over everything.
	*/
	DocEndpoints()
	DocIntegrations()
	DocArchitecture()
	DocResources()

	return 0
}

func (c *DocCommand) Synopsis() string {
	return "Generates documentation for your polka app"
}

func DocEndpoints() {
	//check that we're in the app directory
}

func DocIntegrations() {

}

func DocArchitecture() {

}

//DocResources emits a summary table
func DocResources() {
	//let's just do endpoints initially
	appDirectory, err := GetAppDir()
	if err != nil {
		log.Fatal(err)
	}

	location := fmt.Sprintf("%v/app", appDirectory)

	var configLoader = func(path string, fileInfo os.FileInfo, _ error) (err error) {
		if fileInfo, err := os.Lstat(path); err == nil {
			if fileInfo.Name() == "lambda_config.json" {
				log.Printf("found %v in %v", fileInfo.Name(), path)
				/*

					var resouce ResourceConfig
					var jsontype jsonobject
					json.Unmarshal(file, &resource)
					fmt.Printf("Results: %v\n", resource)
				*/
			}
			return nil
		} else {
			return err
		}
	}

	err = filepath.Walk(location, configLoader)
	log.Println(err)
}
