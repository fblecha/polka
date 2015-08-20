package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"strings"
	//"text/template"
)

type DeployCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *DeployCommand) Help() string {
	helpText := `
Usage: polka deploy [subcommand]

Deploys all or some of the app to AWS

If no subcommand is specified, polka attempts to deploy everything.

Subcommands:
	
`
	return strings.TrimSpace(helpText)
}

func (c *DeployCommand) Run(args []string) int {

	log.Printf("new command # of args = %v\n", len(args))
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	log.Fatal("deploy command is currently just a stub")

	return 0
}

func (c *DeployCommand) Synopsis() string {
	return "deploys a polka app to AWS"
}
