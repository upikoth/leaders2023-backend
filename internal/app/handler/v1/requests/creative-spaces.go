package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

type getCreativeSpaceRequestData struct {
	ID string `json:"id" uri:"id" binding:"required"`
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
	WorkDayIndexes string                                    `json:"workDayIndexes" binding:"required"`
	Events         []createCreativeSpaceRequestCalendarEvent `json:"events" binding:"required"`
	Link           string                                    `json:"link"`
}

type createCreativeSpaceRequestCoordinate struct {
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type createCreativeSpaceRequestMetroStation struct {
	ID                string `json:"id" binding:"required"`
	DistanceInMinutes int    `json:"distanceInMinutes" binding:"required"`
}

type createCreativeSpaceRequestData struct {
	SpaceType     string                                   `json:"spaceType" binding:"required"`
	Area          int                                      `json:"area" binding:"required"`
	Capacity      int                                      `json:"capacity" binding:"required"`
	Title         string                                   `json:"title" binding:"required"`
	Address       string                                   `json:"address" binding:"required"`
	Description   string                                   `json:"description" binding:"required"`
	Photos        string                                   `json:"photos" binding:"required"`
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
	WorkDayIndexes string                                   `json:"workDayIndexes"`
	Events         []patchCreativeSpaceRequestCalendarEvent `json:"events" default:"nil"`
	Link           string                                   `json:"link"`
}

type patchCreativeSpaceRequestCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type patchCreativeSpaceRequestMetroStation struct {
	ID                string `json:"id"`
	DistanceInMinutes int    `json:"distanceInMinutes"`
}

type patchCreativeSpaceRequestURI struct {
	ID string `uri:"id" binding:"required"`
}

type patchCreativeSpaceRequestBody struct {
	SpaceType     string                                  `json:"spaceType"`
	Area          int                                     `json:"area"`
	Capacity      int                                     `json:"capacity"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Status        model.CreativeSpaceStatus               `json:"status"`
	Description   string                                  `json:"description"`
	Photos        string                                  `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
	Calendar      patchCreativeSpaceRequestCalendar       `json:"calendar"`
}

type patchCreativeSpaceRequestData struct {
	ID            string                                  `json:"id"`
	SpaceType     string                                  `json:"spaceType"`
	Area          int                                     `json:"area"`
	Capacity      int                                     `json:"capacity"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Status        model.CreativeSpaceStatus               `json:"status"`
	Description   string                                  `json:"description"`
	Photos        string                                  `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []patchCreativeSpaceRequestMetroStation `json:"metroStations"`
	Coordinate    patchCreativeSpaceRequestCoordinate     `json:"coordinate"`
	Calendar      patchCreativeSpaceRequestCalendar       `json:"calendar"`
}

func PatchCreativeSpaceDataFromRequest(c *gin.Context) (patchCreativeSpaceRequestData, error) {
	dataFromURI := patchCreativeSpaceRequestURI{}
	dataFromBody := patchCreativeSpaceRequestBody{}

	uriErr := c.BindUri(&dataFromURI)

	if uriErr != nil {
		return patchCreativeSpaceRequestData{}, uriErr
	}

	bodyErr := c.BindJSON(&dataFromBody)

	if bodyErr != nil {
		return patchCreativeSpaceRequestData{}, bodyErr
	}

	data := patchCreativeSpaceRequestData{}

	data.ID = dataFromURI.ID

	data.SpaceType = dataFromBody.SpaceType
	data.Area = dataFromBody.Area
	data.Capacity = dataFromBody.Capacity
	data.Title = dataFromBody.Title
	data.Address = dataFromBody.Address
	data.Status = dataFromBody.Status
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
	ID string `uri:"id" binding:"required"`
}

func DeleteCreativeSpaceDataFromRequest(c *gin.Context) (deleteCreativeSpaceRequestData, error) {
	data := deleteCreativeSpaceRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteCreativeSpaceRequestData{}, err
	}

	return data, nil
}
