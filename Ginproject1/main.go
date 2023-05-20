package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Weather struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AirTemparature string `json:"airtempareture"`
	Humidity       string `json:"humidity"`
	windspeed      string `json:"windspeed"`
}

var weatherdevice = []Weather{
	{ID: "1", Name: "vijayawada", AirTemparature: "40", Humidity: "25%", windspeed: "10 km/h"},
	{ID: "2", Name: "kurnool", AirTemparature: "35", Humidity: "20%", windspeed: "11 km/h"},
	{ID: "3", Name: "vizag", AirTemparature: "40", Humidity: "30%", windspeed: "20 km/h"},
}

func getweatherdevice(context *gin.Context) {
	context.JSON(http.StatusOK, weatherdevice)
}

func addweatherdevice(context *gin.Context) {
	var newdevice Weather
	if err := context.BindJSON(&newdevice); err != nil {
		return
	}
	weatherdevice = append(weatherdevice, newdevice)

	context.IndentedJSON(http.StatusCreated, newdevice)
}

func deletebyId(context *gin.Context) {
	id := context.Param("id")
	for i, p := range weatherdevice {
		if p.ID == id {
			weatherdevice = append(weatherdevice[:i], weatherdevice[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}
func main() {
	router := gin.Default()
	router.GET("/", getweatherdevice)
	router.POST("/", addweatherdevice)
	router.DELETE("/:id", deletebyId)

	router.Run("localhost:8080")
}
