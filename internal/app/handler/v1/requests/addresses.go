package requests

import (
	"github.com/gin-gonic/gin"
)

type getAddressesRequestData struct {
	Query string `form:"query" binding:"required"`
}

func GetAddressesDataFromRequest(c *gin.Context) (getAddressesRequestData, error) {
	data := getAddressesRequestData{}

	err := c.BindQuery(&data)

	if err != nil {
		return getAddressesRequestData{}, err
	}

	return data, nil
}
