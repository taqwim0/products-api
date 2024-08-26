package view

import (
	"toggle-features-api/controller"
	"toggle-features-api/utils"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public
	router.HandleFunc("/api/login", utils.Login).Methods("POST")

	// Private, protected by JWT middleware
	api := router.PathPrefix("/api").Subrouter()
	api.Use(utils.Auth)

	api.HandleFunc("/feature-toggles", controller.GetFeatureToggles).Methods("GET")
	api.HandleFunc("/feature-toggles/{id}", controller.GetFeatureToggleByID).Methods("GET")
	api.HandleFunc("/feature-toggles/add", controller.CreateFeatureToggle).Methods("POST")
	api.HandleFunc("/feature-toggles/{id}", controller.UpdateFeatureToggle).Methods("PUT")
	api.HandleFunc("/feature-toggles/{id}", controller.DeleteFeatureToggle).Methods("DELETE")

	return router
}
