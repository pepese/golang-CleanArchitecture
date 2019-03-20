package server

import (
	"github.com/pepese/golang-CleanArchitecture/web/interface/controller"
	"github.com/spf13/cobra"
)

type GinServer interface {
	Run(cmd *cobra.Command, args []string)
}

type ginServer struct{}

func (hs *ginServer) Run(cmd *cobra.Command, args []string) {
	router := controller.NewGinRouter()
	RunWithGracefulStop(router)
}

func NewGinServer() *ginServer {
	return &ginServer{}
}
