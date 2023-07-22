package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type WeatherResponse struct {
	Base   string `json:"base"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Cod   int `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int `json:"dt"`
	ID   int `json:"id"`
	Main struct {
		FeelsLike float64 `json:"feels_like"`
		GrndLevel int     `json:"grnd_level"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
		SeaLevel  int     `json:"sea_level"`
		Temp      float64 `json:"temp"`
		TempMax   float64 `json:"temp_max"`
		TempMin   float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		ID      int    `json:"id"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
		Type    int    `json:"type"`
	} `json:"sys"`
	Timezone   int `json:"timezone"`
	Visibility int `json:"visibility"`
	Weather    []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int    `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func getWeatherResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favorites []WeatherResponse
	// Replace this with code to get favorites from MongoDB
	json.NewEncoder(w).Encode(favorites)
}

func addWeatherResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var favorite WeatherResponse
	_ = json.NewDecoder(r.Body).Decode(&favorite)
	// Replace this with code to save the favorite to MongoDB
	json.NewEncoder(w).Encode(favorite)
}

func getWether(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apiKey := "2f930de950a60bc0999bc26c836eb3cd"
	city := "Kolar"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// var data map[string]interface{}
	var d WeatherResponse
	_ = json.NewDecoder(response.Body).Decode(&d)
	json.NewEncoder(w).Encode(d)
	var a []WeatherResponse
	a = append(a, d)
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/favorites", getWeatherResponse).Methods("GET")
	router.HandleFunc("/favorites", addWeatherResponse).Methods("POST")
	router.HandleFunc("/getwether", getWether).Methods("GET")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
