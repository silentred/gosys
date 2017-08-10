package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	var err error
	cmdName := os.Args[1]

	cmd := exec.Command(cmdName, os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
}
