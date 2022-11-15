package database

import (
	"fmt"
	"goroutine-optimize/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClientDb() *gorm.DB {
	db_host := "localhost"
	db_user := "postgres"
	db_password := "dua_lima25"
	db_dbname := "mini_project_test"
	db_port := "5432"

	fmt.Println("Host", db_dbname)

	// set dsn
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_dbname, db_port)
	// connect to db postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if there error connect to db, show error
	if err != nil {
		logger.Fatal("error connect to database : " + err.Error())
		return nil
	}

	// if connect success give info
	logger.Info("connect to database success")
	return db
}
