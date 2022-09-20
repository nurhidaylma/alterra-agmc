package config

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	c "github.com/nurhidaylma/alterra-agmc/day-3/constants"
	"github.com/nurhidaylma/alterra-agmc/day-3/models"
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
	return fmt.Sprintf(GetValue(c.DATABASE_CONNECTION_STRING), GetValue(c.DATABASE_USER), GetValue(c.DATABASE_PASS), GetValue(c.DATABASE_HOST), GetValue(c.DATABASE_PORT), GetValue(c.DATABASE_NAME))
}
