package main

import (
	"os"
	"time"
	"strconv"
	"net/http"
	"runtime/pprof"
	_ "net/http/pprof"
	"github.com/buduchail/catrina"
)

func setProfiling(profilePath string, port int, logger catrina.Logger) {

	if profilePath != "" {

		if false {
			// For a command line application, run this code:
			f, _ := os.Create(profilePath)
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
			// .. then do something
		}

		go func() {
			err := http.ListenAndServe(":"+strconv.Itoa(port+1), nil)
			if err == nil {
				logger.Info("Runnig profile server", &catrina.LoggerContext{"port": port + 1})
			} else {
				logger.Warn("Could not start profiling server", &catrina.LoggerContext{"error": err})
			}
		}()

		// Allow the go routine for the debug server to start
		time.Sleep(time.Millisecond * 1000)
	}
}
