package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-CleanArchitecture/web/usecase"
)

func NewEngine() *gin.Engine {
	e := gin.Default()
	uc := usecase.NewHelloUsecase("Go")
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(200, uc.Say())
	})
	return e
}
