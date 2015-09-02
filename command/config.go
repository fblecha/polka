package command

import (
	"fmt"
	"github.com/PolkaBand/polka/config"
	"github.com/PolkaBand/polka/utils"
	"github.com/mitchellh/cli"
	"log"
	//"os"
	"strings"
	//"text/template"
)

type ConfigCommand struct {
	Name   string
	Ui     cli.Ui
	Config config.AppConfig
}

func (c *ConfigCommand) Help() string {
	helpText := `
Usage: polka config [item]

Configure a global polka item that applies to the entire app.

Items:

	s3	the root bucket for this app \n

`
	return strings.TrimSpace(helpText)
}

func (c *ConfigCommand) Run(args []string) int {
	if appDir, err := utils.AreWeInAppRootDir(); err != nil {
		fmt.Printf("%v", NotInAppDirectoryMessage())
		return 1
	} else {
		config.CreateAppConfigAsNeeded(appDir)
	}

	if len(args) < 2 {
		fmt.Printf("%v", c.Help())
		return 1
	}

	subcommand := args[0]
	//assume that we're itn the base polka directory
	//confirm that by checking to see if ./app/ exists

	switch subcommand {
		case "s3":
		//e.g. polka config s3 s3://SomeUrlHere/Etc
		c.ConfigureS3(args[1])
		return 0
	}
	return 1
}

func (c *ConfigCommand) Synopsis() string {
	return "configure a global polka app item"
}

func (c *ConfigCommand) ConfigureS3(s3url string) {
	log.Printf("s3url = %s \n", s3url)
	c.Config.S3 = s3url
	c.Config.Save()
}
