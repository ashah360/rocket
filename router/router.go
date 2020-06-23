package router

import (
	"github.com/ashah360/cngo/server/app"
	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}
