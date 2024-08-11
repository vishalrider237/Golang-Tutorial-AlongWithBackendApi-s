package CrudWithMySql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DbConnect() {
	dsn := "root:vishal9743@tcp(127.0.0.1:3306)/employee?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected!")

	// Auto migrate the Employee model
	DB.AutoMigrate(Employee{})
}
