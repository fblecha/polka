package main

import (
	"log"
	"os"
	"fmt"
	"text/template"
	"github.com/mitchellh/cli"
)

type Concept struct {
	Name	string
}

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

	concept := Concept{"accordion"}
	endpoints := [...]string { "create", "read", "update", "delete", "health" }

	for _, value := range endpoints {
		templateFilename := fmt.Sprintf("templates/endpoint/%v.js", value)
		t, _ := template.ParseFiles(templateFilename)
		err = t.Execute(os.Stdout, concept)
		if err != nil { panic(err) }
	}


	os.Exit(exitStatus)
}
