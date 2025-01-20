package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sushimig/listfiles-cli/internal/utils"
)

func TestGetAbsPath(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exmaple")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	expected, _ := filepath.Abs(tempDir)

	absPath, err := utils.GetAbsPath(tempDir)
	if err != nil {
		t.Errorf("failed to read directory: %v", err)
	}
	if absPath != expected {
		t.Errorf("failed to get absolute path: %v", err)
	}
}

func TestNotGetAbsPath(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exmaple")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	expected, _ := filepath.Abs(tempDir)

	diffDir := tempDir + "-different"

	absPath, err := utils.GetAbsPath(diffDir)
	if err != nil {
		t.Errorf("failed to read directory: %v", err)
	}
	if absPath == expected {
		t.Errorf("failed to get absolute path: %v", err)
	}
}

func TestGetDirName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exmaple")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	absPath, err := utils.GetAbsPath(tempDir)
	if err != nil {
		t.Errorf("failed to get absolute path: %v", err)
	}
	expected := filepath.Base(absPath)

	dirName := utils.GetDirName(tempDir)
	if dirName != expected {
		t.Errorf("Expected %s, but got %s", expected, dirName)
	}
}

func TestNotGetDirName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "exmaple")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	absPath, err := utils.GetAbsPath(tempDir)
	if err != nil {
		t.Errorf("failed to get absolute path: %v", err)
	}
	expected := filepath.Base(absPath)

	diffDir := tempDir + "-different"

	dirName := utils.GetDirName(diffDir)
	if dirName == expected {
		t.Errorf("Expected %s, but got %s", dirName, expected)
	}
}
