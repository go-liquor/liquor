package commons

import (
	"strings"

	"github.com/golang-cz/textcase"
)

// ToFilename converts a given name into a valid Go filename by converting it to snake_case
// and appending optional additional strings and the ".go" extension.
//
// Parameters:
//   - name: The base name to be converted to a filename
//   - adds: Optional variadic parameter of additional strings to append before the extension
//
// Returns:
//   - string: The resulting filename in the format "snake_case[additions].go"
func ToFilename(name string, adds ...string) string {
	return textcase.SnakeCase(name) + strings.Join(adds, "") + ".go"
}
