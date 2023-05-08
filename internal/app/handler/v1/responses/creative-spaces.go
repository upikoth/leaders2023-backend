package responses

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

// type createCreativeSpaceResponseWorkingHours struct {
// 	StartAt string `json:"startAt"`
// 	EndAt   string `json:"endAt"`
// }

// type createCreativeSpaceResponseCoordinate struct {
// 	Latitude  float32 `json:"latitude"`
// 	Longitude float32 `json:"longitude"`
// }

// type createCreativeSpaceResponseMetroStation struct {
// 	Id                string `json:"id"`
// 	DistanceInMinutes int    `json:"distanceInMinutes"`
// }

// type createCreativeSpaceResponseCreativeSpace struct {
// 	Id            int                                       `json:"id"`
// 	LandlordId    int                                       `json:"landlordId"`
// 	Description   string                                    `json:"description"`
// 	Photos        []string                                  `json:"photos"`
// 	PricePerHour  int                                       `json:"pricePerHour"`
// 	MetroStations []createCreativeSpaceResponseMetroStation `json:"metroStations"`
// 	Coordinate    createCreativeSpaceResponseCoordinate     `json:"coordinate"`
// 	WorkingHours  createCreativeSpaceResponseWorkingHours   `json:"workingHours"`
// }
