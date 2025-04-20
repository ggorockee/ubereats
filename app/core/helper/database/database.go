package database

import (
	"fmt"
	"log"
	"net/url"
	"ubereats/app/core/entity"
	"ubereats/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(config *config.Config) *gorm.DB {
	gormConfig := newGormConfig()

	values := url.Values{}
	values.Add("charset", "utf8mb4")
	values.Add("parseTime", "true")
	query := values.Encode()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.Infra.DB.User,
		config.Infra.DB.Password,
		config.Infra.DB.Host,
		config.Infra.DB.Port,
		config.Infra.DB.DBName,
		query,
	)

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	db.AutoMigrate(

		new(entity.User),
		new(entity.Category),
		new(entity.Restaurant),
	)

	return db
}

func newGormConfig() *gorm.Config {
	return &gorm.Config{}
}
