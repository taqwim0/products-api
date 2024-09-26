package view

import (
	"toggle-features-api/controller"

	"github.com/gorilla/mux"
)

// notes

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	router.HandleFunc("/product/detail/{id}", controller.GetProductByID).Methods("GET")
	router.HandleFunc("/product/add", controller.InsertProduct).Methods("POST")
	router.HandleFunc("/product/update/{id}", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/delete/{id}", controller.DeleteProduct).Methods("DELETE")

	return router
}
