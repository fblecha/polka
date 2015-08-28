package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
	"github.com/PolkaBand/polka/config"
)

func main() {

	args := os.Args[1:]

	config.CreatePolkaHomeConfigAsNeeded()

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
