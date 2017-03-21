package resources

import (
	"fmt"
	"errors"
	"runtime"
	"net/http"
	"math/rand"
	"encoding/json"
	"skel/app"
	"skel/domain"
	"skel/infrastructure/repository"
)

type (
	StatusHandler struct {
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

func NewStatusHandler(prefix string, port int, repo *repository.DayOfTheDeadMemoryRepository) (handler StatusHandler) {
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

func (s StatusHandler) Post(parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (s StatusHandler) Get(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	return http.StatusNotFound, app.EmptyBody, errors.New("Use " + s.prefix + "/status to query for service status")
}

func (s StatusHandler) GetMany(parentIds []app.ResourceID, params app.QueryParameters) (code int, body app.Payload, err error) {
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

func (s StatusHandler) Put(id app.ResourceID, parentIds []app.ResourceID, payload app.Payload) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}

func (s StatusHandler) Delete(id app.ResourceID, parentIds []app.ResourceID) (code int, body app.Payload, err error) {
	return http.StatusMethodNotAllowed, app.EmptyBody, nil
}
