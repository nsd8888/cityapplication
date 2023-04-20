package Database

import (
	"fmt"
	"log"
        "cityapplication/Model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Getconnection() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", "nilesh", "Hemantnilesh88", "database-app-golang.cetjcletfp74.us-east-1.rds.amazonaws.com", "citydatabase")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: connection}), &gorm.Config{})
	if err != nil {
		log.Println("DB connection Failed")
	}

	log.Println("DB Connection established...")
	Db = db
}
func Migratedatabase() {
	Db.AutoMigrate(&Model.City_info_structs{})
}
