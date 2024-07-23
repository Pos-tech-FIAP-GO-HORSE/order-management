package db_gorm

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	/*var (
		dbUser     = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbHost     = os.Getenv("POSTGRES_DB")
		dbPort     = os.Getenv("POSTGRES_PORT")
		dbName     = os.Getenv("DB_NAME")
	)*/

	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Deu ruim ao conectar")
	}
	DB.AutoMigrate(&models.Product{}, &models.User{})

}
