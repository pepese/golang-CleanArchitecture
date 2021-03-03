package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/domain/model"
)

type HelloRepository interface {
	Create(model *model.Hello) *model.Hello
}
