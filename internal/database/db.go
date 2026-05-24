package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"my-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
	}
	db.AutoMigrate(
		&models.Usuario{},
		&models.Empresa{},
	)

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Error obteniendo instancia SQL DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()

	if err != nil {
		log.Fatalf("Error haciendo ping a PostgreSQL: %v", err)
	}

	DB = db

	log.Println("Conexión exitosa a PostgreSQL")

	return db
}
