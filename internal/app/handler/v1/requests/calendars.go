package requests

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type convertCalendarRequestData struct {
	Calendar *multipart.FileHeader `form:"calendar" binding:"required"`
}

func ConvertCalendarDataFromRequest(c *gin.Context) (convertCalendarRequestData, error) {
	data := convertCalendarRequestData{}

	err := c.Bind(&data)

	if err != nil {
		return convertCalendarRequestData{}, err
	}

	return data, nil
}

type convertCalendarFromLinkRequestData struct {
	Link string `json:"link" binding:"required"`
}

func ConvertCalendarFromLinkDataFromRequest(c *gin.Context) (convertCalendarFromLinkRequestData, error) {
	data := convertCalendarFromLinkRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return convertCalendarFromLinkRequestData{}, err
	}

	return data, nil
}
