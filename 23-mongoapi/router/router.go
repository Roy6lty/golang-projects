package router

import (
	"github.com/gorilla/mux"
	controller "github.com/olowoleru/contoller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies ).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("POST")
	router.HandleFunc("api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router



}