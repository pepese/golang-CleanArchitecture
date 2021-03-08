package controller

import (
	"context"
	"reflect"

	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	upb "github.com/pepese/golang-CleanArchitecture/app/interface/controller/grpc_gen/user/v1"
	db "github.com/pepese/golang-CleanArchitecture/app/interface/gateway"
	"github.com/pepese/golang-CleanArchitecture/app/usecase"
)

type grpcRouter struct {
	UsrServ userService
}

//////////////////
// User UseCase
//////////////////

type userService struct {
	userUc usecase.UserUsecaser
}

func (us *userService) ListUsers(c context.Context, req *upb.UserRequest) (*upb.UsersResponse, error) {
	m := model.User{
		ID:        int(req.ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	result, err := us.userUc.List(c.Value("ctx").(context.Context), &m)
	if err != nil {
		return mapUsersResponse(result), nil
	}
	return nil, err
}

func (us *userService) GetUser(c context.Context, req *upb.UserRequest) (*upb.UserResponse, error) {
	m := model.User{
		ID:        int(req.ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	result, err := us.userUc.Get(c.Value("ctx").(context.Context), &m)
	if err != nil {
		return mapUserResponse(*result), nil
	}
	return nil, err
}

func (us *userService) CreateUser(c context.Context, req *upb.UserRequest) (*upb.UserResponse, error) {
	m := model.User{
		ID:        int(req.ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	result, err := us.userUc.Create(c.Value("ctx").(context.Context), &m)
	if err != nil {
		return mapUserResponse(*result), nil
	}
	return nil, err
}

func (us *userService) UpdateUsers(c context.Context, req *upb.UserRequest) (*upb.UserResponse, error) {
	m := model.User{
		ID:        int(req.ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	result, err := us.userUc.Update(c.Value("ctx").(context.Context), &m)
	if err != nil {
		return mapUserResponse(*result), nil
	}
	return nil, err
}

func (us *userService) DeleteUsers(c context.Context, req *upb.UserRequest) (*upb.UserResponse, error) {
	m := model.User{
		ID:        int(req.ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	result, err := us.userUc.Delete(c.Value("ctx").(context.Context), &m)
	if err != nil {
		return mapUserResponse(*result), nil
	}
	return nil, err
}

func mapUsersResponse(m model.Users) *upb.UsersResponse {
	us := []*upb.UserResponse{}
	for _, u := range ([]model.User)(m) {
		us = append(us, mapUserResponse(u))
	}
	return &upb.UsersResponse{
		Users: us,
	}
}

func mapUserResponse(m model.User) *upb.UserResponse {
	return &upb.UserResponse{
		ID:        int64(m.ID),
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}

func NewGrpcRouter(userUc usecase.UserUsecaser) *grpcRouter {
	if (userUc == nil) || reflect.ValueOf(userUc).IsNil() {
		userUc = &usecase.UserUsecase{
			UserRepo: db.NewUserRepository(),
		}
	}
	return &grpcRouter{
		UsrServ: userService{
			userUc: userUc,
		},
	}
}
