package domain

import (
	"github.com/pepese/golang-CleanArchitecture/web/domain/model"
	"github.com/pepese/golang-CleanArchitecture/web/infrastructure/datastore"
)

/*
Hello ドメインのインターフェース
*/
type HelloLogic interface {
	Create() string
}

type helloLogic struct {
	model *model.Hello
}

/*
ドメインロジック
*/
func (d *helloLogic) Create() string {
	rep := datastore.NewHelloRepository(d.model)
	rep.Create(d.model)
	return d.model.Say
}

/*
Hello ドメインの作成
*/
func NewHelloLogic(model *model.Hello) *helloLogic {
	return &helloLogic{model: model}
}
