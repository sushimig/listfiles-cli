package display

import (
	"log"
	"os"

	"github.com/sushimig/listfiles-cli/internal/display/format"
	"github.com/sushimig/listfiles-cli/internal/utils"
)

func Display(dir string, exts []string, isJson bool, isMarkdown bool) {
	var (
		files []os.DirEntry
		err   error
	)

	absPath, err := utils.GetAbsPath(dir)
	if err != nil {
		log.Printf("Warning: %v. Using current directory instead.", err)
		absPath, err = utils.GetAbsPath(".")
		if err != nil {
			log.Fatalf("Critical error: Failed to get current directory: %v", err)
		}
	}

	if exts[0] == "" {
		files, err = utils.GetFileName(absPath)
	} else {
		files, err = utils.GetSpecifiedExtFileName(absPath, exts)
	}

	if err != nil {
		log.Printf("Error reading files from %s: %v", absPath, err)
		return
	}

	dirName := utils.GetDirName(dir)
	if dirName == "unknown" {
		log.Printf("Warning: Failed to get directory name for %s. Using 'unknown'.", dir)
	}

	if isJson {
		format.DisplayJson(dirName, files)
	} else if isMarkdown {
		format.DisplayText(dirName, files)
	} else {
		format.DisplayText(dirName, files)
	}
}
