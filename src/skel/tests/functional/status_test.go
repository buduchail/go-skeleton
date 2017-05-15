package functional

import (
	"testing"
	"net/http"
	"github.com/gavv/httpexpect"
)

var (
	statusPath string = "/status"
)

func TestStatusOK(t *testing.T) {

	e := httpexpect.New(t, baseUrl)

	e.GET(statusPath).
		Expect().
		Status(http.StatusOK).
		JSON().Object().Value("Status").Equal("OK")
}
