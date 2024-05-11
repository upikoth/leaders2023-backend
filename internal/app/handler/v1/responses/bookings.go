package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getBookingsResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getBookingsResponseCreativeSpace struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Photos      string `json:"photos"`
}

type getBookingsResponseLandlordInfo struct {
	ID              string     `json:"id"`
	Phone           string     `json:"phone"`
	Role            model.Role `json:"role"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
}

type getBookingsResponseTenantInfo struct {
	ID         string     `json:"id"`
	Phone      string     `json:"phone"`
	Role       model.Role `json:"role"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Patronymic string     `json:"patronymic"`
	Email      string     `json:"email"`
}

type getBookingsResponseBooking struct {
	ID             string                             `json:"id"`
	Status         model.BookingStatus                `json:"status"`
	FullPrice      int                                `json:"fullPrice"`
	CalendarEvents []getBookingsResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingsResponseCreativeSpace   `json:"creativeSpace"`
	LandlordInfo   getBookingsResponseLandlordInfo    `json:"landlordInfo"`
	TenantInfo     getBookingsResponseTenantInfo      `json:"tenantInfo"`
	ScoreID        string                             `json:"scoreId"`
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
			ID:          booking.CreativeSpace.ID,
			Title:       booking.CreativeSpace.Title,
			Address:     booking.CreativeSpace.Address,
			Description: booking.CreativeSpace.Description,
			Photos:      booking.CreativeSpace.Photos,
		}

		scoreID := ""

		if booking.Score != nil {
			scoreID = booking.Score.ID
		}

		res.Bookings = append(res.Bookings, getBookingsResponseBooking{
			ID:             booking.ID,
			Status:         model.BookingStatus(booking.Status),
			FullPrice:      booking.FullPrice,
			CalendarEvents: resCalendarEvents,
			CreativeSpace:  resCreativeSpace,
			ScoreID:        scoreID,
			LandlordInfo: getBookingsResponseLandlordInfo{
				ID:              booking.LandlordInfo.ID,
				Phone:           booking.LandlordInfo.Phone,
				Role:            model.Role(booking.LandlordInfo.Role),
				Name:            booking.LandlordInfo.Name,
				Surname:         booking.LandlordInfo.Surname,
				Patronymic:      booking.LandlordInfo.Patronymic,
				Email:           booking.LandlordInfo.Email,
				Inn:             booking.LandlordInfo.Inn,
				LegalEntityName: booking.LandlordInfo.LegalEntityName,
			},
			TenantInfo: getBookingsResponseTenantInfo{
				ID:         booking.TenantInfo.ID,
				Phone:      booking.TenantInfo.Phone,
				Role:       model.Role(booking.TenantInfo.Role),
				Name:       booking.TenantInfo.Name,
				Surname:    booking.TenantInfo.Surname,
				Patronymic: booking.TenantInfo.Patronymic,
				Email:      booking.TenantInfo.Email,
			},
		})
	}

	return res
}

type getBookingResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getBookingResponseCreativeSpace struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Photos      string `json:"photos"`
}

type getBookingResponseLandlordInfo struct {
	ID              string     `json:"id"`
	Phone           string     `json:"phone"`
	Role            model.Role `json:"role"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
}

type getBookingResponseTenantInfo struct {
	ID         string     `json:"id"`
	Phone      string     `json:"phone"`
	Role       model.Role `json:"role"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Patronymic string     `json:"patronymic"`
	Email      string     `json:"email"`
}

type getBookingResponseBooking struct {
	ID             string                            `json:"id"`
	Status         model.BookingStatus               `json:"status"`
	FullPrice      int                               `json:"fullPrice"`
	CalendarEvents []getBookingResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingResponseCreativeSpace   `json:"creativeSpace"`
	LandlordInfo   getBookingResponseLandlordInfo    `json:"landlordInfo"`
	TenantInfo     getBookingResponseTenantInfo      `json:"tenantInfo"`
	ScoreID        string                            `json:"scoreId"`
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
		ID:          booking.CreativeSpace.ID,
		Title:       booking.CreativeSpace.Title,
		Address:     booking.CreativeSpace.Address,
		Description: booking.CreativeSpace.Description,
		Photos:      booking.CreativeSpace.Photos,
	}

	scoreID := ""

	if booking.Score != nil {
		scoreID = booking.Score.ID
	}

	res.Booking = getBookingResponseBooking{
		ID:             booking.ID,
		ScoreID:        scoreID,
		Status:         model.BookingStatus(booking.Status),
		FullPrice:      booking.FullPrice,
		CalendarEvents: resCalendarEvents,
		CreativeSpace:  resCreativeSpace,
		LandlordInfo: getBookingResponseLandlordInfo{
			ID:              booking.LandlordInfo.ID,
			Phone:           booking.LandlordInfo.Phone,
			Role:            model.Role(booking.LandlordInfo.Role),
			Name:            booking.LandlordInfo.Name,
			Surname:         booking.LandlordInfo.Surname,
			Patronymic:      booking.LandlordInfo.Patronymic,
			Email:           booking.LandlordInfo.Email,
			Inn:             booking.LandlordInfo.Inn,
			LegalEntityName: booking.LandlordInfo.LegalEntityName,
		},
		TenantInfo: getBookingResponseTenantInfo{
			ID:         booking.TenantInfo.ID,
			Phone:      booking.TenantInfo.Phone,
			Role:       model.Role(booking.TenantInfo.Role),
			Name:       booking.TenantInfo.Name,
			Surname:    booking.TenantInfo.Surname,
			Patronymic: booking.TenantInfo.Patronymic,
			Email:      booking.TenantInfo.Email,
		},
	}

	return res
}

type createBookingResponseBooking struct {
	ID string `json:"id"`
}

type createBookingResponseData struct {
	Booking createBookingResponseBooking `json:"booking"`
}

func CreateBookingResponseFromStoreData(bookingID string) createBookingResponseData {
	res := createBookingResponseData{}

	res.Booking = createBookingResponseBooking{
		ID: bookingID,
	}

	return res
}
