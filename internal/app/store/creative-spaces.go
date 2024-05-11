package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreativeSpace struct {
	ID                     string `gorm:"primarykey"`
	SpaceType              string
	Area                   int
	Capacity               int
	Title                  string
	Status                 string
	Address                string
	LandlordID             string
	PricePerDay            int
	Latitude               float32
	Longitude              float32
	Description            string
	CalendarLink           string
	Photos                 string
	CalendarWorkDayIndexes string
	LandlordInfo           *User                        `gorm:"foreignKey:ID;references:LandlordID"`
	MetroStations          []*CreativeSpaceMetroStation `gorm:"foreignKey:CreativeSpaceID;references:ID"`
	CalendarEvents         []*CalendarEvent             `gorm:"foreignKey:CreativeSpaceID;references:ID"`
	Scores                 []*Score                     `gorm:"foreignKey:CreativeSpaceID;references:ID"`
}

func (s *Store) GetCreativeSpaces() ([]CreativeSpace, error) {
	creativeSpaces := []CreativeSpace{}

	res := s.db.
		Preload(clause.Associations).
		Preload("Scores.User").
		Preload("MetroStations.MetroStation").
		Find(&creativeSpaces)

	if res.Error != nil {
		return nil, res.Error
	}

	return creativeSpaces, nil
}

func (s *Store) GetCreativeSpaceByID(id string) (CreativeSpace, error) {
	creativeSpace := CreativeSpace{ID: id}

	res := s.db.
		Preload(clause.Associations).
		Preload("Scores.User").
		Preload("MetroStations.MetroStation").
		First(&creativeSpace)

	if res.Error != nil {
		return CreativeSpace{}, res.Error
	}

	return creativeSpace, nil
}

func (s *Store) CreateCreativeSpace(
	creativeSpace CreativeSpace,
) (string, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&creativeSpace)

		if res.Error != nil {
			return res.Error
		}

		if len(creativeSpace.MetroStations) > 0 {
			for i := range creativeSpace.MetroStations {
				creativeSpace.MetroStations[i].CreativeSpaceID = creativeSpace.ID
			}

			resMetroStation := tx.Create(&creativeSpace.MetroStations)

			if resMetroStation.Error != nil {
				return res.Error
			}
		}

		if len(creativeSpace.CalendarEvents) > 0 {
			for i := range creativeSpace.CalendarEvents {
				creativeSpace.CalendarEvents[i].CreativeSpaceID = creativeSpace.ID
			}

			resCalendarEvents := tx.
				Create(&creativeSpace.CalendarEvents)

			if resCalendarEvents.Error != nil {
				return resCalendarEvents.Error
			}
		}

		return nil
	})

	if err != nil {
		return "", constants.ErrCreativeSpacePostDbError
	}

	return creativeSpace.ID, nil
}

func (s *Store) PatchCreativeSpace(
	creativeSpace CreativeSpace,
) error {
	_, err := s.GetCreativeSpaceByID(creativeSpace.ID)

	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Updates(&creativeSpace)

		if res.Error != nil {
			return res.Error
		}

		res = tx.
			Where("creative_space_id = ?", creativeSpace.ID).
			Delete(&CreativeSpaceMetroStation{})

		if res.Error != nil {
			return res.Error
		}

		if len(creativeSpace.MetroStations) > 0 {
			resMetroStation := tx.Create(&creativeSpace.MetroStations)

			if resMetroStation.Error != nil {
				return res.Error
			}
		}

		res = tx.
			Where("creative_space_id = ?", creativeSpace.ID).
			Delete(&CalendarEvent{})

		if res.Error != nil {
			return res.Error
		}

		if len(creativeSpace.CalendarEvents) > 0 {
			resCalendarEvents := tx.
				Create(&creativeSpace.CalendarEvents)

			if resCalendarEvents.Error != nil {
				return resCalendarEvents.Error
			}
		}

		return nil
	})

	if err != nil {
		return constants.ErrCreativeSpacePatchDbError
	}

	return nil
}

func (s *Store) DeleteCreativeSpace(id string) error {
	creativeSpace := CreativeSpace{
		ID: id,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&creativeSpace)

		if res.Error != nil {
			return res.Error
		}

		res = tx.
			Where("creative_space_id = ?", creativeSpace.ID).
			Delete(&CreativeSpaceMetroStation{})

		if res.Error != nil {
			return res.Error
		}

		res = tx.
			Where("creative_space_id = ?", creativeSpace.ID).
			Delete(&CalendarEvent{})

		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	return err
}
