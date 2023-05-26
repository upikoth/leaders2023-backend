package requests

import "github.com/gin-gonic/gin"

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
