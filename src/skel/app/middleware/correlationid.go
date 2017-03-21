package middleware

import (
	"net/http"
	"github.com/satori/go.uuid"
)

type (
	CorrelationID struct {
		headerName string
	}
)

func NewCorrelationID(headerName string) *CorrelationID {
	return &CorrelationID{headerName}
}

func (m CorrelationID) Handle(w http.ResponseWriter, r *http.Request) (err *error) {

	_, exists := r.Header[m.headerName]
	if !exists {
		r.Header[m.headerName] = []string{uuid.NewV4().String()}
	}

	return
}
