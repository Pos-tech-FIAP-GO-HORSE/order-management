package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/create_product"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/delete_product_by_id"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_all_products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_product_by_category"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_product_by_id"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/update_product"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/update_product_availability"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	createProductUseCase             create_product.ICreateProductUseCase
	findAllProductsUseCase           find_all_products.IFindAllProducts
	findProductByIDUseCase           find_product_by_id.IFindProductByID
	findProductByCategoryUseCase     find_product_by_category.IFindProductByCategory
	updateProductUseCase             update_product.IUpdateProductUseCase
	updateProductAvailabilityUseCase update_product_availability.IUpdateProductAvailabilityUseCase
	deleteProductUseCase             delete_product_by_id.IDeleteProductByIDUseCase
}

type ResponseMessage struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewProductHandler(productRepository repositories.IProductRepository) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:             products.NewCreateProductUseCase(productRepository),
		findAllProductsUseCase:           products.NewFindAllProductsUseCase(productRepository),
		findProductByIDUseCase:           products.NewFindProductByIDUseCase(productRepository),
		findProductByCategoryUseCase:     products.NewFindProductByCategoryUseCase(productRepository),
		updateProductUseCase:             products.NewUpdateProductUseCase(productRepository),
		updateProductAvailabilityUseCase: products.NewUpdateProductAvailabilityUseCase(productRepository),
		deleteProductUseCase:             products.NewDeleteProductByIDUseCase(productRepository),
	}
}

// CreateProduct godoc
// @Summary      Create new Product
// @Description  Add a new product to the inventory
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      create_product.Input  true  "Product Data"
// @Success      201      {object}  ResponseMessage
// @Failure      400      {object}  ResponseMessage
// @Failure      500      {object}  ResponseMessage
// @Router       /api/v1/products [post]
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

// FindAllProducts godoc
// @Summary      Get all products
// @Description  Retrieve a list of all products in the inventory
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        query  query     find_all_products.Input  false  "Query Parameters"
// @Success      200    {array}   find_all_products.Product
// @Failure      400    {object}  ResponseMessage
// @Failure      500    {object}  ResponseMessage
// @Router       /api/v1/products [get]
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

// FindProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieve a product by its unique ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "Product ID"
// @Success      200    {object}  find_all_products.Product
// @Failure      400    {object}  ResponseMessage
// @Failure      500    {object}  ResponseMessage
// @Router       /api/v1/products/{id} [get]
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

// FindProductByCategory godoc
// @Summary      Get products by category
// @Description  Retrieve products by a specific category
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        category  path      string  true  "Product Category"
// @Success      200       {array}   find_all_products.Product
// @Failure      400       {object}  ResponseMessage
// @Failure      500       {object}  ResponseMessage
// @Router       /api/v1/products/category/{category} [get]
func (h *ProductHandler) FindProductByCategory(c *gin.Context) {
	var input find_product_by_category.Input
	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	products, err := h.findProductByCategoryUseCase.Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary      Update product details
// @Description  Update the details of an existing product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "Product ID"
// @Param        product body      update_product.Input  true  "Updated Product Data"
// @Success      200     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/products/{id} [patch]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var input update_product.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

// UpdateProductAvailability godoc
// @Summary      Update product availability
// @Description  Update the availability status of an existing product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id         path      string                               true  "Product ID"
// @Success      200        {object}  ResponseMessage
// @Failure      400        {object}  ResponseMessage
// @Failure      500        {object}  ResponseMessage
// @Router       /api/v1/products/{id}/availability [patch]
func (h *ProductHandler) UpdateProductAvailability(c *gin.Context) {
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

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by its unique ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "Product ID"
// @Success      200    {object}  ResponseMessage
// @Failure      400    {object}  ResponseMessage
// @Failure      500    {object}  ResponseMessage
// @Router       /api/v1/products/{id} [delete]
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
