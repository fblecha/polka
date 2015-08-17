package main

import (
	"log"
	"os"
	"github.com/mitchellh/cli"
)

func main() {

	args := os.Args[1:]

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
