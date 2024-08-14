package webserver

import (
	"go-ebanx/internal/infra/web"
	"net/http"
)

type WebServer struct {
	handler *web.Handler
}

func NewWebServer(handler *web.Handler) *WebServer {
	return &WebServer{handler: handler}
}

func (ws *WebServer) Start() {
	http.HandleFunc("/reset", ws.handler.Reset)
	http.HandleFunc("/balance", ws.handler.GetBalance)
	http.HandleFunc("/event", ws.handler.Event)
	http.ListenAndServe(":8080", nil)
}
