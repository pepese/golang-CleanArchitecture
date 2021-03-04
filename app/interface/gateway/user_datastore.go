package gateway

import (
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	"github.com/pepese/golang-CleanArchitecture/app/infrastructure/datastore"
)

type userRepository struct{}

func (repo *userRepository) List(m *model.User) (model.Users, error) {
	rdb := datastore.Gorm()
	searched := []model.User{}
	if result := rdb.Find(&searched, m); result.Error != nil {
		return nil, result.Error
	}
	return searched, nil
}

/*
検索でヒットしない場合は nil が返る
*/
func (repo *userRepository) Get(m *model.User) (*model.User, error) {
	rdb := datastore.Gorm()
	searched := &model.User{}
	if result := rdb.First(searched, m); result.Error != nil {
		return nil, result.Error
	}
	return searched, nil
}

/*
primary key が重複した場合は nil, err が返却される（Error 1062: Duplicate entry '3' for key 'PRIMARY'）
*/
func (repo *userRepository) Create(m *model.User) (*model.User, error) {
	rdb := datastore.Gorm()
	if result := rdb.Create(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func (repo *userRepository) Update(m *model.User) (*model.User, error) {
	rdb := datastore.Gorm()
	if result := rdb.Omit("created_at").Save(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func (repo *userRepository) Delete(m *model.User) (*model.User, error) {
	rdb := datastore.Gorm()
	if result := rdb.Delete(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}
