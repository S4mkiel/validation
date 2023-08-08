package src

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("src", fx.Provide(NewSQLSourcers))

type Sources struct {
	UserSQL *UserSQLRepository
}

func NewSQLSourcers(db *gorm.DB) *Sources {
	var src = Sources{
		UserSQL: NewUserRepository(db),
	}
	return &src
}