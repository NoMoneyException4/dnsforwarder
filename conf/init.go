package conf

import (
	"github.com/mailgun/cfg"
)

var (
	Conf Configuration
)

func LoadConf(path string) {
	Conf = Configuration{}
	err := cfg.LoadConfig(path, &Conf)

	if err != nil {
		panic(err)
	}
}
