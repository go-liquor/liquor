package stdout

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

// Success prints a success message to stdout with a green checkmark emoji.
// The message can include format specifiers that will be replaced with the provided arguments.
//
// Parameters:
//   - format: a format string that follows fmt.Printf conventions
//   - a: variadic arguments to be formatted according to the format string
//
// Example:
//
//	stdout.Success("Server started on port %d", 8080)
func Success(format string, a ...any) {
	fmt.Printf("✅ %v\n", aurora.Green(fmt.Sprintf(format, a...)))
}

// Error prints an error message to stdout with a red prohibition emoji.
// The message can include format specifiers that will be replaced with the provided arguments.
//
// Parameters:
//   - format: a format string that follows fmt.Printf conventions
//   - a: variadic arguments to be formatted according to the format string
//
// Example:
//
//	stdout.Error("Failed to connect to database: %s", err)
func Error(format string, a ...any) {
	fmt.Printf("🚫 %v\n", aurora.Red(fmt.Sprintf(format, a...)))
}
