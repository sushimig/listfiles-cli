package cmd

import (
	"flag"
	"strings"

	"github.com/sushimig/listfiles-cli/internal/display"
)

func Execute() {
	dir := flag.String("dir", ".", "specify the directory name you want.")
	exts := flag.String("ext", "", "specify extentions you want. Separate by slash.")
	isJson := flag.Bool("json", false, "select ture or false if you want to convert to json.")
	isMarkdown := flag.Bool("md", false, "select ture or false if you want to convert to markdown.")
	flag.Parse()

	extList := strings.Split(*exts, "/")
	display.Display(*dir, extList, *isJson, *isMarkdown)
}
