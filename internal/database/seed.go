package database

import (
	"fmt"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(
		&model.MenuItem{},
		&model.Order{},
		&model.ProductOrder{},
		&model.User{},
	)

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("Seeding db data...")
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
