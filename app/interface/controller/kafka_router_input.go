package controller

import "github.com/pepese/golang-CleanArchitecture/app/domain/model"

type UsersTopic struct {
	Method  string     `json:"method"`
	Message model.User `json:"message"`
}
