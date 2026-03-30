package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductReprository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	//criando uma rota ver se retorna status ok
	// server.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//mapeando a rota para acessar os produtos
	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.CreateProduct)

	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")

}
