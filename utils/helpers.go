package utils

import (
	"os"
)

// FileExists checks if the file exists
func FileExists(file_path string) bool {
	i, err := os.Stat(file_path)
	if os.IsNotExist(err) || i.IsDir() {
		return false
	}
	return true
}

// IsDir checks if the path is a directory
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	
	return file.IsDir()
}