package db

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_like/config"
	"github.com/ITu-CloudWeGo/itu_rpc_like/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
)

func init() {
	instance = createDB()
}

func createDB() *gorm.DB {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.PostgresSQL.Host,
		conf.PostgresSQL.User,
		conf.PostgresSQL.Password,
		conf.PostgresSQL.DBName,
		conf.PostgresSQL.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	err = db.AutoMigrate(&model.PidCount{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	return db
}

func GetDB() *gorm.DB {
	return instance
}
