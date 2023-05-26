package requests

import "github.com/gin-gonic/gin"

type getBookingRequestData struct {
	Id int `uri:"id" binding:"required"`
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
	CreativeSpaceId int                                 `json:"creativeSpaceId" binding:"required"`
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

type patchBookingRequestUri struct {
	Id int `uri:"id" binding:"required"`
}

type patchBookingRequestCalendarEvent struct {
	Date string `json:"date" binding:"required"`
}

type patchBookingRequestDataBody struct {
	CalendarEvents []patchBookingRequestCalendarEvent `json:"calendarEvents"`
}

type patchBookingRequestData struct {
	Id             int                                `json:"id"`
	CalendarEvents []patchBookingRequestCalendarEvent `json:"calendarEvents"`
}

func PatchBookingDataFromRequest(c *gin.Context) (patchBookingRequestData, error) {
	dataFromUri := patchBookingRequestUri{}
	dataFromBody := patchBookingRequestDataBody{}

	uriErr := c.BindUri(&dataFromUri)

	if uriErr != nil {
		return patchBookingRequestData{}, uriErr
	}

	bodyErr := c.BindJSON(&dataFromBody)

	if bodyErr != nil {
		return patchBookingRequestData{}, bodyErr
	}

	data := patchBookingRequestData{
		Id:             dataFromUri.Id,
		CalendarEvents: dataFromBody.CalendarEvents,
	}

	return data, nil
}

type deleteBookingRequestData struct {
	Id int `uri:"id" binding:"required"`
}

func DeleteBookingDataFromRequest(c *gin.Context) (deleteBookingRequestData, error) {
	data := deleteBookingRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteBookingRequestData{}, err
	}

	return data, nil
}
