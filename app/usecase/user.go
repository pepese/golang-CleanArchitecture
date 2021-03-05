package usecase

import (
	"context"

	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
)

/*
UserUsecaser interface.
*/
type UserUsecaser interface {
	List(c context.Context, m *model.User) (model.Users, error)
	Get(c context.Context, m *model.User) (*model.User, error)
	Create(c context.Context, m *model.User) (*model.User, error)
	Update(c context.Context, m *model.User) (*model.User, error)
	Delete(c context.Context, m *model.User) (*model.User, error)
}

/*
UserUsecase mode.
*/
type UserUsecase struct {
	UserRepo UserRepository
}

/*
List func is UserUsecase implement.
*/
func (u UserUsecase) List(c context.Context, m *model.User) (model.Users, error) {
	logger := app.LoggerFromContext(c)
	result, err := u.UserRepo.List(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, nil
}

func (u UserUsecase) Get(c context.Context, m *model.User) (*model.User, error) {
	logger := app.LoggerFromContext(c)
	result, err := u.UserRepo.Get(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, nil
}

func (u UserUsecase) Create(c context.Context, m *model.User) (*model.User, error) {
	logger := app.LoggerFromContext(c)
	result, err := u.UserRepo.Create(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, nil
}

func (u UserUsecase) Update(c context.Context, m *model.User) (*model.User, error) {
	logger := app.LoggerFromContext(c)
	result, err := u.UserRepo.Update(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, nil
}

func (u UserUsecase) Delete(c context.Context, m *model.User) (*model.User, error) {
	logger := app.LoggerFromContext(c)
	result, err := u.UserRepo.Delete(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, nil
}
