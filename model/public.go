package model

type Category struct {
    Id       uint   `json:"id"`
    Pid      uint   `json:"pid"`
    Notation string `json:"sn"`
    Caption  string `json:"title"`
}
