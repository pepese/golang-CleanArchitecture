package controller

import (
	"fmt"
	"net/http"

	"github.com/pepese/golang-CleanArchitecture/web/usecase"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewHelloUC("Go")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, uc.Say())
	})
	/*
		mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
			h := usecase.NewUser("Go")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, h.Get())
		})
	*/
	return mux
}
