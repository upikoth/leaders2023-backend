package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getBookingsResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getBookingsResponseBooking struct {
	Id              int                                `json:"id"`
	TenantId        int                                `json:"tenantId"`
	LandlordId      int                                `json:"landlordId"`
	CreativeSpaceId int                                `json:"creativeSpaceId"`
	Status          model.BookingStaus                 `json:"status"`
	FullPrice       int                                `json:"fullPrice"`
	CalendarEvents  []getBookingsResponseCalendarEvent `json:"calendarEvents"`
}

type getBookingsResponseData struct {
	Bookings []getBookingsResponseBooking `json:"bookings"`
}

func GetBookingsResponseFromStoreData(bookings []store.Booking) getBookingsResponseData {
	res := getBookingsResponseData{
		Bookings: []getBookingsResponseBooking{},
	}

	for _, booking := range bookings {
		resCalendarEvents := []getBookingsResponseCalendarEvent{}

		for _, calendarEvent := range booking.CalendarEvents {
			resCalendarEvents = append(resCalendarEvents, getBookingsResponseCalendarEvent{
				Date: calendarEvent.Date,
			})
		}

		res.Bookings = append(res.Bookings, getBookingsResponseBooking{
			Id:              booking.Id,
			TenantId:        booking.TenantId,
			LandlordId:      booking.LandlordId,
			CreativeSpaceId: booking.CreativeSpaceId,
			Status:          booking.Status,
			FullPrice:       booking.FullPrice,
			CalendarEvents:  resCalendarEvents,
		})
	}

	return res
}

type createBookingResponseBooking struct {
	Id int `json:"id"`
}

type createBookingResponseData struct {
	Booking createBookingResponseBooking `json:"booking"`
}

func CreateBookingResponseFromStoreData(bookingId int) createBookingResponseData {
	res := createBookingResponseData{}

	res.Booking = createBookingResponseBooking{
		Id: bookingId,
	}

	return res
}
