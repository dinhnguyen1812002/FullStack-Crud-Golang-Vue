package config

import (
	"github.com/dinhnguyen1812002/go-crud/Entity"
	"github.com/dinhnguyen1812002/go-crud/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:123456789@tcp(localhost:3306)/Book?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })
	helper.ErrorPanic(err)

    Database.AutoMigrate(&entity.Book{})

    return nil
}