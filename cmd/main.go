package main

import (
	"github.com/alissoncorsair/golang-api/controller"
	"github.com/alissoncorsair/golang-api/db"
	"github.com/alissoncorsair/golang-api/repository"
	"github.com/alissoncorsair/golang-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	repository := repository.NewProductRepository(dbConnection)
	usecase := usecase.NewProductUseCase(repository)
	ProductController := controller.NewProductController(usecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductByID)

	server.Run(":8000")
}
