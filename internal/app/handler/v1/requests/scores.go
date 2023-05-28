package requests

import "github.com/gin-gonic/gin"

type createScoreRequestData struct {
	Comment         string `json:"comment"`
	Rating          int    `json:"rating" binding:"required"`
	BookingId       int    `json:"bookingId" binding:"required"`
	CreativeSpaceId int    `json:"creativeSpaceId" binding:"required"`
}

func CreateScoregDataFromRequest(c *gin.Context) (createScoreRequestData, error) {
	data := createScoreRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createScoreRequestData{}, err
	}

	return data, nil
}
