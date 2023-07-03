package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func Connect() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	switch Dbdriver {
	case "sqlsvr":
		SqlsvrDev(DbUser, DbPassword, DbHost, DbName, Dbdriver)
	case "mysql":
		MysqlDev(DbUser, DbPassword, DbHost, DbPort, DbName, Dbdriver)
	}

	return Db
}

func MysqlDev(DbUser string, DbPassword string, DbHost string, DbPort string, DbName string, Dbdriver string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("Database Connection Error")
	}

}

func SqlsvrDev(DbUser string, DbPassword string, DbHost string, DbName string, Dbdriver string) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable&connection+timeout=30", DbUser, DbPassword, DbHost, DbName)
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("Database Connection Error")
	}
}
