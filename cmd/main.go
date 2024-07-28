package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mongo_migration "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/db/migrations/mongo"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	products_inmemorydb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/inmemorydb/products"
	users_inmemorydb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/inmemorydb/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/mongodb"
	products_mongodb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/mongodb/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb"
	products_postgresdb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb/products"
	users_postgresdb "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories/postgresdb/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")
		repo       = os.Getenv("DB_STORAGE")
	)

	var (
		productRepository repositories.IProductRepository
		userRepository    repositories.IUserRepository
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch repo {
	case "postgres":
		uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		conn, err := postgresdb.Connect(ctx, "postgres", uri)
		if err != nil {
			log.Fatalf("error to connect to database: %v", err)
		}

		productRepository = products_postgresdb.NewProductRepository(conn)
		userRepository = users_postgresdb.NewUserRepository(conn)

	case "mongo":
		uri := fmt.Sprintf("mongodb://%s:%s/%s", dbHost, dbPort, dbName)
		mongoClient, err := mongodb.Connect(ctx, uri, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("error to connect to database: %v", err)
		}

		migrate, err := mongo_migration.NewMongoMigration(mongoClient, dbName, "file://internal/db/migrations/mongo")
		if err != nil {
			log.Fatalf("error to init mongo migration: %v", err)
		}

		if err = migrate.Up(); err != nil {
			log.Fatalf("error to execute migrations: %v", err)
		}

		productsCollection := mongoClient.Database(dbName).Collection("products")
		productRepository = products_mongodb.NewProductRepository(productsCollection)

	case "in-memory":
		productRepository = products_inmemorydb.NewProductRepository()
		userRepository = users_inmemorydb.NewUserRepository()

	default:
		log.Fatal("invalid DB_STORAGE provided, must be one of [postgres, mongo, in-memory]")
	}

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Handlers
	productHandler := handlers.NewProductHandler(productRepository)
	userHandler := handlers.NewUserHandler(userRepository)

	app := gin.Default()
	routes.AddProductsRoutes(app, productHandler)
	routes.AddUserRoutes(app, userHandler)

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
