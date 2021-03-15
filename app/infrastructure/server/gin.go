package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/interface/controller"
	uuid "github.com/satori/go.uuid"
)

var (
	gEn     *gin.Engine
	gEnOnce sync.Once
)

type ginServer struct{}

func (hs *ginServer) Run() {
	NewGinServer()
	RunWithGracefulStop(gEn)
}

func NewGinServer() *ginServer {
	gEnOnce.Do(func() {
		e := gin.New()
		e.Use(accessLogging())
		e.Use(gin.Recovery())
		r := controller.NewGinRouter(e, nil) // ここで UC を設定する
		gEn = r.GinRouter()
	})
	return &ginServer{}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func accessLogging() gin.HandlerFunc {
	// Logging Skip Logic
	notlogged := []string{"/healthcheck"}
	var skip map[string]struct{}
	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		// context.Context
		ctx := context.Background()

		// Request ID
		reqId := c.Request.Header.Get("X-Request-Id")
		if reqId == "" {
			reqId = c.Request.Header.Get("x-amzn-trace-id")
		}
		if reqId == "" {
			reqId = uuid.NewV4().String()
		}
		c.Set("reqId", reqId)
		c.Writer.Header().Set("X-Request-Id", reqId)

		// Logger
		logger := app.LoggerWithKeyValue("reqId", reqId)
		ctx = app.SetLoggerToContext(ctx, logger)
		c.Set("ctx", ctx)

		// Access Log
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		if query != "" {
			path = path + "?" + query
		}
		header := c.Request.Header
		// リクエスト取得
		req := c.Request
		requestBody := make([]byte, 0)
		if req.Body != nil {
			body, err := ioutil.ReadAll(req.Body)
			req.Body.Close()
			if err != nil {
				logger.Errorw("read request body is fail:", err)
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(body)) // streamを読み尽くすので入れ直す
			requestBody = body
		}
		// レスポンス取得
		rbw := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rbw
		// アクセスログ（リクエスト）出力
		logger.Infow("===>request", "reqId", reqId, "method", method, "path", path, "header", fmt.Sprintf("%#v", header), "request", string(requestBody))

		c.Next()

		if _, ok := skip[path]; !ok {
			responseTime := time.Now().Sub(start)
			statusCode := c.Writer.Status()
			// アクセスログ（レスポンス）出力
			logger.Infow("<==response", "reqId", reqId, "statusCode", statusCode, "responseTime", responseTime, "response", rbw.body.String())
		}
	}
}
