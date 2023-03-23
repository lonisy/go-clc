package tools

import (
    "encoding/csv"
    "os"
)


var Idata [][]string

func SaveTocsv() {
    //data := [][]string{
    //    {"Name", "Age", "Gender"},
    //    {"John", "25", "Male"},
    //    {"Alice", "30", "Female"},
    //    {"Bob", "40", "Male"},
    //}

    file, err := os.Create("data.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()
    for _, row := range Idata {
        err := writer.Write(row)
        if err != nil {
            panic(err)
        }
    }
}
