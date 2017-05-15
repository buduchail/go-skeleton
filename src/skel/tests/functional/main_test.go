package functional

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"testing"

	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/rest"

	"skel/app"
	"skel/domain"
	"skel/infrastructure/repository"
	"github.com/gavv/httpexpect"
)

type (
	DummyLogger struct {
	}
)

const (
	DATA_FILE = "../../../../data/data.json"
)

var (
	port     int = 1234
	baseUrl  string
	rootPath string = "/"
)

func (l DummyLogger) Debug(message string, context *catrina.LoggerContext)   {}
func (l DummyLogger) Info(message string, context *catrina.LoggerContext)    {}
func (l DummyLogger) Print(message string, context *catrina.LoggerContext)   {}
func (l DummyLogger) Warn(message string, context *catrina.LoggerContext)    {}
func (l DummyLogger) Warning(message string, context *catrina.LoggerContext) {}
func (l DummyLogger) Error(message string, context *catrina.LoggerContext)   {}
func (l DummyLogger) Fatal(message string, context *catrina.LoggerContext)   {}
func (l DummyLogger) Panic(message string, context *catrina.LoggerContext)   {}

func postAltar(t *testing.T) int {

	payload := map[string]interface{}{
		"MexicanID": 1,
		"Levels":    3,
	}

	e := httpexpect.New(t, baseUrl)

	val := e.POST(altaresPath).WithJSON(payload).
		Expect().
		JSON().Object().Value("ID").Raw()

	return int(val.(float64))
}

func bootstrap() (app.Config, catrina.RestAPI, domain.DayOfTheDeadRepository, catrina.Logger) {

	_, filename, _, _ := runtime.Caller(0)
	baseDir := path.Dir(filename) + "/"

	config := app.Config{}

	api := rest.NewNetHTTP(rootPath)

	repo, err := repository.NewDayOfTheDeadMemoryRepository(baseDir + DATA_FILE)
	if err != nil {
		panic(err)
	}

	return config, api, repo, DummyLogger{}
}

func TestMain(m *testing.M) {

	config, api, repo, logger := bootstrap()
	server := app.NewServer(config, api, repo, logger)
	go server.Run(port)

	baseUrl = "http://localhost:" + strconv.Itoa(port) + rootPath

	code := m.Run()

	os.Exit(code)
}
