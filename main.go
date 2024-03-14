package main

import (
	"second-assignment/controller"
	"second-assignment/lib"
	"second-assignment/model"
	"second-assignment/repository"
	"second-assignment/route"

	"github.com/gin-gonic/gin"
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

	ginEngine := gin.Default()
	route.NewOrderController(*controller.NewOrderController(repository.NewOrderRepository(db))).OrderRouter(ginEngine)

	err = ginEngine.Run("localhost:8000")

	if err != nil {
		panic(err)
	}

}