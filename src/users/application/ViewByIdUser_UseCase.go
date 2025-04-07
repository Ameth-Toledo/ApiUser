package application

import (
	"UsersFree/src/users/domain"
	"UsersFree/src/users/domain/entities"
	"errors"
)

type ViewByIdUser struct {
	db domain.IUser
}

func NewUserById(db domain.IUser) *ViewByIdUser {
	return &ViewByIdUser{db: db}
}

func (vc *ViewByIdUser) Execute(id int) (entities.User, error) {
	user, err := vc.db.GetById(id)
	if err != nil {
		return entities.User{}, errors.New("User not found")
	}
	return user, nil
}
