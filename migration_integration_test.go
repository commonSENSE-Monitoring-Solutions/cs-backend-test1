package main

import (
	"cs-backend-test1/config"
	"cs-backend-test1/internal/converter"
	"cs-backend-test1/internal/migrator"
	"cs-backend-test1/internal/model"
	"cs-backend-test1/internal/storage"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Migrator(t *testing.T) {
	cfg, err := config.LoadAndParse()
	if err != nil {
		assert.NoError(t, err)
		return
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.Logger.LogLevel))

	userConverter := new(converter.UserConverter)
	m := migrator.NewUserMigrator(userConverter)

	sourceStorage := storage.NewDatabase("cs_old", &model.User{})
	err = sourceStorage.OpenWithConfig(&cfg.Db)
	if err != nil {
		assert.NoError(t, err)
	}
	sourceStorage.SetMigrator(m)

	destStorage := storage.NewDatabase("cs_new", &model.UserAccount{})
	err = destStorage.OpenWithConfig(&cfg.Db)
	if err != nil {
		assert.NoError(t, err)
		return
	}

	err = destStorage.DeleteAll()
	if err != nil {
		assert.NoError(t, err)
		return
	}

	var users []model.User
	err = sourceStorage.MigrateDataBatches(destStorage, &users, 2)
	assert.NoError(t, err)
}
