package commons

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func PrintCode(code string) {
	code = fmt.Sprintf("```go\n%v\n```", code)
	out, _ := glamour.Render(code, "dark")
	fmt.Println(out)
}
