package config

import (
	"flag"
	"github.com/zasuchilas/gophkeeper/pkg/envflags"
	"log"
)

// Variables
var (
	// ServerAddress is the address and port to run server.
	ServerAddress        string
	defaultServerAddress = "localhost:9999"

	// Config is config filename.
	Config string
)

func ParseFlags() {

	// getting basic flags
	flag.StringVar(&ServerAddress, "s", "", "gophkeeper server address")
	// getting config.json file flag
	flag.StringVar(&Config, "config", "", "config filename")
	flag.StringVar(&Config, "c", "", "config filename")
	// parsing flags
	flag.Parse()

	// replacing from env
	envflags.TryUseEnvString(&ServerAddress, "SERVER_ADDRESS")

	// using config file or set default values
	if Config != "" {
		conf, er := getJSONConfig(Config)
		if er != nil {
			log.Panicf("error getting json config %s, error: %s", Config, er.Error())
		}
		// checking all config variables
		envflags.TryConfigStringFlag(&ServerAddress, conf.ServerAddress)
	}

	// setting defaults
	envflags.TryDefaultStringFlag(&ServerAddress, defaultServerAddress)
}
