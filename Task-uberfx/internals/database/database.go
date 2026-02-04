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
func PostgresDB(cfg config.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	log.Println("Postgress Sqlm connected")
	return db, nil
}

// mysql commection

func MysqlConnection(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.MYSQL_DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	log.Println("Mysql db connected")
	return db, err
}
