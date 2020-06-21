package lib

import (
	"os/exec"
	"runtime"
	"strings"
)

var blackList = []string{
	"unreachable", "unknown", "100% packet loss", "failed", "failure", "100% loss", "timed out",
}

func Ping(addr string) bool {

	var cmd *exec.Cmd

	// todo review commands & support more systems & check for superuser rights

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", addr, "-n", "1", "-w", "2")
	case "linux":
		cmd = exec.Command("ping", addr, "-c", "1", "-w", "2")
	case "android":
		cmd = exec.Command("ping", addr, "-c", "1", "-w", "2")
	}

	if cmd == nil {
		panic("os is not supported")
	}

	out, _ := cmd.Output()
	for _, word := range blackList {
		if strings.Contains(string(out), word) {
			return false
		}
	}
	return true

}
