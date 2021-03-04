package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
)

type UserRepository interface {
	List(m *model.User) (model.Users, error)
	Get(m *model.User) (*model.User, error)
	Create(m *model.User) (*model.User, error)
	Update(m *model.User) (*model.User, error)
	Delete(m *model.User) (*model.User, error)
}
