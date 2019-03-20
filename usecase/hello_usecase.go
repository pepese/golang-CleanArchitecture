package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/domain"
	"github.com/pepese/golang-CleanArchitecture/domain/model"
)

/*
HelloUC interface contains Usecase methods.
*/
type HelloUsecase interface {
	Say() string
}

type helloUsecase struct {
	name string
}

/*
ユースケースは1つ以上のドメインモデル・ロジックを用いて実装される。
*/
func (uc *helloUsecase) Say() string {
	// ドメインモデルの作成
	m := &model.Hello{Name: uc.name}
	// ユースケースで実行するドメインロジックの作成
	d := domain.NewHelloLogic(m)
	return d.Create()
}

/*
NewHelloUC returns *helloUC.
*/
func NewHelloUsecase(name string) *helloUsecase {
	return &helloUsecase{name: name}
}
