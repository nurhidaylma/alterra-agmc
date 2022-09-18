package config

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := getMySQLDSN()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}

func getMySQLDSN() string {
	return fmt.Sprintf(GetValue(DATABASE_CONNECTION_STRING), GetValue(DATABASE_USER), GetValue(DATABASE_PASS), GetValue(DATABASE_HOST), GetValue(DATABASE_PORT), GetValue(DATABASE_NAME))
}
