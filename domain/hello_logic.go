package domain

import (
	"github.com/pepese/golang-CleanArchitecture/domain/model"
)

/*
Hello ドメインのインターフェース
*/
type HelloLogic interface {
	Say() string
}

type helloLogic struct {
	model *model.Hello
}

/*
ドメインロジック
*/
func (d *helloLogic) Say() string {
	return "Hello " + d.model.Name + "!"
}

/*
Hello ドメインの作成
*/
func NewHelloLogic(model *model.Hello) *helloLogic {
	return &helloLogic{model: model}
}
