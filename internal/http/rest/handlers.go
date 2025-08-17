package rest

import (
	"net/http"
)

func handlers() *http.ServeMux {
	routes := http.NewServeMux()
<<<<<<< HEAD
	routes.HandleFunc("/", indexHandler)
=======
	routes.HandleFunc("/", baseURI)

>>>>>>> d2d8cc2 (:sparkles: improvements)
	return routes
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"status": "API server up and running"}`))
}
