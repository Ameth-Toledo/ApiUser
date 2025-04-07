package application

import "UsersFree/src/users/domain"

type DeleteUser struct {
	db domain.IUser
}

func NewDeleteUser(db domain.IUser) *DeleteUser {
	return &DeleteUser{db: db}
}

func (dc DeleteUser) Execute(id int) error {
	return dc.db.Delete(id)
}
