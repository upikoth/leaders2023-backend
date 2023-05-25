package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type CreativeSpace struct {
	tableName              struct{}                     `pg:"creative_spaces"` //nolint:unused // Имя таблицы
	Id                     int                          `pg:"id"`
	Title                  string                       `pg:"title"`
	Address                string                       `pg:"address"`
	LandlordId             int                          `pg:"landlord_id"`
	Photos                 []string                     `pg:"photos,array"`
	PricePerDay            int                          `pg:"price_per_day"`
	Latitude               float32                      `pg:"latitude"`
	Longitude              float32                      `pg:"longitude"`
	Description            string                       `pg:"description"`
	CalendarLink           string                       `pg:"calendar_link"`
	CalendarWorkDayIndexes []int                        `pg:"calendar_work_day_indexes,array"`
	CalendarEvents         []*CalendarEvent             `pg:"rel:has-many"`
	MetroStations          []*CreativeSpaceMetroStation `pg:"rel:has-many"`
}

func (s *Store) GetCreativeSpaces() ([]CreativeSpace, error) {
	creativeSpaces := []CreativeSpace{}

	err := s.db.
		Model(&creativeSpaces).
		Relation("MetroStations", func(q *pg.Query) (*pg.Query, error) {
			return q.Relation("MetroStation"), nil
		}).
		Relation("CalendarEvents").
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
		Relation("CalendarEvents").
		SelectAndCount()

	if err != nil {
		return CreativeSpace{}, err
	}

	if count == 0 {
		return CreativeSpace{}, constants.ErrCreativeSpaceGetNotFoundById
	}

	return creativeSpace, nil
}

//nolint:gocognit // Длина функции. Добавил в игнор, потому что одна транзакция.
func (s *Store) CreateCreativeSpace(
	creativeSpace CreativeSpace,
) (int, error) {
	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// Записываем в таблицу creative_spaces.
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

		if len(creativeSpace.MetroStations) > 0 {
			for i := range creativeSpace.MetroStations {
				creativeSpace.MetroStations[i].CreativeSpaceId = creativeSpace.Id
			}

			// Записываем в таблицу creative_space_metro_station.
			creativeSpaceMetroStationResult, creativeSpaceMetroStationErr := tx.
				Model(&creativeSpace.MetroStations).
				OnConflict("DO NOTHING").
				Insert()

			if creativeSpaceMetroStationErr != nil {
				return creativeSpaceMetroStationErr
			}

			if creativeSpaceMetroStationResult.RowsAffected() == 0 {
				return constants.ErrCreativeSpacePostDbError
			}
		}

		if len(creativeSpace.CalendarEvents) > 0 {
			for i := range creativeSpace.CalendarEvents {
				creativeSpace.CalendarEvents[i].CreativeSpaceId = creativeSpace.Id
			}

			// Записываем в таблицу calendar_events.
			creativeSpaceCalendarEventsResult, creativeSpaceCalendarEventsErr := tx.
				Model(&creativeSpace.CalendarEvents).
				OnConflict("DO NOTHING").
				Insert()

			if creativeSpaceCalendarEventsErr != nil {
				return creativeSpaceCalendarEventsErr
			}

			if creativeSpaceCalendarEventsResult.RowsAffected() == 0 {
				return constants.ErrCreativeSpacePostDbError
			}
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

		if len(creativeSpace.MetroStations) > 0 {
			_, creativeSpaceMetroStationsInsertErr := tx.
				Model(&creativeSpace.MetroStations).
				OnConflict("DO NOTHING").
				Insert()

			if creativeSpaceMetroStationsInsertErr != nil {
				return creativeSpaceMetroStationsInsertErr
			}
		}

		_, creativeSpaceCalendarEventsDeleteErr := tx.
			Model(&CalendarEvent{}).
			Where("creative_space_id = ?", creativeSpace.Id).
			Delete()

		if creativeSpaceCalendarEventsDeleteErr != nil {
			return creativeSpaceCalendarEventsDeleteErr
		}

		if len(creativeSpace.CalendarEvents) > 0 {
			_, creativeSpaceCalendarEventsInsertErr := tx.
				Model(&creativeSpace.CalendarEvents).
				OnConflict("DO NOTHING").
				Insert()

			if creativeSpaceCalendarEventsInsertErr != nil {
				return creativeSpaceCalendarEventsInsertErr
			}
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

	creativeSpaceCalendarEvent := CalendarEvent{
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

		_, calendarEventErr := tx.
			Model(&creativeSpaceCalendarEvent).
			Where("creative_space_id = ?", creativeSpaceCalendarEvent.CreativeSpaceId).
			Delete()

		if calendarEventErr != nil {
			return calendarEventErr
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
