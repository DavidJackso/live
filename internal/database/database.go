package database

import (
	"fmt"
	"live/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDD(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Address,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)
	fmt.Print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
