package router

import (
	controller "olowoleru/pams/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/create/catalogue", controller.CreateCatalogue).Methods("POST")
	// router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	// router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("POST")
	// router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("POST")
	// router.HandleFunc("api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router

}
