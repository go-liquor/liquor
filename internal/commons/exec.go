package commons

import (
	"os"
	"os/exec"
)

// Command execute a new command
func Command(dir string, command string, args ...string) error {
	cm := exec.Command(command, args...)
	cm.Dir = dir
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	cm.Stdin = os.Stdin
	return cm.Run()
}
