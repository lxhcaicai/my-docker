package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 注: 运行时需要 root 权限。
func main() {
	cmd := exec.Command("bash")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWUTS,
		// Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
