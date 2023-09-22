package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    if DB != nil {
        return
    }

    database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
    if err != nil {
        log.Fatalf("Gagal terhubung ke database: %v", err)
    }

    if err := database.AutoMigrate(&Product{}, &User{}); err != nil {
        log.Fatalf("Gagal melakukan migrasi: %v", err)
    }

    DB = database
}
