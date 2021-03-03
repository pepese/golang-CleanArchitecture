package controller

import (
	"fmt"
	"net/http"

	"github.com/pepese/golang-CleanArchitecture/interface/presenter"
	"github.com/pepese/golang-CleanArchitecture/usecase"
)

/*
NewHttpRouter returns *http.ServeMux as a http handler.
*/
func NewHttpRouter() *http.ServeMux {
	mux := http.NewServeMux()

	//////////////////
	// Hello UseCase
	//////////////////
	// input
	hin := usecase.NewHelloInput("Go")
	// usecase
	huc := usecase.NewHelloUsecase(hin)
	// router
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// exec usecase
		huc.Say()
		// exec presenter
		rep := presenter.HelloRender(huc.Out)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, rep)
	})

	return mux
}
