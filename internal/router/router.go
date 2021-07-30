package router

import (
	"day5/internal/Ws"
	"day5/internal/ctrl"
	"net/http"
)

func WebsokectRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctrl.ServerWs(w, r, Ws.H)
	})
}
