package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	inmemorydb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/inmemorydb/products"
	postgresdb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	var (
		repo = os.Getenv("DB_STORAGE")
		conn *sql.DB
		err  error
	)

	var (
		productRepository repositories.IProductRepository
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch repo {
	case "postgres":
		var (
			dbUser     = os.Getenv("DB_USER")
			dbPassword = os.Getenv("DB_PASSWORD")
			dbHost     = os.Getenv("DB_HOST")
			dbPort     = os.Getenv("DB_PORT")
			dbName     = os.Getenv("DB_NAME")
		)
		uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		conn, err = repositories.Connect(ctx, "postgres", uri)
		if err != nil {
			log.Fatalf("error to connect to database: %v", err)
		}

		productRepository = postgresdb.NewProductRepository(conn)

	case "in-memory":
		productRepository = inmemorydb.NewProductRepository()

	default:
		log.Fatal("invalid DB_PROVIDER provided, must be one of [postgres, in-memory]")
	}

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Handlers
	productHandler := handlers.NewProductHandler(productRepository)

	app := gin.Default()
	routes.AddProductsRoutes(app, productHandler)

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
