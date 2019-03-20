package controller

import (
	"fmt"
	"net/http"

	"github.com/pepese/golang-CleanArchitecture/web/usecase"
)

/*
NewHttpRouter returns *http.ServeMux as a http handler.
*/
func NewHttpRouter() *http.ServeMux {
	mux := http.NewServeMux()
	uc := usecase.NewHelloUsecase("Go")
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, uc.Say())
	})
	return mux
}
