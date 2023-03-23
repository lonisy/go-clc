package main

import (
    "flag"
    "go-clc/class"
    "go-clc/tools"
)

var dir string
var debug bool

func main() {
	flag.StringVar(&dir, "dir", "data", "directory to export")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()
	initCategorys()
	class.ExportCategoryList(false)
    tools.SaveTocsv()

    // class.ExportJSON(dir, debug)
	// fmt.Println(topClass("A").notation)
	// fmt.Println(topClass("A").caption)
}




