package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Booking struct {
	ID              string `gorm:"primarykey"`
	TenantID        string
	LandlordID      string
	CreativeSpaceID string
	Status          string
	FullPrice       int
	CalendarEvents  []*CalendarEvent `gorm:"foreignKey:BookingID;references:ID"`
	CreativeSpace   *CreativeSpace   `gorm:"foreignKey:ID;references:CreativeSpaceID"`
	TenantInfo      *User            `gorm:"foreignKey:ID;references:TenantID"`
	LandlordInfo    *User            `gorm:"foreignKey:ID;references:LandlordID"`
	Score           *Score           `gorm:"foreignKey:BookingID;references:ID"`
}

type BookingsFilter struct {
	TenantID   string
	LandlordID string
}

func (s *Store) GetBookings(filters BookingsFilter) ([]Booking, error) {
	bookings := []Booking{}

	res := s.db.
		Preload(clause.Associations).
		Where("bookings.tenant_id = ? OR ?", filters.TenantID, filters.TenantID == "").
		Where("bookings.landlord_id = ? OR ?", filters.LandlordID, filters.LandlordID == "").
		Find(&bookings)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookings, nil
}

func (s *Store) GetBookingByID(id string) (Booking, error) {
	booking := Booking{ID: id}

	res := s.db.
		Preload(clause.Associations).
		First(&booking)

	if res.Error != nil {
		return Booking{}, res.Error
	}

	return booking, nil
}

func (s *Store) CreateBooking(booking Booking) (string, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&booking)

		if res.Error != nil {
			return res.Error
		}

		if len(booking.CalendarEvents) > 0 {
			for i := range booking.CalendarEvents {
				booking.CalendarEvents[i].BookingID = booking.ID
			}

			resCalendarEvents := tx.
				Create(&booking.CalendarEvents)

			if resCalendarEvents.Error != nil {
				return resCalendarEvents.Error
			}
		}

		return nil
	})

	if err != nil {
		return "", constants.ErrBookingPostDbError
	}

	return booking.ID, nil
}

func (s *Store) PatchBooking(booking Booking) error {
	_, err := s.GetBookingByID(booking.ID)

	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Updates(&booking)

		if res.Error != nil {
			return res.Error
		}

		res = tx.
			Where("booking_id = ?", booking.ID).
			Delete(&CalendarEvent{})

		if res.Error != nil {
			return res.Error
		}

		if len(booking.CalendarEvents) > 0 {
			resCalendarEvents := tx.
				Create(&booking.CalendarEvents)

			if resCalendarEvents.Error != nil {
				return resCalendarEvents.Error
			}
		}

		return nil
	})

	if err != nil {
		return constants.ErrBookingPatchDbError
	}

	return nil
}

func (s *Store) DeleteBooking(id string) error {
	booking := Booking{
		ID: id,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&booking)

		if res.Error != nil {
			return res.Error
		}

		res = tx.
			Where("booking_id = ?", booking.ID).
			Delete(&CreativeSpaceMetroStation{})

		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	return err
}
