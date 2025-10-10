package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {

	items := []models.Item{
		{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
		{ID: 2, Name: "商品２", Price: 2000, Description: "説明２", SoldOut: true},
		{ID: 3, Name: "商品３", Price: 3000, Description: "説明３", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
