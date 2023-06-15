package handler

import (
	"assignment-2/database"
	"assignment-2/handler/httphandler"
	"assignment-2/repository/itemrepository/item_pg"
	"assignment-2/repository/orderrepository/order_pg"
	"assignment-2/service"
	"log"

	"github.com/gin-gonic/gin"
)

func StartAssignment2() {
	db := database.GetPostgresInstance()

	itemRepo := item_pg.NewItemPG(db)
	orderRepo := order_pg.NewOrderPG(db, itemRepo)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := httphandler.NewOrderHandler(orderService)

	r := gin.Default()

	// Create
	r.POST("/orders", orderHandler.Create)
	// Get All
	r.GET("/orders", orderHandler.GetAll)
	// Get One
	r.GET("/orders/:orderID", orderHandler.GetOrderByID)
	// Update
	r.PATCH("/orders/:orderID", orderHandler.UpdateByID)
	// Delete
	r.DELETE("/orders/:orderID", orderHandler.DeleteByID)
	

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
