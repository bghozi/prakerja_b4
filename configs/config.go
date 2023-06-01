package configs

import (
	"fmt"
	"os"
	"prakerja_b4/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("gagal loading .env file")
	// }
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectDatabase() {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&models.User{})
}
