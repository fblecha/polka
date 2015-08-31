package command

import (
	//"fmt"
	//"github.com/PolkaBand/polka/config"
	"github.com/PolkaBand/polka/utils"
	"github.com/mitchellh/cli"
	//"log"
	//"os"
	"strings"
	//"text/template"
)

type DestroyCommand struct {
	Name string
	Ui   cli.Ui
}

func (c *DestroyCommand) Help() string {
	helpText := `
Usage: polka destroy [options] [subcommand] concept

Destroy the result of 'polka generate [subcommand] concept'

Options:
	-f 	Force

Subcommands:

	endpoint	Not implemented
`
	return strings.TrimSpace(helpText)
}

func (c *DestroyCommand) Run(args []string) int {
	if _, err := utils.AreWeInAppRootDir(); err != nil {

		return 1
	}

	return 1
}

func (c *DestroyCommand) Synopsis() string {
	return "Destroys generated files associated with the given subcommand"
}
