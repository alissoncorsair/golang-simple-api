package controller

import (
	"net/http"
	"strconv"

	"github.com/alissoncorsair/golang-api/model"
	"github.com/alissoncorsair/golang-api/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	product, err := p.productUseCase.GetProductByID(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if product != nil {
		ctx.JSON(http.StatusOK, product)

		return
	}

	response := model.Response{
		Message: "product was not found",
	}

	ctx.JSON(http.StatusNotFound, response)

}
