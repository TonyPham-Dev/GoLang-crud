package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movies

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var movies Movies
	json.Unmarshal(body, &movies)
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func router() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/update/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("server is running port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	movies = append(movies, Movies{ID: "1", Title: "title one", Isbn: "123123", Director: &Director{FirstName: "Tony", LastName: "Pham"}})
	movies = append(movies, Movies{ID: "2", Title: "title two", Isbn: "123112323", Director: &Director{FirstName: "Pham", LastName: "Cong"}})
	router()
}
