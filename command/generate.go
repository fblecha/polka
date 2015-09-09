package command

/*
 Cribbed **heavily** from https://github.com/hashicorp/consul
*/

import (
	"fmt"
	"github.com/PolkaBand/polka/config"
	"github.com/PolkaBand/polka/utils"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"strings"
	"text/template"
)

type GenerateCommand struct {
	Name string
	Ui   cli.Ui
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

func MakeAllEndpoints(currentDir string, conceptName string) {
	concept := Concept{conceptName}
	for _, templateName := range GetEndpoints() {
		FindAndExecuteTemplate(currentDir, templateName, concept)
		dir := strings.Split(templateName, ".")[0]
		log.Printf("%q", strings.Split(templateName, "."))
		config.Create(concept.Name, dir)
	}
}

func MakeEndpointDir(currentDir string, concept Concept) (string, error) {
	endpointDir := fmt.Sprintf("%v/app/endpoint/%v", currentDir, concept.Name)
	log.Printf("trying to make %v\n", endpointDir)
	return endpointDir, os.MkdirAll(endpointDir, 0777)
}

func FindAndExecuteTemplate(currentDir string, templateName string, concept Concept) error {

	templateDir, err := utils.FindTemplateDir()
	if err != nil {
		return err
	}

	endpointDir, err := MakeEndpointDir(currentDir, concept)
	if err != nil {
		return err
	}
	//TODO assume that js is the only implementation language for now
	templatesGlob := fmt.Sprintf("%v/js/endpoint/*.js", templateDir)

	t, err := template.ParseGlob(templatesGlob)
	if err != nil {
		return err
	}
	outputFileName := fmt.Sprintf("%v/%v", endpointDir, templateName)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(outputFile, templateName, concept)
}

func (c *GenerateCommand) Run(args []string) int {
	if _, err := utils.AreWeInAppRootDir(); err != nil {
		return 1
	}

	if len(args) < 2 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	subcommand := args[0]

	//assume that we're in the base polka directory
	//confirm that by checking to see if ./app/ exists
	if currentDir, err := os.Getwd(); err == nil {
		//TODO check to see if ./app/ exists
		switch subcommand {
		case "endpoint":
			MakeAllEndpoints(currentDir, args[1])
		case "integration":
			log.Println("generating integrations are not implemented yet")
		}
		return 0
	}
	return 1
}

func (c *GenerateCommand) Synopsis() string {
	return "generates lambda functions for endpoints"
}
