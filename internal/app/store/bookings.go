package store

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

type Booking struct {
	tableName       struct{}           `pg:"bookings"` //nolint:unused // Имя таблицы
	Id              int                `pg:"id"`
	TenantId        int                `pg:"tenant_id"`
	LandlordId      int                `pg:"landlord_id"`
	CreativeSpaceId int                `pg:"creative_space_id"`
	Status          model.BookingStaus `pg:"status"`
	FullPrice       int                `pg:"full_price"`
	CalendarEvents  []*CalendarEvent   `pg:"rel:has-many"`
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
