package config
import (
	"github.com/PolkaBand/polka/hidden_home"
)

type PolkaConfig struct {
	Name string
}

//CreatePolkaHomeConfig creates ~/.polka if it doesn't exist already.
//It then places the contents of PolkaConfig into it as ~/.polka/polka_config.json
func CreatePolkaHomeConfigAsNeeded() {
	homeDir := "~/.polka"
	config := PolkaConfig{
		Name: "blah",
	}
	hidden_home.Save(homeDir, "~/.polka/polka_config.json", config)
}
