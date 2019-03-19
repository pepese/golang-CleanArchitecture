package datastore

import "github.com/pepese/golang-CleanArchitecture/web/domain/model"

type helloRepository struct{}

func (r *helloRepository) Create(model *model.Hello) {
	model.Say = "Hello " + model.Name + "!"
}

func NewHelloRepository() *helloRepository {
	return &helloRepository{}
}
