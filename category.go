package main

import (
	"github.com/spf13/cast"
	"go-clc/class"
	"go-clc/model"
    "go-clc/tools"
)

var Categorys []model.Category
var CateId uint



func initCategorys() {
    tools.Idata = append(tools.Idata, []string{
        cast.ToString(makeCateId()),
        "0",
        "",
        "",
        "中图法分类",
    })
	for _, clc := range TopClass {
		Categorys = append(Categorys, model.Category{
			Id:       makeCateId(),
			Pid:      0,
			Notation: clc.notation,
			Caption:  clc.caption,
		})

	}
	class.Categorys = Categorys

	for _, category := range class.Categorys {
		tools.Idata = append(tools.Idata, []string{
			cast.ToString(category.Id),
			cast.ToString(category.Pid),
			"",
			cast.ToString(category.Notation),
			cast.ToString(category.Caption),
		})
	}
}

func makeCateId() uint {
	CateId = CateId + 1
	class.CateId = CateId
	return CateId
}
