package routers

import (
	"github.com/DNahar74/golang-playlist/23.mongoDB_API/controllers"
	"github.com/gorilla/mux"
)

// Router is used for routing
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/add", controllers.CreateMovie).Methods("PATCH")
	router.HandleFunc("/api/watched/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/delete", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}