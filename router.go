package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	lua "github.com/yuin/gopher-lua"
)

// API is the main REST API
type API struct {
	ScriptEnabled bool  `json:"enabled"`
	Error         error `json:"error,omitempty"`
	luaState      *lua.LState
	handler       http.Handler
}

type state struct {
	ScriptEnabled bool `json:"enabled"`
}

// NewAPI instantiates a new REST API
func NewAPI(L *lua.LState) *API {
	api := &API{
		// TODO: add `RUN_SCRIPT_AT_LOAD` env config
		ScriptEnabled: true,
		luaState:      L,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/automata", func(r chi.Router) {
		r.Get("/state", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, api)
		})
		r.Get("/enable", func(w http.ResponseWriter, r *http.Request) {
			api.ScriptEnabled = true
			api.Error = api.callLuaMethod("onEnable")
			render.JSON(w, r, api)
		})
		r.Get("/disable", func(w http.ResponseWriter, r *http.Request) {
			api.ScriptEnabled = false
			api.Error = api.callLuaMethod("onDisable")
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

func (api API) callLuaMethod(method string) error {
	return api.luaState.CallByParam(lua.P{
		Fn:      api.luaState.GetGlobal(method), // name of Lua function
		Protect: true,                           // return err or panic
	})
}
