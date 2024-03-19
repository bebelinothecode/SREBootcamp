package db

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB;



func Connect()  {
	dsn := os.Getenv("dsn")
	
	d, err := gorm.Open("mysql",dsn);
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	DB = d;
}

func GetDB() *gorm.DB {
	return DB;
}