package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cjoudrey/gluahttp"
	"github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"

	"github.com/smarthut/automata/module"
)

// TODO: add this to env config
const (
	defaultHost = ""
	defaultPort = 8080
	pollingTime = 5 * time.Minute
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

var scriptsPath = "scripts"

func main() {
	if os.Getenv("AUTOMATA_SCRIPTS_PATH") != "" {
		scriptsPath = os.Getenv("AUTOMATA_SCRIPTS_PATH")
	}

	L := lua.NewState(lua.Options{SkipOpenLibs: false})
	// lua.OpenBase(L)
	defer L.Close()

	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("json", luajson.Loader)
	L.PreloadModule("automata", module.Loader)

	// setting base path for scripts
	if err := L.DoString(`package.path = [[` + scriptsPath + `/?.lua;]] .. package.path`); err != nil {
		log.Println(err)
	}

	if err := L.DoFile(scriptsPath + "/main.lua"); err != nil {
		log.Println(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("start"), // name of Lua function
		Protect: true,                 // return err or panic
	}); err != nil {
		log.Println(err)
	}

	go startPolling(L, pollingTime)

	api := NewAPI()
	l := fmt.Sprintf("%s:%d", defaultHost, defaultPort)
	log.Printf("Starting SmartHut Automata %s on %s\n", version, l)
	api.Start(l)
}

func startPolling(L *lua.LState, t time.Duration) {
	for {
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("update"), // name of Lua function
			Protect: true,                  // return err or panic
		}); err != nil {
			log.Println(err)
		}
		<-time.After(t)
	}
}
