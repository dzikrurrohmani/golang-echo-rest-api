package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	seedDB(db)

	return db
}
