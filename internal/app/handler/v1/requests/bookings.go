package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

type getBookingRequestData struct {
	ID string `uri:"id" binding:"required"`
}

func GetBookingDataFromRequest(c *gin.Context) (getBookingRequestData, error) {
	data := getBookingRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return getBookingRequestData{}, err
	}

	return data, nil
}

type createBookingRequestCalendarEvent struct {
	Date string `json:"date" binding:"required"`
}

type createBookingRequestData struct {
	CreativeSpaceID string                              `json:"creativeSpaceId" binding:"required"`
	CalendarEvents  []createBookingRequestCalendarEvent `json:"calendarEvents" binding:"required"`
}

func CreateBookingDataFromRequest(c *gin.Context) (createBookingRequestData, error) {
	data := createBookingRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createBookingRequestData{}, err
	}

	return data, nil
}

type patchBookingRequestURI struct {
	ID string `uri:"id" binding:"required"`
}

type patchBookingRequestCalendarEvent struct {
	Date string `json:"date" binding:"required"`
}

type patchBookingRequestDataBody struct {
	CalendarEvents []patchBookingRequestCalendarEvent `json:"calendarEvents"`
	Status         model.BookingStatus                `json:"status"`
}

type patchBookingRequestData struct {
	ID             string                             `json:"id"`
	CalendarEvents []patchBookingRequestCalendarEvent `json:"calendarEvents"`
	Status         model.BookingStatus                `json:"status"`
}

func PatchBookingDataFromRequest(c *gin.Context) (patchBookingRequestData, error) {
	dataFromURI := patchBookingRequestURI{}
	dataFromBody := patchBookingRequestDataBody{}

	uriErr := c.BindUri(&dataFromURI)

	if uriErr != nil {
		return patchBookingRequestData{}, uriErr
	}

	bodyErr := c.BindJSON(&dataFromBody)

	if bodyErr != nil {
		return patchBookingRequestData{}, bodyErr
	}

	data := patchBookingRequestData{
		ID:             dataFromURI.ID,
		CalendarEvents: dataFromBody.CalendarEvents,
		Status:         dataFromBody.Status,
	}

	return data, nil
}

type deleteBookingRequestData struct {
	ID string `uri:"id" binding:"required"`
}

func DeleteBookingDataFromRequest(c *gin.Context) (deleteBookingRequestData, error) {
	data := deleteBookingRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteBookingRequestData{}, err
	}

	return data, nil
}
