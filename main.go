package main

import (
	"github.com/marchellodev/sharic/cmd"
	"os/exec"
	"runtime"
)

func main() {
	// setting english language in the terminal
	switch runtime.GOOS {
	case "windows":
		exec.Command("chcp", "437").Output()
	case "linux":
		exec.Command("export", "LC_ALL=C").Output()
	}

	cmd.Execute()
}
