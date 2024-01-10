package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	listenAddr        = "localhost:3000"
	readTimeout       = 1 * time.Second
	writeTimeout      = 2 * time.Second
	readHeaderTimeout = 1 * time.Second
)

func StartBusBookingServer() {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	server := &http.Server{
		Addr:              listenAddr,
		Handler:           handlers(),
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("Bus booking API running at %s", listenAddr)
		log.Println(server.ListenAndServe())
	}()

	<-stopChan
	log.Println("Shutting down bus booking API server")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("Bus booking API server gracefully stopped")
}
