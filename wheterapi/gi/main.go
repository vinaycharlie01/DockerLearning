package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func getWeatherResponse(c *gin.Context) {
	var favorites []WeatherResponse
	// Replace this with code to get favorites from MongoDB
	c.JSON(http.StatusOK, favorites)
}

func addWeatherResponse(c *gin.Context) {
	var favorite WeatherResponse
	if err := c.BindJSON(&favorite); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Replace this with code to save the favorite to MongoDB
	c.JSON(http.StatusCreated, favorite)
}

func getWeather(c *gin.Context) {

	apiKey := "2f930de950a60bc0999bc26c836eb3cd"
	city := "Kolar"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	response, err := http.Get(url)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer response.Body.Close()

	var d WeatherResponse
	// if err := c.BindJSON(&d); err != nil {
	//
	// }
	if err := json.NewDecoder(response.Body).Decode(&d); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var detalis []WeatherResponse
	detalis = append(detalis, d)
	c.JSON(http.StatusOK, detalis)
}

func main() {
	fmt.Println("Starting the application...")
	// Replace this with code to connect to MongoDB
	// client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))

	router := gin.Default()
	router.GET("/favorites", getWeatherResponse)
	router.POST("/favorites", addWeatherResponse)
	router.GET("/weather", getWeather)

	if err := router.Run(":3000"); err != nil {
		fmt.Println(err)
	}
}
