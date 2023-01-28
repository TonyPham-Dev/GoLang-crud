package routers

import (
	"github.com/gorilla/mux"
	"github.com/TonyPham-Dev/Golang-Crud/pkg/controllers"
	
)

var RegisterRouterBooks = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.getBook).Methods("GET")
	router.HandleFunc("/books", controllers.createBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE")
}