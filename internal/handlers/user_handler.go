package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/user"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"net/http"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/create_user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase create_user.ICreateUserUseCase
}

func NewUserHandler(
	createUserRepository repositories.IUserRepository) *UserHandler {
	return &UserHandler{
		createUserUseCase: user.NewCreateUserUseCase(createUserRepository),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var newUser create_user.Input
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if err := h.createUserUseCase.Execute(ctx, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})

}
