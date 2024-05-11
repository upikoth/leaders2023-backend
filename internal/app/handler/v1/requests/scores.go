package requests

import "github.com/gin-gonic/gin"

type createScoreRequestData struct {
	Comment         string `json:"comment"`
	Rating          int    `json:"rating" binding:"required"`
	BookingID       string `json:"bookingId" binding:"required"`
	CreativeSpaceID string `json:"creativeSpaceId" binding:"required"`
}

func CreateScoreDataFromRequest(c *gin.Context) (createScoreRequestData, error) {
	data := createScoreRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createScoreRequestData{}, err
	}

	return data, nil
}
