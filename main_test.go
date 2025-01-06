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

	absPath := getAbsPath(tempDir)
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

	expected := filepath.Base(getAbsPath(tempDir))

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

/**
func getSpecifiedExtFileName(absPath string, exts []string) []os.DirEntry {
	var filesFiltered []os.DirEntry
	files := getFileName(absPath)

	for _, file := range files {
		if containExt(file.Name(), exts) {
			filesFiltered = append(filesFiltered, file)
		}
	}

	return filesFiltered
}

func containExt(fileName string, exts []string) bool {
	for _, ext := range exts {
		if strings.Contains(fileName, ext) {
			return true
		}
	}

	return false
}

func display(dir string, exts []string) {
	var files []os.DirEntry

	if exts[0] == "" {
		files = getFileName(getAbsPath(dir))
	} else {
		files = getSpecifiedExtFileName(getAbsPath(dir), exts)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{getDirName(dir)})
	table.SetBorder(true)

	for _, file := range files {
		table.Append([]string{file.Name()})
	}

	table.Render()
}

func main() {
	dir := flag.String("dir", ".", "specify the directory name you want.")
	exts := flag.String("ext", "", "specify extentions you want. Separate by slash.")
	flag.Parse()

	extList := strings.Split(*exts, "/")

	display(*dir, extList)
}
**/
