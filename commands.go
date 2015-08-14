package main

/*
 Cribbed **heavily** from https://github.com/hashicorp/consul
*/

import (
	"os"
	//"fmt"
	"github.com/PolkaBand/polka/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Polka commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{

		"generate": func() (cli.Command, error) {
			return &command.GenerateCommand{
				Name: "blah",
				Ui:   ui,
			}, nil
		},

		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Name: "new",
				Ui:   ui,
			}, nil
		},


		"version": func() (cli.Command, error) {
			ver := "0.0.1"
			rel := "a"

			return &command.VersionCommand{
				Revision:          "",
				Version:           ver,
				VersionPrerelease: rel,
				Ui:                ui,
			}, nil
		},
	}

}
