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
	//h := "0.0.0.0"
	u := "mariadbuser"
	pwd := "exam_engine_password"
	//p := "5000"
	d := "ktp-db"

	dsn := u + ":" + pwd + "@tcp(" + "ktp-database" + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)

	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = dbConnection
	return nil
}
