package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type BaseAction struct {
	method   string
	route    string
	isPublic bool
}

func NewBaseAction(method, route string, isPublic bool) *BaseAction {
	return &BaseAction{
		method:   method,
		route:    route,
		isPublic: isPublic,
	}
}

func (a BaseAction) JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func (a BaseAction) Error(w http.ResponseWriter, status int, err error) {
	a.JSON(w, status, map[string]any{"error": err.Error()})
}

func (a BaseAction) Body(r *http.Request, data interface{}) error {
	// var f interface{}
	// json.NewDecoder(r.Body).Decode(&f)

	// fmt.Printf("form: %+v", f)
	return json.NewDecoder(r.Body).Decode(&data)
}

func (a BaseAction) Method() string {
	return a.method
}

func (a BaseAction) Route() string {
	return a.route
}

func (a BaseAction) IsPublic() bool {
	return a.isPublic
}

func (a BaseAction) GetAuthUserID(r *http.Request) int64 {
	_, claims, _ := jwtauth.FromContext(r.Context())

	return int64(claims["userId"].(float64))
}
