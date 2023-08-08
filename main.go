package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/S4mkiel/validation/domain/service"
	"github.com/S4mkiel/validation/infra/db"
	"github.com/S4mkiel/validation/infra/env"
	"github.com/S4mkiel/validation/infra/http/fiber"
	logger "github.com/S4mkiel/validation/infra/log"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if os.Getenv("ENV") != "production" {
		LoadConfig()
	}
	fx.New(
		fiber.Module,
		env.Module,
		logger.Module,
		service.Module,
		db.Module,

	).Run()
}

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
