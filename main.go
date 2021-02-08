package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		return
	}
	t, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("Couldn't read time, first arg should be time:", err)
		return
	}
	timeout := time.Duration(t) * time.Second

	args = args[1:]
	if len(args) <= 0 {
		fmt.Println("No commmand found")
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	time.AfterFunc(timeout, func() {
		fmt.Println("Command reached timeout")
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(res) > 0 {
		fmt.Println(">", string(res))
	}
	fmt.Println("DONE")
}
