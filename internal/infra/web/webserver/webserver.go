package webserver

import (
	"context"
	"go-ebanx/internal/infra/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type WebServer struct {
	server  *http.Server
	handler *web.Handler
}

func NewWebServer(handler *web.Handler) *WebServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/reset", handler.Reset)
	mux.HandleFunc("/balance", handler.GetBalance)
	mux.HandleFunc("/event", handler.Event)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &WebServer{
		server:  server,
		handler: handler,
	}
}

func (ws *WebServer) Start() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		log.Println("Starting server on port 8080...")
		if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ws.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
