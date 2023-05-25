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

type createCreativeSpaceRequestCoordinate struct {
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type createCreativeSpaceRequestMetroStation struct {
	Id                int `json:"id" binding:"required"`
	DistanceInMinutes int `json:"distanceInMinutes" binding:"required"`
}

type createCreativeSpaceRequestData struct {
	Title         string                                   `json:"title" binding:"required"`
	Address       string                                   `json:"address" binding:"required"`
	Description   string                                   `json:"description" binding:"required"`
	Photos        []string                                 `json:"photos" binding:"required"`
	PricePerHour  int                                      `json:"pricePerHour" binding:"required"`
	MetroStations []createCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    createCreativeSpaceRequestCoordinate     `json:"coordinate" binding:"required"`
}

func CreateCreativeSpaceDataFromRequest(c *gin.Context) (createCreativeSpaceRequestData, error) {
	data := createCreativeSpaceRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createCreativeSpaceRequestData{}, err
	}

	return data, nil
}

type patchCreativeSpaceRequestCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type patchCreativeSpaceRequestMetroStation struct {
	Id                int `json:"id"`
	DistanceInMinutes int `json:"distanceInMinutes"`
}

type patchCreativeSpaceRequestUri struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

type patchCreativeSpaceRequestBody struct {
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerHour  int                                     `json:"pricePerHour"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
}

type patchCreativeSpaceRequestData struct {
	Id            int                                     `json:"id"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerHour  int                                     `json:"pricePerHour"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
}

func PatchCreativeSpaceDataFromRequest(c *gin.Context) (patchCreativeSpaceRequestData, error) {
	dataFromUri := patchCreativeSpaceRequestUri{}
	dataFromBody := patchCreativeSpaceRequestBody{}

	uriErr := c.BindUri(&dataFromUri)

	if uriErr != nil {
		return patchCreativeSpaceRequestData{}, uriErr
	}

	bodyErr := c.BindJSON(&dataFromBody)

	if bodyErr != nil {
		return patchCreativeSpaceRequestData{}, bodyErr
	}

	data := patchCreativeSpaceRequestData{}

	data.Id = dataFromUri.Id

	data.Title = dataFromBody.Title
	data.Address = dataFromBody.Address
	data.Description = dataFromBody.Description
	data.Photos = dataFromBody.Photos
	data.PricePerHour = dataFromBody.PricePerHour
	data.MetroStations = dataFromBody.MetroStations
	data.Coordinate = dataFromBody.Coordinate

	return data, nil
}

type deleteCreativeSpaceRequestData struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

func DeleteCreativeSpaceDataFromRequest(c *gin.Context) (deleteCreativeSpaceRequestData, error) {
	data := deleteCreativeSpaceRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteCreativeSpaceRequestData{}, err
	}

	return data, nil
}
