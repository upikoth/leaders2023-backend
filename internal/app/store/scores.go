package store

type Score struct {
	ID              string `gorm:"primarykey"`
	UserID          string
	CreativeSpaceID string
	BookingID       string
	Comment         string
	Rating          int
	User            *User `gorm:"foreignKey:ID;references:UserID"`
}

func (s *Store) CreateScore(score Score) (string, error) {
	res := s.db.Create(&score)

	if res.Error != nil {
		return "", res.Error
	}

	return score.ID, nil
}
