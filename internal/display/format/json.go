package format

import (
	"encoding/json"
	"log"
	"os"
)

type Json struct {
	Directory string   `json:"directory"`
	Files     []string `json:"files"`
}

func DisplayJson(dirName string, files []os.DirEntry) {
	var filesNames []string
	for _, file := range files {
		filesNames = append(filesNames, file.Name())
	}

	dirInfo := &Json{
		Directory: dirName,
		Files:     filesNames,
	}
	jsonData, err := json.MarshalIndent(&dirInfo, "", "\t\t")
	if err != nil {
		log.Printf("Critical error: Failed to convert to json: %v", err)
	}

	os.Stdout.Write(jsonData)
}
