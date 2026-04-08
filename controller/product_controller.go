package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//recebe a requisiçao
//trata a entrada da requisicao e retorna

type ProductController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {

	return ProductController{
		productUseCase: usecase,
	}

}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)

}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product) //pega o json que recebemos e transforma no product
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Product Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Product Id needs to be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found in DB",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Corpo vazio ou inválido: " + err.Error()})
		return
	}

	product.ID = productId

	updatedProduct, err := p.productUseCase.UpdateProduct(productId, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductController) DeleteProduct (ctx *gin.Context) {
	id := ctx.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}
