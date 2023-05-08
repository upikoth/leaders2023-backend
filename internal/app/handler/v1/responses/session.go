package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type createSessionResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type createSessionResponseData struct {
	User createSessionResponseUser `json:"user"`
}

func CreateSessionResponseFromStoreData(user store.User) createSessionResponseData {
	res := createSessionResponseData{}

	res.User = createSessionResponseUser{
		Id:    user.Id,
		Email: user.Email,
	}

	return res
}
