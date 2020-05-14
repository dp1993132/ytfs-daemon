package lmos

import (
	lua "github.com/yuin/gopher-lua"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

func kill(L *lua.LState) int {
	pidstr := L.Get(1).String()
	pid, err := strconv.ParseInt(pidstr, 10, 32)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	err = syscall.Kill(int(pid), 9)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

var exports = map[string]lua.LGFunction{
	"kill": kill,
	"killSelf": func(state *lua.LState) int {
		fl, err := os.OpenFile(".pid", os.O_RDONLY, 0644)
		if err == nil {
			pidbuf, err := ioutil.ReadAll(fl)
			defer fl.Close()
			if err == nil {
				pid, err := strconv.ParseUint(string(pidbuf), 10, 64)
				if err == nil {
					syscall.Kill(int(pid), 9)
				}
			}
		}
		return 0
	},
}

func Load(L *lua.LState) int {
	tb := L.NewTable()
	L.SetFuncs(tb, exports)

	L.Push(tb)
	return 1
}
