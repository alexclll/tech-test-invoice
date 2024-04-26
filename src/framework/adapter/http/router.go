package http

import (
	"net/http"
)

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

type Route interface {
	http.Handler
	Pattern() string
}
