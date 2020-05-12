package cmd

import (
	lua "github.com/yuin/gopher-lua"
	"os"
	"os/exec"
	"strings"
)

func parseCmd(L *lua.LState) *exec.Cmd {
	cmdstr := L.Get(1).String()

	cmdArgs := strings.Split(cmdstr, " ")
	if len(cmdArgs) == 0 {
		return nil
	}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = os.Environ()
	return cmd
}

func Command(L *lua.LState) int {
	cmd := parseCmd(L)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdTb := L.NewTable()

	L.SetFuncs(cmdTb, map[string]lua.LGFunction{
		"run": func(state *lua.LState) int {
			err := cmd.Run()
			if err != nil {
				state.Push(lua.LString(err.Error()))
				return 1
			}

			return 0
		},
		"setStdout": func(state *lua.LState) int {
			flpath := state.Get(1).String()
			if flpath != "" {
				fl, err := os.OpenFile(flpath, os.O_CREATE|os.O_WRONLY, 0644)
				if err == nil {
					cmd.Stdout = fl
				}
			}
			return 0
		},
		"setSterr": func(state *lua.LState) int {
			flpath := state.Get(1).String()
			if flpath != "" {
				fl, err := os.OpenFile(flpath, os.O_CREATE|os.O_WRONLY, 0644)
				if err == nil {
					cmd.Stderr = fl
				}
			}
			return 0
		},
		"setStdin": func(state *lua.LState) int {
			flpath := state.Get(1).String()
			if flpath != "" {
				fl, err := os.OpenFile(flpath, os.O_CREATE|os.O_WRONLY, 0644)
				if err == nil {
					cmd.Stdin = fl
				}
			}
			return 0
		},
		"pid": func(state *lua.LState) int {
			state.Push(lua.LString(cmd.Process.Pid))
			return 1
		},
		"kill": func(state *lua.LState) int {
			err := cmd.Process.Kill()
			if err != nil {
				L.Push(lua.LString(err.Error()))
				return 1
			}
			return 0
		},
	})

	L.Push(cmdTb)

	return 1
}

func ExecCmd(L *lua.LState) int {
	cmd := parseCmd(L)

	buf, err := cmd.Output()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(buf))
	L.Push(lua.LNil)
	return 2
}

var exports = map[string]lua.LGFunction{
	"command": Command,
	"exec":    ExecCmd,
}

func Load(L *lua.LState) int {
	tb := L.NewTable()
	L.SetFuncs(tb, exports)

	L.Push(tb)
	return 1
}
