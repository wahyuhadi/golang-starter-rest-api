/**
	Rahmat wahyu hadi
**/

package application

import (
	"./controller"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true

	api := mainRouter.PathPrefix("/v1").Subrouter()
	api.Path("/info").HandlerFunc(controller.GetAPIInfo)

	// user
	user := mainRouter.PathPrefix("/v1/user").Subrouter()
	user.Path("/").HandlerFunc(controller.CreateUser).Methods("POST")
	user.Path("/").HandlerFunc(controller.GetAllUsers).Methods("GET")
	user.Path("/{id}").HandlerFunc(controller.GetUserById).Methods("GET")
	user.Path("/{id}").HandlerFunc(controller.UpdateUser).Methods("PUT")
	user.Path("/{id}").HandlerFunc(controller.DeleteUser).Methods("DELETE")

	// tambahkan disini sesuai keperluan seperti contoh di atas

	return mainRouter

}
