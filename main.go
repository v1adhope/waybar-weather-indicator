package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type weather struct {
	temp, likeFeels string
}

func main() {
	uri := "https://wttr.in/Penza?format=j1"

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	var wx weather
	wx.temp = fmt.Sprintf("%v", data["current_condition"].([]interface{})[0].(map[string]interface{})["temp_C"])
	wx.likeFeels = fmt.Sprintf("%v", data["current_condition"].([]interface{})[0].(map[string]interface{})["FeelsLikeF"])
	fmt.Printf("%v(%v)", wx.temp, wx.likeFeels)
}
