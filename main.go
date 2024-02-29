package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"ISBN"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response will be in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Extract the route parameters from the request using Gorilla Mux
	params := mux.Vars(r)

	// Iterate over the movies slice to find the movie with the specified ID
	for index, movie := range movies {
		if movie.ID == params["id"] {
			// If the movie with the specified ID is found, remove it from the movies slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			// Print the ID before encoding the movie information
			fmt.Println("Received request for movie ID:", params["id"])

			// Encode and send the movie information in the response
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func creteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		ISBN:     "448743",
		Title:    "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		ISBN:     "783405",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Charles", Lastname: "Chaplin"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", creteMovie).Methods("POST")
	// r.HandleFunc("/movies/[id]", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", r))
}
