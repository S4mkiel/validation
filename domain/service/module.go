package service

import (
	"github.com/S4mkiel/validation/domain/service/user"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(service.NewUserService),
)