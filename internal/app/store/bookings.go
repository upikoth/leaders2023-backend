package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

type Booking struct {
	tableName       struct{}            `pg:"bookings,alias:bookings"` //nolint:unused // Имя таблицы
	Id              int                 `pg:"id"`
	TenantId        int                 `pg:"tenant_id"`
	LandlordId      int                 `pg:"landlord_id"`
	CreativeSpaceId int                 `pg:"creative_space_id"`
	Status          model.BookingStatus `pg:"status"`
	FullPrice       int                 `pg:"full_price"`
	CalendarEvents  []*CalendarEvent    `pg:"rel:has-many"`
	CreativeSpace   *CreativeSpace      `pg:"rel:has-one"`
	TenantInfo      *User               `pg:"rel:has-one,fk:tenant_id"`
	LandlordInfo    *User               `pg:"rel:has-one,fk:landlord_id"`
}

type BookingsFilter struct {
	TenantId   int `pg:"tenant_id"`
	LandlordId int `pg:"landlord_id"`
}

func (s *Store) GetBookings(filters BookingsFilter) ([]Booking, error) {
	bookings := []Booking{}

	err := s.db.
		Model(&bookings).
		Where("bookings.tenant_id = ? OR ?", filters.TenantId, filters.TenantId == 0).
		Where("bookings.landlord_id = ? OR ?", filters.LandlordId, filters.LandlordId == 0).
		Relation("CalendarEvents").
		Relation("CreativeSpace").
		Relation("TenantInfo").
		Relation("LandlordInfo").
		Select()

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *Store) GetBookingById(bookingId int) (Booking, error) {
	booking := Booking{
		Id: bookingId,
	}

	err := s.db.
		Model(&booking).
		WherePK().
		Relation("CalendarEvents").
		Relation("CreativeSpace").
		Relation("TenantInfo").
		Relation("LandlordInfo").
		Select()

	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}

func (s *Store) CreateBooking(booking Booking) (int, error) {
	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// Записываем в таблицу bookings.
		result, bookingErr := tx.
			Model(&booking).
			OnConflict("DO NOTHING").
			Insert()

		if bookingErr != nil {
			return bookingErr
		}

		if result.RowsAffected() == 0 {
			return constants.ErrBookingPostDbError
		}

		if len(booking.CalendarEvents) > 0 {
			for i := range booking.CalendarEvents {
				booking.CalendarEvents[i].BookingId = booking.Id
			}

			// Записываем в таблицу calendar_events.
			bookingCalendarEventsResult, bookingCalendarEventsErr := tx.
				Model(&booking.CalendarEvents).
				OnConflict("DO NOTHING").
				Insert()

			if bookingCalendarEventsErr != nil {
				return bookingCalendarEventsErr
			}

			if bookingCalendarEventsResult.RowsAffected() == 0 {
				return constants.ErrBookingPostDbError
			}
		}

		return nil
	})

	if storeErr != nil {
		return 0, storeErr
	}

	return booking.Id, nil
}

func (s *Store) PatchBooking(booking Booking) error {
	count, err := s.db.
		Model(&booking).
		WherePK().
		Count()

	if err != nil {
		return err
	}

	if count == 0 {
		return constants.ErrBookingPatchNotFoundById
	}

	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, bookingUpdateErr := tx.
			Model(&booking).
			WherePK().
			OnConflict("DO NOTHING").
			UpdateNotZero()

		if bookingUpdateErr != nil {
			return bookingUpdateErr
		}

		_, bookingCalendarEventsDeleteErr := tx.
			Model(&CalendarEvent{}).
			Where("booking_id = ?", booking.Id).
			Delete()

		if bookingCalendarEventsDeleteErr != nil {
			return bookingCalendarEventsDeleteErr
		}

		if len(booking.CalendarEvents) > 0 {
			// Записываем в таблицу calendar_events.
			bookingCalendarEventsResult, bookingCalendarEventsErr := tx.
				Model(&booking.CalendarEvents).
				OnConflict("DO NOTHING").
				Insert()

			if bookingCalendarEventsErr != nil {
				return bookingCalendarEventsErr
			}

			if bookingCalendarEventsResult.RowsAffected() == 0 {
				return constants.ErrBookingPatchDbError
			}
		}

		return nil
	})

	return storeErr
}

func (s *Store) DeleteBooking(id int) error {
	booking := Booking{
		Id: id,
	}

	creativeSpaceCalendarEvent := CalendarEvent{
		BookingId: id,
	}

	storeErr := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		_, calendarEventErr := tx.
			Model(&creativeSpaceCalendarEvent).
			Where("booking_id = ?", creativeSpaceCalendarEvent.BookingId).
			Delete()

		if calendarEventErr != nil {
			return calendarEventErr
		}

		result, err := tx.
			Model(&booking).
			WherePK().
			Delete()

		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return constants.ErrBookingDeleteNotFoundById
		}

		return nil
	})

	if storeErr != nil {
		return storeErr
	}

	return nil
}
