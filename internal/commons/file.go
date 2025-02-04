package commons

import (
	"strings"

	"github.com/golang-cz/textcase"
)

// ToFilename makes name to filename
func ToFilename(name string, adds ...string) string {
	return textcase.SnakeCase(name) + strings.Join(adds, "") + ".go"
}
