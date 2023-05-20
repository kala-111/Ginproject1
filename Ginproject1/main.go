package main
import (
	"net/http"
	"gin.com/gin-gonic/gin"
)
type Weather struct{

	 ID             string 'json:"id"'
	 Name           string 'json:"name"'
	 AirTemparature string 'json:"airtempareture"'
	 Humidity       string 'json:"humidity"'
	 windspeed      string 'json:"windspeed"'  
	}
	var weatherdevice =[]weather{
		{ID:"1",Name:"vijayawada",AirTemparature:"40",Humidity:"25%",windspeed:"10 km/h"}
		{ID:"2",Name:"kurnool",AirTemparature:"35",Humidity:"20%",windspeed:"11 km/h"}
		{ID:"3",Name:"vizag",AirTemparature:"40",Humidity:"30%",windspeed:"20 km/h"}
	}
	func getweatherdevice(context *gin.context){
		context.JOSN(http.StatusOK,weather)
	}

	func addweatherdevice(context *gin.context){
		var newdevice Weather  
		if err :=context.BindJSON(&newdevice);err!=nil{
			return
		}
		weatherdevice=append(weatherdevice,newdevice)

		context.IndentedJSON(http.StatusCreated,newdevice)
	}

	func deleteby Id(context *gin.context){
		id:=c.Param("id")
		for i,p:=range newdevice{
			if p.ID==id{
				newdevice=append(newdevice[:i]newdevice[i+1:]...)
				c.IndentedJSON(http.StatusOK, gin.H{"message":"deleted"})
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"not found"})
	}
	func main(){
		router:=gin.Deafult()
		router.GET("/",getweatherdevice)
		router.POST("/",addweatherdevice)
		router.DELETE("/",deleteby id)
		
		router.Run("localhost:8080")
	}