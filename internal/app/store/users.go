package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	modelStore "github.com/upikoth/leaders2023-backend/internal/app/model/store"
)

func (s *Store) GetUsers() ([]modelStore.User, error) {
	users := []modelStore.User{}

	err := s.db.Model(&users).Select()

	if err != nil {
		return nil, constants.ErrUsersGetDbError
	}

	return users, nil
}

func (s *Store) GetUserById(id int) (modelStore.User, error) {
	user := modelStore.User{
		Id: id,
	}

	count, err := s.db.Model(&user).WherePK().SelectAndCount()

	if count == 0 {
		return modelStore.User{}, constants.ErrUserGetNotFoundById
	}

	if err != nil {
		return modelStore.User{}, constants.ErrUserGetDbError
	}

	return user, nil
}

func (s *Store) CreateUser(user modelStore.User) (modelStore.User, error) {
	result, err := s.db.Model(&user).OnConflict("DO NOTHING").Insert()

	if result.RowsAffected() == 0 {
		return modelStore.User{}, constants.ErrUserPostEmailExist
	}

	if err != nil {
		return modelStore.User{}, constants.ErrUserPostDbError
	}

	return user, nil
}

func (s *Store) DeleteUser(id int) error {
	user := modelStore.User{
		Id: id,
	}

	result, err := s.db.Model(&user).WherePK().Delete()

	if result.RowsAffected() == 0 {
		return constants.ErrUserDeleteNotFoundById
	}

	if err != nil {
		return constants.ErrUserDeleteDbError
	}

	return nil
}

func (s *Store) PatchUser(user modelStore.User) (modelStore.User, error) {
	count, err := s.db.Model(&user).WherePK().Count()

	if err != nil {
		return modelStore.User{}, constants.ErrUserPatchDbError
	}

	if count == 0 {
		return modelStore.User{}, constants.ErrUserPatchNotFoundById
	}

	result, err := s.db.Model(&user).WherePK().OnConflict("DO NOTHING").UpdateNotZero()

	if result == nil {
		return modelStore.User{}, constants.ErrUserPatchEmailExist
	}

	if err != nil {
		return modelStore.User{}, constants.ErrUserPatchDbError
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (modelStore.User, error) {
	user := modelStore.User{
		Email: email,
	}

	count, err := s.db.Model(&user).Where("email = ?", user.Email).SelectAndCount()

	if err != nil {
		return modelStore.User{}, constants.ErrUserGetByEmailDbError
	}

	if count == 0 {
		return modelStore.User{}, constants.ErrUserGetByEmailUserNotExist
	}

	return user, err
}
