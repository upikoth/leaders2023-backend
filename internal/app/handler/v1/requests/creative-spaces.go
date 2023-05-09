package requests

import "github.com/gin-gonic/gin"

type getCreativeSpaceRequestData struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

func GetCreativeSpaceDataFromRequest(c *gin.Context) (getCreativeSpaceRequestData, error) {
	data := getCreativeSpaceRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return getCreativeSpaceRequestData{}, err
	}

	return data, nil
}

type createCreativeSpaceRequestWorkingHours struct {
	StartAt string `json:"startAt" binding:"required"`
	EndAt   string `json:"endAt" binding:"required"`
}

type createCreativeSpaceRequestCoordinate struct {
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type createCreativeSpaceRequestMetroStation struct {
	Id                int `json:"id" binding:"required"`
	DistanceInMinutes int `json:"distanceInMinutes" binding:"required"`
}

type createCreativeSpaceRequestData struct {
	Description   string                                   `json:"description" binding:"required"`
	Photos        []string                                 `json:"photos" binding:"required"`
	PricePerHour  int                                      `json:"pricePerHour" binding:"required"`
	MetroStations []createCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    createCreativeSpaceRequestCoordinate     `json:"coordinate" binding:"required"`
	WorkingHours  createCreativeSpaceRequestWorkingHours   `json:"workingHours" binding:"required"`
}

func CreateCreativeSpaceDataFromRequest(c *gin.Context) (createCreativeSpaceRequestData, error) {
	data := createCreativeSpaceRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createCreativeSpaceRequestData{}, err
	}

	return data, nil
}
