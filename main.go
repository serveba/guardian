package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const roleAdmin = "ADMIN"
const roleUser = "USER"

// User Model
type User struct {
	ID          uint
	Login       string `gorm:"size:50"`
	Password    string `gorm:"size:64"`
	Fullname    string `gorm:"size:75"`
	Role        string `gorm:"size:75"`
	Appenders   []Appender
	Datasources []Datasource
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Appender Model
type Appender struct {
	ID uint
	// User          User
	Metadatas     []Metadata
	Notifications []Notification
	Description   string `gorm:"size:200"`
	Type          string `gorm:"size:50"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Notification Model
type Notification struct {
	ID          uint
	CreatedAt   time.Time
	Notified    bool
	Appender    Appender
	Healthcheck Healthcheck
}

// Healthcheck Model
type Healthcheck struct {
	ID            uint
	CreatedAt     time.Time
	Result        bool
	Datasource    Datasource
	Notifications []Notification
}

// Metadata Model
type Metadata struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Datasource Datasource
	Appender   Appender
	Key        string `gorm:"size:50"`
	Value      string `gorm:"size:1000"`
}

// Datasource Model
type Datasource struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	// User             User
	Metadatas        []Metadata
	Description      string `gorm:"size:200"`
	Type             string `gorm:"size:50"`
	PollingFrequency int
	PollingType      string `gorm:"size:50"`
}

func checkPassword(login, password string) {
	byteArray := sha256.Sum256([]byte(password))
	encryptedPwd := string(byteArray[:32])
	fmt.Printf("encriptedPwd: %x", encryptedPwd)

	// db.Where("login = ? AND password = ?", login, encryptedPwd).Find(&user)
	// TODO
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "db/guardian.db")
	if err != nil {
		return nil, err
	}

	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)

	// relationships
	// db.Model(&User{}).Related(&[]Appender{})
	// db.Model(&User{}).Related(&[]Datasource{})

	// db.Model(&Appender{}).Related(&[]Metadata{})
	// db.Model(&Appender{}).Related(&[]Notification{})

	// db.Model(&Notification{}).Related(&Appender{})
	// db.Model(&Notification{}).Related(&Healthcheck{})

	// db.Model(&Healthcheck{}).Related(&Datasource{})
	// db.Model(&Healthcheck{}).Related(&[]Notification{})

	// db.Model(&Metadata{}).Related(&Datasource{})
	// db.Model(&Metadata{}).Related(&Appender{})

	// db.Model(&Datasource{}).Related(&[]Metadata{})
	// db.Model(&Datasource{}).Related(&User{})

	db.AutoMigrate(&User{}, &Metadata{}, &Appender{}, &Notification{},
		&Healthcheck{}, &Datasource{})

	db.LogMode(true)
	return db, nil
}

func main() {

	byteArray := sha256.Sum256([]byte("12345"))

	var user = User{
		Login:    "serveba@gmail.com",
		Password: string(byteArray[:32]),
		Fullname: "Sergio Velasco",
		Role:     roleAdmin}

	db, err := initDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.Where("").Delete(User{})
	db.Create(&user)

	datasources := [2]Datasource{}

	datasources[0] = Datasource{
		Description:      "Service1 Prod",
		Type:             "REST",
		PollingFrequency: 20,
		PollingType:      "SECONDS"}

	datasources[1] = Datasource{
		Description:      "Service1 Test",
		Type:             "REST",
		PollingFrequency: 20,
		PollingType:      "SECONDS"}

	db.Model(&user).Related(&datasources)

	fmt.Printf("main executed")
}
