package commons

import (
	"os"
	"path"
	"strings"

	"github.com/go-liquor/liquor/internal/message"
)

// ReplacePackage read firstDir and replace package
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
