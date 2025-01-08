package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func getAbsPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %v", err)
	}

	return absPath, nil
}

func getDirName(path string) string {
	absPath, err := getAbsPath(path)
	if err != nil {
		return "unkonow"
	}

	return filepath.Base(absPath)
}

func getFileName(absPath string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	return files, nil
}

func getSpecifiedExtFileName(absPath string, exts []string) ([]os.DirEntry, error) {
	var filesFiltered []os.DirEntry
	files, err := getFileName(absPath)
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

func display(dir string, exts []string) {
	var (
		files []os.DirEntry
		err   error
	)

	absPath, err := getAbsPath(dir)
	if err != nil {
		log.Printf("Warning: %v. Using current directory instead.", err)
		absPath, err = getAbsPath(".")
		if err != nil {
			log.Fatalf("Critical error: Failed to get current directory: %v", err)
		}
	}

	if exts[0] == "" {
		files, err = getFileName(absPath)
	} else {
		files, err = getSpecifiedExtFileName(absPath, exts)
	}

	if err != nil {
		log.Printf("Error reading files from %s: %v", absPath, err)
		return
	}

	dirName := getDirName(dir)
	if dirName == "unknown" {
		log.Printf("Warning: Failed to get directory name for %s. Using 'unknown'.", dir)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{dirName})
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
