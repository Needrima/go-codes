package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

//splits form input to information to call Api on and actual data
//eg Weather Lagos will be split to "weather" and "lagos"
func getInfoTypeAndInfo(formInput string) (string, string) {
	trimString := strings.Trim(formInput, " ")

	infoParts := strings.Split(trimString, " ")

	infoType := infoParts[0]
	infoType = strings.ToLower(infoType)

	info := infoParts[1]
	info = strings.ToLower(info)

	return infoType, info
}

func HanldeRequest(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	if r.Method == http.MethodPost {
		formInput := r.FormValue("do")
		infoType, info := getInfoTypeAndInfo(formInput)

		switch infoType {
		case "weather":
			type WeatherInfo struct {
				Temparature float64 `json:"temp"`
				FeelsLike   float64 `json:"feels_like"`
				MinTemp     float64 `json:"temp_min"`
				MaxTemp     float64 `json:"temp_max"`
				Pressure    float64 `json:"pressure"`
				Humidity    float64 `json:"humidity"`
				SeaLevel    float64 `json:"sea_level"`
				GroundLevel float64 `json:"grnd_level"`
			}
			
			type Weather struct {
				Location    string `json:"name"`
				WeatherInfo `json:"main"`
			}

			var we Weather
			resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=da8d44e4505f5153cf700b5eeeb1885d&units=metric", info))
			if err != nil {
				http.Error(w, "Something went wrong", 500)
				return
			}
			defer resp.Body.Close()

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "Something went wrong", 500)
				return
			}

			err = json.Unmarshal(respBody, &we)
			if err != nil {
				http.Error(w, "Something went wrong", 500)
				return
			}

			tpl.Execute(w, we)
		default:
			http.Error(w, "Invalid input", 400)
		}
		return
	}

	tpl.Execute(w, nil)
}

func main() {
	fmt.Println("Hello world!")
}
