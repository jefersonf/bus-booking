package rest

import (
	"fmt"
	"net/http"
)

func handlers() *http.ServeMux {
	routes := http.NewServeMux()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "fleet_size=%v", 10)
	})

	return routes
}
