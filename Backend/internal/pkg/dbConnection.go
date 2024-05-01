package pkg

import (
	"fmt"

	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBaseInstance *gorm.DB

func DBConnection(conf models.DBConfig) (db *gorm.DB, err error) {

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", conf.DATABASE_HOST, conf.POSTGRES_USER, conf.POSTGRES_PASSWORD, conf.POSTGRES_DB, conf.DATABASE_PORT, conf.DATABASE_SSLMODE, conf.DATABASE_TIMEZONE)

	if DataBaseInstance, err = gorm.Open(postgres.Open(dns), &gorm.Config{}); err != nil {
		return db, err
	}

	return db, nil
}
