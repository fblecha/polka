package main

import (
	"log"
	"os"
	"fmt"
	"github.com/mitchellh/cli"
)

func main() {
	fmt.Println("Let's Polka!")

	args := os.Args[1:]

	fmt.Println("args = ", args)


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
