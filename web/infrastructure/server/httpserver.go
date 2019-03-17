package httpserver

import (
	"net/http"

	httphandler "github.com/pepese/golang-CleanArchitecture/web/interface/controller"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	mux := httphandler.NewMux()
	srv := &http.Server{Addr: ":8080", Handler: mux}
	srv.ListenAndServe()
}
