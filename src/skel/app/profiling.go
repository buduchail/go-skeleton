package app

import (
	"os"
	"time"
	"strconv"
	"net/http"
	"runtime/pprof"
	_ "net/http/pprof"
)

func SetUpProfiling(profilePath string, port int, logger Logger) {

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
				logger.Info("Runnig profile server", &LoggerContext{"port": port + 1})
			} else {
				logger.Warn("Could not start profiling server", &LoggerContext{"error": err})
			}
		}()

		// Allow the go routine for the debug server to start
		time.Sleep(time.Millisecond * 1000)
	}
}
