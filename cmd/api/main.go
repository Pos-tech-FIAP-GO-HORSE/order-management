package main

import (
	"fmt"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/db/db_gorm"
	inmemorydb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/inmemorydb/products"
	postgresdb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb/user"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/factories"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	var (
		repo = "gorm"
	)

	var (
		productRepository repositories.IProductRepository
		userRepository    repositories.IUserRepository
	)

	switch repo {
	case "gorm":

		db_gorm.ConectaComBancoDeDados()

	case "in-memory":

		productRepository = inmemorydb.NewProductRepository()

	default:
		log.Fatal("invalid DB_PROVIDER provided, must be one of [postgres, in-memory]")
	}

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Repositories
	productRepository = postgresdb.NewProductRepository()
	userRepository = user.NewUserRepository()

	// Factories
	productHandler := factories.MakeProductFactory(productRepository)
	userHandler := factories.MakeUserFactory(userRepository)

	app := gin.Default()
	routes.AddProductsRoutes(app, productHandler)
	routes.AddUseRoutes(app, userHandler)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("[x] HTTP server is running on port %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error to listen and serve: %v", err)
	}
}
