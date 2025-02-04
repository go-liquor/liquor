package commons

import (
	"os"
	"path"
	"strings"

	"github.com/go-liquor/liquor/internal/message"
	"golang.org/x/mod/modfile"
)

// ReplacePackage recursively traverses through directories starting from firstDir
// and replaces all occurrences of originalPackage with newPackage in all files.
// It updates the package declarations and import paths in Go files.
//
// Parameters:
//   - firstDir: The starting directory path for the recursive search
//   - originalPackage: The package name/path to be replaced
//   - newPackage: The new package name/path to replace with
//
// Returns:
//   - error: nil if successful, otherwise returns an error if directory reading fails
func ReplacePackage(firstDir string, originalPackage string, newPackage string) error {
	paths := []string{
		firstDir,
	}

	for p := 0; p < len(paths); p++ {
		dirname := paths[p]
		rd, err := os.ReadDir(dirname)
		if err != nil {
			return err
		}
		for _, file := range rd {
			filepath := path.Join(dirname, file.Name())
			if file.IsDir() {
				paths = append(paths, filepath)
				continue
			}
			content, err := os.ReadFile(filepath)
			if err != nil {
				message.Error("failed to read %v: %v", filepath, err)
				continue
			}
			contentStr := strings.ReplaceAll(string(content), originalPackage, newPackage)
			if err := os.WriteFile(filepath, []byte(contentStr), 0755); err != nil {
				message.Error("failed to writefile: %v", err)
				continue
			}
			message.Success("updated file %v", filepath)
		}
	}
	return nil
}

// GetModFile reads and parses the go.mod file from the specified directory.
// It returns the parsed module file that contains information about the Go module.
//
// Parameters:
//   - origin: The directory path containing the go.mod file
//
// Returns:
//   - *modfile.File: Parsed module file structure
//   - error: nil if successful, otherwise returns an error if reading or parsing fails
func GetModFile(origin string) (*modfile.File, error) {
	content, err := os.ReadFile(path.Join(origin, "go.mod"))
	if err != nil {
		return nil, err
	}
	modFile, err := modfile.Parse("go.mod", content, nil)
	return modFile, err
}
