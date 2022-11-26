package database

import (
	"fmt"

	"github.com/jgcaceres97/goly/app/model"
	"github.com/jgcaceres97/goly/app/settings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		*settings.DB.User,
		*settings.DB.Password,
		*settings.DB.Host,
		*settings.DB.Port,
		*settings.DB.Name,
	)

	conn, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)

	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	DB = conn
	conn.AutoMigrate(&model.Goly{})
}
