package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/web/domain"
	"github.com/pepese/golang-CleanArchitecture/web/domain/model"
)

type HelloUC interface {
	Say() string
}

type helloUC struct {
	name string
}

func (uc *helloUC) Say() string {
	m := model.Hello{name: uc.name}
	d := domain.NewHello(m)
	return d.Create()
}

func NewHelloUC(name string) *helloUC {
	return &helloUC{name: name}
}
