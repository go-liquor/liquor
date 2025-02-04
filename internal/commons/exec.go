package commons

import (
	"os"
	"os/exec"
)

// Command executes a shell command in the specified directory with the given arguments.
// It pipes the standard output, standard error, and standard input to the current process.
//
// Parameters:
//   - dir: The working directory where the command will be executed
//   - command: The name or path of the command to execute
//   - args: Variable number of arguments to pass to the command
//
// Returns:
//   - error: nil if the command executes successfully, otherwise returns the error
func Command(dir string, command string, args ...string) error {
	cm := exec.Command(command, args...)
	cm.Dir = dir
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	cm.Stdin = os.Stdin
	return cm.Run()
}
