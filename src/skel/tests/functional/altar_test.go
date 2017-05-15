package functional

import (
	"strconv"
	"testing"
	"net/http"
	"github.com/gavv/httpexpect"
)

var (
	altaresPath string = "/altares/"
	altaresID   string = "1"
)

func TestAltarStatusCodes(t *testing.T) {

	cases := []struct {
		method, path string
		code         int
	}{
		{"POST", altaresPath, http.StatusBadRequest },
		// POST without body
		{"POST", altaresPath + altaresID, http.StatusBadRequest },
		{"GET", altaresPath, http.StatusMethodNotAllowed },
		{"GET", altaresPath + altaresID, http.StatusNotFound },
		{"GET", altaresPath + "x", http.StatusBadRequest },
		{"PUT", altaresPath, http.StatusBadRequest },
		{"PUT", altaresPath + altaresID, http.StatusMethodNotAllowed },
		{"DELETE", altaresPath, http.StatusBadRequest },
		// DELETE non-existent shrine
		{"DELETE", altaresPath + altaresID, http.StatusInternalServerError },
		{"DELETE", altaresPath + "x", http.StatusBadRequest },
	}

	e := httpexpect.New(t, baseUrl)

	for _, c := range cases {
		e.Request(c.method, c.path).
			Expect().
			Status(c.code)
	}
}

func TestPostAltar(t *testing.T) {

	payload := map[string]interface{}{
		"MexicanID": 1,
		"Levels":    3,
	}

	e := httpexpect.New(t, baseUrl)

	o := e.POST(altaresPath).WithJSON(payload).
		Expect().
		JSON().Object()

	o.Value("MexicanID").Equal(1)
	o.Value("Levels").Equal(3)
}

func TestPostAltarErrors(t *testing.T) {

	cases := []map[string]interface{}{
		// invalid ID field
		map[string]interface{}{
			"ID":        1,
			"MexicanID": 1,
			"Levels":    3,
		},
		// unknown mexican
		map[string]interface{}{
			"MexicanID": 100,
			"Levels":    3,
		},
		// wrong level number
		map[string]interface{}{
			"MexicanID": 1,
			"Levels":    30,
		},
	}

	e := httpexpect.New(t, baseUrl)

	for _, template := range cases {
		e.POST(altaresPath).WithJSON(template).
			Expect().
			Status(http.StatusInternalServerError)
	}
}

func TestGetAltar(t *testing.T) {

	id := postAltar(t)

	e := httpexpect.New(t, baseUrl)

	e.GET(altaresPath + strconv.Itoa(id)).
		Expect().
		JSON().Object().Value("ID").Equal(id)

}

func TestDeleteAltar(t *testing.T) {

	id := postAltar(t)

	e := httpexpect.New(t, baseUrl)

	e.DELETE(altaresPath + strconv.Itoa(id)).
		Expect().
		Status(http.StatusOK)
}
