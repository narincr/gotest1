package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "simsupport"      // Username
	password = "4Superpoint*#77" // Password
	hostname = "localhost"       // Hostname
	port     = "3306"            // MySQL Port
	dbname   = "gotest"          // Database Name
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

func dsn(dbName string) string {
	// gentool -dsn "simsupport:4Superpoint*#77@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local" -db "mysql" -tables "Member" -fieldNullable "true"  -fieldWithIndexTag "true" -fieldWithTypeTag "true" -withUnitTest "false"
	dsnText := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, dbname)
	fmt.Println(dsnText)
	return dsnText
}

// SetupDB - call this method to connect db
func SetupDB() {

	database, err := gorm.Open(mysql.Open(dsn(dbname)), &gorm.Config{PrepareStmt: false})

	if err != nil {
		fmt.Println("connect database fail")
	}
	db = database
}
