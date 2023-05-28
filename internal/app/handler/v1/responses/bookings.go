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

type getBookingsResponseLandlordInfo struct {
	Id              int        `json:"id"`
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
	Id         int        `json:"id"`
	Phone      string     `json:"phone"`
	Role       model.Role `json:"role"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Patronymic string     `json:"patronymic"`
	Email      string     `json:"email"`
}

type getBookingsResponseBooking struct {
	Id             int                                `json:"id"`
	Status         model.BookingStatus                `json:"status"`
	FullPrice      int                                `json:"fullPrice"`
	CalendarEvents []getBookingsResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingsResponseCreativeSpace   `json:"creativeSpace"`
	LandlordInfo   getBookingsResponseLandlordInfo    `json:"landlordInfo"`
	TenantInfo     getBookingsResponseTenantInfo      `json:"tenantInfo"`
	ScoreId        int                                `json:"scoreId"`
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

		scoreId := 0

		if booking.Score != nil {
			scoreId = booking.Score.Id
		}

		res.Bookings = append(res.Bookings, getBookingsResponseBooking{
			Id:             booking.Id,
			Status:         booking.Status,
			FullPrice:      booking.FullPrice,
			CalendarEvents: resCalendarEvents,
			CreativeSpace:  resCreativeSpace,
			ScoreId:        scoreId,
			LandlordInfo: getBookingsResponseLandlordInfo{
				Id:              booking.LandlordInfo.Id,
				Phone:           booking.LandlordInfo.Phone,
				Role:            booking.LandlordInfo.Role,
				Name:            booking.LandlordInfo.Name,
				Surname:         booking.LandlordInfo.Surname,
				Patronymic:      booking.LandlordInfo.Patronymic,
				Email:           booking.LandlordInfo.Email,
				Inn:             booking.LandlordInfo.Inn,
				LegalEntityName: booking.LandlordInfo.LegalEntityName,
			},
			TenantInfo: getBookingsResponseTenantInfo{
				Id:         booking.TenantInfo.Id,
				Phone:      booking.TenantInfo.Phone,
				Role:       booking.TenantInfo.Role,
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
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Address     string   `json:"address"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type getBookingResponseLandlordInfo struct {
	Id              int        `json:"id"`
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
	Id         int        `json:"id"`
	Phone      string     `json:"phone"`
	Role       model.Role `json:"role"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Patronymic string     `json:"patronymic"`
	Email      string     `json:"email"`
}

type getBookingResponseBooking struct {
	Id             int                               `json:"id"`
	Status         model.BookingStatus               `json:"status"`
	FullPrice      int                               `json:"fullPrice"`
	CalendarEvents []getBookingResponseCalendarEvent `json:"calendarEvents"`
	CreativeSpace  getBookingResponseCreativeSpace   `json:"creativeSpace"`
	LandlordInfo   getBookingResponseLandlordInfo    `json:"landlordInfo"`
	TenantInfo     getBookingResponseTenantInfo      `json:"tenantInfo"`
	ScoreId        int                               `json:"scoreId"`
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

	scoreId := 0

	if booking.Score != nil {
		scoreId = booking.Score.Id
	}

	res.Booking = getBookingResponseBooking{
		Id:             booking.Id,
		ScoreId:        scoreId,
		Status:         booking.Status,
		FullPrice:      booking.FullPrice,
		CalendarEvents: resCalendarEvents,
		CreativeSpace:  resCreativeSpace,
		LandlordInfo: getBookingResponseLandlordInfo{
			Id:              booking.LandlordInfo.Id,
			Phone:           booking.LandlordInfo.Phone,
			Role:            booking.LandlordInfo.Role,
			Name:            booking.LandlordInfo.Name,
			Surname:         booking.LandlordInfo.Surname,
			Patronymic:      booking.LandlordInfo.Patronymic,
			Email:           booking.LandlordInfo.Email,
			Inn:             booking.LandlordInfo.Inn,
			LegalEntityName: booking.LandlordInfo.LegalEntityName,
		},
		TenantInfo: getBookingResponseTenantInfo{
			Id:         booking.TenantInfo.Id,
			Phone:      booking.TenantInfo.Phone,
			Role:       booking.TenantInfo.Role,
			Name:       booking.TenantInfo.Name,
			Surname:    booking.TenantInfo.Surname,
			Patronymic: booking.TenantInfo.Patronymic,
			Email:      booking.TenantInfo.Email,
		},
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
