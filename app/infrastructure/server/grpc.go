package server

import (
	"log"
	"net"
	"sync"

	"github.com/pepese/golang-CleanArchitecture/app/interface/controller"
	upb "github.com/pepese/golang-CleanArchitecture/app/interface/controller/grpc_gen/user/v1"
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
		gSrv = grpc.NewServer()
		c := controller.NewGrpcRouter(nil) // ここで UC を設定する
		upb.RegisterUserServiceServer(gSrv, &c.UsrServ)
	})
	return &grpcServer{}
}
