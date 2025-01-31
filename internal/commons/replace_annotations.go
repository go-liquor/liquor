package commons

import (
	"fmt"
	"os"
	"strings"
)

// ReplaceAnnotations replaces occurrences of a given annotation in a file with a new value,
// while keeping the annotation intact for future modifications.
//
// It reads the file content, replaces all instances of the annotation by inserting the value
// before each occurrence, and writes the updated content back to the file.
//
// Parameters:
//   - file: The path to the file where the replacements will be made.
//   - annotation: The string annotation that should be replaced by inserting a new value before it.
//   - value: The string value to be inserted before each annotation occurrence.
//
// Returns:
//   - An error if reading or writing to the file fails, otherwise nil.
//
// Example Usage:
//
//	err := ReplaceAnnotations("config.txt", "//go:liquor:migrate", "New Value")
//	if err != nil {
//	    log.Fatal(err)
//	}
func ReplaceAnnotations(file string, annotation string, value string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	contentStr := strings.ReplaceAll(string(content), annotation, fmt.Sprintf("%v\n%v", value, annotation))
	return os.WriteFile(file, []byte(contentStr), 0755)
}
