package controller

import (
	"fmt"
	"net/http"

	"github.com/pepese/golang-CleanArchitecture/web/usecase"
)

/*
NewMux returns *http.ServeMux as a http handler.
*/
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewHelloUsecase("Go")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, uc.Say())
	})
	return mux
}
