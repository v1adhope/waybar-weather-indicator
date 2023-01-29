package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	_city         = "Penza"
	_timeModel24h = "15:04"
	_timeModel12h = "03:04PM"
)

// NOTE: Output structure. Read more
// https://github.com/Alexays/Waybar/wiki/Module:-Custom#return-type
type waybar struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}

func main() {
	uri := fmt.Sprintf("https://wttr.in/%s?format=j1", _city)

	var (
		resp *http.Response
		err  error

		attempts = 5
	)

	for attempts > 0 { //If the server temporarily does not respond
		resp, err = http.Get(uri)
		if err == nil {
			break
		}

		log.Printf("attempts left: %d", attempts)
		time.Sleep(1 * time.Minute)

		attempts--
	}
	if err != nil {
		log.Fatalf("request failed: %s", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("could not decode: %s", err)
	}

	var w waybar

	// Current weather block
	w.Text = fmt.Sprintf("%s°", data["current_condition"].([]interface{})[0].(map[string]interface{})["temp_C"])
	w.Text += fmt.Sprintf("(%s°)", data["current_condition"].([]interface{})[0].(map[string]interface{})["FeelsLikeC"])

	w.Tooltip = fmt.Sprintf("<b>Weather</b>\n")
	w.Tooltip += fmt.Sprintf("Current temp: %s°\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["temp_C"])
	w.Tooltip += fmt.Sprintf("Feels like: %s°\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["FeelsLikeC"])
	w.Tooltip += fmt.Sprintf("Humidity: %s%%\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["humidity"])
	w.Tooltip += fmt.Sprintf("Pressure: %s hPa\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["pressure"])
	w.Tooltip += fmt.Sprintf("Wind speed: %s Km/h\n\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["windspeedKmph"])

	// Solar block
	notProcessedTime := data["weather"].([]interface{})[0].(map[string]interface{})["astronomy"].([]interface{})[0].(map[string]interface{})["sunrise"].(string)
	sunriseTime, err := timeConvert(notProcessedTime)
	if err != nil {
		log.Fatalf("could not convert time: %s", err)
	}

	notProcessedTime = data["weather"].([]interface{})[0].(map[string]interface{})["astronomy"].([]interface{})[0].(map[string]interface{})["sunset"].(string)
	sunsetTime, err := timeConvert(notProcessedTime)
	if err != nil {
		log.Fatalf("could not convert time: %s", err)
	}

	w.Tooltip += fmt.Sprintf("<b>Solar cycle</b>\n")
	w.Tooltip += fmt.Sprintf("Sunrise at %s\n", sunriseTime)
	w.Tooltip += fmt.Sprintf("Sunset at %s", sunsetTime)

	// 3 days weather block
	w.Tooltip += fmt.Sprint("\n")
	weatherDays := 3

	hours := time.Now().Hour()

	for i := 0; i < weatherDays; i++ {
		switch i {
		case 0:
			w.Tooltip += fmt.Sprint("\n<b>Today</b>\n")
		case 1:
			w.Tooltip += fmt.Sprint("\n<b>Tomorrow</b>\n")
		case 2:
			w.Tooltip += fmt.Sprint("\n<b>After a day</b>\n")
		}

		for k := range data["weather"].([]interface{})[i].(map[string]interface{})["hourly"].([]interface{}) {
			wttrTime := k * 3 // Conversion into hours

			if hours > wttrTime+2 && i == 0 { // Exit if the watch is overdue
				continue
			}

			temp := data["weather"].([]interface{})[i].(map[string]interface{})["hourly"].([]interface{})[k].(map[string]interface{})["tempC"]
			feelsLike := data["weather"].([]interface{})[i].(map[string]interface{})["hourly"].([]interface{})[k].(map[string]interface{})["FeelsLikeC"]

			if wttrTime < 10 {
				w.Tooltip += fmt.Sprintf("At 0%d:00 %s°(%s°)\n", wttrTime, temp, feelsLike)
			} else {
				w.Tooltip += fmt.Sprintf("At %d:00 %s°(%s°)\n", wttrTime, temp, feelsLike)
			}
		}
	}

	w.Tooltip = strings.TrimSuffix(w.Tooltip, "\n")

	jsonOut, err := json.Marshal(w)
	if err != nil {
		log.Fatalf("could not encode: %s", err)
	}

	fmt.Print(string(jsonOut))
}

func timeConvert(target string) (string, error) {
	target = strings.ReplaceAll(target, " ", "")

	normalTime, err := time.Parse(_timeModel12h, target)
	if err != nil {
		return "", fmt.Errorf("could not convert time: %s", err)
	}

	return normalTime.Format(_timeModel24h), nil
}
