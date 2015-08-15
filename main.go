package main

import (
	"log"
	"os"
	"github.com/mitchellh/cli"
)

func main() {

	templateDir, err := FindTemplateDir()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("template dir = %v\n", templateDir)

	args := os.Args[1:]

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("polka"),
	}

	log.Printf("topflags = %v \n", cli.Args)

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
