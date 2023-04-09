package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/v1adhope/waybar-weather-indicator/wttrstruct"
)

const (
	_timeModel12h = "03:04 PM"
	_timeModel24h = "15:04"
)

// NOTE: Output structure. Read more
// https://github.com/Alexays/Waybar/wiki/Module:-Custom#return-type
type waybar struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}

func main() {
	var city = "" // use ip location by default

	if envCity := os.Getenv("CITY_WEATHER"); envCity != "" {
		city = envCity
	}

	var (
		resp *http.Response
		err  error
		u    *url.URL
	)

	u, err = url.Parse(fmt.Sprintf("https://wttr.in/%s?format=j1", city))
	if err != nil {
		log.Fatalf("parse failed: %s", err)
	}

	//If the server temporarily does not respond
	for attempts := 5; attempts > 0; attempts-- {
		resp, err = http.Get(u.String())
		if attempts == 1 || err == nil {
			break
		}

		log.Printf("attempts left: %d", attempts-1)
		time.Sleep(time.Minute)
	}
	if err != nil {
		log.Fatalf("request failed: %s", err)
	}
	defer resp.Body.Close()

	var data wttrstruct.Data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("could not decode: %s", err)
	}

	var (
		w waybar
		b strings.Builder
	)

	// Current weather block
	fmt.Fprintf(&b, "%s°", data.CurrentCondition[0].TempC)
	fmt.Fprintf(&b, "(%s°)", data.CurrentCondition[0].FeelsLikeC)
	w.Text = b.String()
	b.Reset()

	fmt.Fprintf(&b, "<b>Weather</b>\n")
	fmt.Fprintf(&b, "Current temp: %s°\n", data.CurrentCondition[0].TempC)
	fmt.Fprintf(&b, "Feels like: %s°\n", data.CurrentCondition[0].FeelsLikeC)
	fmt.Fprintf(&b, "Humidity: %s%%\n", data.CurrentCondition[0].Humidity)
	fmt.Fprintf(&b, "Pressure: %s hPa\n", data.CurrentCondition[0].Pressure)
	fmt.Fprintf(&b, "Wind speed: %s Km/h\n\n", data.CurrentCondition[0].WindspeedKmph)

	// Solar block
	notProcessedTime := data.Weather[0].Astronomy[0].SunRise
	sunriseTime, err := timeConvert(notProcessedTime)
	if err != nil {
		log.Fatalf("could not convert time: %s", err)
	}

	notProcessedTime = data.Weather[0].Astronomy[0].SunSet
	sunsetTime, err := timeConvert(notProcessedTime)
	if err != nil {
		log.Fatalf("could not convert time: %s", err)
	}

	fmt.Fprintf(&b, "<b>Solar cycle</b>\n")
	fmt.Fprintf(&b, "Sunrise at %s\n", sunriseTime)
	fmt.Fprintf(&b, "Sunset at %s\n", sunsetTime)

	// 3 days weather block
	weatherDays := 3

	hours := time.Now().Hour()

	for i := 0; i < weatherDays; i++ {
		switch i {
		case 0:
			fmt.Fprintf(&b, "\n<b>Today</b>\n")
		case 1:
			fmt.Fprintf(&b, "\n<b>Tomorrow</b>\n")
		case 2:
			fmt.Fprintf(&b, "\n<b>After a day</b>\n")
		}

		for k, v := range data.Weather[i].Hourly {
			wttrTime := k * 3 // Conversion into hours

			if hours > wttrTime+2 && i == 0 { // Exit if the watch is overdue
				continue
			}

			temp := v.TempC
			feelsLike := v.FeelsLikeC
			temp, err = alignment(temp)
			if err != nil {
				log.Fatal(err)
			}

			feelsLike, err = alignment(feelsLike)
			if err != nil {
				log.Fatal(err)
			}

			if wttrTime < 10 {
				fmt.Fprintf(&b, "At 0%d:00 %s°(%s°)\n", wttrTime, temp, feelsLike)
			} else {
				fmt.Fprintf(&b, "At %d:00 %s°(%s°)\n", wttrTime, temp, feelsLike)
			}
		}
	}

	w.Tooltip = strings.TrimSuffix(b.String(), "\n")

	json, err := json.Marshal(w)
	if err != nil {
		log.Fatalf("could not encode: %s", err)
	}

	fmt.Print(string(json))
}

func timeConvert(target string) (string, error) {
	time, err := time.Parse(_timeModel12h, target)
	if err != nil {
		return "", err
	}

	return time.Format(_timeModel24h), nil
}

func alignment(target string) (string, error) {
	switch len(target) {
	default:
		return "", fmt.Errorf("bad data")
	case 1:
		return fmt.Sprintf(" %s", target), nil
	case 2:
		return fmt.Sprintf("%s", target), nil
	case 3:
		return target, nil
	}
}
