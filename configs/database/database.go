package database

import (
	//"os"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is reusable gorm sql connection.
	DB *gorm.DB
)

func ConnectDB() error {
	h := "ktp-database"
	u := "mariadbuser"
	pwd := "exam_engine_password"
	p := "3306"
	d := "ktp-db"

	dsn := u + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)

	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = dbConnection
	return nil
}
