package mongo

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/db/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMigration struct {
	migrateClient *migrate.Migrate
}

func NewMongoMigration(mongoClient *mongo.Client, databaseName string, filePath string) (migrations.IMigration, error) {
	driver, err := mongodb.WithInstance(mongoClient, &mongodb.Config{
		DatabaseName: databaseName,
	})
	if err != nil {
		return nil, err
	}

	migrate, err := migrate.NewWithDatabaseInstance(filePath, databaseName, driver)
	if err != nil {
		return nil, err
	}

	return &MongoMigration{
		migrateClient: migrate,
	}, nil
}

func (m *MongoMigration) Up() error {
	return m.migrateClient.Up()
}
func (m *MongoMigration) Down() error {
	return m.migrateClient.Down()
}
