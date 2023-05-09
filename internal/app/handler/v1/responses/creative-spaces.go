package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type getCreativeSpacesResponseWorkingHours struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
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
	LandlordId    int                                     `json:"landlordId"`
	Description   string                                  `json:"description"`
	Photos        []string                                `json:"photos"`
	PricePerHour  int                                     `json:"pricePerHour"`
	MetroStations []getCreativeSpacesResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpacesResponseCoordinate     `json:"coordinate"`
	WorkingHours  getCreativeSpacesResponseWorkingHours   `json:"workingHours"`
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

		res.CreativeSpaces = append(res.CreativeSpaces, getCreativeSpacesResponseCreativeSpace{
			Id:           creativeSpace.Id,
			LandlordId:   creativeSpace.LandlordId,
			Description:  creativeSpace.Description,
			Photos:       creativeSpace.Photos,
			PricePerHour: creativeSpace.PricePerHour,
			Coordinate: getCreativeSpacesResponseCoordinate{
				Latitude:  creativeSpace.Latitude,
				Longitude: creativeSpace.Longitude,
			},
			WorkingHours: getCreativeSpacesResponseWorkingHours{
				StartAt: creativeSpace.WorkingHoursStartAt,
				EndAt:   creativeSpace.WorkingHoursEndAt,
			},
			MetroStations: resMetroStations,
		})
	}

	return res
}

type getCreativeSpaceResponseWorkingHours struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
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
	LandlordId    int                                    `json:"landlordId"`
	Description   string                                 `json:"description"`
	Photos        []string                               `json:"photos"`
	PricePerHour  int                                    `json:"pricePerHour"`
	MetroStations []getCreativeSpaceResponseMetroStation `json:"metroStations"`
	Coordinate    getCreativeSpaceResponseCoordinate     `json:"coordinate"`
	WorkingHours  getCreativeSpaceResponseWorkingHours   `json:"workingHours"`
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

	res.CreativeSpace = getCreativeSpaceResponseCreativeSpace{
		Id:           creativeSpace.Id,
		LandlordId:   creativeSpace.LandlordId,
		Description:  creativeSpace.Description,
		Photos:       creativeSpace.Photos,
		PricePerHour: creativeSpace.PricePerHour,
		Coordinate: getCreativeSpaceResponseCoordinate{
			Latitude:  creativeSpace.Latitude,
			Longitude: creativeSpace.Longitude,
		},
		WorkingHours: getCreativeSpaceResponseWorkingHours{
			StartAt: creativeSpace.WorkingHoursStartAt,
			EndAt:   creativeSpace.WorkingHoursEndAt,
		},
		MetroStations: resMetroStations,
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
