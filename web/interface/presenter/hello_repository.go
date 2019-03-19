package presenter

import (
	"github.com/pepese/golang-CleanArchitecture/web/domain/model"
)

type HelloRepository interface {
	Create(model *model.Hello) *model.Hello
}
