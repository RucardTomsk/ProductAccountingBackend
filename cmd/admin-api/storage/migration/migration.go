package migration

import (
	"gorm.io/gorm"
	"productAccounting-v1/internal/domain/entity"
	"strings"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Component{},
		&entity.Chapter{},
	)

	if err != nil {
		if strings.Contains(err.Error(), "42P07") {
			return nil
		}
		return err
	}

	return nil

}
