package json

import (
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"decode": func(state *lua.LState) int {

		return 1
	},
}

func Load(L *lua.LState) int {
	tb := L.NewTable()
	L.SetFuncs(tb, exports)

	L.Push(tb)
	return 1
}
