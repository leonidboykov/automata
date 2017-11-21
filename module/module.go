package module

import (
	"github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"inTimeSpan": inTimeSpan,
}

// Loader loads lua module
func Loader(L *lua.LState) int {
	m := L.SetFuncs(L.NewTable(), exports)
	L.SetField(m, "name", lua.LString("value"))
	L.Push(m)
	return 1
}
