package main

import (
	"os"
	"path"
	"errors"
	"strings"
	"runtime"

	"github.com/sirupsen/logrus"

	"skel/app"
	"skel/infrastructure/rest"
	"skel/infrastructure/repository"
	"skel/infrastructure/logger"
)

var (
	logLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	routers = map[string]string{
		"n": "nethttp",
		"i": "iris",
		"h": "httprouter",
		"e": "echo",
		"f": "fasthttp",
		"g": "gin",
		"r": "go-restful",
	}
)

func getConfig(path string) (config Config, err error) {
	config = Config{}
	err = app.LoadConfig(path, &config)
	return config, err
}

func getLogger(config Config) (app.Logger, error) {

	l := logger.NewLogrus(app.LoggerContext{
		"App": "go-skeleton",
	})

	l.SetLevel(logrus.InfoLevel)
	l.SetOutput(os.Stdout)

	if strings.ToLower(config.LogFormat) == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}

	level, ok := logLevels[strings.ToLower(config.LogLevel)]
	if ok {
		l.SetLevel(level)
	} else {
		l.Warn("Ignoring unknown log level", &app.LoggerContext{"ignored": config.LogLevel})
	}

	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err == nil {
			l.SetOutput(file)
		}
	}

	return l, nil
}

func getRepository(path string) (repo *repository.DayOfTheDeadMemoryRepository, err error) {
	repo = repository.NewDayOfTheDeadMemoryRepository()
	err = repo.LoadData(path)
	return repo, err
}

func getApi(prefix string, apiType string) (app.RestAPI, error) {

	switch apiType {
	case "n", routers["n"]:
		return rest.NewNetHTTP(prefix), nil
	case "i", routers["i"]:
		return rest.NewIris(prefix), nil
	case "h", routers["h"]:
		return rest.NewHttpRouter(prefix), nil
	case "e", routers["e"]:
		return rest.NewEcho(prefix), nil
	case "f", routers["f"]:
		return rest.NewFast(prefix), nil
	case "g", routers["g"]:
		return rest.NewGin(prefix), nil
	case "r", routers["r"]:
		return rest.NewGoRestful(prefix), nil
	}

	return nil, errors.New("Unknow router type: " + apiType)
}

func bootstrap(configFile string, dataFile string) (Config, app.Logger, *repository.DayOfTheDeadMemoryRepository, app.RestAPI) {

	_, filename, _, _ := runtime.Caller(0)
	baseDir := path.Dir(filename)

	config, err := getConfig(baseDir + "/" + configFile)
	if err != nil {
		panic("[BOOTSTRAP] Could not load configuration (" + err.Error() + ")")
	}

	logger_, err := getLogger(config)
	if err != nil {
		panic("[BOOTSTRAP] Could not configure logger (" + err.Error() + ")")
	}

	repo, err := getRepository(baseDir + "/" + dataFile)
	if err != nil {
		panic("[BOOTSTRAP] Could not load repository data (" + err.Error() + ")")
	}

	api, err := getApi(config.Prefix, config.Router)
	if err != nil {
		panic("[BOOTSTRAP] Could not provision api (" + err.Error() + ")")
	}

	return config, logger_, repo, api
}
