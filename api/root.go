package api

import (
	"fmt"
	"net/http"
)

var rootRoutes = Routes{
	Route{"RootIndex", "GET", "/", RootIndex},
}

func init() {
	for i := range rootRoutes {
		routes = append(routes, rootRoutes[i])
	}
}

func RootIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! This is a demo using golang to implement RESTful API.")
}
