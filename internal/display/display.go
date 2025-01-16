package display

import (
	"listfiles-cli/internal/dirutils"
	"listfiles-cli/internal/fileutils"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Display(dir string, exts []string) {
	var (
		files []os.DirEntry
		err   error
	)

	absPath, err := dirutils.GetAbsPath(dir)
	if err != nil {
		log.Printf("Warning: %v. Using current directory instead.", err)
		absPath, err = dirutils.GetAbsPath(".")
		if err != nil {
			log.Fatalf("Critical error: Failed to get current directory: %v", err)
		}
	}

	if exts[0] == "" {
		files, err = fileutils.GetFileName(absPath)
	} else {
		files, err = fileutils.GetSpecifiedExtFileName(absPath, exts)
	}

	if err != nil {
		log.Printf("Error reading files from %s: %v", absPath, err)
		return
	}

	dirName := dirutils.GetDirName(dir)
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
