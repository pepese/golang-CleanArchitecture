package server

import (
	"github.com/pepese/golang-CleanArchitecture/interface/controller"
	"github.com/spf13/cobra"
)

type GinServer interface {
	Run(cmd *cobra.Command, args []string)
}

type ginServer struct{}

func (hs *ginServer) Run() {
	router := controller.NewGinRouter()
	RunWithGracefulStop(router)
}

func NewGinServer() *ginServer {
	return &ginServer{}
}
