package functional

import (
	"testing"
	"net/http"
	"github.com/gavv/httpexpect"
)

var (
	ofrendasPath string = "/ofrendas/"
	ofrendasID   string = "1"
)

// TODO: test PUT/GET, POST altar

func TestOfrendaStatusCodes(t *testing.T) {

	cases := []struct {
		method, path string
		code         int
	}{
		{"POST", ofrendasPath, http.StatusMethodNotAllowed },
		{"POST", ofrendasPath + ofrendasID, http.StatusBadRequest },
		// GET list requires "type" filter
		{"GET", ofrendasPath, http.StatusBadRequest },
		{"GET", ofrendasPath + ofrendasID, http.StatusOK},
		{"GET", ofrendasPath + ofrendasID + "123", http.StatusNotFound},
		{"GET", ofrendasPath + "x", http.StatusBadRequest},
		{"PUT", ofrendasPath, http.StatusBadRequest },
		{"PUT", ofrendasPath + ofrendasID, http.StatusMethodNotAllowed },
		{"DELETE", ofrendasPath, http.StatusBadRequest },
		{"DELETE", ofrendasPath + ofrendasID, http.StatusMethodNotAllowed },
	}

	e := httpexpect.New(t, baseUrl)

	for _, c := range cases {
		e.Request(c.method, c.path).
			Expect().
			Status(c.code)
	}
}

func TestGetOfrenda(t *testing.T) {

	e := httpexpect.New(t, baseUrl)

	o := e.GET(ofrendasPath + ofrendasID).
		Expect().
		JSON().Object()

	o.Value("Name").Equal("Chocolate")
}

func TestGetOfrendas(t *testing.T) {

	cases := []struct {
		type_ string
		count int
	}{
		{"food", 6},
		{"drink", 3},
		{"flower", 4},
		{"x", 0},
	}

	e := httpexpect.New(t, baseUrl)

	for _, c := range cases {
		e.GET(ofrendasPath).WithQueryString("type=" + c.type_).
			Expect().
			JSON().Array().Length().Equal(c.count)
	}
}
