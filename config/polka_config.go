package config
import (
	"github.com/fblecha/homeplate"
)

type PolkaConfig struct {
	Name string
}

func (p *PolkaConfig) Save() {
	homeDir := "polka"
	homeplate.Save(homeDir, "polka", p, false)
}


//CreatePolkaHomeConfig creates ~/.polka if it doesn't exist already.
//It then places the contents of PolkaConfig into it as ~/.polka/polka_config.json
func CreatePolkaHomeConfigAsNeeded() {
	config := PolkaConfig{
		Name: "polka",
	}
	config.Save()
}
