package app

import "net/http"

var (
	EmptyBody = Payload([]byte(""))
)

type (
	RestAPI interface {
		AddResource(name string, handler ResourceHandler)
		AddMiddleware(m Middleware)
		Run(port int)
	}

	ResourceHandler interface {
		Post(parentIds []ResourceID, payload Payload) (
			code int, body Payload, err error,
		)
		Get(id ResourceID, parentIds []ResourceID) (
			code int, body Payload, err error,
		)
		GetMany(parentIds []ResourceID, query QueryParameters) (
			code int, body Payload, err error,
		)
		Put(id ResourceID, parentIds []ResourceID, payload Payload) (
			code int, body Payload, err error,
		)
		Delete(id ResourceID, parentIds []ResourceID) (
			code int, body Payload, err error,
		)
	}

	// Some syntactic sugar
	ResourceID string
	Payload []byte
	QueryParameters map[string][]string

	Middleware interface {
		Handle(w http.ResponseWriter, r *http.Request) *error
	}
)
