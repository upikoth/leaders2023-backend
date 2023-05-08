package requests

import "github.com/gin-gonic/gin"

type createSessionRequestData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func CreateSessionDataFromRequest(c *gin.Context) (createSessionRequestData, error) {
	data := createSessionRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createSessionRequestData{}, err
	}

	return data, nil
}
