package server

import "github.com/pepese/golang-CleanArchitecture/app/interface/controller"

type httpServer struct{}

func (hs *httpServer) Run() {
	router := controller.NewHttpRouter()
	RunWithGracefulStop(router)
}

func NewHttpServer() *httpServer {
	return &httpServer{}
}
