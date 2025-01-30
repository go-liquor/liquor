package commons

import (
	"os"
	"os/exec"
)

// Command execute a new command
func Command(command string, args ...string) error {
	cm := exec.Command(command, args...)
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	cm.Stdin = os.Stdin
	return cm.Run()
}
