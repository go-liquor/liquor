package message

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
)

// Error print message error
func Error(format string, a ...any) {
	fmt.Println(aurora.Red(fmt.Sprintf(format, a...)))
}

// Fatal print message error with exit 1
func Fatal(format string, a ...any) {
	Error(format, a...)
	os.Exit(1)
}

// Success print message success
func Success(format string, a ...any) {
	fmt.Println(aurora.Green(fmt.Sprintf(format, a...)))
}

// Info print message info
func Info(format string, a ...any) {
	fmt.Println(aurora.Blue(fmt.Sprintf(format, a...)))
}

// Warning print message warning
func Warning(format string, a ...any) {
	fmt.Println(aurora.Yellow(fmt.Sprintf(format, a...)))
}
