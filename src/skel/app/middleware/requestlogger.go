package middleware

import (
	"net/http"
	"skel/app"
)

type (
	RequestLogger struct {
		logger    app.Logger
		logHeader string
	}
)

func NewRequestLogger(logger app.Logger, correlationIdHeader string) *RequestLogger {
	return &RequestLogger{logger, correlationIdHeader}
}

func (m RequestLogger) Handle(w http.ResponseWriter, r *http.Request) (err *error) {

	m.logger.Info(
		r.Method+" "+r.URL.String(),
		&app.LoggerContext{m.logHeader: r.Header[m.logHeader]},
	)

	return
}
