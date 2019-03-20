package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-CleanArchitecture/web/usecase"
)

func NewGinRouter() *gin.Engine {
	e := gin.Default()
	uc := usecase.NewHelloUsecase("Go")
	e.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, uc.Say())
	})
	return e
}
