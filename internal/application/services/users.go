package services

import (
	"errors"

	"github.com/kabanasy1/app/internal/application/db"
	"github.com/kabanasy1/app/internal/application/models"
)

type UserService struct {
	db *db.UserStorage
}

func NewUserSvc(udb *db.UserStorage) *UserService {
	svc := new(UserService)
	svc.db = udb
	return svc
}

func (svc *UserService) CreateUser(user models.User) error {
	if user.Firstname == "" {
		return errors.New("name filed should not be empty")
	}
	if user.Lastname == "" {
		return errors.New("surname filed should not be empty")
	}

	return svc.db.CreateUser(user)
}

func (svc *UserService) FindUser(id int64) (models.User, error) {
	user := svc.db.GetUserById(id)

	if user.Id != id {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (svc *UserService) ListUsers(username string) ([]models.User, error) {
	return svc.db.GetListUsers(username), nil
}
