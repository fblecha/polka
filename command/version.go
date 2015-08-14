package command

/*
 Cribbed **heavily** from https://github.com/hashicorp/consul
*/

import (
	"bytes"
	"fmt"
	"log"
	//"github.com/hashicorp/consul/consul"
	"github.com/mitchellh/cli"
)

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	Revision          string
	Version           string
	VersionPrerelease string
	Ui                cli.Ui
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Run(_ []string) int {
	log.Println("in version.Run()")

	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "Polka %s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.Ui.Output(versionString.String())
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the Polka version"
}
