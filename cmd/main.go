package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/v1adhope/waybar-weather-indicator/structure"
)

const (
	_reqAttempts = 10

	_timeModel12h = "03:04 PM"
	_timeModel24h = "15:04"
)

// Read more https://github.com/Alexays/Waybar/wiki/Module:-Custom#return-type
type module struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}

func main() {
	u, err := url.Parse("https://wttr.in/" + os.Getenv("CITY_WEATHER") + "?format=j1")
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	data := structure.Data{}
	func() {
		resp, err := &http.Response{}, error(nil)
		for range _reqAttempts {
			resp, err = http.Get(u.String())
			if resp.StatusCode == 200 {
				break
			}
			time.Sleep(30 * time.Second)
		}
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Fatalf("get: %v", err)
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			log.Fatalf("decode: %v", err)
		}
	}()

	module := module{}
	module.Text = fmt.Sprintf(
		"%s°(%s°)",
		data.CurrentCondition[0].TempC,
		data.CurrentCondition[0].FeelsLikeC,
	)

	buf := bytes.Buffer{}
	for _, row := range []struct{ Title, Value string }{
		{"Current temp:", data.CurrentCondition[0].TempC + "°"},
		{"Feels like:", data.CurrentCondition[0].FeelsLikeC + "°"},
		{"Humidity:", data.CurrentCondition[0].Humidity + "%"},
		{"Pressure:", data.CurrentCondition[0].Pressure + "hPa"},
		{"Wind speed:", data.CurrentCondition[0].WindspeedKmph + "km/h"},
		{"Description:", data.CurrentCondition[0].WeatherDesc[0].Value},
	} {
		fmt.Fprintln(&buf, row.Title+" "+row.Value)
	}

	sunriseTime, err := timeConvertFrom12to24H(data.Weather[0].Astronomy[0].SunRise)
	if err != nil {
		log.Fatalf("sunrise convert: %v", err)
	}
	sunsetTime, err := timeConvertFrom12to24H(data.Weather[0].Astronomy[0].SunSet)
	if err != nil {
		log.Fatalf("sunset convert: %v", err)
	}
	fmt.Fprint(&buf, "\n<b>Solar cycle</b>\n")
	fmt.Fprintln(&buf, "Sunrise at "+sunriseTime)
	fmt.Fprintln(&buf, "Sunset at "+sunsetTime)

	days, nowHour := 3, time.Now().Hour()
	for day := range days {
		switch day {
		case 0:
			fmt.Fprintf(&buf, "\n<b>Today</b>\n")
		case 1:
			fmt.Fprintf(&buf, "\n<b>Tomorrow</b>\n")
		case 2:
			fmt.Fprintf(&buf, "\n<b>After a day</b>\n")
		}
		for idx, weather := range data.Weather[day].Hourly {
			tWindow := idx * 3
			// Skip if time is overdue
			if day == 0 && nowHour > tWindow+2 {
				continue
			}
			if tWindow < 10 {
				fmt.Fprintf(&buf, "At 0%d:00 %2s°(%2s°) %s\n", tWindow, weather.TempC, weather.FeelsLikeC, checkDescription(weather.WeatherDesc[0].Value))
			} else {
				fmt.Fprintf(&buf, "At %d:00 %2s°(%2s°) %s\n", tWindow, weather.TempC, weather.FeelsLikeC, checkDescription(weather.WeatherDesc[0].Value))
			}
		}
	}

	fmt.Fprint(&buf, "\n<b>Timestamp</b>\n")
	fmt.Fprint(&buf, time.Now().Format(time.DateTime))

	module.Tooltip = buf.String()

	json, err := json.Marshal(module)
	if err != nil {
		log.Fatalf("marshal: %v", err)
	}
	fmt.Println(string(json))
}

func checkDescription(target string) string {
	if strings.Contains(strings.ToLower(target), "rain") {
		return "r"
	}

	if strings.Contains(strings.ToLower(target), "snow") {
		return "s"
	}

	return ""
}

func timeConvertFrom12to24H(target string) (string, error) {
	time, err := time.Parse(_timeModel12h, target)
	if err != nil {
		return "", err
	}

	return time.Format(_timeModel24h), nil
}
