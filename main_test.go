
/***
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetAbsPath(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	expected, _ := filepath.Abs(tempDir)

	absPath, err := getAbsPath(tempDir)
	if err != nil {
		t.Errorf("failed to get absolute path: %v", err)
	}
	if absPath != expected {
		t.Errorf("Expected %s, but got %s", expected, absPath)
	}
}

func TestGetDirName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	absPath, err := getAbsPath(tempDir)
	if err != nil {
		t.Errorf("failed to get absolute path: %v", err)
	}
	expected := filepath.Base(absPath)

	dirName := getDirName(tempDir)
	if dirName != expected {
		t.Errorf("Expected %s, but got %s", expected, dirName)
	}
}

func TestGetFileName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tempDir, "file.go"), []byte("package main"), 0644)

	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Getting files name failed: %v\n", err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, but got %d", len(files))
	}
}

func TestGetSpecifiedExtFileName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("text file"), 0644)
	os.WriteFile(filepath.Join(tempDir, "file.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tempDir, "image.jpg"), []byte("jpg image"), 0644)
	os.WriteFile(filepath.Join(tempDir, "README.md"), []byte("markdown"), 0644)

	exts := []string{".txt", ".go"}
	result, err := getSpecifiedExtFileName(tempDir, exts)
	if err != nil {
		t.Errorf("Error reading files from %s: %v", tempDir, err)
	}

	expectedFiles := map[string]bool{
		"file.txt": true,
		"file.go":  true,
	}

	for _, file := range result {
		if !expectedFiles[file.Name()] {
			t.Errorf("Unexpected file: %s", file.Name())
		}
		delete(expectedFiles, file.Name())
	}

	if len(expectedFiles) != 0 {
		t.Errorf("Some expected files were not found: %v", expectedFiles)
	}
}

func TestContainExt(t *testing.T) {
	tests := []struct {
		fileName string
		exts     []string
		expected bool
	}{
		{"example.txt", []string{".txt"}, true},
		{"example.txt", []string{".t"}, false},
		{"example.txt", []string{".exe"}, false},
		{"example.jpg", []string{".png", ".jpg"}, true},
		{"example.txt", []string{}, false},
		{"file.go", []string{".go"}, true},
	}

	for _, tt := range tests {
		result := containExt(tt.fileName, tt.exts)
		if result != tt.expected {
			t.Errorf("containExt(%q, %v) = %v, want %v",
				tt.fileName, tt.exts, result, tt.expected)
		}
	}
}

func TestDisplay(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	os.WriteFile(filepath.Join(tempDir, "file.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("sample text"), 0644)
	os.WriteFile(filepath.Join(tempDir, "README.md"), []byte("# Readme"), 0644)

	exts := []string{".go", ".md"}
	display(tempDir, exts)
}
