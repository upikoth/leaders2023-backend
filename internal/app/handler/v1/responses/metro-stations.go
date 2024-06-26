package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type getMetroStationsResponseMetroStation struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type getMetroStationsResponseData struct {
	MetroStations []getMetroStationsResponseMetroStation `json:"metroStations"`
}

func GetMetroStationsResponseFromStoreData(metroStations []store.MetroStation) getMetroStationsResponseData {
	res := getMetroStationsResponseData{
		MetroStations: []getMetroStationsResponseMetroStation{},
	}

	for _, station := range metroStations {
		resStation := getMetroStationsResponseMetroStation{
			ID:    station.ID,
			Name:  station.Name,
			Color: station.Color,
		}

		res.MetroStations = append(res.MetroStations, resStation)
	}

	return res
}
