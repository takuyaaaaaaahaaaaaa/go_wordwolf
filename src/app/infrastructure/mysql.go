package infrastructure

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error){
	db, err = gorm.Open("mysql", os.Getenv("DB_USERNAME")+":" + os.Getenv("DB_USERPASS")+
    "@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nill {
		logrus.Fatalf("Error connect DB:%v", err)
	}

	return db, err
}
