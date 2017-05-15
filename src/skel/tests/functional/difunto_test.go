package functional

import (
	"testing"
	"net/http"
	"github.com/gavv/httpexpect"
)

var (
	difuntosPath string = "/difuntos/"
	difuntosID   string = "1"
)

func TestDifuntoStatusCodes(t *testing.T) {

	cases := []struct {
		method, path string
		code         int
	}{
		{"POST", difuntosPath, http.StatusMethodNotAllowed },
		{"POST", difuntosPath + difuntosID, http.StatusBadRequest },
		{"GET", difuntosPath, http.StatusOK },
		{"GET", difuntosPath + difuntosID, http.StatusOK },
		{"GET", difuntosPath + difuntosID + "123", http.StatusNotFound },
		{"GET", difuntosPath + "x", http.StatusBadRequest },
		{"PUT", difuntosPath, http.StatusBadRequest },
		{"PUT", difuntosPath + difuntosID, http.StatusMethodNotAllowed },
		{"DELETE", difuntosPath, http.StatusBadRequest },
		{"DELETE", difuntosPath + difuntosID, http.StatusMethodNotAllowed },
	}

	e := httpexpect.New(t, baseUrl)

	for _, c := range cases {
		e.Request(c.method, c.path).
			Expect().
			Status(c.code)
	}
}

func TestGetDifunto(t *testing.T) {

	e := httpexpect.New(t, baseUrl)

	o := e.GET(difuntosPath + difuntosID).
		Expect().
		JSON().Object()

	o.Value("Name").Equal("Frida Kahlo")
}

func TestGetDifuntos(t *testing.T) {

	e := httpexpect.New(t, baseUrl)

	e.GET(difuntosPath).
		Expect().
		JSON().Array().Length().Equal(3)
}
