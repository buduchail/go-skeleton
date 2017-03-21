package rest

import (
	"bufio"
	"strings"
	"skel/app"
)

type (
	pathHandler struct {
		handler  app.ResourceHandler
		resource string
		children map[string]*pathHandler
	}
)

func NewPathHandler(resource string) *pathHandler {
	ph := &pathHandler{
		resource: resource,
		children: make(map[string]*pathHandler, 0),
	}
	return ph
}

func (ph *pathHandler) addHandler(path string, handler app.ResourceHandler) {
	var (
		child, p *pathHandler
		exists   bool
	)
	p = ph
	for _, part := range strings.Split(path, "/*/") {
		child, exists = p.children[part]
		if !exists {
			child = NewPathHandler(part)
		}
		p.children[part] = child
		p = child
	}
	p.handler = handler
}

func (ph *pathHandler) findHandler(path string) (handler app.ResourceHandler, id app.ResourceID, parentIds []app.ResourceID) {
	handler = nil
	id = ""
	parentIds = make([]app.ResourceID, 0)

	scanner := bufio.NewScanner(strings.NewReader(path))

	scanner.Split(func(path []byte, atEOF bool) (advance int, token []byte, err error) {

		if atEOF && len(path) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(path), "/"); i >= 0 {
			return i + 1, path[0:i], nil
		}

		if atEOF {
			return len(path), path, nil
		}

		return
	})

	parts := 0
	p := ph
	i := 0
	for scanner.Scan() {
		i++
		if i%2 == 1 {
			parts++
			child, exists := p.children[scanner.Text()]
			if !exists {
				return nil, "", nil
			}
			p = child
		} else {
			parentIds = append(parentIds, app.ResourceID(scanner.Text()))
		}
	}

	if parts == len(parentIds) {
		id = parentIds[parts-1]
		parentIds = parentIds[:parts-1]
	}

	return p.handler, app.ResourceID(id), parentIds
}
