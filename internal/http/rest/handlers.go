package rest

import "net/http"

func handlers() *http.ServeMux {
	routes := http.NewServeMux()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "bus booking server running"}`))
	})

	return routes
}
