package commons

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

// PrintCode formats and displays Go code with syntax highlighting in the terminal.
// It wraps the provided code in Go markdown syntax and renders it using the
// glamour dark theme.
//
// Parameters:
//   - code: A string containing the Go code to be displayed
func PrintCode(code string) {
	code = fmt.Sprintf("```go\n%v\n```", code)
	out, _ := glamour.Render(code, "dark")
	fmt.Println(out)
}
