package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileName(absPath string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	return files, nil
}

func GetSpecifiedExtFileName(absPath string, exts []string) ([]os.DirEntry, error) {
	var filesFiltered []os.DirEntry
	files, err := GetFileName(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file name: %v", err)
	}

	for _, file := range files {
		if containExt(file.Name(), exts) {
			filesFiltered = append(filesFiltered, file)
		}
	}

	return filesFiltered, nil
}

func containExt(fileName string, exts []string) bool {
	fileExt := filepath.Ext(fileName)
	for _, ext := range exts {
		if fileExt == ext {
			return true
		}
	}

	return false
}
