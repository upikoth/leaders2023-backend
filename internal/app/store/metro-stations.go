package store

type MetroStation struct {
	ID    string `gorm:"primarykey"`
	Name  string
	Color string
}

func (s *Store) GetMetroStations() ([]MetroStation, error) {
	metroStations := []MetroStation{}

	res := s.db.
		Find(&metroStations).
		Order("name")

	if res.Error != nil {
		return nil, res.Error
	}

	return metroStations, nil
}
