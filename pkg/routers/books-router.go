package routers

import (
	"github.com/TonyPham-Dev/Golang-Crud/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRouterBooks(router *mux.Router) (routers *mux.Router) {
	// router.HandleFunc("/books", controllers.getBooks).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.getBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT")
	// router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE")
	return router
}
