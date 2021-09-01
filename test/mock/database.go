package mock

import (
	"flag"
	"github.com/ao-concepts/storage"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	flag.Set("tags", "json1") // needed for sqlite to support the json data type
}

const Anything = mock.Anything

func AnythingOfType(t string) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

type Database struct {
	mock.Mock
}

func (db *Database) UseTransaction(fn storage.HandlerFunc) error {
	err := db.Called(fn).Error(0)

	if err != nil {
		return err
	}

	tx, rollback := Transaction()
	defer rollback()

	return fn(tx)
}

func (db *Database) Gorm() *gorm.DB {
	return db.Called().Get(0).(*gorm.DB)
}

func (db *Database) Begin() (tx *storage.Transaction, err error) {
	args := db.Called()
	return args.Get(0).(*storage.Transaction), args.Error(1)
}

func Transaction() (tx *storage.Transaction, rollback func()) {
	gormLogger := &GormLogger{}
	gormLogger.On("Trace", Anything, Anything, Anything, Anything)

	log := &Log{}
	log.On("CreateGormLogger").Return(gormLogger)
	log.On("Trace")

	db, err := storage.New(sqlite.Open(":memory:"), log)

	if err != nil {
		panic(err)
	}

	gormDB := db.Gorm()
	gormDB.SkipDefaultTransaction = true

	sqldb, _ := db.Gorm().DB()
	sqldb.SetMaxOpenConns(1) // needed for sqlite
	sqldb.SetMaxIdleConns(1) // needed for sqlite

	t, err := db.Begin()

	if err != nil {
		panic(err)
	}

	if err := t.Gorm().AutoMigrate(
		&Entity{},
	); err != nil {
		panic(err)
	}

	return t, func() {
		if err := t.Rollback(); err != nil {
			panic(err)
		}
	}
}
