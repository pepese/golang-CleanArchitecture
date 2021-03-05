package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	"github.com/pepese/golang-CleanArchitecture/app/usecase/mock_usecase"
	"gopkg.in/go-playground/assert.v1"
)

func TestUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("List", func(t *testing.T) {
		m := mock_usecase.NewMockUserRepository(ctrl)
		in := &model.User{}
		m.EXPECT().List(in).Return(model.Users{model.User{FirstName: "first", LastName: "last"}}, nil)
		userUc := &UserUsecase{UserRepo: m}
		result, _ := userUc.List(nil, in)
		assert.Equal(t, result, model.Users{model.User{FirstName: "first", LastName: "last"}})
	})
}
