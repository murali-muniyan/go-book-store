package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murali-muniyan/go-book-store/pkg/controllers"
)

func RegisterRouter(path string, rtr *mux.Router) {
	rtr.Path(path).HandlerFunc(controllers.CreateBook).Methods(http.MethodPost)
	rtr.Path(path).HandlerFunc(controllers.GetBooks).Methods(http.MethodGet)
	rtr.Path(path + "/{id}").HandlerFunc(controllers.GetBook).Methods(http.MethodGet)
	rtr.Path(path + "/{id}").HandlerFunc(controllers.UpdateBook).Methods(http.MethodPut)
	rtr.Path(path + "/{id}").HandlerFunc(controllers.DeleteBook).Methods(http.MethodDelete)
}
