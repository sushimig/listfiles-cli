package format

import (
	"fmt"
	"os"
)

func DisplayMarkdown(dirName string, files []os.DirEntry) {
	table := "| " + dirName + " |\n|------|\n"

	for _, file := range files {
		table += fmt.Sprintf("| %s |\n", file.Name())
	}

	fmt.Println(table)
}
