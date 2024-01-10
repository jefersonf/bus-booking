package rest

import (
	"net/http"
)

func handlers() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("/", baseURI)
	return routes
}

func baseURI(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"status": "running"}`))
}
