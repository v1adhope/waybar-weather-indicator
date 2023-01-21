package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const city = "Penza"

// NOTE: Output structure. Read more
// https://github.com/Alexays/Waybar/wiki/Module:-Custom#return-type
type weather struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}

func main() {
	uri := fmt.Sprintf("https://wttr.in/%s?format=j1", city)

	var (
		resp *http.Response
		err  error

		attempts = 5
	)

	//NOTE: If the server temporarily does not respond
	for attempts > 0 {
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

	var w weather
	w.Text = fmt.Sprintf("%s°", data["current_condition"].([]interface{})[0].(map[string]interface{})["temp_C"])
	w.Text += fmt.Sprintf("(%s°)", data["current_condition"].([]interface{})[0].(map[string]interface{})["FeelsLikeC"])

	w.Tooltip = fmt.Sprintf("<b>Weather</b>\n")
	w.Tooltip += fmt.Sprintf("Current temperature: %s°\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["temp_C"])
	w.Tooltip += fmt.Sprintf("Feels like: %s°\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["FeelsLikeC"])
	w.Tooltip += fmt.Sprintf("Humidity: %s%%\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["humidity"])
	w.Tooltip += fmt.Sprintf("Pressure: %s hPa\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["pressure"])
	w.Tooltip += fmt.Sprintf("Wind speed: %s Km/h\n\n", data["current_condition"].([]interface{})[0].(map[string]interface{})["windspeedKmph"])

	w.Tooltip += fmt.Sprintf("<b>Solar cycle</b>\n")
	w.Tooltip += fmt.Sprintf("Sunrise at %s\n", data["weather"].([]interface{})[0].(map[string]interface{})["astronomy"].([]interface{})[0].(map[string]interface{})["sunrise"])
	w.Tooltip += fmt.Sprintf("Sunset at %s", data["weather"].([]interface{})[0].(map[string]interface{})["astronomy"].([]interface{})[0].(map[string]interface{})["sunset"])

	jsonOut, err := json.Marshal(w)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(jsonOut))
}
