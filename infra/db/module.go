package db

import (
	"github.com/S4mkiel/validation/domain/repository"
	"github.com/S4mkiel/validation/infra/db/src"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	PostgresModule,
	src.Module,
	fx.Provide(func(s *src.Sources) repository.UserRepository { return s.UserSQL }),
)