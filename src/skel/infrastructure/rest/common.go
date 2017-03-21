package rest

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"net/http"
)

var (
	unknownErr = errors.New("Unknown error")
)

func getHttpError(code int) error {

	status := http.StatusText(code)
	if status != "" {
		return errors.New(status)
	}

	return unknownErr
}

func normalizePrefix(prefix string) string {
	if prefix == "" {
		return prefix
	}
	return "/" + strings.TrimLeft(strings.TrimRight(prefix, "/"), "/")
}

func expandPath(path, idTemplate string) (fullPath string, parentIds []string, idParam string) {
	var i = 0
	parts := strings.Split(path, "/*/")
	parentIds = make([]string, 0, len(parts))
	fullPath = "/" + parts[0]
	l := len(parts)
	if l > 1 {
		for i = range parts[1:] {
			id := "id" + strconv.Itoa(i+1)
			parentIds = append(parentIds, id)
			fullPath += "/" + fmt.Sprintf(idTemplate, id) + "/" + parts[i+1]
		}
		i += 1
	}
	return fullPath, parentIds, "id" + strconv.Itoa(i+1)
}
