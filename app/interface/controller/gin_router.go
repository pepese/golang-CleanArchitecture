package controller

import (
	"context"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	db "github.com/pepese/golang-CleanArchitecture/app/interface/gateway"
	"github.com/pepese/golang-CleanArchitecture/app/interface/presenter"
	"github.com/pepese/golang-CleanArchitecture/app/usecase"
)

func NewGinRouter(e *gin.Engine, userUc usecase.UserUsecaser) *gin.Engine {
	router := e.Group("/api/v1")

	//////////////////
	// Hello UseCase
	//////////////////
	hello := router.Group("/hello")
	// input
	hin := usecase.NewHelloInput("Go")
	// usecase
	huc := usecase.NewHelloUsecase(hin)
	// router
	hello.GET("", func(c *gin.Context) {
		// exec usecase
		huc.Say()
		// exec presenter
		r := presenter.HelloRender(huc.Out)

		c.String(http.StatusOK, r)
	})

	//////////////////
	// User UseCase
	//////////////////
	user := router.Group("/users")
	if (userUc == nil) || reflect.ValueOf(userUc).IsNil() {
		userUc = &usecase.UserUsecase{
			UserRepo: db.NewUserRepository(),
		}
	}
	// GET /api/v1/users
	user.GET("", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		result, err := userUc.List(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// GET /api/v1/users/:user_id
	user.GET("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := userUc.Get(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// POST /api/v1/users
	user.POST("", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		result, err := userUc.Create(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// PUT /api/v1/users/:user_id
	user.PUT("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := userUc.Update(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// DELETE /api/v1/users/:user_id
	user.DELETE("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := userUc.Delete(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})

	return e
}
