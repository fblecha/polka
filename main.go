package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
	"github.com/PolkaBand/polka/config"
)

var Config config.PolkaConfig

func main() {

	args := os.Args[1:]

	Config = config.CreatePolkaHomeConfigAsNeeded()

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("polka"),
	}

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
