package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
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

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		return
	}
	number, _ := strconv.Atoi(movies[len(movies)-1].ID)
	movie.ID = strconv.Itoa(number + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			var newMovie Movie
			json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = movie.ID
			movies[index] = newMovie
			json.NewEncoder(w).Encode(newMovie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	port := "8080"
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "15012011",
		Title: "Li Quang Tien",
		Director: &Director{
			Firstname: "Tien",
			Lastname:  "Li",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "16081980",
		Title: "Nguyen Thi Du",
		Director: &Director{
			Firstname: "Du",
			Lastname:  "Nguyen",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
