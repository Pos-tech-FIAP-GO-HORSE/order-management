package factories

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/user"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

func MakeUserFactory(userRepository repositories.IUserRepository) *handlers.UserHandler {
	createUserUseCase := user.NewCreateUserUseCase(userRepository)

	return handlers.NewUserHandler(
		createUserUseCase,
	)
}
