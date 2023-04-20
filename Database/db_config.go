package Database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Getconnection() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", "admin", "Hemantnilesh88", "database-1.cetjcletfp74.us-east-1.rds.amazonaws.com", "database-1")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: connection}), &gorm.Config{})
	if err != nil {
		log.Println("DB connection Failed")
	}

	log.Println("DB Connection established...")
	Db = db
}
