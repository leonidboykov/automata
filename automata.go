package main

import (
	"fmt"

	"github.com/smarthut/automata/module"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("automata", module.Loader)
	if err := L.DoFile("scripts/main.lua"); err != nil {
		fmt.Println(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("start"), // name of Lua function
		Protect: true,                 // return err or panic
	}); err != nil {
		fmt.Println(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("update"), // name of Lua function
		Protect: true,                  // return err or panic
	}); err != nil {
		fmt.Println(err)
	}
}
