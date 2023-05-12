package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type createSessionResponseUser struct {
	Id    int        `json:"id"`
	Phone string     `json:"phone"`
	Role  model.Role `json:"role"`
}

type createSessionResponseData struct {
	User createSessionResponseUser `json:"user"`
}

func CreateSessionResponseFromStoreData(user store.User) createSessionResponseData {
	res := createSessionResponseData{}

	res.User = createSessionResponseUser{
		Id:    user.Id,
		Phone: user.Phone,
		Role:  user.Role,
	}

	return res
}

type getSessionResponseUser struct {
	Id    int        `json:"id"`
	Phone string     `json:"phone"`
	Role  model.Role `json:"role"`
}

type getSessionResponseData struct {
	User getSessionResponseUser `json:"user"`
}

func GetSessionResponseFromStoreData(user store.User) getSessionResponseData {
	res := getSessionResponseData{}

	res.User = getSessionResponseUser{
		Id:    user.Id,
		Phone: user.Phone,
		Role:  user.Role,
	}

	return res
}
