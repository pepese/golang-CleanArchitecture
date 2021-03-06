package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	"github.com/pepese/golang-CleanArchitecture/app/usecase/mock_usecase"
	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////
// TEST /api/v1/users
////////////////////////////////////////////////////////////////////////

func TestUsersAPI(t *testing.T) {
	// mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app.Config()

	t.Run("Get empty users", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().List(gomock.Any(), gomock.Any()).Return(model.Users{}, nil)

		gRouter := NewGinRouter(getGinEngine(), m)
		testRouter := gRouter.GinRouter()
		req, _ := http.NewRequest("GET", "/api/v1/users", nil)

		res := httptest.NewRecorder()
		testRouter.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "[]\n", getBodyString(res))
	})

	t.Run("Post user", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		gRouter := NewGinRouter(getGinEngine(), m)
		testRouter := gRouter.GinRouter()
		req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(`
		{
			"first_name": "First",
			"last_name": "Last"
		}`))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		testRouter.ServeHTTP(res, req)
		assert.Equal(t, res.Code, 200)
	})

	t.Run("Get user", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		gRouter := NewGinRouter(getGinEngine(), m)
		testRouter := gRouter.GinRouter()
		req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)

		res := httptest.NewRecorder()
		testRouter.ServeHTTP(res, req)
		assert.Equal(t, res.Code, 200)
	})

	t.Run("Put user", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		gRouter := NewGinRouter(getGinEngine(), m)
		testRouter := gRouter.GinRouter()
		req, _ := http.NewRequest("PUT", "/api/v1/users/1", strings.NewReader(`
		{
			"first_name": "FirstX",
			"last_name": "LastX"
		}`))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		testRouter.ServeHTTP(res, req)
		assert.Equal(t, res.Code, 200)
	})

	t.Run("Create user", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		gRouter := NewGinRouter(getGinEngine(), m)
		testRouter := gRouter.GinRouter()
		req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)

		res := httptest.NewRecorder()
		testRouter.ServeHTTP(res, req)
		assert.Equal(t, res.Code, 200)
	})
}

////////////////////////////////////////////////////////////////////////
// TEST Utilities
////////////////////////////////////////////////////////////////////////

func getGinEngine() *gin.Engine {
	// setting Default gin.Engine
	var ctxF func() gin.HandlerFunc = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Set("ctx", context.Background())
		}
	}
	e := gin.Default()
	e.Use(ctxF())

	return e
}

func getBodyString(res *httptest.ResponseRecorder) string {
	defer res.Result().Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Result().Body)
	return buf.String()
}

func getBodyBytes(res *httptest.ResponseRecorder) []byte {
	defer res.Result().Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Result().Body)
	return buf.Bytes()
}

func getStruct(entity interface{}, res *httptest.ResponseRecorder) interface{} {
	jsonBytes := getBodyBytes(res)
	if err := json.Unmarshal(jsonBytes, entity); err != nil {
		return nil
	}
	return entity
}
