package main

import (
	"net/http"
)

const (
	playersPath = "/players/"
)

func (p *PlayerServer) NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(playersPath, replyWithServer(p))
	return mux
}

func replyWithServer(p *PlayerServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
