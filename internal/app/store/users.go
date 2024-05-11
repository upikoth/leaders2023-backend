package store

import (
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type User struct {
	ID              string `gorm:"primarykey"`
	Name            string
	Surname         string
	Patronymic      string
	Email           string
	Inn             string
	LegalEntityName string
	Phone           string
	Role            string
	PasswordHash    string
}

func (s *Store) GetUsers() ([]User, error) {
	users := []User{}

	res := s.db.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (s *Store) GetUserByID(id string) (User, error) {
	user := User{ID: id}

	res := s.db.First(&user)

	if res.Error != nil {
		return User{}, res.Error
	}

	return user, nil
}

func (s *Store) CreateUser(user User) (string, error) {
	res := s.db.Create(&user)

	if res.Error != nil {
		return "", res.Error
	}

	return user.ID, nil
}

func (s *Store) DeleteUser(id string) error {
	res := s.db.Delete(&User{ID: id})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) PatchUser(user User) error {
	_, err := s.GetUserByID(user.ID)

	if err != nil {
		return err
	}

	res := s.db.Updates(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) GetUserByPhone(phone string) (User, error) {
	user := User{}

	res := s.db.
		Where("phone = ?", phone).
		First(&user)

	if res.RowsAffected == 0 {
		return user, nil
	}

	if res.Error != nil {
		return user, constants.ErrUserGetByPhoneDbError
	}

	return user, nil
}
