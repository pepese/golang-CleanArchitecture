package domain

import "github.com/pepese/golang-CleanArchitecture/app/domain/model"

/*
UserLogicer Interface
*/
type UserLogicer interface {
	FullName() string
}

type userLogic struct {
	model *model.User
}

func (d *userLogic) FullName() string {
	return d.model.FirstName + " " + d.model.LastName
}
