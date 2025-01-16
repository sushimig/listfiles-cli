package format

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func DisplayText(dirName string, files []os.DirEntry) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{dirName})
	table.SetBorder(true)

	for _, file := range files {
		table.Append([]string{file.Name()})
	}

	table.Render()
}
