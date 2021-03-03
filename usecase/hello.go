package usecase

import (
	"github.com/pepese/golang-CleanArchitecture/domain"
	"github.com/pepese/golang-CleanArchitecture/domain/model"
	"github.com/pepese/golang-CleanArchitecture/interface/gateway"
)

/*
HelloUC interface contains Usecase methods.
*/
type HelloUsecaser interface {
	Say()
}

type helloUsecase struct {
	in  *helloInput
	Out *HelloOutput
}

/*
ユースケースは1つ以上のドメインモデル・ロジックを用いて実装される。
*/
func (uc *helloUsecase) Say() {
	// ドメインモデルの作成
	m := &model.Hello{Name: uc.in.name}
	// Repository
	var rep HelloRepository = gateway.NewHelloRepository()
	rep.Create(m)
	// ユースケースで実行するドメインロジックの作成
	d := domain.NewHelloLogic(m)
	// ドメインロジックの実行
	r := d.Say()
	// output
	o := NewHelloOutput(r)
	uc.Out = o

	return
}

/*
NewHelloUC returns *helloUC.
*/
func NewHelloUsecase(in *helloInput) *helloUsecase {
	return &helloUsecase{in: in}
}
