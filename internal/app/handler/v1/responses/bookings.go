package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getBookingsResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getBookingsResponseCreativeSpace struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Address     string   `json:"address"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type getBookingsResponseBooking struct {
	Id             int                                `json:"id"`
	TenantId       int                                `json:"tenantId"`
	LandlordId     int                                `json:"landlordId"`
	Status         model.BookingStaus                 `json:"status"`
	FullPrice      int                                `json:"fullPrice"`
	CalendarEvents []getBookingsResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingsResponseCreativeSpace   `json:"creativeSpace"`
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

		resCreativeSpace := getBookingsResponseCreativeSpace{
			Id:          booking.CreativeSpace.Id,
			Title:       booking.CreativeSpace.Title,
			Address:     booking.CreativeSpace.Address,
			Description: booking.CreativeSpace.Description,
			Photos:      booking.CreativeSpace.Photos,
		}

		res.Bookings = append(res.Bookings, getBookingsResponseBooking{
			Id:             booking.Id,
			TenantId:       booking.TenantId,
			LandlordId:     booking.LandlordId,
			Status:         booking.Status,
			FullPrice:      booking.FullPrice,
			CalendarEvents: resCalendarEvents,
			CreativeSpace:  resCreativeSpace,
		})
	}

	return res
}

type getBookingResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getBookingResponseCreativeSpace struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Address     string   `json:"address"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type getBookingResponseBooking struct {
	Id             int                               `json:"id"`
	TenantId       int                               `json:"tenantId"`
	LandlordId     int                               `json:"landlordId"`
	Status         model.BookingStaus                `json:"status"`
	FullPrice      int                               `json:"fullPrice"`
	CalendarEvents []getBookingResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingResponseCreativeSpace   `json:"creativeSpace"`
}

type getBookingResponseData struct {
	Booking getBookingResponseBooking `json:"booking"`
}

func GetBookingResponseFromStoreData(booking store.Booking) getBookingResponseData {
	res := getBookingResponseData{}

	resCalendarEvents := []getBookingResponseCalendarEvent{}

	for _, calendarEvent := range booking.CalendarEvents {
		resCalendarEvents = append(resCalendarEvents, getBookingResponseCalendarEvent{
			Date: calendarEvent.Date,
		})
	}

	resCreativeSpace := getBookingResponseCreativeSpace{
		Id:          booking.CreativeSpace.Id,
		Title:       booking.CreativeSpace.Title,
		Address:     booking.CreativeSpace.Address,
		Description: booking.CreativeSpace.Description,
		Photos:      booking.CreativeSpace.Photos,
	}

	res.Booking = getBookingResponseBooking{
		Id:             booking.Id,
		TenantId:       booking.TenantId,
		LandlordId:     booking.LandlordId,
		Status:         booking.Status,
		FullPrice:      booking.FullPrice,
		CalendarEvents: resCalendarEvents,
		CreativeSpace:  resCreativeSpace,
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
