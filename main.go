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


	//layout the following:
	// {{thing}} /
	// 		endpoint /
	//	 					thing_endpoint_create.js
	//						thing_endpoint_read.js
	//						thing_endpoint_update.js
	//						thing_endpoint_delete.js
	//						thing_endpoint_list.js
	//						thing_endpoint_health.js

	concept := Concept{"accordion"}
	endpoints := [...]string {
		"_create.js",
		"_read.js",
		"_update.js",
		"_delete.js",
		"_health.js",
	}
	template, err := template.New("blah").Parse("{{.Name}} is a new endpoint\n")

	for _, value := range endpoints {

		fmt.Println(concept.Name, "_endpoint", value )


		if err != nil { panic(err) }
		err = template.Execute(os.Stdout, concept)
		if err != nil { panic(err) }

	}






	os.Exit(exitStatus)
}
