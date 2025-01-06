package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func getAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Getting path failed: %v\n", err)
	}

	return absPath
}

func getDirName(path string) string {
	absPath := getAbsPath(path)

	dirName := filepath.Base(absPath)

	return dirName
}

func getFileName(absPath string) []os.DirEntry {
	files, err := os.ReadDir(absPath)
	if err != nil {
		log.Fatalf("Getting files name failed: %v\n", err)
	}

	return files
}

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
