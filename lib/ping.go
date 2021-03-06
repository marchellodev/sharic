package lib

import (
	"fmt"
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
		cmd = exec.Command("ping", addr, "-n", "1", "-w", "1")
	case "linux":
		cmd = exec.Command("ping", addr, "-c", "1", "-w", "1")
	case "android":
		cmd = exec.Command("/system/bin/ping", "-c", "1", "-w", "1", addr)
	case "darwin":
		cmd = exec.Command("ping", addr, "-c", "1", "-W", "1")
	}

	if cmd == nil {
		fmt.Println("os is not supported, please file an issue")
		return false
	}

	out, err := cmd.Output()
	if err != nil {
		return false
	}

	for _, word := range blackList {
		if strings.Contains(strings.ToLower(string(out)), word) {
			return false
		}
	}
	return true

}
