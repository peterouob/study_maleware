package main

import (
	bgf "bgf/path"
	"flag"
)

func main() {
	var pathName string
	flag.StringVar(&pathName, "pathName", ".", "輸入查詢路徑")
	flag.Parse()
	bgf.PrintFilePath(pathName)
}
