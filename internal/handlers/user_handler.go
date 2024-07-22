package handlers

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/create_user"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	createUserUseCase create_user.ICreateUserUseCase
}

func NewUserHandler(
	createUserUseCase create_user.ICreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var newUser create_user.NewUser
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
