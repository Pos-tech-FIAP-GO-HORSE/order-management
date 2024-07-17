package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/create_product"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/delete_product_by_id"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/find_all_products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/find_product_by_id"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/update_product"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/update_product_availability"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	createProductUseCase             create_product.ICreateProductUseCase
	findAllProductsUseCase           find_all_products.IFindAllProducts
	findProductByIDUseCase           find_product_by_id.IFindProductByID
	updateProductUseCase             update_product.IUpdateProductUseCase
	updateProductAvailabilityUseCase update_product_availability.IUpdateProductAvailabilityUseCase
	deleteProductUseCase             delete_product_by_id.IDeleteProductByIDUseCase
}

func NewProductHandler(
	createProductUseCase create_product.ICreateProductUseCase,
	findAllProductsUseCase find_all_products.IFindAllProducts,
	findProductByIDUseCase find_product_by_id.IFindProductByID,
	updateProductUseCase update_product.IUpdateProductUseCase,
	updateProductAvailabilityUseCase update_product_availability.IUpdateProductAvailabilityUseCase,
	deleteProductUseCase delete_product_by_id.IDeleteProductByIDUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:             createProductUseCase,
		findAllProductsUseCase:           findAllProductsUseCase,
		findProductByIDUseCase:           findProductByIDUseCase,
		updateProductUseCase:             updateProductUseCase,
		updateProductAvailabilityUseCase: updateProductAvailabilityUseCase,
		deleteProductUseCase:             deleteProductUseCase,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input create_product.Input
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
	var input find_all_products.Input
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
	var input find_product_by_id.Input
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

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	convertedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var input update_product.Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.ID = convertedID

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.updateProductUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product updated successfully",
	})
}

func (h *ProductHandler) UpdateProductAvalability(c *gin.Context) {
	var input update_product_availability.Input
	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.updateProductAvailabilityUseCase.Execute(ctx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product availability updated successfully",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var input delete_product_by_id.Input
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
