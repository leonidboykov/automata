package module

import (
	"fmt"
	"time"

	lua "github.com/yuin/gopher-lua"
)

const timeLayout = "15:04"

func inTimeSpan(L *lua.LState) int {
	start, err := time.Parse(timeLayout, L.ToString(1))
	if err != nil {
		fmt.Println(err)
		return -1
	}

	end, err := time.Parse(timeLayout, L.ToString(2))
	if err != nil {
		fmt.Println(err)
		return -1
	}

	check, err := time.Parse(timeLayout, L.ToString(3))
	if err != nil {
		fmt.Println(err)
		return -1
	}

	res := lua.LBool(check.After(start) && check.Before(end))
	L.Push(res)

	return 1
}
