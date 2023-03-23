package main

import (
    "log"
)

func Example() {
	results := SearchByCaption("计算机技术")
	log.Print(results) // [{TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]
	result, err := SearchByNotation("TP3")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)            // {TP3 计算技术、计算机技术}
	log.Print(result.String())   //  TP3 计算技术、计算机技术
	log.Print(result.TopClass()) // { T 工业技术}
	log.Print(result.Classes())  // [{T 工业技术} {TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]
}
