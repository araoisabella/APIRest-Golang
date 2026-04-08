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

	//retorna a lista de produtos em formato json
	server.GET("/products", ProductController.GetProducts)

	//cria um novo produto
	server.POST("/product", ProductController.CreateProduct)

	//retorna um produto pelo id
	server.GET("/product/:productId", ProductController.GetProductById)

	//atualiza um produto pelo id
	server.PUT("/product/:productId", ProductController.UpdateProduct)

	//deleta um produto pelo id
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")

}
