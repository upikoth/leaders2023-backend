package responses

import (
	"math"

	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getCreativeSpacesResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getCreativeSpacesResponseCalendar struct {
	WorkDayIndexes string                                   `json:"workDayIndexes"`
	Events         []getCreativeSpacesResponseCalendarEvent `json:"events"`
	Link           string                                   `json:"link"`
}

type getCreativeSpacesResponseCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type getCreativeSpacesResponseMetroStation struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	DistanceInMinutes int    `json:"distanceInMinutes"`
}

type getCreativeSpacesResponseScoreUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type getCreativeSpacesResponseScore struct {
	ID      string                             `json:"id"`
	Comment string                             `json:"comment"`
	Rating  int                                `json:"rating"`
	User    getCreativeSpacesResponseScoreUser `json:"user"`
}

type getCreativeSpacesResponseCreativeSpace struct {
	ID            string                                  `json:"id"`
	SpaceType     string                                  `json:"spaceType"`
	Area          int                                     `json:"area"`
	Capacity      int                                     `json:"capacity"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Status        model.CreativeSpaceStatus               `json:"status"`
	LandlordID    string                                  `json:"landlordId"`
	Description   string                                  `json:"description"`
	Photos        string                                  `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []getCreativeSpacesResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpacesResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpacesResponseCalendar       `json:"calendar"`
	Scores        []getCreativeSpacesResponseScore        `json:"scores"`
	AverageRating float64                                 `json:"averageRating"`
}

type getCreativeSpacesResponseData struct {
	CreativeSpaces []getCreativeSpacesResponseCreativeSpace `json:"creativeSpaces"`
}

func GetCreativeSpacesResponseFromStoreData(creativeSpaces []store.CreativeSpace) getCreativeSpacesResponseData {
	res := getCreativeSpacesResponseData{
		CreativeSpaces: []getCreativeSpacesResponseCreativeSpace{},
	}

	for _, creativeSpace := range creativeSpaces {
		resMetroStations := []getCreativeSpacesResponseMetroStation{}

		for _, metroStation := range creativeSpace.MetroStations {
			resMetroStations = append(resMetroStations, getCreativeSpacesResponseMetroStation{
				ID:                metroStation.MetroStationID,
				Name:              metroStation.MetroStation.Name,
				Color:             metroStation.MetroStation.Color,
				DistanceInMinutes: metroStation.DistanceInMinutes,
			})
		}

		resScores := []getCreativeSpacesResponseScore{}
		totalRating := 0
		averageRating := 0.0

		for _, score := range creativeSpace.Scores {
			resScores = append(resScores, getCreativeSpacesResponseScore{
				ID:      score.ID,
				Rating:  score.Rating,
				Comment: score.Comment,
				User: getCreativeSpacesResponseScoreUser{
					ID:         score.User.ID,
					Name:       score.User.Name,
					Surname:    score.User.Surname,
					Patronymic: score.User.Patronymic,
				},
			})

			totalRating += score.Rating
		}

		if len(resScores) > 0 {
			averageRating = float64(totalRating) / float64(len(resScores))
			//nolint:gomnd // Оставляем 1 знак после зяпятой.
			averageRating = math.Round(averageRating*10) / 10
		}

		resCalendarEvents := []getCreativeSpacesResponseCalendarEvent{}

		for _, calendarEvent := range creativeSpace.CalendarEvents {
			resCalendarEvents = append(resCalendarEvents, getCreativeSpacesResponseCalendarEvent{
				Date: calendarEvent.Date,
			})
		}

		resCalendar := getCreativeSpacesResponseCalendar{
			WorkDayIndexes: creativeSpace.CalendarWorkDayIndexes,
			Events:         resCalendarEvents,
			Link:           creativeSpace.CalendarLink,
		}

		res.CreativeSpaces = append(res.CreativeSpaces, getCreativeSpacesResponseCreativeSpace{
			ID:          creativeSpace.ID,
			SpaceType:   creativeSpace.SpaceType,
			Area:        creativeSpace.Area,
			Capacity:    creativeSpace.Capacity,
			Title:       creativeSpace.Title,
			Address:     creativeSpace.Address,
			Status:      model.CreativeSpaceStatus(creativeSpace.Status),
			LandlordID:  creativeSpace.LandlordID,
			Description: creativeSpace.Description,
			Photos:      creativeSpace.Photos,
			PricePerDay: creativeSpace.PricePerDay,
			Coordinate: getCreativeSpacesResponseCoordinate{
				Latitude:  creativeSpace.Latitude,
				Longitude: creativeSpace.Longitude,
			},
			MetroStations: resMetroStations,
			Calendar:      resCalendar,
			Scores:        resScores,
			AverageRating: averageRating,
		})
	}

	return res
}

type getCreativeSpaceResponseCalendarEvent struct {
	Date      string `json:"date"`
	BookingID string `json:"bookingId"`
}

type getCreativeSpaceResponseLandlordInfo struct {
	ID              string     `json:"id"`
	Phone           string     `json:"phone"`
	Role            model.Role `json:"role"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
}

type getCreativeSpaceResponseCalendar struct {
	WorkDayIndexes string                                  `json:"workDayIndexes"`
	Events         []getCreativeSpaceResponseCalendarEvent `json:"events"`
	Link           string                                  `json:"link"`
}

type getCreativeSpaceResponseCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type getCreativeSpaceResponseMetroStation struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	DistanceInMinutes int    `json:"distanceInMinutes"`
}

type getCreativeSpaceResponseScoreUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type getCreativeSpaceResponseScore struct {
	ID      string                            `json:"id"`
	Comment string                            `json:"comment"`
	Rating  int                               `json:"rating"`
	User    getCreativeSpaceResponseScoreUser `json:"user"`
}

type getCreativeSpaceResponseCreativeSpace struct {
	ID            string                                 `json:"id"`
	SpaceType     string                                 `json:"spaceType"`
	Area          int                                    `json:"area"`
	Capacity      int                                    `json:"capacity"`
	Title         string                                 `json:"title"`
	Address       string                                 `json:"address"`
	Status        model.CreativeSpaceStatus              `json:"status"`
	Description   string                                 `json:"description"`
	Photos        string                                 `json:"photos"`
	PricePerDay   int                                    `json:"pricePerDay"`
	MetroStations []getCreativeSpaceResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpaceResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpaceResponseCalendar       `json:"calendar"`
	LandlordInfo  getCreativeSpaceResponseLandlordInfo   `json:"landlordInfo"`
	Scores        []getCreativeSpaceResponseScore        `json:"scores"`
	AverageRating float64                                `json:"averageRating"`
}

type getCreativeSpaceResponseData struct {
	CreativeSpace getCreativeSpaceResponseCreativeSpace `json:"creativeSpace"`
}

func GetCreativeSpaceResponseFromStoreData(creativeSpace store.CreativeSpace) getCreativeSpaceResponseData {
	res := getCreativeSpaceResponseData{}
	resMetroStations := []getCreativeSpaceResponseMetroStation{}

	for _, metroStation := range creativeSpace.MetroStations {
		resMetroStations = append(resMetroStations, getCreativeSpaceResponseMetroStation{
			ID:                metroStation.MetroStationID,
			Name:              metroStation.MetroStation.Name,
			Color:             metroStation.MetroStation.Color,
			DistanceInMinutes: metroStation.DistanceInMinutes,
		})
	}

	resCalendarEvents := []getCreativeSpaceResponseCalendarEvent{}

	for _, calendarEvent := range creativeSpace.CalendarEvents {
		resCalendarEvents = append(resCalendarEvents, getCreativeSpaceResponseCalendarEvent{
			Date:      calendarEvent.Date,
			BookingID: calendarEvent.BookingID,
		})
	}

	resScores := []getCreativeSpaceResponseScore{}
	totalRating := 0
	averageRating := 0.0

	for _, score := range creativeSpace.Scores {
		resScores = append(resScores, getCreativeSpaceResponseScore{
			ID:      score.ID,
			Rating:  score.Rating,
			Comment: score.Comment,
			User: getCreativeSpaceResponseScoreUser{
				ID:         score.User.ID,
				Name:       score.User.Name,
				Surname:    score.User.Surname,
				Patronymic: score.User.Patronymic,
			},
		})

		totalRating += score.Rating
	}

	if len(resScores) > 0 {
		averageRating = float64(totalRating) / float64(len(resScores))
		//nolint:gomnd // Оставляем 1 знак после зяпятой.
		averageRating = math.Round(averageRating*10) / 10
	}

	res.CreativeSpace = getCreativeSpaceResponseCreativeSpace{
		ID:          creativeSpace.ID,
		SpaceType:   creativeSpace.SpaceType,
		Area:        creativeSpace.Area,
		Capacity:    creativeSpace.Capacity,
		Title:       creativeSpace.Title,
		Address:     creativeSpace.Address,
		Status:      model.CreativeSpaceStatus(creativeSpace.Status),
		Description: creativeSpace.Description,
		Photos:      creativeSpace.Photos,
		PricePerDay: creativeSpace.PricePerDay,
		Coordinate: getCreativeSpaceResponseCoordinate{
			Latitude:  creativeSpace.Latitude,
			Longitude: creativeSpace.Longitude,
		},
		MetroStations: resMetroStations,
		Calendar: getCreativeSpaceResponseCalendar{
			WorkDayIndexes: creativeSpace.CalendarWorkDayIndexes,
			Events:         resCalendarEvents,
			Link:           creativeSpace.CalendarLink,
		},
		LandlordInfo: getCreativeSpaceResponseLandlordInfo{
			ID:              creativeSpace.LandlordInfo.ID,
			Phone:           creativeSpace.LandlordInfo.Phone,
			Role:            model.Role(creativeSpace.LandlordInfo.Role),
			Name:            creativeSpace.LandlordInfo.Name,
			Surname:         creativeSpace.LandlordInfo.Surname,
			Patronymic:      creativeSpace.LandlordInfo.Patronymic,
			Email:           creativeSpace.LandlordInfo.Email,
			Inn:             creativeSpace.LandlordInfo.Inn,
			LegalEntityName: creativeSpace.LandlordInfo.LegalEntityName,
		},
		Scores:        resScores,
		AverageRating: averageRating,
	}

	return res
}

type createCreativeSpaceResponseCreativeSpace struct {
	ID string `json:"id"`
}

type createCreativeSpaceResponseData struct {
	CreativeSpace createCreativeSpaceResponseCreativeSpace `json:"creativeSpace"`
}

func CreateCreativeSpaceResponseFromStoreData(creativeSpaceID string) createCreativeSpaceResponseData {
	res := createCreativeSpaceResponseData{}

	res.CreativeSpace = createCreativeSpaceResponseCreativeSpace{
		ID: creativeSpaceID,
	}

	return res
}
