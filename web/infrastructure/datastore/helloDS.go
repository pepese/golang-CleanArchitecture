package datastore

import (
	"github.com/pepese/golang-CleanArchitecture/web/domain"
	"github.com/pepese/golang-CleanArchitecture/web/domain/model"
)

type helloRepository struct{}

func (r *helloRepository) Create(model *model.Hello) *domain.Hello {
	model.Say = "Hello " + model.Name + "!"
	return model
}

func NewHelloRepository() *helloRepository {
	return &helloRepository{}
}
