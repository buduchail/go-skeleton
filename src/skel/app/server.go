package app

import (
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/middleware"

	"skel/app/resources"
	"skel/domain"
)

type (
	Server struct {
		api    catrina.RestAPI
		repo   domain.DayOfTheDeadRepository
		config Config
		logger catrina.Logger
	}
)

func NewServer(config Config, api catrina.RestAPI, repo domain.DayOfTheDeadRepository, logger catrina.Logger) Server {

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

	return Server{
		api:    api,
		repo:   repo,
		config: config,
		logger: logger,
	}
}

func (s Server) Run(port int) {
	s.logger.Info("Starting server", &catrina.LoggerContext{"router": s.config.Router, "port": port})
	s.api.Run(port)
}
