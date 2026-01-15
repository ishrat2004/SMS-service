package router

import (
	"fmt"
	"goserver/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("router being started")
	router := mux.NewRouter()

	router.HandleFunc("/v1/{phoneNumber}", controller.GetUser).Methods("GET")
	router.HandleFunc("/v1/get/all", controller.GetAllMessages).Methods("GET")
	return router
}
