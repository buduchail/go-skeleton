package main

import (
	"skel/app"
)

const (
	CONFIG_FILE = "config.toml"
	DATA_FILE   = "data.json"
)

func main() {

	config, logger, repo, api := bootstrap(CONFIG_FILE, DATA_FILE)

	setProfiling(config.Profile, config.Port, logger)

	server := app.NewServer(config, api, repo, logger)

	server.Run(config.Port)
}
