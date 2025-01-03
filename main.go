package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func displayDirName(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("Getting path failed: %v\n", err)
	}

	dirName := filepath.Base(absPath)

	fmt.Printf("%s\n", dirName)
	getFilesNameInDirectory(absPath)
}

func displayDirNameSpecifiedExt()

func getFilesNameInDirectory(absPath string) {
	files, err := os.ReadDir(absPath)
	if err != nil {
		log.Fatalf("Getting files name failed: %v\n", err)
	}

	for _, file := range files {
		fmt.Printf("%s\n", file)
	}
}

func main() {
	dir := flag.String("dir", ".", "specify a directory name you want")
	flag.Parse()

	displayDirName(*dir)
}
