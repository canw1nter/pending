package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MySQLDB *gorm.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MySQLDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Can't connect to mysql database! err: %s\n", err.Error())
	}

	err = MySQLDB.AutoMigrate(&User{})
	if err != nil {
		log.Printf("Can't auto migrate DB! err: %s\n", err.Error())
	}
}
