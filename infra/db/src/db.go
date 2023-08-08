package src

import (
	"context"
	"errors"

	"github.com/S4mkiel/validation/domain/entity"
	"gorm.io/gorm"
)

type UserSQLRepository struct {
	orm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserSQLRepository {
	return &UserSQLRepository{orm: db}
}

func (db UserSQLRepository) Create(ctx context.Context, u entity.User) (*entity.User, error) {
	err := db.orm.Create(&u).Error
	if err != nil {
		if db.orm.Where(entity.User{Email: u.Email}).Take(&entity.User{}).Error == nil {
			return nil, errors.New("user alredy exists")
		}
	}
	return &u, nil
}