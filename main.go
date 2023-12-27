package main

import (
	"cs-backend-test1/config"
	"cs-backend-test1/internal/converter"
	"cs-backend-test1/internal/migrator"
	"cs-backend-test1/internal/model"
	"cs-backend-test1/internal/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadAndParse()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.Logger.LogLevel))

	userConverter := new(converter.UserConverter)
	m := migrator.NewUserMigrator(userConverter)

	sourceStorage := storage.NewDatabase("cs_old", &model.User{})
	if err = sourceStorage.OpenWithConfig(&cfg.Db); err != nil {
		log.Fatal().Err(err).Send()
	}
	sourceStorage.SetMigrator(m)

	destStorage := storage.NewDatabase("cs_new", &model.UserAccount{})
	if err = destStorage.OpenWithConfig(&cfg.Db); err != nil {
		log.Fatal().Err(err).Send()
	}

	var users []model.User
	err = sourceStorage.MigrateDataBatches(destStorage, &users, 5)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
