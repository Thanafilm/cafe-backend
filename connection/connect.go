package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"cafe-backend/model"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(db:3306)/cafe?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("cant connect DB")
	}
	DB = db
	db.AutoMigrate(&model.User{}, &model.Food{})
}
