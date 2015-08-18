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
	for _, templateName := range GetEndpoints() {
		FindAndExecuteTemplate(templateName, concept)
	}
}

func MakeEndpointDir() (string, error) {
	if appDir, err := GetAppDir(); err == nil {
		endpointDir := fmt.Sprintf("%v/app/endpoint", appDir )
		log.Printf("trying to make %v\n", endpointDir )
		if err := os.MkdirAll(endpointDir, 0777); err == nil {
			log.Printf("should have made %v\n", endpointDir )
			return endpointDir, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func FindAndExecuteTemplate(templateName string, concept Concept) {

	if templateDir, err := utils.FindTemplateDir(); err == nil {

		if endpointDir, err := MakeEndpointDir(); err == nil {
			//TODO assume that js is the only implementation language for now
			templatesGlob := fmt.Sprintf("%v/js/endpoint/*.js", templateDir )

			if t, err := template.ParseGlob(templatesGlob); err == nil {

				outputFileName := fmt.Sprintf("%v/%v", endpointDir, templateName )
				log.Println(outputFileName)

				outputFilename, err := os.Create(outputFileName)

				if err = t.ExecuteTemplate(outputFilename, templateName, concept); err != nil {
					log.Fatal(err)
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
