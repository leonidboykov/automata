package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cjoudrey/gluahttp"
	"github.com/yuin/gopher-lua"
	json "layeh.com/gopher-json"

	"github.com/smarthut/automata/module"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("json", json.Loader)
	L.PreloadModule("automata", module.Loader)

	// setting base path for scripts
	if err := L.DoString(`package.path = [[scripts/?.lua;]] .. package.path`); err != nil {
		fmt.Println(err)
	}

	if err := L.DoFile("scripts/main.lua"); err != nil {
		fmt.Println(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("start"), // name of Lua function
		Protect: true,                 // return err or panic
	}); err != nil {
		fmt.Println(err)
	}

	for {
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("update"), // name of Lua function
			Protect: true,                  // return err or panic
		}); err != nil {
			fmt.Println(err)
		}
		<-time.After(20 * time.Second) // edit this
	}
}
