package commons

import "os"

// IsNotExist checks if a file or directory does not exist at the specified path.
//
// Parameters:
//   - p: The path to check
//
// Returns:
//   - bool: true if the path does not exist, false otherwise
func IsNotExist(p string) bool {
	_, err := os.Stat(p)
	return os.IsNotExist(err)
}

// IsExist checks if a file or directory exists at the specified path.
//
// Parameters:
//   - p: The path to check
//
// Returns:
//   - bool: true if the path exists, false otherwise
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
