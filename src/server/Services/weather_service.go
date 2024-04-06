package Services

import (
	"net/http"
	"io/ioutil"
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
)

func Initialize_weather_service(r *gin.Context, token *oauth2.Token) string {

	url := "https://visual-crossing-weather.p.rapidapi.com/history?startDateTime=2023-02-05T00%3A00%3A00&aggregateHours=24&location=Paris&endDateTime=2023-02-08T00%3A00%3A00&unitGroup=us&dayStartTime=8%3A00%3A00&contentType=csv&dayEndTime=17%3A00%3A00&shortColumnNames=true"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "04d653f49bmsh1ebffe69731acabp11ea51jsn2474e9e338a2")
	req.Header.Add("X-RapidAPI-Host", "visual-crossing-weather.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func Weather_1_area(r *gin.Context, token *oauth2.Token) {
	println("ON EST DANS L'EXECUTION DE L'ACTION DE L'AREA !!")

	weather_service := Initialize_weather_service(r, token)
	println(weather_service)
}

func Init_weather_service_object(r *gin.Context) *Service {
	s := new(Service)
	s.Areas = map[string]func(r *gin.Context, token *oauth2.Token){
		"Weather_1": Weather_1_area,
	}
	return s
}