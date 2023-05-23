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
