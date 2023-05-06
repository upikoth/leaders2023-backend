package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

func (s *Store) GetMetroStations() ([]model.MetroStation, error) {
	metroStations := []model.MetroStation{}

	err := s.db.Model(&metroStations).Order("name").Select()

	if err != nil {
		return nil, constants.ErrMetroStationsGetDbError
	}

	return metroStations, nil
}
