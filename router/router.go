package router

import (
	"github.com/Vzaldat/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.Deleteam).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.Deleteallms).Methods("DELETE")

	return router
}
