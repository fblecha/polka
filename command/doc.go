package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"strings"
	//"text/template"
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

`
	return strings.TrimSpace(helpText)
}

func (c *DocCommand) Run(args []string) int {

	log.Printf("new command # of args = %v\n", len(args))
	if len(args) < 1 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	log.Fatal("doc command is currently just a stub")

	return 0
}

func (c *DocCommand) Synopsis() string {
	return "Generates documentation for your polka app"
}
