package db_config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()

	user := getEnv("DB_USER")
	pass := getEnv("DB_PASS")
	host := getEnv("DB_HOST")
	port := getEnv("DB_PORT")
	name := getEnv("DB_NAME")

	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port)
	sqlDB, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		log.Fatal("Gagal terkoneksi ke MySQL server: ", err)
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		log.Fatal("Gagal membuat database: ", err)
	}
	log.Println("Database sudah ada:", name)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database: ", err)
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}
