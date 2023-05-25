package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

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

type getCreativeSpacesResponseCreativeSpace struct {
	Id            int                                     `json:"id"`
	Title         string                                  `json:"title"`
	Address       string                                  `json:"address"`
	LandlordId    int                                     `json:"landlordId"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerDay   int                                     `json:"pricePerDay"`
	MetroStations []getCreativeSpacesResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpacesResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpacesResponseCalendar       `json:"calendar"`
}

type getCreativeSpacesResponseData struct {
	CreativeSpaces []getCreativeSpacesResponseCreativeSpace `json:"creativeSpaces"`
}

func GetCreativeSpacesResponseFromStoreData(creativeSpaces []store.CreativeSpace) getCreativeSpacesResponseData {
	res := getCreativeSpacesResponseData{}

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
			Title:       creativeSpace.Title,
			Address:     creativeSpace.Address,
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
		})
	}

	return res
}

type getCreativeSpaceResponseCalendarEvent struct {
	Date string `json:"date"`
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

type getCreativeSpaceResponseCreativeSpace struct {
	Id            int                                    `json:"id"`
	Title         string                                 `json:"title"`
	Address       string                                 `json:"address"`
	LandlordId    int                                    `json:"landlordId"`
	Description   string                                 `json:"description"`
	Photos        []string                               `json:"photos"`
	PricePerDay   int                                    `json:"pricePerDay"`
	MetroStations []getCreativeSpaceResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpaceResponseCoordinate     `json:"coordinate"`
	Calendar      getCreativeSpaceResponseCalendar       `json:"calendar"`
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
			Date: calendarEvent.Date,
		})
	}

	res.CreativeSpace = getCreativeSpaceResponseCreativeSpace{
		Id:          creativeSpace.Id,
		Title:       creativeSpace.Title,
		Address:     creativeSpace.Address,
		LandlordId:  creativeSpace.LandlordId,
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
