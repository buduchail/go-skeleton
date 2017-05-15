package functional

import (
	"testing"
	"net/http"
	"github.com/gavv/httpexpect"
	"strconv"
)

var (
	nivelesPath string                 = "/altares/10/niveles/"
	nivelesID   string                 = "1"
	gift        map[string]interface{} = map[string]interface{}{
		"ID":   10,
		"Name": "Cempasúchil",
		"Type": "flower",
	}
)

func TestNivelStatusCodes(t *testing.T) {

	cases := []struct {
		method, path string
		code         int
	}{
		{"POST", nivelesPath, http.StatusMethodNotAllowed },
		{"POST", nivelesPath + nivelesID, http.StatusBadRequest },
		{"GET", nivelesPath, http.StatusMethodNotAllowed },
		{"GET", nivelesPath + nivelesID, http.StatusNotFound },
		{"PUT", nivelesPath, http.StatusBadRequest },
		// PUT without body
		{"PUT", nivelesPath + nivelesID, http.StatusBadRequest },
		{"DELETE", nivelesPath, http.StatusBadRequest },
		{"DELETE", nivelesPath + nivelesID, http.StatusMethodNotAllowed },
	}

	e := httpexpect.New(t, baseUrl)

	for _, c := range cases {
		e.Request(c.method, c.path).
			Expect().
			Status(c.code)
	}
}

func TestGetNivel(t *testing.T) {

	id := postAltar(t)

	e := httpexpect.New(t, baseUrl)

	// shelve is empty when created
	e.GET("/altares/" + strconv.Itoa(id) + "/niveles/1").
		Expect().
		JSON().Equal(nil)
}

func TestGetNivelErrors(t *testing.T) {
	cases := []struct {
		path string
		code int
	}{
		{"/altares/x/niveles/1", http.StatusBadRequest},
		{"/altares/100/niveles/1", http.StatusNotFound},
		{"/altares/1/niveles/0", http.StatusOK},
		{"/altares/1/niveles/x", http.StatusBadRequest},
	}

	e := httpexpect.New(t, baseUrl)
	for _, c := range cases {
		e.GET(c.path).
			Expect().
			Status(c.code)
	}
}

func TestPutNivel(t *testing.T) {

	id := postAltar(t)

	e := httpexpect.New(t, baseUrl)

	// PUT returns shrine with modified shelves, each shelve is an array of gifts
	e.PUT("/altares/" + strconv.Itoa(id) + "/niveles/1").WithJSON(gift).
		Expect().
		JSON().Object().Value("Shelves").
		Array().Element(0).
		Array().Element(0).Equal(gift)

}

func TestPutNivelErrors(t *testing.T) {

	missing := map[string]interface{}{
		"ID":   100,
		"Name": "Peonia",
		"Type": "flower",
	}

	cases := []struct {
		path string
		gift map[string]interface{}
		code int
	}{
		{"/altares/x/niveles/1", gift, http.StatusBadRequest},
		{"/altares/100/niveles/1", gift, http.StatusInternalServerError},
		{"/altares/1/niveles/x", gift, http.StatusBadRequest},
		{"/altares/1/niveles/0", gift, http.StatusInternalServerError},
		{"/altares/1/niveles/1", missing, http.StatusInternalServerError},
	}

	e := httpexpect.New(t, baseUrl)
	for _, c := range cases {
		e.PUT(c.path).WithJSON(c.gift).
			Expect().
			Status(c.code)
	}
}
