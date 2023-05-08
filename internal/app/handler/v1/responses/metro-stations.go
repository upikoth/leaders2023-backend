package responses

import (
	modelStore "github.com/upikoth/leaders2023-backend/internal/app/model/store"
)

type getMetroStationsResponseMetroStation struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type getMetroStationsResponseData struct {
	MetroStations []getMetroStationsResponseMetroStation `json:"metroStations"`
}

func GetMetroStationsResponseFromStoreData(metroStations []modelStore.MetroStation) getMetroStationsResponseData {
	res := getMetroStationsResponseData{}

	for _, station := range metroStations {
		resStation := getMetroStationsResponseMetroStation{
			Id:    station.Id,
			Name:  station.Name,
			Color: station.Color,
		}

		res.MetroStations = append(res.MetroStations, resStation)
	}

	return res
}
