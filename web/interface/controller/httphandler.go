package httphandler

import (
	"fmt"
	"net/http"

	hello "github.com/pepese/golang-CleanArchitecture/web/usecase"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, hello.Exec())
	})
	return mux
}
