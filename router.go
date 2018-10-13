package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// API is the main REST API
type API struct {
	ScriptEnabled bool `json:"enabled"`
	handler       http.Handler
}

type state struct {
	ScriptEnabled bool `json:"enabled"`
}

// NewAPI instantiates a new REST API
func NewAPI() *API {
	api := &API{ScriptEnabled: false}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/automata", func(r chi.Router) {
		r.Get("/state", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, api)
		})
		r.Post("/state", func(w http.ResponseWriter, r *http.Request) {
			var s state
			if err := render.DecodeJSON(r.Body, &s); err != nil {
				log.Println(err)
				return
			}
			api.ScriptEnabled = s.ScriptEnabled
			render.JSON(w, r, api)
		})
	})

	api.handler = r

	return api
}

// Start starts API at address
func (api API) Start(addr string) {
	http.ListenAndServe(addr, api.handler)
}
