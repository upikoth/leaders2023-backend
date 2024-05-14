package store

import (
	"os"

	ydb "github.com/ydb-platform/gorm-driver"
	"gorm.io/gorm"
)

type Store struct {
	db     *gorm.DB
	config *Config
}

func New() *Store {
	config, _ := NewConfig()

	return &Store{
		config: config,
	}
}

func (s *Store) Connect() error {
	if len(s.config.YdbAuthInfo) > 0 {
		err := os.WriteFile(s.config.YdbAuthFileName, s.config.YdbAuthInfo, 0600)

		if err != nil {
			return err
		}
	}

	os.Setenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS", s.config.YdbAuthFileName)
	db, err := gorm.Open(
		ydb.Open(s.config.YdbDsn),
	)
	os.Remove(s.config.YdbAuthFileName)

	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) AutoMigrate() error {
	return s.db.AutoMigrate(
		&User{},
		&MetroStation{},
		&CalendarEvent{},
		&Score{},
		&CreativeSpaceMetroStation{},
		&CreativeSpace{},
		&Booking{},
	)
}

func (s *Store) Disconnect() error {
	db, _ := s.db.DB()
	return db.Close()
}
