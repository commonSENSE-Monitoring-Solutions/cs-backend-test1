package storage

import (
	"cs-backend-test1/config"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	name    string
	session *gorm.DB
	schema  interface{}

	migrator Migrator
}

type Migrator interface {
	Migrate(dest Storage, records interface{}) error
}

type Storage interface {
	IsSessionActive() bool
	Close() error
	SetMigrator(migrator Migrator)
	OpenWithConfig(cfg *config.Database) error
	ReadAll(dest interface{}) error
	DeleteAll() error
	MigrateSchema(schema interface{}) error
	MigrateDataBatches(dest Storage, records interface{}, size int) error
	Write(record interface{}) error
}

func NewDatabase(name string, schema interface{}) Storage {
	return &database{name: name, schema: schema}
}

func (d *database) SetMigrator(migrator Migrator) {
	d.migrator = migrator
}

func (d *database) IsSessionActive() bool {
	if d.session == nil {
		return false
	}

	return true
}

func (d *database) Close() error {
	db, err := d.session.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (d *database) OpenWithConfig(cfg *config.Database) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Host, cfg.User, cfg.Password, d.name, cfg.Port)

	session, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.session = session

	return nil
}

func (d *database) ReadAll(dest interface{}) error {
	res := d.session.Find(dest)
	log.Info().Msg(fmt.Sprintf("num of rows: %d", res.RowsAffected))

	return res.Error
}

func (d *database) DeleteAll() error {
	return d.session.Exec(`DELETE FROM "public"."user_accounts"`).Error
}

func (d *database) MigrateSchema(schema interface{}) error {
	return d.session.AutoMigrate(schema)
}

func (d *database) Write(dest interface{}) error {
	return d.session.Transaction(func(tx *gorm.DB) error {
		return tx.Create(dest).Error
	})
}

func (d *database) MigrateDataBatches(dest Storage, records interface{}, size int) error {
	if d.migrator == nil {
		return errors.New("migrator was not set during db initialization")
	}

	err := dest.MigrateSchema(d.schema)
	if err != nil {
		return fmt.Errorf("sync schema error: %w", err)
	}

	return d.session.FindInBatches(records, size, func(tx *gorm.DB, batch int) error {
		return d.migrator.Migrate(dest, records)
	}).Error
}
