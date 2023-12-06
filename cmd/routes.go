package main

import (
	"github.com/bmizerany/pat"
	"learning-web-socket/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
