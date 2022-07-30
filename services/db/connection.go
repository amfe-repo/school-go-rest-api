package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "192.168.64.129"
	user     = "user"
	password = "root"
	dbname   = "dev_school_system"
	port     = "5432"
)

var DB *gorm.DB

func generateDBUri() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
}

func ConnectDB() {
	var err error

	DB, err = gorm.Open(postgres.Open(generateDBUri()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected!")
}
