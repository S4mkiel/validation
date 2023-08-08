package repository

import (
	"context"

	"github.com/S4mkiel/validation/domain/entity"
)


type UserRepository interface {
	Create (context.Context, entity.User) (*entity.User, error)
}