package commons

import "os"

// IsNotExist return true if path or file not exists
func IsNotExist(p string) bool {
	_, err := os.Stat(p)
	return os.IsNotExist(err)
}

// IsExists return true if path or file exists
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
