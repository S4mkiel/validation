package table

import (
	"github.com/S4mkiel/validation/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
	)
	return err
}
