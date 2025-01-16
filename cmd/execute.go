package cmd

import (
	"flag"
	"listfiles-cli/internal/display"
	"strings"
)

func Execute() {
	dir := flag.String("dir", ".", "specify the directory name you want.")
	exts := flag.String("ext", "", "specify extentions you want. Separate by slash.")
	flag.Parse()

	extList := strings.Split(*exts, "/")
	display.Display(*dir, extList)
}
