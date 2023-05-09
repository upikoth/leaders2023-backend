package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type User struct {
	tableName    struct{} `pg:"users"` //nolint:unused // Имя таблицы
	Id           int      `pg:"id"`
	Email        string   `pg:"email"`
	PasswordHash string   `pg:"password_hash"`
}

func (s *Store) GetUsers() ([]User, error) {
	users := []User{}

	err := s.db.Model(&users).Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) GetUserById(id int) (User, error) {
	user := User{
		Id: id,
	}

	count, err := s.db.Model(&user).WherePK().SelectAndCount()

	if count == 0 {
		return User{}, constants.ErrUserGetNotFoundById
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *Store) CreateUser(user User) (int, error) {
	result, err := s.db.Model(&user).OnConflict("DO NOTHING").Insert()

	if result.RowsAffected() == 0 {
		return 0, constants.ErrUserPostEmailExist
	}

	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (s *Store) DeleteUser(id int) error {
	user := User{
		Id: id,
	}

	result, err := s.db.Model(&user).WherePK().Delete()

	if result.RowsAffected() == 0 {
		return constants.ErrUserDeleteNotFoundById
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) PatchUser(user User) error {
	count, err := s.db.Model(&user).WherePK().Count()

	if err != nil {
		return err
	}

	if count == 0 {
		return constants.ErrUserPatchNotFoundById
	}

	result, err := s.db.Model(&user).WherePK().OnConflict("DO NOTHING").UpdateNotZero()

	if result == nil {
		return constants.ErrUserPatchEmailExist
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (User, error) {
	user := User{
		Email: email,
	}

	count, err := s.db.Model(&user).Where("email = ?", user.Email).SelectAndCount()

	if err != nil {
		return User{}, constants.ErrUserGetByEmailDbError
	}

	if count == 0 {
		return User{}, constants.ErrUserGetByEmailUserNotExist
	}

	return user, err
}
