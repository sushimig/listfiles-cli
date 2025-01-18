package utils

import (
	"fmt"
	"path/filepath"
)

func GetAbsPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %v", err)
	}

	return absPath, nil
}

func GetDirName(path string) string {
	absPath, err := GetAbsPath(path)
	if err != nil {
		return "unkonow"
	}

	return filepath.Base(absPath)
}
