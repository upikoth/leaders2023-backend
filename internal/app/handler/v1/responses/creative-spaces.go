package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type createCreativeSpaceResponseWorkingHours struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

type createCreativeSpaceResponseCoordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type createCreativeSpaceResponseMetroStation struct {
	Id                int `json:"id"`
	DistanceInMinutes int `json:"distanceInMinutes"`
}

type getCreativeSpacesResponseCreativeSpace struct {
	Id            int                                       `json:"id"`
	LandlordId    int                                       `json:"landlordId"`
	Description   string                                    `json:"description"`
	Photos        []string                                  `json:"photos"`
	PricePerHour  int                                       `json:"pricePerHour"`
	MetroStations []createCreativeSpaceResponseMetroStation `json:"metroStations"`
	Coordinate    createCreativeSpaceResponseCoordinate     `json:"coordinate"`
	WorkingHours  createCreativeSpaceResponseWorkingHours   `json:"workingHours"`
}

type getCreativeSpacesResponseData struct {
	CreativeSpaces []getCreativeSpacesResponseCreativeSpace `json:"creativeSpaces"`
}

func GetCreativeSpacesResponseFromStoreData(creativeSpaces []store.CreativeSpace) getCreativeSpacesResponseData {
	res := getCreativeSpacesResponseData{}

	for _, creativeSpace := range creativeSpaces {
		resMetroStations := []createCreativeSpaceResponseMetroStation{}

		for _, metroStation := range creativeSpace.MetroStations {
			resMetroStations = append(resMetroStations, createCreativeSpaceResponseMetroStation{
				Id:                metroStation.MetroStationId,
				DistanceInMinutes: metroStation.DistanceInMinutes,
			})
		}

		res.CreativeSpaces = append(res.CreativeSpaces, getCreativeSpacesResponseCreativeSpace{
			Id:           creativeSpace.Id,
			LandlordId:   creativeSpace.LandlordId,
			Description:  creativeSpace.Description,
			Photos:       creativeSpace.Photos,
			PricePerHour: creativeSpace.PricePerHour,
			Coordinate: createCreativeSpaceResponseCoordinate{
				Latitude:  creativeSpace.Latitude,
				Longitude: creativeSpace.Longitude,
			},
			WorkingHours: createCreativeSpaceResponseWorkingHours{
				StartAt: creativeSpace.WorkingHoursStartAt,
				EndAt:   creativeSpace.WorkingHoursEndAt,
			},
			MetroStations: resMetroStations,
		})
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
