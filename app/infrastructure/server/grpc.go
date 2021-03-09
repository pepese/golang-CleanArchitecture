package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/pepese/golang-CleanArchitecture/app/interface/controller"
	upb "github.com/pepese/golang-CleanArchitecture/app/interface/controller/grpc_gen/user/v1"
	"github.com/pepese/golang-gin-sample/app"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	gSrv     *grpc.Server
	gSrvOnce sync.Once
)

type grpcServer struct{}

func (gs *grpcServer) Run() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := gSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewGrpcServer() *grpcServer {
	gSrvOnce.Do(func() {
		gSrv = grpc.NewServer(
			grpc.UnaryInterceptor(NewUnaryServierInterceptor()),
		)
		c := controller.NewGrpcRouter(nil) // ここで UC を設定する
		upb.RegisterUserServiceServer(gSrv, &c.UsrServ)
	})
	return &grpcServer{}
}

func NewUnaryServierInterceptor() grpc.UnaryServerInterceptor {
	return func(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		// context.Context
		ctx := context.Background()

		// Request ID
		reqId := uuid.NewV4().String()

		// Logger
		logger := app.GetLoggerWithKeyValue("reqId", reqId)
		ctx = app.SetLoggerToContext(ctx, logger)
		c = context.WithValue(c, "ctx", ctx)

		// Access Log
		start := time.Now()
		// アクセスログ（リクエスト）出力
		logger.Infow("===>request", zap.Reflect("request", req), "reqId", reqId, "method", info.FullMethod)

		reply, err = handler(c, req)

		// アクセスログ（レスポンス）出力
		responseTime := time.Now().Sub(start)
		if err != nil {
			logger.Infow("<==response", zap.Reflect("response", reply), "reqId", reqId, "method", info.FullMethod, "responseTime", responseTime, "error", err)
		} else {
			logger.Infow("<==response", "reqId", reqId, "method", info.FullMethod, "responseTime", responseTime)
		}

		return
	}
}
