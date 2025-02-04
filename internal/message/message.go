package message

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
)

// Error prints a message in red color to indicate an error condition.
// It accepts a format string and optional arguments that will be formatted
// according to the format specifier.
//
// Parameters:
//   - format: A format string that follows fmt.Printf conventions
//   - a: Optional arguments to be formatted into the string
func Error(format string, a ...any) {
	fmt.Println(aurora.Red(fmt.Sprintf(format, a...)))
}

// Fatal prints an error message in red color and terminates the program
// with exit code 1. Use this function for unrecoverable errors.
//
// Parameters:
//   - format: A format string that follows fmt.Printf conventions
//   - a: Optional arguments to be formatted into the string
func Fatal(format string, a ...any) {
	Error(format, a...)
	os.Exit(1)
}

// Success prints a message in green color to indicate successful completion
// of an operation.
//
// Parameters:
//   - format: A format string that follows fmt.Printf conventions
//   - a: Optional arguments to be formatted into the string
func Success(format string, a ...any) {
	fmt.Println(aurora.Green(fmt.Sprintf(format, a...)))
}

// Info prints a message in bold blue color to provide informational content
// to the user.
//
// Parameters:
//   - format: A format string that follows fmt.Printf conventions
//   - a: Optional arguments to be formatted into the string
func Info(format string, a ...any) {
	fmt.Println(aurora.Bold(aurora.Blue(fmt.Sprintf(format, a...))))
}

// Warning prints a message in yellow color to indicate a warning condition
// that requires attention but is not necessarily an error.
//
// Parameters:
//   - format: A format string that follows fmt.Printf conventions
//   - a: Optional arguments to be formatted into the string
func Warning(format string, a ...any) {
	fmt.Println(aurora.Yellow(fmt.Sprintf(format, a...)))
}
