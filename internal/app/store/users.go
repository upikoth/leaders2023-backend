package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

type User struct {
	tableName    struct{}   `pg:"users"` //nolint:unused // Имя таблицы
	Id           int        `pg:"id"`
	Phone        string     `pg:"phone"`
	Role         model.Role `pg:"role"`
	PasswordHash string     `pg:"password_hash"`
}

func (s *Store) GetUsers() ([]User, error) {
	users := []User{}

	err := s.db.
		Model(&users).
		Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) GetUserById(id int) (User, error) {
	user := User{
		Id: id,
	}

	count, err := s.db.
		Model(&user).
		WherePK().
		SelectAndCount()

	if err != nil {
		return User{}, err
	}

	if count == 0 {
		return User{}, constants.ErrUserGetNotFoundById
	}

	return user, nil
}

func (s *Store) CreateUser(user User) (int, error) {
	result, err := s.db.
		Model(&user).
		OnConflict("DO NOTHING").
		Insert()

	if err != nil {
		return 0, err
	}

	if result.RowsAffected() == 0 {
		return 0, constants.ErrUserPostPhoneExist
	}

	return user.Id, nil
}

func (s *Store) DeleteUser(id int) error {
	user := User{
		Id: id,
	}

	result, err := s.db.
		Model(&user).
		WherePK().
		Delete()

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return constants.ErrUserDeleteNotFoundById
	}

	return nil
}

func (s *Store) PatchUser(user User) error {
	count, err := s.db.
		Model(&user).
		WherePK().
		Count()

	if err != nil {
		return err
	}

	if count == 0 {
		return constants.ErrUserPatchNotFoundById
	}

	result, err := s.db.
		Model(&user).
		WherePK().
		OnConflict("DO NOTHING").
		UpdateNotZero()

	if err != nil {
		return err
	}

	if result == nil {
		return constants.ErrUserPatchPhoneExist
	}

	return nil
}

func (s *Store) GetUserByPhone(phone string) (User, error) {
	user := User{
		Phone: phone,
	}

	count, err := s.db.
		Model(&user).
		Where("phone = ?", user.Phone).
		SelectAndCount()

	if err != nil {
		return User{}, constants.ErrUserGetByPhoneDbError
	}

	if count == 0 {
		return User{}, constants.ErrUserGetByPhoneUserNotExist
	}

	return user, err
}
