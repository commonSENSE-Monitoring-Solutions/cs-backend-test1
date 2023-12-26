package migrator

import (
	"cs-backend-test1/config"
	"cs-backend-test1/internal/model"
	"cs-backend-test1/internal/storage"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync/atomic"
	"testing"
	"time"
)

type MockDB struct {
	mock.Mock
	migratedUsersCount atomic.Uint64
}

func (db *MockDB) SetMigrator(migrator storage.Migrator) {}

func (db *MockDB) SyncSchema(schema interface{}) error {
	return nil
}

func (db *MockDB) MigrateWithDataBatches(dest storage.Storage, records interface{}, size int) error {
	return nil
}

func (db *MockDB) IsSessionActive() bool {
	return true
}

func (db *MockDB) Close() error {
	return nil
}

func (db *MockDB) OpenWithConfig(cfg *config.Database) error {
	return nil
}

func (db *MockDB) ReadAll(dest interface{}) error {
	args := db.Called(dest)
	return args.Error(0)
}

func (db *MockDB) Write(record interface{}) error {
	args := db.Called(record)

	err := args.Error(0)

	if err == nil {
		db.migratedUsersCount.Add(1)
	}

	return err
}

type MockConverter struct {
	mock.Mock
}

func (mc *MockConverter) convertToTimeUnix(t *int64) *time.Time {
	if t == nil {
		return nil
	}

	timeUnix := time.Unix(*t, 0)

	return &timeUnix
}

func (mc *MockConverter) Convert(val interface{}) (interface{}, error) {
	args := mc.Called(val)

	switch v := val.(type) {
	case model.User:
		return model.UserAccount{
			OldID:     int64(v.ID),
			Email:     v.EmailAddress,
			Password:  v.Password,
			CreatedAt: mc.convertToTimeUnix(v.CreatedAt),
		}, args.Error(0)
	}

	return nil, errors.New("invalid passed data type")
}

type MigratorSuite struct {
	suite.Suite
	dest      *MockDB
	converter *MockConverter
}

func (suite *MigratorSuite) SetupTest() {
	suite.dest = new(MockDB)
	suite.converter = new(MockConverter)
}

func (suite *MigratorSuite) AfterTest(suiteName, testName string) {
	suite.dest.AssertExpectations(suite.T())
	suite.converter.AssertExpectations(suite.T())
}

func (suite *MigratorSuite) TestMigrateDataWithNoEmailDuplicates() {
	testTimestamp := time.Now().Unix()

	users := []model.User{
		{
			ID:           1, // primary key
			EmailAddress: "test@domain1.com",
			Password:     "test_password",
			CreatedAt:    &testTimestamp,
		},
		{
			ID:           2,
			EmailAddress: "test@domain2.com",
			Password:     "test_password",
			CreatedAt:    &testTimestamp,
		},
	}

	for _, user := range users {
		suite.converter.On("Convert", user).Return(nil)
	}

	suite.dest.On("Write", mock.Anything).Return(nil)

	userMigrator := UserMigrator{
		converter: suite.converter,
	}

	err := userMigrator.Migrate(suite.dest, &users)

	assert.NoError(suite.T(), err)
}

func (suite *MigratorSuite) TestMigrateDataWithErrorOnDuplicatedEmail() {
	testTimestamp := time.Now().Unix()

	users := []model.User{
		{
			ID:           1, // primary key
			EmailAddress: "duplicate@domain.com",
			Password:     "test_password",
			CreatedAt:    &testTimestamp,
		},
		{
			ID:           2,
			EmailAddress: "duplicate@domain.com",
			Password:     "test_password",
			CreatedAt:    &testTimestamp,
		},
	}

	for _, user := range users {
		suite.converter.On("Convert", user).Return(nil)
	}

	suite.dest.On("Write", mock.Anything).Return(errors.New("duplicates error"))

	userMigrator := UserMigrator{
		converter: suite.converter,
	}

	err := userMigrator.Migrate(suite.dest, &users)

	assert.Error(suite.T(), err)
}

func TestUserMigrator_Migrate(t *testing.T) {
	suite.Run(t, new(MigratorSuite))
}
