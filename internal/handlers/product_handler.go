package handlers

import (
	"context"
	"net/http"
	"time"

	createproduct "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/create_product"
	deleteproductbyid "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/delete_product_by_id"
	findallproducts "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/find_all_products"
	findproductbyid "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/find_product_by_id"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	createProductUseCase   createproduct.ICreateProductUseCase
	findAllProductsUseCase findallproducts.IFindAllProducts
	findProductByIDUseCase findproductbyid.IFindProductByID
	deleteProductUseCase   deleteproductbyid.IDeleteProductByIDUseCase
}

func NewProductHandler(
	createProductUseCase createproduct.ICreateProductUseCase,
	findAllProductsUseCase findallproducts.IFindAllProducts,
	findProductByIDUseCase findproductbyid.IFindProductByID,
	deleteProductUseCase deleteproductbyid.IDeleteProductByIDUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:   createProductUseCase,
		findAllProductsUseCase: findAllProductsUseCase,
		findProductByIDUseCase: findProductByIDUseCase,
		deleteProductUseCase:   deleteProductUseCase,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input createproduct.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.createProductUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
	})
}

func (h *ProductHandler) FindAllProducts(c *gin.Context) {
	var input findallproducts.Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	products, err := h.findAllProductsUseCase.Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FindProductByID(c *gin.Context) {
	var input findproductbyid.Input
	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	product, err := h.findProductByIDUseCase.Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var input deleteproductbyid.Input
	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.deleteProductUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
