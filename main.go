package main
import(
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:OpenWeatherMapApiKey`
}

type weatherData struct{
	Name struct `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}

func loadApiConfig(filename string)(apiConfigData,error){
	bytes,err:ioutil.ReadFile(filename)
	if err != nil{
		return apiConfigData{},err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes,&c)

	if err !=nil{
		return apiConfigData{},err
	}
	return c , nil
}
func hello(w http.ResponseWriter, r *http.Request){
	w.write([]byte("hello from go !\n"))
}

func query(city string)(weatherData,error){
	apiConfig,err:= loadApiConfig(".apiConfig")
	if err!= nil{
		return weatherData{},err
	}
	http.Get("http://ap.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey)
}
func main() {
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/weather/",
func(w http.ResponseWriter, r *http.Request){
	city := strings.SplitN(e.Url.Path,"/",3)[2]
	date.err := query(city)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
})
	http.ListenAndServe(":8080",nil)

}
