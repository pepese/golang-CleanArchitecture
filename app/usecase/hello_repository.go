package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
)

type HelloRepository interface {
	Create(model *model.Hello) *model.Hello
}
