package server

import (
	"net/http"

	"github.com/pepese/golang-CleanArchitecture/web/interface/controller"
	"github.com/spf13/cobra"
)

type HttpServer interface {
	Run(cmd *cobra.Command, args []string)
}

type httpServer struct{}

func (hs *httpServer) Run(cmd *cobra.Command, args []string) {
	mux := controller.NewMux()
	srv := &http.Server{Addr: ":8080", Handler: mux}
	srv.ListenAndServe()
}

func NewHttpServer() *httpServer {
	return &httpServer{}
}
