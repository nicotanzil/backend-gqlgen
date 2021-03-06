package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const DB_NAME = "staem"
const DB_HOST = "127.0.0.1"
const DB_PORT = "5432"     //port on installation
const DB_USER = "postgres" //default is postgres
const DB_PASS = "password" //password on installation

func Connect() (*gorm.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL")
	fmt.Println(dbUrl)
	if dbUrl != "" {
		return gorm.Open(postgres.Open(dbUrl), &gorm.Config{
			PrepareStmt: true,
			Logger:      logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					Colorful:      true,
					LogLevel:      logger.Info,
				},
			),
		})
	} else {
		dsn := "host=127.0.0.1 user=postgres password=password dbname=staem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
}
