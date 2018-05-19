package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const roleAdmin = "ADMIN"
const roleUser = "USER"

// User Model
type User struct {
	gorm.Model
	Login    string "gorm:size:50"
	Password string "gorm:size:50"
	Fullname string "gorm:size:75"
	Role     string "gorm:size:75"
}

// Datasource Model
type Datasource struct {
	gorm.Model
	User             User
	Description      string "gorm:size:200"
	Type             string "gorm:size:50"
	PollingFrequency int
	PollingType      string ""
}

func checkPassword(login, password string) {

}

func main() {
	db, err := gorm.Open("sqlite3", "db/guardian.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	var user = User{
		Login:    "serveba@gmail.com",
		Password: "12345",
		Fullname: "Sergio Velasco",
		Role:     roleAdmin}

	db.Create(&user)
}
