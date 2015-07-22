package main

import (
	"log"
	"os"
	"fmt"
	"github.com/mitchellh/cli"
)
func main() {
	fmt.Println("Let's Polka!")
  c := cli.NewCLI("polka", "0.0.1")
  c.Args = os.Args[1:]
  c.Commands = map[string]cli.CommandFactory{
      "foo": fooCommandFactory,
      "bar": barCommandFactory,
  }

  exitStatus, err := c.Run()
  if err != nil {
      log.Println(err)
  }

  os.Exit(exitStatus)
}
