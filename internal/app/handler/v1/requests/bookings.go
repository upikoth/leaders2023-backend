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
