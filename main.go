package main

import (
	"github.com/PolkaBand/polka/config"
	"github.com/PolkaBand/polka/utils"
	"github.com/mitchellh/cli"
	"log"
	"os"
	//"fmt"
)

var Config config.PolkaConfig
var MyAppConfig config.AppConfig

func main() {

	args := os.Args[1:]
	var err error
	var appDir string

	Config, err = config.CreatePolkaHomeConfigAsNeeded()
	if err != nil {
		log.Panic(err)
	}

	if len(args) > 0 {
		switch args[0] {
		case "new":
			//do nothing in the case of the "new" command
		default:
			if appDir, err = utils.AreWeInAppRootDir(); err == nil {
				//otherwise, make sure we have the AppConfig setup
				MyAppConfig, err = config.CreateAppConfigAsNeeded(appDir)
				if err != nil {
					log.Panic(err)
				}
			} else {
				log.Panic(err)
			}

		}
	}

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
