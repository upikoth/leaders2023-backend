package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type MetroStation struct {
	tableName struct{} `pg:"metro_stations"` //nolint:unused // Имя таблицы
	Id        int      `pg:"id"`
	Name      string   `pg:"name"`
	Color     string   `pg:"color"`
}

func (s *Store) GetMetroStations() ([]MetroStation, error) {
	metroStations := []MetroStation{}

	err := s.db.Model(&metroStations).Order("name").Select()

	if err != nil {
		return nil, constants.ErrMetroStationsGetDbError
	}

	return metroStations, nil
}
