package main

import (
	"second-assignment/lib"
	"second-assignment/model"
)

func main() {
	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Order{},&model.Item{});
	if err != nil {
		panic(err)
	}

}