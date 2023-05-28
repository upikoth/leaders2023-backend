package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getCreativeSpacesResponseCalendarEvent struct {
	Date string `json:"date"`
}

type getCreativeSpacesResponseCalendar struct {
	WorkDayIndexes []int                                    `json:"workDayIndexes"`
	Events         []getCreativeSpacesResponseCalendarEvent `json:"events"`
	Link           string                                   `json:"link"`
}

type getCreativeSpacesResponseCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type getCreativeSpacesResponseMetroStation struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	DistanceInMinutes int    `json:"distanceInMinutes"`
}

type getCreativeSpacesResponseScoreUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type getCreativeSpacesResponseScore struct {
	Id      int                                `json:"id"`
	Comment string                             `json:"comment"`
	Rating  int                                `json:"rating"`
	User    getCreativeSpacesResponseScoreUser `json:"user"`
}

type getCreativeSpacesResponseCreativeSpace struct {
	Id            int                                     `json:"id"`
	SpaceType     string                                  `json:"spaceType"`
	Area          int                                     `json:"area"`
	Capacity      int                                     `json:"capacity"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	Status        model.CreativeSpaceStatus               `json:"status"`
	LandlordId    int                                     `json:"landlordId"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []getCreativeSpacesResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpacesResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpacesResponseCalendar       `json:"calendar"`
	Scores        []getCreativeSpacesResponseScore        `json:"scores"`
	AverageRating int                                     `json:"averageRating"`
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
				Id:                metroStation.MetroStationId,
				Name:              metroStation.MetroStation.Name,
				Color:             metroStation.MetroStation.Color,
				DistanceInMinutes: metroStation.DistanceInMinutes,
			})
		}

		resScores := []getCreativeSpacesResponseScore{}
		totalRating := 0
		averageRating := 0

		for _, score := range creativeSpace.Scores {
			resScores = append(resScores, getCreativeSpacesResponseScore{
				Id:      score.Id,
				Rating:  score.Rating,
				Comment: score.Comment,
				User: getCreativeSpacesResponseScoreUser{
					Id:         score.User.Id,
					Name:       score.User.Name,
					Surname:    score.User.Surname,
					Patronymic: score.User.Patronymic,
				},
			})

			totalRating += score.Rating
		}

		if len(resScores) > 0 {
			averageRating = totalRating / len(resScores)
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
			Id:          creativeSpace.Id,
			SpaceType:   creativeSpace.SpaceType,
			Area:        creativeSpace.Area,
			Capacity:    creativeSpace.Capacity,
			Title:       creativeSpace.Title,
			Address:     creativeSpace.Address,
			Status:      creativeSpace.Status,
			LandlordId:  creativeSpace.LandlordId,
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
	BookingId int    `json:"bookingId"`
}

type getCreativeSpaceResponseLandlordInfo struct {
	Id              int        `json:"id"`
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
	WorkDayIndexes []int                                   `json:"workDayIndexes"`
	Events         []getCreativeSpaceResponseCalendarEvent `json:"events"`
	Link           string                                  `json:"link"`
}

type getCreativeSpaceResponseCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type getCreativeSpaceResponseMetroStation struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	DistanceInMinutes int    `json:"distanceInMinutes"`
}

type getCreativeSpaceResponseScoreUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type getCreativeSpaceResponseScore struct {
	Id      int                               `json:"id"`
	Comment string                            `json:"comment"`
	Rating  int                               `json:"rating"`
	User    getCreativeSpaceResponseScoreUser `json:"user"`
}

type getCreativeSpaceResponseCreativeSpace struct {
	Id            int                                    `json:"id"`
	SpaceType     string                                 `json:"spaceType"`
	Area          int                                    `json:"area"`
	Capacity      int                                    `json:"capacity"`
	Title         string                                 `json:"title"`
	Address       string                                 `json:"address"`
	Status        model.CreativeSpaceStatus              `json:"status"`
	Description   string                                 `json:"description"`
	Photos        []string                               `json:"photos"`
	PricePerDay   int                                    `json:"pricePerDay"`
	MetroStations []getCreativeSpaceResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpaceResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpaceResponseCalendar       `json:"calendar"`
	LandlordInfo  getCreativeSpaceResponseLandlordInfo   `json:"landlordInfo"`
	Scores        []getCreativeSpaceResponseScore        `json:"scores"`
	AverageRating int                                    `json:"averageRating"`
}

type getCreativeSpaceResponseData struct {
	CreativeSpace getCreativeSpaceResponseCreativeSpace `json:"creativeSpace"`
}

func GetCreativeSpaceResponseFromStoreData(creativeSpace store.CreativeSpace) getCreativeSpaceResponseData {
	res := getCreativeSpaceResponseData{}
	resMetroStations := []getCreativeSpaceResponseMetroStation{}

	for _, metroStation := range creativeSpace.MetroStations {
		resMetroStations = append(resMetroStations, getCreativeSpaceResponseMetroStation{
			Id:                metroStation.MetroStationId,
			Name:              metroStation.MetroStation.Name,
			Color:             metroStation.MetroStation.Color,
			DistanceInMinutes: metroStation.DistanceInMinutes,
		})
	}

	resCalendarEvents := []getCreativeSpaceResponseCalendarEvent{}

	for _, calendarEvent := range creativeSpace.CalendarEvents {
		resCalendarEvents = append(resCalendarEvents, getCreativeSpaceResponseCalendarEvent{
			Date:      calendarEvent.Date,
			BookingId: calendarEvent.BookingId,
		})
	}

	resScores := []getCreativeSpaceResponseScore{}
	totalRating := 0
	averageRating := 0

	for _, score := range creativeSpace.Scores {
		resScores = append(resScores, getCreativeSpaceResponseScore{
			Id:      score.Id,
			Rating:  score.Rating,
			Comment: score.Comment,
			User: getCreativeSpaceResponseScoreUser{
				Id:         score.User.Id,
				Name:       score.User.Name,
				Surname:    score.User.Surname,
				Patronymic: score.User.Patronymic,
			},
		})

		totalRating += score.Rating
	}

	if len(resScores) > 0 {
		averageRating = totalRating / len(resScores)
	}

	res.CreativeSpace = getCreativeSpaceResponseCreativeSpace{
		Id:          creativeSpace.Id,
		SpaceType:   creativeSpace.SpaceType,
		Area:        creativeSpace.Area,
		Capacity:    creativeSpace.Capacity,
		Title:       creativeSpace.Title,
		Address:     creativeSpace.Address,
		Status:      creativeSpace.Status,
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
			Id:              creativeSpace.LandlordInfo.Id,
			Phone:           creativeSpace.LandlordInfo.Phone,
			Role:            creativeSpace.LandlordInfo.Role,
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
	Id int `json:"id"`
}

type createCreativeSpaceResponseData struct {
	CreativeSpace createCreativeSpaceResponseCreativeSpace `json:"creativeSpace"`
}

func CreateCreativeSpaceResponseFromStoreData(creativeSpaceId int) createCreativeSpaceResponseData {
	res := createCreativeSpaceResponseData{}

	res.CreativeSpace = createCreativeSpaceResponseCreativeSpace{
		Id: creativeSpaceId,
	}

	return res
}
