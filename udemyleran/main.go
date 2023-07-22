package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type favorite struct {
	ID   string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

func getFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favorites []favorite
	// Replace this with code to get favorites from MongoDB
	json.NewEncoder(w).Encode(favorites)
}

func addFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favorite favorite
	_ = json.NewDecoder(r.Body).Decode(&favorite)
	// Replace this with code to save the favorite to MongoDB
	json.NewEncoder(w).Encode(favorite)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get("https://swapi.dev/api/films")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&data)
	json.NewEncoder(w).Encode(data)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&data)
	json.NewEncoder(w).Encode(data)
}

func main() {
	fmt.Println("Starting the application...")
	// Replace this with code to connect to MongoDB
	client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	router := mux.NewRouter()
	router.HandleFunc("/favorites", getFavorites).Methods("GET")
	router.HandleFunc("/favorites", addFavorite).Methods("POST")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/people", getPeople).Methods("GET")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
