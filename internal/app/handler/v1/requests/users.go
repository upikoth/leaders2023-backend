package requests

import (
	"github.com/gin-gonic/gin"
)

type getUserRequestData struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

func GetUserDataFromRequest(c *gin.Context) (getUserRequestData, error) {
	data := getUserRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return getUserRequestData{}, err
	}

	return data, nil
}

type createUserRequestData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func CreateUserDataFromRequest(c *gin.Context) (createUserRequestData, error) {
	data := createUserRequestData{}

	err := c.BindJSON(&data)

	if err != nil {
		return createUserRequestData{}, err
	}

	return data, nil
}

type patchUserRequestUri struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

type patchUserRequestBody struct {
	Email string `json:"email,omitempty" binding:"email"`
}

type patchUserRequestData struct {
	Id    int    `json:"id"`
	Email string `json:"email,omitempty"`
}

func PatchUserDataFromRequest(c *gin.Context) (patchUserRequestData, error) {
	dataFromUri := patchUserRequestUri{}
	dataFromBody := patchUserRequestBody{}

	uriErr := c.BindUri(&dataFromUri)

	if uriErr != nil {
		return patchUserRequestData{}, uriErr
	}

	bodyErr := c.BindJSON(&dataFromBody)

	if bodyErr != nil {
		return patchUserRequestData{}, bodyErr
	}

	data := patchUserRequestData{}
	data.Id = dataFromUri.Id
	data.Email = dataFromBody.Email

	return data, nil
}

type deleteUserRequestData struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

func DeleteUserDataFromRequest(c *gin.Context) (deleteUserRequestData, error) {
	data := deleteUserRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteUserRequestData{}, err
	}

	return data, nil
}
