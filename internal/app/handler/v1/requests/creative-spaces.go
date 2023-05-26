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

type createCreativeSpaceRequestCalendarEvent struct {
	Date string `json:"date" binding:"required"`
}

type createCreativeSpaceRequestCalendar struct {
	WorkDayIndexes []int                                     `json:"workDayIndexes" binding:"required"`
	Events         []createCreativeSpaceRequestCalendarEvent `json:"events" binding:"required"`
	Link           string                                    `json:"link"`
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
	PricePerDay   int                                      `json:"pricePerDay" binding:"required"`
	MetroStations []createCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    createCreativeSpaceRequestCoordinate     `json:"coordinate" binding:"required"`
	Calendar      createCreativeSpaceRequestCalendar       `json:"calendar" binding:"required"`
}

func CreateCreativeSpaceDataFromRequest(c *gin.Context) (createCreativeSpaceRequestData, error) {
	data := createCreativeSpaceRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createCreativeSpaceRequestData{}, err
	}

	return data, nil
}

type patchCreativeSpaceRequestCalendarEvent struct {
	Date string `json:"date"`
}

type patchCreativeSpaceRequestCalendar struct {
	WorkDayIndexes []int                                    `json:"workDayIndexes"`
	Events         []patchCreativeSpaceRequestCalendarEvent `json:"events"`
	Link           string                                   `json:"link"`
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
	Id int `uri:"id" binding:"required"`
}

type patchCreativeSpaceRequestBody struct {
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
	Calendar      patchCreativeSpaceRequestCalendar       `json:"calendar"`
}

type patchCreativeSpaceRequestData struct {
	Id            int                                     `json:"id"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
	Calendar      patchCreativeSpaceRequestCalendar       `json:"calendar"`
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
	data.PricePerDay = dataFromBody.PricePerDay
	data.MetroStations = dataFromBody.MetroStations
	data.Coordinate = dataFromBody.Coordinate

	data.Calendar = patchCreativeSpaceRequestCalendar{
		WorkDayIndexes: dataFromBody.Calendar.WorkDayIndexes,
		Events:         dataFromBody.Calendar.Events,
		Link:           dataFromBody.Calendar.Link,
	}

	return data, nil
}

type deleteCreativeSpaceRequestData struct {
	Id int `uri:"id" binding:"required"`
}

func DeleteCreativeSpaceDataFromRequest(c *gin.Context) (deleteCreativeSpaceRequestData, error) {
	data := deleteCreativeSpaceRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteCreativeSpaceRequestData{}, err
	}

	return data, nil
}
