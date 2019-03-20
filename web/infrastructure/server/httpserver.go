package server

import (
	"github.com/pepese/golang-CleanArchitecture/web/interface/controller"
	"github.com/spf13/cobra"
)

type HttpServer interface {
	Run(cmd *cobra.Command, args []string)
}

type httpServer struct{}

func (hs *httpServer) Run(cmd *cobra.Command, args []string) {
	router := controller.NewHttpRouter()
	RunWithGracefulStop(router)
}

func NewHttpServer() *httpServer {
	return &httpServer{}
}
