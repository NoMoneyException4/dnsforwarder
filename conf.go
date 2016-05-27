package main

import (
	"github.com/mailgun/cfg"
)

var (
	//Conf Loaded configurations
	Conf configuration
)

type configuration struct {
	Cache struct {
		Enable bool
		Ttl    int
	}
	Timeout struct {
		Read  int
		Write int
	}
	Hosts struct {
		Enable  bool
		Resolvs []string
	}
	Upstreams []string
	Loggers   struct {
		Console struct {
			Enable bool
		}
		File struct {
			Enable bool
			Level  string
			Path   string
		}
	}
}

// LoadConf Load configuration from given file
func LoadConf(path string) {
	Conf = configuration{}
	err := cfg.LoadConfig(path, &Conf)

	if err != nil {
		panic(err)
	}
}
