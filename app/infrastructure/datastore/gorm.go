package datastore

import (
	"fmt"
	"sync"

	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	"gorm.io/driver/mysql"
	gm "gorm.io/gorm"
)

var (
	gorm  *gm.DB
	gOnce sync.Once
)

func Gorm() *gm.DB {
	initGorm()
	return gorm
}

func initGorm() {
	gOnce.Do(func() {
		conf := app.Config()
		connectTemplate := "%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		con := fmt.Sprintf(connectTemplate, conf.RdbUser, conf.RdbPassword, conf.RdbProtocol, conf.RdbHost, conf.RdbName)
		db, err := gm.Open(mysql.Open(con), &gm.Config{})
		if err != nil {
			panic(err.Error())
		}
		db.AutoMigrate(&model.User{})
		gorm = db
	})
}
