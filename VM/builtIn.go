package VM

import (
	lua "github.com/yuin/gopher-lua"
)

func run(L *lua.LState) int {
	shellPath := L.Get(1).String()
	if shellPath == "" {
		L.Push(lua.LString("shellPath is nil"))
		return 1
	}

	vm := GetVM()
	err := vm.DoFile(shellPath)
	PutVM(vm)
	if err != nil {
		L.Push(lua.LString(err.Error()))
	}

	return 0
}

func load(L *lua.LState) int {
	L.SetGlobal("run", L.NewFunction(run))
	return 0
}
