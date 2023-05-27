package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type getUsersResponseUser struct {
	Id              int        `json:"id"`
	Phone           string     `json:"phone"`
	Role            model.Role `json:"role"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
}

type getUsersResponseData struct {
	Users []getUsersResponseUser `json:"users"`
}

func GetUsersResponseFromStoreData(users []store.User) getUsersResponseData {
	res := getUsersResponseData{
		Users: []getUsersResponseUser{},
	}

	for _, user := range users {
		resUser := getUsersResponseUser{
			Id:              user.Id,
			Phone:           user.Phone,
			Role:            user.Role,
			Name:            user.Name,
			Surname:         user.Surname,
			Patronymic:      user.Patronymic,
			Email:           user.Email,
			Inn:             user.Inn,
			LegalEntityName: user.LegalEntityName,
		}

		res.Users = append(res.Users, resUser)
	}

	return res
}

type getUserResponseUser struct {
	Id              int        `json:"id"`
	Phone           string     `json:"phone"`
	Role            model.Role `json:"role"`
	Name            string     `json:"name"`
	Surname         string     `json:"surname"`
	Patronymic      string     `json:"patronymic"`
	Email           string     `json:"email"`
	Inn             string     `json:"inn"`
	LegalEntityName string     `json:"legalEntityName"`
}

type getUserResponseData struct {
	User getUserResponseUser `json:"user"`
}

func GetUserResponseFromStoreData(user store.User) getUserResponseData {
	res := getUserResponseData{}

	res.User = getUserResponseUser{
		Id:              user.Id,
		Phone:           user.Phone,
		Role:            user.Role,
		Name:            user.Name,
		Surname:         user.Surname,
		Patronymic:      user.Patronymic,
		Email:           user.Email,
		Inn:             user.Inn,
		LegalEntityName: user.LegalEntityName,
	}

	return res
}

type createUserResponseUser struct {
	Id int `json:"id"`
}

type createUserResponseData struct {
	User createUserResponseUser `json:"user"`
}

func CreateUserResponseFromStoreData(id int) createUserResponseData {
	res := createUserResponseData{}

	res.User = createUserResponseUser{
		Id: id,
	}

	return res
}
