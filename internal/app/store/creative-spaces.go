package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type CreativeSpace struct {
	tableName           struct{}                     `pg:"creative_spaces"` //nolint:unused // Имя таблицы
	Id                  int                          `pg:"id"`
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

	if count == 0 {
		return CreativeSpace{}, constants.ErrCreativeSpaceGetNotFoundById
	}

	if err != nil {
		return CreativeSpace{}, err
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

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpacePostDbError
		}

		if creativeSpaceErr != nil {
			return creativeSpaceErr
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

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpacePostDbError
		}

		if creativeSpaceMetroStationErr != nil {
			return creativeSpaceMetroStationErr
		}

		return nil
	})

	if storeErr != nil {
		return 0, storeErr
	}

	return creativeSpace.Id, nil
}

func (s *Store) DeleteCreativeSpace(id int) error {
	creativeSpace := CreativeSpace{
		Id: id,
	}

	creativeSpaceMetroStation := CreativeSpaceMetroStation{
		CreativeSpaceId: id,
	}

	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, err := s.db.
			Model(&creativeSpaceMetroStation).
			Where("creative_space_id = ?", creativeSpaceMetroStation.CreativeSpaceId).
			Delete()

		if err != nil {
			return err
		}

		result, err := s.db.
			Model(&creativeSpace).
			WherePK().
			Delete()

		if result.RowsAffected() == 0 {
			return constants.ErrCreativeSpaceDeleteNotFoundById
		}

		if err != nil {
			return err
		}

		return nil
	})

	if storeErr != nil {
		return storeErr
	}

	return nil
}
