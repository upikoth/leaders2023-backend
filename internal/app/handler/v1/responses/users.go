package responses

import (
	modelStore "github.com/upikoth/leaders2023-backend/internal/app/model/store"
)

type getUsersResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type getUsersResponseData struct {
	Users []getUsersResponseUser `json:"users"`
}

func GetUsersResponseFromStoreData(users []modelStore.User) getUsersResponseData {
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

func GetUserResponseFromStoreData(user modelStore.User) getUserResponseData {
	res := getUserResponseData{}

	res.User = getUserResponseUser{
		Id:    user.Id,
		Email: user.Email,
	}

	return res
}

type createUserResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type createUserResponseData struct {
	User createUserResponseUser `json:"user"`
}

func CreateUserResponseFromStoreData(user modelStore.User) createUserResponseData {
	res := createUserResponseData{}

	res.User = createUserResponseUser{
		Id:    user.Id,
		Email: user.Email,
	}

	return res
}

type patchUserResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type patchUserResponseData struct {
	User patchUserResponseUser `json:"user"`
}

func PatchUserResponseFromStoreData(user modelStore.User) patchUserResponseData {
	res := patchUserResponseData{}

	res.User = patchUserResponseUser{
		Id:    user.Id,
		Email: user.Email,
	}

	return res
}
