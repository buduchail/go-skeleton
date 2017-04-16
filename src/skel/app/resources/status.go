package resources

import (
	"fmt"
	"errors"
	"runtime"
	"net/http"
	"math/rand"
	"encoding/json"
	"github.com/buduchail/catrina"
	"github.com/buduchail/catrina/rest"

	"skel/domain"
)

type (
	StatusHandler struct {
		rest.ResourceHandler
		prefix string
		port   int
		routes []string
		repo   domain.DayOfTheDeadRepository
	}

	Status struct {
		Status string
		Memory struct {
			Heap string
			Mem  string
			Sys  string
		}
		Goroutines int
		Routes     []string
		Stats      map[string]int
	}
)

func getMemStr(bytes uint64) string {
	var n float64 = float64(bytes)
	var d float64 = 1024

	for _, u := range []string{"B!", "KB", "MB", "GB"} {
		if n < d {
			return fmt.Sprintf("%.1f %s", n, u)
		}
		n = n / d
	}

	return fmt.Sprintf("%.2f TB", n)
}

func NewStatusHandler(prefix string, port int, repo domain.DayOfTheDeadRepository) (handler StatusHandler) {
	handler = StatusHandler{}
	handler.prefix = prefix
	handler.port = port
	handler.routes = make([]string, 0)
	handler.repo = repo
	return handler
}

func (s *StatusHandler) SetRoutes(resources []string) {
	s.routes = make([]string, 0, len(resources))
	for _, resource := range resources {
		s.routes = append(s.routes, fmt.Sprintf("http://localhost:%d%s/%s", s.port, s.prefix, resource))
	}
}

func (s StatusHandler) Get(id catrina.ResourceID, parentIds []catrina.ResourceID) (code int, body catrina.Payload, err error) {
	return http.StatusNotFound, catrina.EmptyBody, errors.New("Use " + s.prefix + "/status to query for service status")
}

func (s StatusHandler) GetMany(parentIds []catrina.ResourceID, params catrina.QueryParameters) (code int, body catrina.Payload, err error) {
	var (
		status Status
		mem    runtime.MemStats
	)

	status.Status = "OK"

	if rand.Intn(10) == 0 {
		status.Status = "KO"
	}

	runtime.ReadMemStats(&mem)

	status.Memory.Mem = getMemStr(mem.Alloc)
	status.Memory.Heap = getMemStr(mem.HeapAlloc)
	status.Memory.Sys = getMemStr(mem.Sys)

	status.Goroutines = runtime.NumGoroutine()

	status.Routes = s.routes
	status.Stats = s.repo.GetStats()

	str, _ := json.MarshalIndent(status, "", "    ")

	return http.StatusOK, str, nil
}
