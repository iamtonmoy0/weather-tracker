package main
import(
	"net/http"
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
	err=json.Unmarshal(bytes,&c)
	if err !=nil{
		return apiConfigData{},err
	}
	return c , nil
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
}
)
	http.ListenAndServe(":8080",nil)

}
