package module

import (
	"fmt"
	"time"

	"github.com/yuin/gopher-lua"
)

const timeLayout = "15:04"

func inTimeSpan(L *lua.LState) int {
	start, err := time.Parse(timeLayout, L.ToString(1))
	if err != nil {
		fmt.Println(err)
	}

	end, err := time.Parse(timeLayout, L.ToString(2))
	if err != nil {
		fmt.Println(err)
	}

	check, err := time.Parse(timeLayout, L.ToString(3))
	if err != nil {
		fmt.Println(err)
	}

	res := lua.LBool(check.After(start) && check.Before(end))
	L.Push(res)

	return 1
}

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
