package responses

import (
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type createSessionResponseUser struct {
	ID    string     `json:"id"`
	Phone string     `json:"phone"`
	Role  model.Role `json:"role"`
	Token string     `json:"token"`
}

type createSessionResponseData struct {
	User createSessionResponseUser `json:"user"`
}

func CreateSessionResponseFromStoreData(user store.User, jwtToken string) createSessionResponseData {
	res := createSessionResponseData{}

	res.User = createSessionResponseUser{
		ID:    user.ID,
		Phone: user.Phone,
		Role:  model.Role(user.Role),
		Token: jwtToken,
	}

	return res
}

type getSessionResponseUser struct {
	ID    string     `json:"id"`
	Phone string     `json:"phone"`
	Role  model.Role `json:"role"`
}

type getSessionResponseData struct {
	User getSessionResponseUser `json:"user"`
}

func GetSessionResponseFromStoreData(user store.User) getSessionResponseData {
	res := getSessionResponseData{}

	res.User = getSessionResponseUser{
		ID:    user.ID,
		Phone: user.Phone,
		Role:  model.Role(user.Role),
	}

	return res
}
