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
	filePath := s.config.YdbAuthDirName + "/" + s.config.YdbAuthFileName

	if len(s.config.YdbAuthInfo) > 0 {
		_, err := os.Stat(s.config.YdbAuthDirName)

		if err != nil {
			mkdirErr := os.Mkdir(s.config.YdbAuthDirName, 0777)

			if mkdirErr != nil {
				return mkdirErr
			}
		}

		err = os.WriteFile(filePath, s.config.YdbAuthInfo, 0600)

		if err != nil {
			return err
		}
	}

	os.Setenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS", filePath)
	db, err := gorm.Open(ydb.Open(s.config.YdbDsn))

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
