package VM

import (
	"crypto/md5"
	"encoding/hex"
	lua "github.com/yuin/gopher-lua"
	"io"
	"os"
	"runtime"
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

func md5sum(L *lua.LState) int {
	filename := L.Get(1).String()

	fl, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}

	m5 := md5.New()

	io.Copy(m5, fl)

	resbuf := m5.Sum(nil)
	resstr := hex.EncodeToString(resbuf)

	L.Push(lua.LString(resstr))
	return 1
}

func load(L *lua.LState) int {
	// 运行新的脚本
	L.SetGlobal("run", L.NewFunction(run))
	L.SetGlobal("md5sum", L.NewFunction(md5sum))
	// 常量---------------------------------
	L.SetGlobal("L_OS", lua.LString(runtime.GOOS))
	L.SetGlobal("L_ARCH", lua.LString(runtime.GOARCH))
	// 常量---------------------------------

	return 0
}
