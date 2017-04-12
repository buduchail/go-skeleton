package main

import (
	"os"
	"path"
	"errors"
	"strings"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/rest"
	"github.com/buduchail/catrina/config"
	"github.com/buduchail/catrina/logger"

	"skel/infrastructure/repository"
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

func getConfig(path string) (config_ Config, err error) {
	config_ = Config{}
	err = config.Load(path, &config_)
	return config_, err
}

func getLogger(config_ Config) (catrina.Logger, error) {

	l := logger.NewLogrus(catrina.LoggerContext{
		"App": "go-skeleton",
	})

	l.SetLevel(logrus.InfoLevel)
	l.SetOutput(os.Stdout)

	if strings.ToLower(config_.LogFormat) == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}

	level, ok := logLevels[strings.ToLower(config_.LogLevel)]
	if ok {
		l.SetLevel(level)
	} else {
		l.Warn("Ignoring unknown log level", &catrina.LoggerContext{"ignored": config_.LogLevel})
	}

	if config_.LogFile != "" {
		file, err := os.OpenFile(config_.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
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

func getApi(prefix string, apiType string) (catrina.RestAPI, error) {

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

func bootstrap(configFile string, dataFile string) (Config, catrina.Logger, *repository.DayOfTheDeadMemoryRepository, catrina.RestAPI) {

	_, filename, _, _ := runtime.Caller(0)
	baseDir := path.Dir(filename)

	config_, err := getConfig(baseDir + "/" + configFile)
	if err != nil {
		panic("[BOOTSTRAP] Could not load configuration (" + err.Error() + ")")
	}

	logger_, err := getLogger(config_)
	if err != nil {
		panic("[BOOTSTRAP] Could not configure logger (" + err.Error() + ")")
	}

	repo, err := getRepository(baseDir + "/" + dataFile)
	if err != nil {
		panic("[BOOTSTRAP] Could not load repository data (" + err.Error() + ")")
	}

	api, err := getApi(config_.Prefix, config_.Router)
	if err != nil {
		panic("[BOOTSTRAP] Could not provision api (" + err.Error() + ")")
	}

	return config_, logger_, repo, api
}
