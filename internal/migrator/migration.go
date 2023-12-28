package migrator

import (
	"cs-backend-test1/internal/converter"
	"cs-backend-test1/internal/model"
	"cs-backend-test1/internal/storage"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"sync"
)

type UserMigrator struct {
	converter converter.Converter
	mu        sync.Mutex
}

func NewUserMigrator(converter converter.Converter) *UserMigrator {
	return &UserMigrator{converter: converter}
}

func (um *UserMigrator) Migrate(dest storage.Storage, records interface{}) error {
	var wg sync.WaitGroup
	var userAccounts []model.UserAccount

	users, ok := records.(*[]model.User)
	if !ok {
		err := errors.New("invalid records type for migration")
		return fmt.Errorf("%w: %T", err, records)
	}

	for _, user := range *users {
		user := user

		wg.Add(1)

		go func() {
			defer wg.Done()

			result, err := um.converter.Convert(user)
			if err != nil {
				log.Error().Err(err).Send()
				return
			}

			userAccount := result.(model.UserAccount)

			um.mu.Lock()
			userAccounts = append(userAccounts, userAccount)
			um.mu.Unlock()
		}()
	}

	wg.Wait()

	// create a transaction and insert record to db
	return dest.Write(&userAccounts)
}
