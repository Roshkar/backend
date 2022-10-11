package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-rest-shop-backend/pkg/database"
	"github.com/golang-rest-shop-backend/pkg/handler"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"

	_ "github.com/golang-rest-shop-backend/pkg/swagger"
)

// @title           Golang Rest Shop Backend
// @version         1.0
// @description     This is a sample rest backend for an online shop.
// @accept json
// @produce json

// @contact.name   Alexandar Naydenov
// @contact.email  alexandar.naydenov99@gmail.com

// @host      localhost:8080

func init() {
	err := database.InitMySqlConnection()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/product", handler.GetAllProductHandler)
	r.GET("/order", handler.GetAllOrdersHandler)
	r.GET("/product/:productId", handler.GetProductHandler)
	r.GET("/order/:orderId", handler.GetOrderHandler)

	r.POST("/order", handler.AddOrderHandler)
	r.POST("/product", handler.AddProductHandler)

	r.PUT("/order/:orderId", handler.UpdateOrderHandler)
	r.PUT("/product/:productId", handler.UpdateProductHandler)

	r.DELETE("/delete/product/:productId", handler.DeleteProductHandler)
	r.DELETE("/delete/order/:orderId", handler.DeleteOrderHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Listening to port 8080...")
	log.Fatal(r.Run(":8080"))
}
