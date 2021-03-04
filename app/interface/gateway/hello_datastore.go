package gateway

import (
	"fmt"

	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
)

type helloRepository struct{}

func (r *helloRepository) Create(m *model.Hello) *model.Hello {
	fmt.Print(m)
	fmt.Println(" saved !")
	return m
}

func NewHelloRepository() *helloRepository {
	return &helloRepository{}
}
