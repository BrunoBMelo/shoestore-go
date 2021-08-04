package databases

import (
	"fmt"
	"os"
	"store/core/accounts/entities"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Load() *gorm.DB {

	db, err := gorm.Open(postgres.Open(os.Getenv("db_connection")), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed during connection database.")
		panic(err)
	}

	if enableMigrations, err := strconv.ParseBool(os.Getenv("enable_migrations")); err == nil && enableMigrations {
		if err := db.AutoMigrate(&entities.Account{}); err != nil {
			fmt.Println("Failed during migration executation.")
			panic(err)
		}
		fmt.Println("Migration executed has success.")
	}

	return db
}
