package router

import (
	"github.com/Yashsharma1911/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// home serve
	router.HandleFunc("/", controller.HomeServe).Methods("GET")

	// get all movies
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")

	// create a movie
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")

	// mark movie as watched
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	// delete one movie
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")

	// delete all movies
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
