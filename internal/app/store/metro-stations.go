package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	modelStore "github.com/upikoth/leaders2023-backend/internal/app/model/store"
)

func (s *Store) GetMetroStations() ([]modelStore.MetroStation, error) {
	metroStations := []modelStore.MetroStation{}

	err := s.db.Model(&metroStations).Order("name").Select()

	if err != nil {
		return nil, constants.ErrMetroStationsGetDbError
	}

	return metroStations, nil
}
