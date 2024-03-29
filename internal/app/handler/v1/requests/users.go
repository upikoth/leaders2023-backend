package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
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
	Phone           string     `json:"phone" binding:"required"`
	Password        string     `json:"password" binding:"required"`
	Role            model.Role `json:"role" binding:"required"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
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
	Phone           string `json:"phone,omitempty"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Patronymic      string `json:"patronymic"`
	Email           string `json:"email"`
	Inn             string `json:"inn"`
	LegalEntityName string `json:"legalEntityName"`
}

type patchUserRequestData struct {
	Id              int    `json:"id"`
	Phone           string `json:"phone,omitempty"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Patronymic      string `json:"patronymic"`
	Email           string `json:"email"`
	Inn             string `json:"inn"`
	LegalEntityName string `json:"legalEntityName"`
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

	data := patchUserRequestData{
		Id:              dataFromUri.Id,
		Phone:           dataFromBody.Phone,
		Name:            dataFromBody.Name,
		Surname:         dataFromBody.Surname,
		Patronymic:      dataFromBody.Patronymic,
		Email:           dataFromBody.Email,
		Inn:             dataFromBody.Inn,
		LegalEntityName: dataFromBody.LegalEntityName,
	}

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
