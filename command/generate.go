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
	"github.com/PolkaBand/polka/utils"
	"text/template"
)

type GenerateCommand struct {
	Name string
	Ui   cli.Ui
}

type Concept struct {
	Name string
}

func (c *GenerateCommand) Help() string {
	helpText := `
Usage: polka generate [subcommand] concept

Generate a lambda function of type [subcommand] to implement the concept.

Subcommands:

	endpoint	an endpoint that supports CRUD and a health check on (concept)\n
`
	return strings.TrimSpace(helpText)
}

func GetEndpoints() []string {
	endpoints := []string{
		"create.js",
		"read.js",
		"update.js",
		"delete.js",
		"health.js",
	}
	return endpoints
}


func GenerateEndpoint(name string) {
	concept := Concept{name}
	endpoints := GetEndpoints()

	if templateDir, err := utils.FindTemplateDir(); err == nil {

		appDir, err := GetAppDir()
		if err != nil {
			log.Fatal(err)
		}
		endpointDir := fmt.Sprintf("%v/app/endpoint", appDir )

		log.Printf("trying to make %v\n", endpointDir )

		if err := os.MkdirAll(endpointDir, 0777); err == nil  {

			log.Printf("should have made %v\n", endpointDir )

			for _, value := range endpoints {
				//TODO assume that the implementation language will be JavaScript (js) for now


				templatesGlob := fmt.Sprintf("%v/js/endpoint/*.js", templateDir )
				t, err := template.ParseGlob(templatesGlob)
				if err != nil {
					log.Panic(err)
				}


				outputFileName := fmt.Sprintf("%v/%v", endpointDir, value )
				//log.Println("generate endpoint---=-")
				log.Println(outputFileName)

				outputFile, err := os.Create(outputFileName)

				if err = t.ExecuteTemplate(outputFile, value, concept); err != nil {
					log.Panic(err)
				}
			}
		}
	}

}

func (c *GenerateCommand) Run(args []string) int {
	if len(args) < 2 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	subcommand := args[0]

	switch subcommand {
	case "endpoint":
		GenerateEndpoint(args[1])
	case "integration":
		log.Println("generating integrations are not implemented yet")
	}

	return 0
}

func (c *GenerateCommand) Synopsis() string {
	return "generates lambda functions for endpoints"
}
