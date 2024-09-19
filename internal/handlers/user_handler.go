package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/create_user"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/find_user_by_cpf"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/user"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UserHandler struct {
	createUserUseCase    create_user.ICreateUserUseCase
	findUserByCpfUseCase find_user_by_cpf.IFindUserByCpfUseCase
}

func NewUserHandler(
	createUserRepository repositories.IUserRepository) *UserHandler {
	return &UserHandler{
		createUserUseCase:    user.NewCreateUserUseCase(createUserRepository),
		findUserByCpfUseCase: user.NewFindUserByCpfUseCase(createUserRepository),
	}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Add a new user to the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user    body      create_user.Input  true  "User Data"
// @Success      201     {object}  ResponseMessage
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/users [post]
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

// FindUserByCpf godoc
// @Summary      Find user by CPF
// @Description  Retrieve a user by their CPF
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        cpf     path      string  true  "CPF"
// @Success      200     {object}  find_user_by_cpf.User
// @Failure      400     {object}  ResponseMessage
// @Failure      500     {object}  ResponseMessage
// @Router       /api/v1/users/{cpf} [get]
func (h *UserHandler) FindUserByCpf(c *gin.Context) {

	var input find_user_by_cpf.Input

	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.findUserByCpfUseCase.Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
