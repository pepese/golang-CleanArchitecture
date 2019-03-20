package datastore

import "github.com/pepese/golang-CleanArchitecture/domain/model"

type helloRepository struct{}

func (r *helloRepository) Create(model *model.Hello) {
	model.Say = "Hello " + model.Name + "!"
}

func NewHelloRepository() *helloRepository {
	return &helloRepository{}
}
