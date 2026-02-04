package database

import (
	"fmt"
	"log"

	"github.com/roshith/uber-fx/internals/config"
	"github.com/roshith/uber-fx/internals/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// psql database connection
func PostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	log.Println("Postgres SQL connected")
	return db, nil
}

// mysql commection
func MysqlConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		cfg.MY_USER, cfg.MY_PASS, cfg.MY_HOST, cfg.MY_PORT, cfg.MY_DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	log.Println("MySQL connected")
	return db, nil
}
