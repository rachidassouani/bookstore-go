package configs

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	result, err := gorm.Open("mysql", "rachid:password/bookstore?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}
	db = result
}

func GetDb() *gorm.DB {
	return db
}
