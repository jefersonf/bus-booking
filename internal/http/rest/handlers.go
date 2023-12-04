package rest

import (
	"fmt"
	"net/http"

	"github.com/jefersonf/bus-booking/pkg/fleet"
)

func handlers() *http.ServeMux {
	routes := http.NewServeMux()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "fleet_size=%v", fleet.Size)
	})

	return routes
}
