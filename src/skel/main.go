package main

import (
	"skel/app"
	"skel/app/resources"
	"skel/app/middleware"
)

const (
	CONFIG_FILE = "config.toml"
	DATA_FILE   = "data.json"
)

type (
	Config struct {
		Prefix    string
		Router    string
		Port      int
		CorrID    string
		Profile   string
		LogLevel  string
		LogFile   string
		LogFormat string
	}
)

func main() {

	config, logger, repo, api := bootstrap(CONFIG_FILE, DATA_FILE)

	app.SetUpProfiling(config.Profile, config.Port, logger)

	// middleware is applied in the order in which it is added
	api.AddMiddleware(middleware.NewCorrelationID(config.CorrID))
	api.AddMiddleware(middleware.NewRequestLogger(logger, config.CorrID))

	status := resources.NewStatusHandler(config.Prefix, config.Port, repo)

	status.SetRoutes([]string{"status", "ofrendas", "altares", "altares/*/niveles", "difuntos"})

	api.AddResource("status", status)
	api.AddResource("ofrendas", resources.NewOfrendaHandler(repo))
	api.AddResource("altares", resources.NewAltarHandler(repo))
	api.AddResource("altares/*/niveles", resources.NewNivelHandler(repo))
	api.AddResource("difuntos", resources.NewDifuntoHandler(repo))

	logger.Info("Starting server", &app.LoggerContext{"router": config.Router, "port": config.Port})

	api.Run(config.Port)
}
