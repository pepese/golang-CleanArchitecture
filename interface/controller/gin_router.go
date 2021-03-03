package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-CleanArchitecture/interface/presenter"
	"github.com/pepese/golang-CleanArchitecture/usecase"
)

func NewGinRouter() *gin.Engine {
	e := gin.Default()

	//////////////////
	// Hello UseCase
	//////////////////
	// input
	hin := usecase.NewHelloInput("Go")
	// usecase
	huc := usecase.NewHelloUsecase(hin)
	// router
	e.GET("/hello", func(c *gin.Context) {
		// exec usecase
		huc.Say()
		// exec presenter
		r := presenter.HelloRender(huc.Out)

		c.String(http.StatusOK, r)
	})

	return e
}
