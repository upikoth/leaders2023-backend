package responses

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type getUsersResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type getUsersResponseData struct {
	Users []getUsersResponseUser `json:"users"`
}

func GetUsersResponseFromStoreData(users []store.User) getUsersResponseData {
	res := getUsersResponseData{}

	for _, user := range users {
		resUser := getUsersResponseUser{
			Id:    user.Id,
			Email: user.Email,
		}

		res.Users = append(res.Users, resUser)
	}

	return res
}

type getUserResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type getUserResponseData struct {
	User getUserResponseUser `json:"user"`
}

func GetUserResponseFromStoreData(user store.User) getUserResponseData {
	res := getUserResponseData{}

	res.User = getUserResponseUser{
		Id:    user.Id,
		Email: user.Email,
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
