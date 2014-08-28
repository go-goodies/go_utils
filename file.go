package go_utils

import "os"

// FileExists returns true if the file exists
func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}
