package rest

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type APIAction interface {
	Method() string
	Route() string
	IsPublic() bool
	Handle(w http.ResponseWriter, r *http.Request)
}
