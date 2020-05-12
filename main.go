package main

import (
	"flag"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
	"os/exec"
	"path"
	"yottachain/ytfs-daemon/VM"
)

var isDaemon bool

var FileLogger = &lumberjack.Logger{
	Filename:   path.Join("output.log"),
	MaxSize:    128,
	Compress:   false,
	MaxAge:     7,
	MaxBackups: 30,
}

func main() {
	flag.BoolVar(&isDaemon, "d", false, "是否以守护进程启动")
	flag.Parse()

	if isDaemon {
		log.Println("日志文件:output.log")
		log.SetOutput(FileLogger)
		log.Println("守护进程已启动")
		for {
			cmd := exec.Command(os.Args[0])
			cmd.Stdout = log.Writer()
			cmd.Stderr = log.Writer()
			cmd.Env = os.Environ()
			cmd.Run()
		}
		log.Println("守护进程退出")
	} else {
		VM.Run("update.lua", "boot.lua")
	}
}
