package util

import (
	"os"
	"path/filepath"
)

// returns the current exec path
func GetCurrentExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	return exPath
}
