package application

import (
	"UsersFree/src/core/security"
	"UsersFree/src/users/domain"
	"UsersFree/src/users/domain/entities"
	"fmt"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cc *CreateUser) Execute(client entities.User) error {
	existingUser, err := cc.db.GetByEsp32Serial(*client.Id_esp32)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("el número de serie del ESP32 ya está en uso, por favor ingrese otro")
	}

	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword

	if err := cc.db.Save(client); err != nil {
		return err
	}

	return nil
}
