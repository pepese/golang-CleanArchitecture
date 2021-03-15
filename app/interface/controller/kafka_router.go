package controller

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	db "github.com/pepese/golang-CleanArchitecture/app/interface/gateway"
	"github.com/pepese/golang-CleanArchitecture/app/usecase"
)

type kafkaRouter struct {
	userUc usecase.UserUsecaser
}

func (k *kafkaRouter) KafkaRouter(c context.Context, in UsersTopic) (*model.User, error) {
	log.Println("kafkaRouter exec.")
	log.Printf("input: %v\n", in)

	//////////////////
	// User UseCase
	//////////////////
	if (k.userUc == nil) || reflect.ValueOf(k.userUc).IsNil() {
		k.userUc = &usecase.UserUsecase{
			UserRepo: db.NewUserRepository(),
		}
	}
	switch in.Method {
	case "create":
		result, err := k.userUc.Create(c.Value("ctx").(context.Context), &in.Message)
		if err != nil {
			return nil, err
		}
		return result, nil
	case "update":
		result, err := k.userUc.Update(c.Value("ctx").(context.Context), &in.Message)
		if err != nil {
			return nil, err
		}
		return result, nil
	case "delete":
		result, err := k.userUc.Delete(c.Value("ctx").(context.Context), &in.Message)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, errors.New("no method to match: " + in.Method)
	}
}

func NewKafkaRouter() *kafkaRouter {
	return &kafkaRouter{}
}
