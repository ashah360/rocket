package app

import "net/http"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenxt-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.Write([]byte("Handler index"))
}
