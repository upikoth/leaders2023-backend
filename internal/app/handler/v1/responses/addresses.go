package responses

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
)

type getAddressesResponseUser struct {
	Value     string `json:"value"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type getAddressesResponseData struct {
	Addresses []getAddressesResponseUser `json:"addresses"`
}

func GetAddressesResponseFromStoreData(addresses []*suggest.AddressSuggestion) getAddressesResponseData {
	res := getAddressesResponseData{
		Addresses: []getAddressesResponseUser{},
	}

	for _, address := range addresses {
		resAddress := getAddressesResponseUser{
			Value:     address.Value,
			Latitude:  address.Data.GeoLat,
			Longitude: address.Data.GeoLon,
		}

		res.Addresses = append(res.Addresses, resAddress)
	}

	return res
}
