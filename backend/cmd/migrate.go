package main

import (
	"log"
	db_config "test-backend-sharing-vision/config"
	"test-backend-sharing-vision/model"
)

func main() {
	db_config.InitDB()
	db := db_config.GetDB()

	log.Println("Menjalankan migration...")
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		log.Fatal("Migration gagal: ", err)
	}
	log.Println("Migration sukses")
}
