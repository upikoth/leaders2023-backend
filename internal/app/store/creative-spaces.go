package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type CreativeSpace struct {
	tableName           struct{}                     `pg:"creative_spaces"` //nolint:unused // Имя таблицы
	Id                  int                          `pg:"id"`
	Title               string                       `pg:"title"`
	Address             string                       `pg:"address"`
	LandlordId          int                          `pg:"landlord_id"`
	Photos              []string                     `pg:"photos,array"`
	PricePerHour        int                          `pg:"price_per_hour"`
	Latitude            float32                      `pg:"latitude"`
	Longitude           float32                      `pg:"longitude"`
	WorkingHoursStartAt string                       `pg:"working_hours_start_at"`
	WorkingHoursEndAt   string                       `pg:"working_hours_end_at"`
	Description         string                       `pg:"description"`
	MetroStations       []*CreativeSpaceMetroStation `pg:"rel:has-many"`
}

type CreativeSpaceMetroStation struct {
	tableName         struct{}      `pg:"creative_space_metro_station"` //nolint:unused // Имя таблицы
	MetroStationId    int           `pg:"metro_station_id"`
	CreativeSpaceId   int           `pg:"creative_space_id"`
	DistanceInMinutes int           `pg:"distance_in_minutes"`
	MetroStation      *MetroStation `pg:"rel:has-one"`
}

func (s *Store) GetCreativeSpaces() ([]CreativeSpace, error) {
	creativeSpaces := []CreativeSpace{}

	err := s.db.
		Model(&creativeSpaces).
		Relation("MetroStations", func(q *pg.Query) (*pg.Query, error) {
			return q.Relation("MetroStation"), nil
		}).
		Select()

	if err != nil {
		return nil, err
	}

	return creativeSpaces, nil
}

func (s *Store) GetCreativeSpaceById(id int) (CreativeSpace, error) {
	creativeSpace := CreativeSpace{
		Id: id,
	}

	count, err := s.db.
		Model(&creativeSpace).
		WherePK().
		Relation("MetroStations", func(q *pg.Query) (*pg.Query, error) {
			return q.Relation("MetroStation"), nil
		}).
		SelectAndCount()

	if err != nil {
		return CreativeSpace{}, err
	}

	if count == 0 {
		return CreativeSpace{}, constants.ErrCreativeSpaceGetNotFoundById
	}

	return creativeSpace, nil
}

func (s *Store) CreateCreativeSpace(
	creativeSpace CreativeSpace,
	creativeSpaceMetroStations []CreativeSpaceMetroStation,
) (int, error) {
	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		result, creativeSpaceErr := tx.
			Model(&creativeSpace).
			OnConflict("DO NOTHING").
			Insert()

		if creativeSpaceErr != nil {
			return creativeSpaceErr
		}

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpacePostDbError
		}

		if len(creativeSpaceMetroStations) == 0 {
			return nil
		}

		for i := range creativeSpaceMetroStations {
			creativeSpaceMetroStations[i].CreativeSpaceId = creativeSpace.Id
		}

		result, creativeSpaceMetroStationErr := tx.
			Model(&creativeSpaceMetroStations).
			OnConflict("DO NOTHING").
			Insert()

		if creativeSpaceMetroStationErr != nil {
			return creativeSpaceMetroStationErr
		}

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpacePostDbError
		}

		return nil
	})

	if storeErr != nil {
		return 0, storeErr
	}

	return creativeSpace.Id, nil
}

func (s *Store) PatchCreativeSpace(
	creativeSpace CreativeSpace,
	creativeSpaceMetroStations []CreativeSpaceMetroStation,
) error {
	count, err := s.db.
		Model(&creativeSpace).
		WherePK().
		Count()

	if err != nil {
		return err
	}

	if count == 0 {
		return constants.ErrCreativeSpacePatchNotFoundById
	}

	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, creativeSpaceUpdateErr := tx.
			Model(&creativeSpace).
			WherePK().
			OnConflict("DO NOTHING").
			UpdateNotZero()

		if creativeSpaceUpdateErr != nil {
			return creativeSpaceUpdateErr
		}

		_, creativeSpaceMetroStationsDeleteErr := tx.
			Model(&CreativeSpaceMetroStation{}).
			Where("creative_space_id = ?", creativeSpace.Id).
			Delete()

		if creativeSpaceMetroStationsDeleteErr != nil {
			return creativeSpaceMetroStationsDeleteErr
		}

		if len(creativeSpaceMetroStations) == 0 {
			return nil
		}

		_, creativeSpaceMetroStationsInsertErr := tx.
			Model(&creativeSpaceMetroStations).
			OnConflict("DO NOTHING").
			Insert()

		if creativeSpaceMetroStationsInsertErr != nil {
			return creativeSpaceMetroStationsInsertErr
		}

		return nil
	})

	if storeErr != nil {
		return storeErr
	}

	return nil
}

func (s *Store) DeleteCreativeSpace(id int) error {
	creativeSpace := CreativeSpace{
		Id: id,
	}

	creativeSpaceMetroStation := CreativeSpaceMetroStation{
		CreativeSpaceId: id,
	}

	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, err := tx.
			Model(&creativeSpaceMetroStation).
			Where("creative_space_id = ?", creativeSpaceMetroStation.CreativeSpaceId).
			Delete()

		if err != nil {
			return err
		}

		result, err := tx.
			Model(&creativeSpace).
			WherePK().
			Delete()

		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpaceDeleteNotFoundById
		}

		return nil
	})

	if storeErr != nil {
		return storeErr
	}

	return nil
}
