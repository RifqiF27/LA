package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen membersihkan layar terminal
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // untuk Windows
	} else {
		cmd = exec.Command("clear") // untuk Linux dan macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
